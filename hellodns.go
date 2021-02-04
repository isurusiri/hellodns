// An example CoreDNS plugin.
package hellodns

import (
	"context"

	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"

	"github.com/miekg/dns"
)

// Define a logger with the plugin name.
var log = clog.NewWithPlugin("hellodns")

type HelloDNS struct {
	Next plugin.Handler
}

func (h HelloDNS) ServeDNS(ctx conext.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	log.Debug("Receive response")

	pw := NewResponsePrinter(w)

	requestCount.WithLabelValues(metrics.WithServer(ctx)).Inc()

	return plugin.NextOrFailure(h.Name(), h.Next, ctx, pw, r)
}

func (h HelloDNS) Name() string {
	return "Hello"
}

type ResponsePrinter struct {
	dns.ResponseWriter
}

func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}

func (r *ResponsePrinter) WriteMsg(res *dns.Msg) error {
	log.Info("Hello")
	return r.ResponseWriter.WriteMsg(res)
}
