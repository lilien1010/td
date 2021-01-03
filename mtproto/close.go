package mtproto

// Close closes underlying connection.
func (c *Conn) Close() error {
	c.cancel()

	c.rpc.Close()

	if err := c.conn.Close(); err != nil {
		return err
	}

	// Probably we should wait with timeout, but it is unclear
	// whether we can try to save session or should hard fail on timeout.
	c.wg.Wait()

	return nil
}
