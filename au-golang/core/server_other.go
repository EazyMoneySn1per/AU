// +build !windows

package core

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 50 * time.Millisecond
	s.WriteTimeout = 50 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
