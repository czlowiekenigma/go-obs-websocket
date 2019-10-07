package requests

import (
	"github.com/czlowiekenigma/go-obs-websocket/responses"
	"reflect"
)

type Request interface {
	SetMessageID(uid string)
	GetResponseChannel() chan responses.Response
	SetResponseChannel(rchan chan responses.Response)
	GetResponseType() responses.Response
}

type RequestBase struct {
	MessageID   string `json:"message-id"`
	RequestType string `json:"request-type"`
	response    chan responses.Response
	rType       responses.Response
}

func (r *RequestBase) SetMessageID(ID string) {
	r.MessageID = ID
}

func (r *RequestBase) GetResponseChannel() chan responses.Response {
	return r.response
}

func (r *RequestBase) SetResponseChannel(rchan chan responses.Response) {
	r.response = rchan
}

func (r *RequestBase) GetResponseType() responses.Response {
	return r.rType
}

func ForgeRequest(req Request) (new Request) {
	ps := reflect.ValueOf(req)
	s := ps.Elem()

	if s.Kind() != reflect.Struct {
		return
	}

	f := s.FieldByName("RequestBase")
	if !f.IsValid() {
		return
	}

	if !f.CanSet() {
		return
	}

	f.Set(reflect.ValueOf(RequestBase{
		RequestType: reflect.TypeOf(req).Elem().Name(),
		rType: &responses.ResponseBase{},
	}))
	new = req
	return
}

func ForgeRequestWithExpectedResponse(resp responses.Response, reqs... Request) (new Request) {
	if len(reqs) != 0 {
		req := reqs[0]

		ps := reflect.ValueOf(req)
		s := ps.Elem()

		if s.Kind() != reflect.Struct {
			return
		}

		f := s.FieldByName("RequestBase")
		if !f.IsValid() {
			return
		}

		if !f.CanSet() {
			return
		}

		f.Set(reflect.ValueOf(RequestBase{
			RequestType: reflect.TypeOf(req).Elem().Name(),
			rType: resp,
		}))
		new = req
		return
	}

	new = &RequestBase{
		RequestType: reflect.TypeOf(resp).Elem().Name(),
		rType:       resp,
	}
	return
}