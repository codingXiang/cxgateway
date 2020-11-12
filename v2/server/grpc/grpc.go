package grpc

import (
	"fmt"
	"github.com/codingXiang/configer/v2"
	protocol2 "github.com/codingXiang/cxgateway/v2/util/protocol"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

var (
	Gateway *Server
)

type Server struct {
	config   *viper.Viper
	listener net.Listener
	server   *grpc.Server
}

func New(listener net.Listener, config *viper.Viper) *Server {
	s := new(Server)
	s.config = config

	var (
		err      error
		protocol = protocol2.New(config.GetString(configer.GetConfigPath(GRPC, PROTOCOL)))
		port     = config.GetInt(configer.GetConfigPath(GRPC, PORT))
	)

	if listener == nil {
		s.listener, err = net.Listen(protocol.String(), fmt.Sprintf(":%d", port))
		if err != nil {
			panic(err)
		}
	} else {
		s.listener = listener
	}

	s.server = grpc.NewServer()

	return s
}

func (s *Server) Run() {
	s.server.Serve(s.listener)
}

func (s *Server) GetServer() *grpc.Server {
	return s.server
}

func (s *Server) RunBackground() {
	go s.Run()
}
