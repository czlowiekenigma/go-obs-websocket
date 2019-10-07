package client

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/czlowiekenigma/go-obs-websocket/requests"
	"github.com/czlowiekenigma/go-obs-websocket/responses"
)

func (c *Client) GetVersion() (resp *responses.GetVersion, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetVersion{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetVersion)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetAuthRequired() (resp *responses.GetAuthRequired, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetAuthRequired{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetAuthRequired)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) Authenticate(password string, resp *responses.GetAuthRequired) (err error) {
	h := sha256.New()
	h.Write([]byte(password))
	h.Write([]byte(resp.Salt))
	secret := base64.StdEncoding.EncodeToString(h.Sum([]byte(nil)))

	h = sha256.New()
	h.Write([]byte(secret))
	h.Write([]byte(resp.Challenge))
	auth := base64.StdEncoding.EncodeToString(h.Sum([]byte(nil)))

	req := &requests.Authenticate{
		Auth: auth,
	}

	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) SetHeartbeat(req *requests.SetHeartbeat) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) SetFilenameFormatting(req *requests.SetFilenameFormatting) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetFilenameFormatting() (resp *responses.GetFilenameFormatting, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetFilenameFormatting{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetFilenameFormatting)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetStats() (resp *responses.GetStats, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetStats{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetStats)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) BroadcastCustomMessage(req *requests.BroadcastCustomMessage) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetVideoInfo() (resp *responses.GetVideoInfo, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetVideoInfo{}))
	if err != nil {
		return
	}

	var ok bool
	resp, ok = raw.(*responses.GetVideoInfo)
	if !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}