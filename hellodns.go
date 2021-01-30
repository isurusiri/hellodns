package hellodns

import (
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
)

// Define a logger with the plugin name.
var log = clog.NewWithPlugin("hellodns")

type HelloDNS struct {
	Next plugin.Handler
}
