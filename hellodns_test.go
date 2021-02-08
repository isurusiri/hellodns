package hellodns

import (
	"bytes"
	"context"
	golog "log"
	"strings"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	"github.com/miekg/dns"
)

func TestExample(t *testing.T) {
	x := HelloDNS{Next: test.ErrorHandler()}

	b := &bytes.Buffer{}
	golog.SetOutput(b)

	ctx := context.TODO()
	r := new(dns.Msg)
	r.SetQuestion("hellodns.org.", dns.TypeA)
	
	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	// Call our plugin directly, and check the result.
	x.ServeDNS(ctx, rec, r)
	if a := b.String(); !strings.Contains(a, "[INFO] plugin/hellodns: hellodns") {
		t.Errorf("Failed to print '%s', got %s", "[INFO] plugin/hellodns: hellodns", a)
	}
}
