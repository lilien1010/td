package mtproto

import "github.com/gotd/td/tg"

// Config returns current config received during setting init.
func (c *Conn) Config() tg.Config {
	cfg := c.cfg

	return cfg
}
