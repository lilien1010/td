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

// EncryptedMessage represents TL type `encryptedMessage#ed18c118`.
// Encrypted message.
//
// See https://core.telegram.org/constructor/encryptedMessage for reference.
type EncryptedMessage struct {
	// Random message ID, assigned by the author of message
	RandomID int64
	// ID of encrypted chat
	ChatID int
	// Date of sending
	Date int
	// TL-serialising of DecryptedMessage¹ type, encrypted with the key creatied at stage of
	// chat initialization
	//
	// Links:
	//  1) https://core.telegram.org/type/DecryptedMessage
	Bytes []byte
	// Attached encrypted file
	File EncryptedFileClass
}

// EncryptedMessageTypeID is TL type id of EncryptedMessage.
const EncryptedMessageTypeID = 0xed18c118

// construct implements constructor of EncryptedMessageClass.
func (e EncryptedMessage) construct() EncryptedMessageClass { return &e }

// Ensuring interfaces in compile-time for EncryptedMessage.
var (
	_ bin.Encoder     = &EncryptedMessage{}
	_ bin.Decoder     = &EncryptedMessage{}
	_ bin.BareEncoder = &EncryptedMessage{}
	_ bin.BareDecoder = &EncryptedMessage{}

	_ EncryptedMessageClass = &EncryptedMessage{}
)

func (e *EncryptedMessage) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.RandomID == 0) {
		return false
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.Bytes == nil) {
		return false
	}
	if !(e.File == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedMessage) String() string {
	if e == nil {
		return "EncryptedMessage(nil)"
	}
	type Alias EncryptedMessage
	return fmt.Sprintf("EncryptedMessage%+v", Alias(*e))
}

// FillFrom fills EncryptedMessage from given interface.
func (e *EncryptedMessage) FillFrom(from interface {
	GetRandomID() (value int64)
	GetChatID() (value int)
	GetDate() (value int)
	GetBytes() (value []byte)
	GetFile() (value EncryptedFileClass)
}) {
	e.RandomID = from.GetRandomID()
	e.ChatID = from.GetChatID()
	e.Date = from.GetDate()
	e.Bytes = from.GetBytes()
	e.File = from.GetFile()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedMessage) TypeID() uint32 {
	return EncryptedMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedMessage) TypeName() string {
	return "encryptedMessage"
}

// TypeInfo returns info about TL type.
func (e *EncryptedMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedMessage",
		ID:   EncryptedMessageTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
		{
			Name:       "File",
			SchemaName: "file",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedMessage) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedMessage#ed18c118 as nil")
	}
	b.PutID(EncryptedMessageTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedMessage) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedMessage#ed18c118 as nil")
	}
	b.PutLong(e.RandomID)
	b.PutInt(e.ChatID)
	b.PutInt(e.Date)
	b.PutBytes(e.Bytes)
	if e.File == nil {
		return fmt.Errorf("unable to encode encryptedMessage#ed18c118: field file is nil")
	}
	if err := e.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedMessage#ed18c118: field file: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedMessage) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedMessage#ed18c118 to nil")
	}
	if err := b.ConsumeID(EncryptedMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedMessage#ed18c118: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedMessage) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedMessage#ed18c118 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field random_id: %w", err)
		}
		e.RandomID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field bytes: %w", err)
		}
		e.Bytes = value
	}
	{
		value, err := DecodeEncryptedFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessage#ed18c118: field file: %w", err)
		}
		e.File = value
	}
	return nil
}

// GetRandomID returns value of RandomID field.
func (e *EncryptedMessage) GetRandomID() (value int64) {
	return e.RandomID
}

// GetChatID returns value of ChatID field.
func (e *EncryptedMessage) GetChatID() (value int) {
	return e.ChatID
}

// GetDate returns value of Date field.
func (e *EncryptedMessage) GetDate() (value int) {
	return e.Date
}

// GetBytes returns value of Bytes field.
func (e *EncryptedMessage) GetBytes() (value []byte) {
	return e.Bytes
}

// GetFile returns value of File field.
func (e *EncryptedMessage) GetFile() (value EncryptedFileClass) {
	return e.File
}

// EncryptedMessageService represents TL type `encryptedMessageService#23734b06`.
// Encrypted service message
//
// See https://core.telegram.org/constructor/encryptedMessageService for reference.
type EncryptedMessageService struct {
	// Random message ID, assigned by the author of message
	RandomID int64
	// ID of encrypted chat
	ChatID int
	// Date of sending
	Date int
	// TL-serialising of DecryptedMessage¹ type, encrypted with the key creatied at stage of
	// chat initialization
	//
	// Links:
	//  1) https://core.telegram.org/type/DecryptedMessage
	Bytes []byte
}

// EncryptedMessageServiceTypeID is TL type id of EncryptedMessageService.
const EncryptedMessageServiceTypeID = 0x23734b06

