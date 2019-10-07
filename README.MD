# go-obs-websocket

[![GoDoc](https://godoc.org/github.com/czlowiekenigma/go-obs-websocket?status.svg)](https://godoc.org/github.com/czlowiekenigma/go-obs-websocket)

go-obs-websocket allows to connect to Open Broadcaster Studio plugin [obs-websocket](https://github.com/Palakis/obs-websocket).

## Installation

`go get github.com/czlowiekenigma/go-obs-websocket`

## Usage

Create a new client using
```gotemplate
client, err := obsws.NewClient(ip, port)
if err != nil {
    panic(err)
}
```

Check if auth is required, if it is - authenticate using password

```gotemplate
authRequired, err := client.GetAuthRequired()
if err != nil {
    panic(err)
}

if authRequired.AuthRequired {
    if err := client.Authenticate(password, authRequired); err != nil {
        panic(err)
    }
}
```

Send request without response
```gotemplate
if err := client.StartStopStreaming(); err != nil {
    panic(err)
}
```

Send request with expected response
```gotemplate
streamingStatusResponse, err := obsClient.GetStreamingStatus()
if err != nil {
	panic(err)
}

log.Printf("Streaming: %v, recording: %v\n", streamingStatusResponse.Streaming, streamingStatusResponse.Recording)
```

Handle events
```gotemplate
for {
    event := <-client.EventChannel()
    
    //Listen for StreamStopped event
    if _, ok := e.(*events.StreamStopped); ok {
        log.Println("Stream has stopped!")
    }
}
```
