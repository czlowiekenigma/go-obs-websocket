package client

import (
	"fmt"
	"github.com/czlowiekenigma/go-obs-websocket/requests"
	"github.com/czlowiekenigma/go-obs-websocket/responses"
)

func (c *Client) GetSourcesList() (resp *responses.GetSourcesList, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSourcesList{}))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetSourcesList); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetSourceTypesList() (resp *responses.GetSourceTypesList, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSourceTypesList{}))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetSourceTypesList); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetVolume(req *requests.GetVolume) (resp *responses.GetVolume, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetVolume{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetVolume); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetVolume(req *requests.SetVolume) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetMute(req *requests.GetMute) (resp *responses.GetMute, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetMute{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetMute); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetMute(req *requests.SetMute) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) SetSyncOffset(req *requests.SetSyncOffset) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetSyncOffset(req *requests.GetSyncOffset) (resp *responses.GetSyncOffset, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSyncOffset{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetSyncOffset); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetSourceSettings(req *requests.GetSourceSettings) (resp *responses.GetSourceSettings, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSourceSettings{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetSourceSettings); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetSourceSettings(req *requests.SetSourceSettings) (resp *responses.SetSourceSettings, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.SetSourceSettings{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.SetSourceSettings); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetTextGDIPlusProperties(req *requests.GetTextGDIPlusProperties) (resp *responses.GetTextGDIPlusProperties, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetTextGDIPlusProperties{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetTextGDIPlusProperties); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetTextGDIPlusProperties(req *requests.SetTextGDIPlusProperties) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetTextFreetype2Properties(req *requests.GetTextFreetype2Properties) (resp *responses.GetTextFreetype2Properties, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetTextFreetype2Properties{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetTextFreetype2Properties); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetTextFreetype2Properties(req *requests.SetTextFreetype2Properties) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetBrowserSourceProperties(req *requests.GetBrowserSourceProperties) (resp *responses.GetBrowserSourceProperties, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetBrowserSourceProperties{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetBrowserSourceProperties); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) SetBrowserSourceProperties(req *requests.SetBrowserSourceProperties) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) GetSpecialSources() (resp *responses.GetSpecialSources, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSpecialSources{}))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetSpecialSources); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) GetSourceFilters(req *requests.GetSourceFilters) (resp *responses.GetSourceFilters, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.GetSourceFilters{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.GetSourceFilters); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}

func (c *Client) AddFilterToSource(req *requests.AddFilterToSource) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) RemoveFilterFromSource(req *requests.RemoveFilterFromSource) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) ReorderSourceFilter(req *requests.ReorderSourceFilter) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) MoveSourceFilter(req *requests.MoveSourceFilter) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) SetSourceFilterSettings(req *requests.SetSourceFilterSettings) (err error) {
	_, err = c.submitRequest(requests.ForgeRequest(req))
	return
}

func (c *Client) TakeSourceScreenshot(req *requests.TakeSourceScreenshot) (resp *responses.TakeSourceScreenshot, err error) {
	raw, err := c.submitRequest(requests.ForgeRequestWithExpectedResponse(&responses.TakeSourceScreenshot{}, req))
	if err != nil {
		return
	}

	var ok bool
	if resp, ok = raw.(*responses.TakeSourceScreenshot); !ok {
		err = fmt.Errorf("obsws: unexpected response from server: %#v", resp)
	}
	return
}