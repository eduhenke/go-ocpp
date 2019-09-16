package cp

import (
	"github.com/eduhenke/go-ocpp/internal/service"
	"github.com/eduhenke/go-ocpp/messages/v1x/csreq"
	"github.com/eduhenke/go-ocpp/messages/v1x/csresp"
	"github.com/eduhenke/go-ocpp/soap"
	"github.com/eduhenke/go-ocpp/ws"
)

type Service interface {
	Send(csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error)
}

type SoapService struct {
	*service.SoapService
}

func NewSoapService(csURL string, options *soap.CallOptions) Service {
	return &SoapService{
		service.NewSoapService(csURL, options),
	}
}

func (service *SoapService) Send(req csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error) {
	rawResp, err := service.SoapService.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(csresp.CentralSystemResponse)
	if !ok {
		return nil, csresp.ErrorNotCentralSystemResponse
	}
	return resp, nil
}

type JsonService struct {
	*service.JsonService
}

func NewJsonService(conn *ws.Conn) Service {
	return &JsonService{service.NewJsonService(conn)}
}

func (service *JsonService) Send(req csreq.CentralSystemRequest) (csresp.CentralSystemResponse, error) {
	rawResp, err := service.JsonService.Send(req)
	if err != nil {
		return nil, err
	}
	resp, ok := rawResp.(csresp.CentralSystemResponse)
	if !ok {
		return nil, csresp.ErrorNotCentralSystemResponse
	}
	return resp, nil
}