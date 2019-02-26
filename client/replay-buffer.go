package client

import "github.com/defabricated/go-obs-websocket/requests"

func (c *Client) StartStopReplayBuffer() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StartStopReplayBuffer{}))
	return
}

func (c *Client) StartReplayBuffer() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StartReplayBuffer{}))
	return
}

func (c *Client) StopReplayBuffer() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.StopReplayBuffer{}))
	return
}

func (c *Client) SaveReplayBuffer() (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(&requests.SaveReplayBuffer{}))
	return
}
