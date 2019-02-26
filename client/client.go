package client

import (
	"encoding/json"
	"fmt"
	"github.com/defabricated/go-obs-websocket/events"
	"github.com/defabricated/go-obs-websocket/requests"
	"github.com/defabricated/go-obs-websocket/responses"
	"golang.org/x/net/websocket"
	"log"
	"sync"
)

type responseData struct {
	channel chan responses.Response
	rType   responses.Response
}

type Client struct {
	eventChannelLock sync.RWMutex
	wg               sync.WaitGroup

	ws *websocket.Conn

	events       chan events.Event
	requests     chan requests.Request
	frames       chan []byte
	responsesMap map[string]responseData

	closed bool
}

// NewClient connects to a websocket instance.
func NewClient(address string, port int) (*Client, error) {

	ws, err := websocket.Dial(fmt.Sprintf("ws://%s:%d/", address, port),
		"",
		fmt.Sprintf("http://%s:%d/", address, port))

	if err != nil {
		return nil, err
	}

	res := &Client{
		ws:           ws,
		requests:     make(chan requests.Request),
		responsesMap: make(map[string]responseData),
	}

	go res.internalLoop()
	return res, nil
}

func (c *Client) IsClosed() bool {
	return c.closed
}

func (c *Client) submitRequest(r requests.Request) (responses.Response, error) {
	rchan := make(chan responses.Response)
	r.SetResponseChannel(rchan)
	c.requests <- r
	resp, ok := <-rchan
	if ok == false {
		return nil, fmt.Errorf("obsws: internal error, response channel closed without response")
	}

	if err := resp.GetError(); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) handleResponse(frame []byte) {
	//check if the message is an event
	ev, err := events.UnmarshalEvent(frame)
	if err == nil {
		//check if use is listening events
		c.eventChannelLock.RLock()
		defer c.eventChannelLock.RUnlock()
		if c.events != nil {
			c.events <- ev
		}
		return
	}
	if _, ok := err.(events.ErrNotEventMessage); ok == false {
		//handle error
		if _, ok := err.(events.ErrUnknownEventType); ok == true {
			//we only log unknown eventype
			log.Printf("%s", err)
			return
		} else {
			log.Println(fmt.Sprintf("obsws: %s", err))
			c.Close()
			return
		}
	}

	// handle response
	var respBase responses.ResponseBase
	err = json.Unmarshal(frame, &respBase)
	if err != nil {
		log.Println(fmt.Sprintf("obsws: %s\n'%s'", err, frame))
		c.Close()
		return
	}

	respData, ok := c.responsesMap[respBase.GetMessageID()]
	if ok == false {
		log.Println(fmt.Sprintf("obsws: unknown message-id '%s'\n'%s'", respBase.GetMessageID(), frame))
		c.Close()
		return
	}
	err = json.Unmarshal(frame, &(respData.rType))
	if err != nil {
		log.Println(fmt.Sprintf("obsws: %s\n'%s'", err, frame))
		c.Close()
		return
	}
	respData.channel <- respData.rType
	delete(c.responsesMap, respBase.GetMessageID())
	close(respData.channel)
}

func (c *Client) internalLoop() {
	c.wg.Add(1)
	defer c.wg.Done()

	c.frames = make(chan []byte)

	requestUID := 0

	for {
		c.wg.Add(1)
		defer c.wg.Done()
		go func() {
			// read a response asynchronously
			frame := make([]byte, 0, 100)
			websocket.Message.Receive(c.ws, &frame)
			// send response to channel
			c.frames <- frame
		}()

		select {
		case f := <-c.frames:
			c.handleResponse(f)
		case r, ok := <-c.requests:
			if ok == false {
				c.requests = nil
				break
			}
			// send the right request, with an UID
			requestUID++
			rUID := fmt.Sprintf("%d", requestUID)
			c.responsesMap[rUID] = responseData{
				channel: r.GetResponseChannel(),
				rType:   r.GetResponseType(),
			}
			r.SetMessageID(rUID)
			websocket.JSON.Send(c.ws, r)
		}

		// we are closing the for loop
		if c.requests == nil {
			break
		}

	}

	//will discard the next response, either error or anything...
	c.frames = nil
}

// Close terminates the connection to the instance
func (c *Client) Close() {
	if c.closed {
		return
	}
	c.closed = true

	close(c.requests)
	// wait to be done
	c.wg.Wait()

	c.eventChannelLock.Lock()
	defer c.eventChannelLock.Unlock()
	if c.events != nil {
		close(c.events)
	}
}

// EventChannel returns a channel to read Event from
func (c *Client) EventChannel() <-chan events.Event {
	c.eventChannelLock.RLock()
	if c.events != nil {
		defer c.eventChannelLock.RUnlock()
		return c.events
	}

	res := make(chan chan events.Event)
	go func() {
		defer close(res)
		c.eventChannelLock.Lock()
		defer c.eventChannelLock.Unlock()
		if c.events != nil {
			res <- c.events
		}
		c.events = make(chan events.Event)
		res <- c.events
	}()
	c.eventChannelLock.RUnlock()
	events := <-res
	return events
}

func (c *Client) eventSender() chan<- events.Event {
	c.eventChannelLock.RLock()
	defer c.eventChannelLock.RUnlock()
	return c.events
}