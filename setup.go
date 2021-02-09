package hellodns

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

// init registers this plugin.
func init() { plugin.Register("hellodns", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // Ignore "example" and give us the next token.
	if c.NextArg() {
		return plugin.Error("hellodns", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return HelloDNS{Next: next}
	})

	return nil
}