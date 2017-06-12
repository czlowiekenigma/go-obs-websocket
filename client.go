package obsws

import (
	"encoding/json"
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

// A Client connects to a obs-studio websocket to get event and
// perform request on OBS instance remotely
type Client struct {
	eventChannelLock sync.RWMutex
	wg               sync.WaitGroup

	ws *websocket.Conn

	events   chan Event
	requests chan request
	frames   chan []byte
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
		ws:       ws,
		requests: make(chan request),
	}

	go res.internalLoop()
	return res, nil
}

func (c *Client) handleResponse(frame []byte) {
	//check if the message is an event
	var ev Event
	err := json.Unmarshal(frame, &ev)
	if err != nil && len(ev.UpdateType()) > 0 {
		// this a nice event
		c.eventChannelLock.RLock()
		defer c.eventChannelLock.RUnlock()
		if c.events != nil {
			c.events <- ev
		}
		return
	}

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
		case _, ok := <-c.requests:
			if ok == false {
				c.requests = nil
				break
			}
			// send the right request, with an UID
			requestUID++
		}

		// we are closing the for loop
		if c.requests == nil {
			break
		}

	}

	//will discard the next response, either error or anything...
	c.frames = nil
}

// Authentify performs the authenfication to this websocket instance.
func (c *Client) Authentify(psswd string) error {
	return NotYetImplemented()
}

// Close terminates the connection to the instance
func (c *Client) Close() {
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
func (c *Client) EventChannel() <-chan Event {
	c.eventChannelLock.RLock()
	if c.events != nil {
		defer c.eventChannelLock.RUnlock()
		return c.events
	}

	res := make(chan chan Event)
	go func() {
		defer close(res)
		c.eventChannelLock.Lock()
		defer c.eventChannelLock.Unlock()
		if c.events != nil {
			res <- c.events
		}
		c.events = make(chan Event)
		res <- c.events
	}()
	c.eventChannelLock.RUnlock()
	events := <-res
	return events
}

func (c *Client) eventSender() chan<- Event {
	c.eventChannelLock.RLock()
	defer c.eventChannelLock.RUnlock()
	return c.events
}
