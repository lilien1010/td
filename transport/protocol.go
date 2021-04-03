package transport

import (
	"net"

	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/proto/codec"
)

// Protocol is MTProto transport protocol.
//
// See https://core.telegram.org/mtproto/mtproto-transports
type Protocol struct {
	codec func() Codec
}

// NewProtocol creates new transport protocol using user Codec constructor.
//
// See https://core.telegram.org/mtproto/mtproto-transports
func NewProtocol(getCodec func() Codec) *Protocol {
	return &Protocol{
		codec: getCodec,
	}
}

// Telegram transport protocols.
//
// See https://core.telegram.org/mtproto/mtproto-transports
var (
	// Abridged is abridged transport protocol.
	//
	// See https://core.telegram.org/mtproto/mtproto-transports#abridged
	Abridged = NewProtocol(func() Codec {
		return codec.Abridged{}
	})

	// Intermediate is intermediate transport protocol.
	//
	// See https://core.telegram.org/mtproto/mtproto-transports#intermediate
	Intermediate = NewProtocol(func() Codec {
		return codec.Intermediate{}
	})

	// PaddedIntermediate is padded intermediate transport protocol.
	//
	// See https://core.telegram.org/mtproto/mtproto-transports#padded-intermediate
	PaddedIntermediate = NewProtocol(func() Codec {
		return codec.PaddedIntermediate{}
	})

	// Full is full transport protocol.
	//
	// See https://core.telegram.org/mtproto/mtproto-transports#full
	Full = NewProtocol(func() Codec {
		return &codec.Full{}
	})
)

// Codec creates new codec using protocol settings.
func (t *Protocol) Codec() Codec {
	return t.codec()
}

// Handshake inits given net.Conn as MTProto connection.
func (t *Protocol) Handshake(conn net.Conn) (Conn, error) {
	connCodec := t.codec()
	if err := connCodec.WriteHeader(conn); err != nil {
		return nil, xerrors.Errorf("write header: %w", err)
	}

	return &connection{
		conn:  conn,
		codec: connCodec,
	}, nil
}

// Pipe creates a in-memory MTProto connection.
func (t *Protocol) Pipe() (a, b Conn) {
	p1, p2 := net.Pipe()

	return &connection{
			conn:  p1,
			codec: t.codec(),
		}, &connection{
			conn:  p2,
			codec: t.codec(),
		}
}