// construct implements constructor of EncryptedMessageClass.
func (e EncryptedMessageService) construct() EncryptedMessageClass { return &e }

// Ensuring interfaces in compile-time for EncryptedMessageService.
var (
	_ bin.Encoder     = &EncryptedMessageService{}
	_ bin.Decoder     = &EncryptedMessageService{}
	_ bin.BareEncoder = &EncryptedMessageService{}
	_ bin.BareDecoder = &EncryptedMessageService{}

	_ EncryptedMessageClass = &EncryptedMessageService{}
)

func (e *EncryptedMessageService) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.RandomID == 0) {
		return false
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.Bytes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedMessageService) String() string {
	if e == nil {
		return "EncryptedMessageService(nil)"
	}
	type Alias EncryptedMessageService
	return fmt.Sprintf("EncryptedMessageService%+v", Alias(*e))
}

// FillFrom fills EncryptedMessageService from given interface.
func (e *EncryptedMessageService) FillFrom(from interface {
	GetRandomID() (value int64)
	GetChatID() (value int)
	GetDate() (value int)
	GetBytes() (value []byte)
}) {
	e.RandomID = from.GetRandomID()
	e.ChatID = from.GetChatID()
	e.Date = from.GetDate()
	e.Bytes = from.GetBytes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EncryptedMessageService) TypeID() uint32 {
	return EncryptedMessageServiceTypeID
}

// TypeName returns name of type in TL schema.
func (*EncryptedMessageService) TypeName() string {
	return "encryptedMessageService"
}

// TypeInfo returns info about TL type.
func (e *EncryptedMessageService) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "encryptedMessageService",
		ID:   EncryptedMessageServiceTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RandomID",
			SchemaName: "random_id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EncryptedMessageService) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedMessageService#23734b06 as nil")
	}
	b.PutID(EncryptedMessageServiceTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EncryptedMessageService) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedMessageService#23734b06 as nil")
	}
	b.PutLong(e.RandomID)
	b.PutInt(e.ChatID)
	b.PutInt(e.Date)
	b.PutBytes(e.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (e *EncryptedMessageService) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedMessageService#23734b06 to nil")
	}
	if err := b.ConsumeID(EncryptedMessageServiceTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedMessageService#23734b06: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EncryptedMessageService) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedMessageService#23734b06 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field random_id: %w", err)
		}
		e.RandomID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedMessageService#23734b06: field bytes: %w", err)
		}
		e.Bytes = value
	}
	return nil
}

// GetRandomID returns value of RandomID field.
func (e *EncryptedMessageService) GetRandomID() (value int64) {
	return e.RandomID
}

// GetChatID returns value of ChatID field.
func (e *EncryptedMessageService) GetChatID() (value int) {
	return e.ChatID
}

// GetDate returns value of Date field.
func (e *EncryptedMessageService) GetDate() (value int) {
	return e.Date
}

// GetBytes returns value of Bytes field.
func (e *EncryptedMessageService) GetBytes() (value []byte) {
	return e.Bytes
}

// EncryptedMessageClass represents EncryptedMessage generic type.
//
// See https://core.telegram.org/type/EncryptedMessage for reference.
//
// Example:
//  g, err := tg.DecodeEncryptedMessage(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.EncryptedMessage: // encryptedMessage#ed18c118
//  case *tg.EncryptedMessageService: // encryptedMessageService#23734b06
//  default: panic(v)
//  }
type EncryptedMessageClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() EncryptedMessageClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Random message ID, assigned by the author of message
	GetRandomID() (value int64)

	// ID of encrypted chat
	GetChatID() (value int)

	// Date of sending
	GetDate() (value int)

	// TL-serialising of DecryptedMessage¹ type, encrypted with the key creatied at stage of
	// chat initialization
	//
	// Links:
	//  1) https://core.telegram.org/type/DecryptedMessage
	GetBytes() (value []byte)
}

// DecodeEncryptedMessage implements binary de-serialization for EncryptedMessageClass.
func DecodeEncryptedMessage(buf *bin.Buffer) (EncryptedMessageClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedMessageTypeID:
		// Decoding encryptedMessage#ed18c118.
		v := EncryptedMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedMessageClass: %w", err)
		}
		return &v, nil
	case EncryptedMessageServiceTypeID:
		// Decoding encryptedMessageService#23734b06.
		v := EncryptedMessageService{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedMessageClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedMessageClass: %w", bin.NewUnexpectedID(id))
	}
}

// EncryptedMessage boxes the EncryptedMessageClass providing a helper.
type EncryptedMessageBox struct {
	EncryptedMessage EncryptedMessageClass
}

// Decode implements bin.Decoder for EncryptedMessageBox.
func (b *EncryptedMessageBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EncryptedMessageBox to nil")
	}
	v, err := DecodeEncryptedMessage(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EncryptedMessage = v
	return nil
}

