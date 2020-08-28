package lisp

import (
	"github.com/aclivo/olap"
)

type server struct {
	olap.Server
}

func newLispServer(srv olap.Server) olap.Server {
	return &server{Server: srv}
}
