package obsws

import (
	"github.com/czlowiekenigma/go-obs-websocket/client"
)

func NewClient(address string, port int) (*client.Client, error) {
	return client.NewClient(address, port)
}

//var cl *client.Client
//
//func main() {
//	var err error
//	cl, err = NewClient("localhost", 4444)
//	if err != nil {
//		log.Error(err)
//		return
//	}
//
//	auth, err := cl.GetAuthRequired()
//	if err != nil {
//		log.Error(err)
//		return
//	}
//
//	if auth.AuthRequired {
//		if err := cl.Authenticate("boczek", auth); err != nil {
//			log.Error(err)
//			return
//		}
//	}
//
//	if err := cl.StartStopStreaming(); err != nil {
//		log.Error(err)
//		return
//	}
//
//	for {
//		e := <-cl.EventChannel()
//
//		go handleEvent(e)
//	}
//}
//
//func handleEvent(e events.Event) {
//	fmt.Println(reflect.TypeOf(e).Elem().Name())
//
//	if _, ok := e.(*events.StreamStopped); ok {
//		fmt.Println("stopped")
//
//		time.Sleep(5 * time.Second)
//
//		fmt.Println("start")
//		if err := cl.StartStreaming(&requests.StartStreaming{}); err != nil {
//			log.Error(err)
//		}
//	}
//
//	if _, ok := e.(*events.StreamStarted); ok {
//		fmt.Println("started")
//
//		time.Sleep(5 * time.Second)
//
//		fmt.Println("stop")
//		if err := cl.StopStreaming(); err != nil {
//			log.Error(err)
//		}
//	}
//}