// Encode implements bin.Encode for EncryptedMessageBox.
func (b *EncryptedMessageBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EncryptedMessage == nil {
		return fmt.Errorf("unable to encode EncryptedMessageClass as nil")
	}
	return b.EncryptedMessage.Encode(buf)
}

// EncryptedMessageClassArray is adapter for slice of EncryptedMessageClass.
type EncryptedMessageClassArray []EncryptedMessageClass

// Sort sorts slice of EncryptedMessageClass.
func (s EncryptedMessageClassArray) Sort(less func(a, b EncryptedMessageClass) bool) EncryptedMessageClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EncryptedMessageClass.
func (s EncryptedMessageClassArray) SortStable(less func(a, b EncryptedMessageClass) bool) EncryptedMessageClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EncryptedMessageClass.
func (s EncryptedMessageClassArray) Retain(keep func(x EncryptedMessageClass) bool) EncryptedMessageClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s EncryptedMessageClassArray) First() (v EncryptedMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EncryptedMessageClassArray) Last() (v EncryptedMessageClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EncryptedMessageClassArray) PopFirst() (v EncryptedMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EncryptedMessageClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EncryptedMessageClassArray) Pop() (v EncryptedMessageClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of EncryptedMessageClass by Date.
func (s EncryptedMessageClassArray) SortByDate() EncryptedMessageClassArray {
	return s.Sort(func(a, b EncryptedMessageClass) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of EncryptedMessageClass by Date.
func (s EncryptedMessageClassArray) SortStableByDate() EncryptedMessageClassArray {
	return s.SortStable(func(a, b EncryptedMessageClass) bool {
		return a.GetDate() < b.GetDate()
	})
}

// AsEncryptedMessage returns copy with only EncryptedMessage constructors.
func (s EncryptedMessageClassArray) AsEncryptedMessage() (to EncryptedMessageArray) {
	for _, elem := range s {
		value, ok := elem.(*EncryptedMessage)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsEncryptedMessageService returns copy with only EncryptedMessageService constructors.
func (s EncryptedMessageClassArray) AsEncryptedMessageService() (to EncryptedMessageServiceArray) {
	for _, elem := range s {
		value, ok := elem.(*EncryptedMessageService)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// EncryptedMessageArray is adapter for slice of EncryptedMessage.
type EncryptedMessageArray []EncryptedMessage

// Sort sorts slice of EncryptedMessage.
func (s EncryptedMessageArray) Sort(less func(a, b EncryptedMessage) bool) EncryptedMessageArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EncryptedMessage.
func (s EncryptedMessageArray) SortStable(less func(a, b EncryptedMessage) bool) EncryptedMessageArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EncryptedMessage.
func (s EncryptedMessageArray) Retain(keep func(x EncryptedMessage) bool) EncryptedMessageArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s EncryptedMessageArray) First() (v EncryptedMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EncryptedMessageArray) Last() (v EncryptedMessage, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EncryptedMessageArray) PopFirst() (v EncryptedMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EncryptedMessage
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EncryptedMessageArray) Pop() (v EncryptedMessage, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of EncryptedMessage by Date.
func (s EncryptedMessageArray) SortByDate() EncryptedMessageArray {
	return s.Sort(func(a, b EncryptedMessage) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of EncryptedMessage by Date.
func (s EncryptedMessageArray) SortStableByDate() EncryptedMessageArray {
	return s.SortStable(func(a, b EncryptedMessage) bool {
		return a.GetDate() < b.GetDate()
	})
}

// EncryptedMessageServiceArray is adapter for slice of EncryptedMessageService.
type EncryptedMessageServiceArray []EncryptedMessageService

// Sort sorts slice of EncryptedMessageService.
func (s EncryptedMessageServiceArray) Sort(less func(a, b EncryptedMessageService) bool) EncryptedMessageServiceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of EncryptedMessageService.
func (s EncryptedMessageServiceArray) SortStable(less func(a, b EncryptedMessageService) bool) EncryptedMessageServiceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of EncryptedMessageService.
func (s EncryptedMessageServiceArray) Retain(keep func(x EncryptedMessageService) bool) EncryptedMessageServiceArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s EncryptedMessageServiceArray) First() (v EncryptedMessageService, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s EncryptedMessageServiceArray) Last() (v EncryptedMessageService, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *EncryptedMessageServiceArray) PopFirst() (v EncryptedMessageService, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero EncryptedMessageService
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *EncryptedMessageServiceArray) Pop() (v EncryptedMessageService, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByDate sorts slice of EncryptedMessageService by Date.
func (s EncryptedMessageServiceArray) SortByDate() EncryptedMessageServiceArray {
	return s.Sort(func(a, b EncryptedMessageService) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of EncryptedMessageService by Date.
func (s EncryptedMessageServiceArray) SortStableByDate() EncryptedMessageServiceArray {
	return s.SortStable(func(a, b EncryptedMessageService) bool {
		return a.GetDate() < b.GetDate()
	})
}
