// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// MessagesGetEmojiKeywordsLanguagesRequest represents TL type `messages.getEmojiKeywordsLanguages#4e9963b2`.
// Get info about an emoji keyword localization
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsLanguages for reference.
type MessagesGetEmojiKeywordsLanguagesRequest struct {
	// Language codes
	LangCodes []string
}

// MessagesGetEmojiKeywordsLanguagesRequestTypeID is TL type id of MessagesGetEmojiKeywordsLanguagesRequest.
const MessagesGetEmojiKeywordsLanguagesRequestTypeID = 0x4e9963b2

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsLanguagesRequest.
var (
	_ bin.Encoder     = &MessagesGetEmojiKeywordsLanguagesRequest{}
	_ bin.Decoder     = &MessagesGetEmojiKeywordsLanguagesRequest{}
	_ bin.BareEncoder = &MessagesGetEmojiKeywordsLanguagesRequest{}
	_ bin.BareDecoder = &MessagesGetEmojiKeywordsLanguagesRequest{}
)

func (g *MessagesGetEmojiKeywordsLanguagesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangCodes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) String() string {
	if g == nil {
		return "MessagesGetEmojiKeywordsLanguagesRequest(nil)"
	}
	type Alias MessagesGetEmojiKeywordsLanguagesRequest
	return fmt.Sprintf("MessagesGetEmojiKeywordsLanguagesRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetEmojiKeywordsLanguagesRequest from given interface.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) FillFrom(from interface {
	GetLangCodes() (value []string)
}) {
	g.LangCodes = from.GetLangCodes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetEmojiKeywordsLanguagesRequest) TypeID() uint32 {
	return MessagesGetEmojiKeywordsLanguagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetEmojiKeywordsLanguagesRequest) TypeName() string {
	return "messages.getEmojiKeywordsLanguages"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getEmojiKeywordsLanguages",
		ID:   MessagesGetEmojiKeywordsLanguagesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCodes",
			SchemaName: "lang_codes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywordsLanguages#4e9963b2 as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsLanguagesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywordsLanguages#4e9963b2 as nil")
	}
	b.PutVectorHeader(len(g.LangCodes))
	for _, v := range g.LangCodes {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywordsLanguages#4e9963b2 to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsLanguagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywordsLanguages#4e9963b2 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: field lang_codes: %w", err)
		}

		if headerLen > 0 {
			g.LangCodes = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getEmojiKeywordsLanguages#4e9963b2: field lang_codes: %w", err)
			}
			g.LangCodes = append(g.LangCodes, value)
		}
	}
	return nil
}

// GetLangCodes returns value of LangCodes field.
func (g *MessagesGetEmojiKeywordsLanguagesRequest) GetLangCodes() (value []string) {
	return g.LangCodes
}

// MessagesGetEmojiKeywordsLanguages invokes method messages.getEmojiKeywordsLanguages#4e9963b2 returning error if any.
// Get info about an emoji keyword localization
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsLanguages for reference.
func (c *Client) MessagesGetEmojiKeywordsLanguages(ctx context.Context, langcodes []string) ([]EmojiLanguage, error) {
	var result EmojiLanguageVector

	request := &MessagesGetEmojiKeywordsLanguagesRequest{
		LangCodes: langcodes,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []EmojiLanguage(result.Elems), nil
}
