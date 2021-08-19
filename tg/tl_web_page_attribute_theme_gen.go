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

// WebPageAttributeTheme represents TL type `webPageAttributeTheme#54b56617`.
// Page theme
//
// See https://core.telegram.org/constructor/webPageAttributeTheme for reference.
type WebPageAttributeTheme struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Theme files
	//
	// Use SetDocuments and GetDocuments helpers.
	Documents []DocumentClass
	// Theme settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings ThemeSettings
}

// WebPageAttributeThemeTypeID is TL type id of WebPageAttributeTheme.
const WebPageAttributeThemeTypeID = 0x54b56617

// Ensuring interfaces in compile-time for WebPageAttributeTheme.
var (
	_ bin.Encoder     = &WebPageAttributeTheme{}
	_ bin.Decoder     = &WebPageAttributeTheme{}
	_ bin.BareEncoder = &WebPageAttributeTheme{}
	_ bin.BareDecoder = &WebPageAttributeTheme{}
)

func (w *WebPageAttributeTheme) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Flags.Zero()) {
		return false
	}
	if !(w.Documents == nil) {
		return false
	}
	if !(w.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebPageAttributeTheme) String() string {
	if w == nil {
		return "WebPageAttributeTheme(nil)"
	}
	type Alias WebPageAttributeTheme
	return fmt.Sprintf("WebPageAttributeTheme%+v", Alias(*w))
}

// FillFrom fills WebPageAttributeTheme from given interface.
func (w *WebPageAttributeTheme) FillFrom(from interface {
	GetDocuments() (value []DocumentClass, ok bool)
	GetSettings() (value ThemeSettings, ok bool)
}) {
	if val, ok := from.GetDocuments(); ok {
		w.Documents = val
	}

	if val, ok := from.GetSettings(); ok {
		w.Settings = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*WebPageAttributeTheme) TypeID() uint32 {
	return WebPageAttributeThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*WebPageAttributeTheme) TypeName() string {
	return "webPageAttributeTheme"
}

// TypeInfo returns info about TL type.
func (w *WebPageAttributeTheme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "webPageAttributeTheme",
		ID:   WebPageAttributeThemeTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Documents",
			SchemaName: "documents",
			Null:       !w.Flags.Has(0),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !w.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *WebPageAttributeTheme) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeTheme#54b56617 as nil")
	}
	b.PutID(WebPageAttributeThemeTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *WebPageAttributeTheme) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webPageAttributeTheme#54b56617 as nil")
	}
	if !(w.Documents == nil) {
		w.Flags.Set(0)
	}
	if !(w.Settings.Zero()) {
		w.Flags.Set(1)
	}
	if err := w.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field flags: %w", err)
	}
	if w.Flags.Has(0) {
		b.PutVectorHeader(len(w.Documents))
		for idx, v := range w.Documents {
			if v == nil {
				return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field documents element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field documents element with index %d: %w", idx, err)
			}
		}
	}
	if w.Flags.Has(1) {
		if err := w.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webPageAttributeTheme#54b56617: field settings: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (w *WebPageAttributeTheme) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeTheme#54b56617 to nil")
	}
	if err := b.ConsumeID(WebPageAttributeThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *WebPageAttributeTheme) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webPageAttributeTheme#54b56617 to nil")
	}
	{
		if err := w.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field flags: %w", err)
		}
	}
	if w.Flags.Has(0) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field documents: %w", err)
		}

		if headerLen > 0 {
			w.Documents = make([]DocumentClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field documents: %w", err)
			}
			w.Documents = append(w.Documents, value)
		}
	}
	if w.Flags.Has(1) {
		if err := w.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode webPageAttributeTheme#54b56617: field settings: %w", err)
		}
	}
	return nil
}

// SetDocuments sets value of Documents conditional field.
func (w *WebPageAttributeTheme) SetDocuments(value []DocumentClass) {
	w.Flags.Set(0)
	w.Documents = value
}

// GetDocuments returns value of Documents conditional field and
// boolean which is true if field was set.
func (w *WebPageAttributeTheme) GetDocuments() (value []DocumentClass, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return w.Documents, true
}

// SetSettings sets value of Settings conditional field.
func (w *WebPageAttributeTheme) SetSettings(value ThemeSettings) {
	w.Flags.Set(1)
	w.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (w *WebPageAttributeTheme) GetSettings() (value ThemeSettings, ok bool) {
	if !w.Flags.Has(1) {
		return value, false
	}
	return w.Settings, true
}

// MapDocuments returns field Documents wrapped in DocumentClassArray helper.
func (w *WebPageAttributeTheme) MapDocuments() (value DocumentClassArray, ok bool) {
	if !w.Flags.Has(0) {
		return value, false
	}
	return DocumentClassArray(w.Documents), true
}
