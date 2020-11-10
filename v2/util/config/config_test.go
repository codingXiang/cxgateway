package config

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

//Suite struct
type Suite struct {
	suite.Suite
}

//初始化 Suite
func (s *Suite) SetupSuite() {

}

//TestStart 為測試程式進入點
func TestStartSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}
