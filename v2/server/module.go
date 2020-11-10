package server

import (
	"github.com/codingXiang/cxgateway/v2/util/i18n"
)

type HttpModule interface {
	SetI18n(moduleName string)
	GetI18n() i18n.I18nMessageHandlerInterface
	SetGateway(s *Server)
	GetGateway() *Server
	Setup()
}

type Http struct {
	i18nMsg i18n.I18nMessageHandlerInterface
	gateway *Server
}

func (h *Http) GetI18n() i18n.I18nMessageHandlerInterface {
	return h.i18nMsg
}
func (h *Http) SetI18n(n string) {
	h.i18nMsg = i18n.NewI18nMessageHandler(n)
}

func (h *Http) SetGateway(s *Server) {
	h.gateway = s
}

func (h *Http) GetGateway() *Server {
	return h.gateway
}

func (h *Http) Setup() {

}
