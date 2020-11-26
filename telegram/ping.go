package telegram

import (
	"context"

	"github.com/ernado/td/bin"
	"github.com/ernado/td/internal/proto"
)

type pingMessage struct {
	id int64
}

func (p pingMessage) Encode(b *bin.Buffer) error {
	b.PutID(0x7abe77ec)
	b.PutLong(p.id)
	return nil
}

func (c Client) Ping(ctx context.Context) error {
	b := new(bin.Buffer)
	if err := c.newEncryptedMessage(&pingMessage{id: 0xafef}, b); err != nil {
		return err
	}
	if err := proto.WriteIntermediate(c.conn, b); err != nil {
		return err
	}

	b.Reset()
	if err := proto.ReadIntermediate(c.conn, b); err != nil {
		return err
	}
	if err := c.checkProtocolError(b); err != nil {
		return err
	}

	encMessage := proto.EncryptedMessage{}
	if err := encMessage.Decode(b); err != nil {
		return err
	}

	// TODO(ernado): decrypt and parse pong response

	return nil
}