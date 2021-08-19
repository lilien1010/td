// Code generated by gotdgen, DO NOT EDIT.

package td

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

// DCOption represents TL type `dcOption#18b7a10d`.
//
// See https://localhost:80/doc/constructor/dcOption for reference.
type DCOption struct {
	// Flags field of DCOption.
	Flags bin.Fields
	// Ipv6 field of DCOption.
	Ipv6 bool
	// MediaOnly field of DCOption.
	MediaOnly bool
	// TCPObfuscatedOnly field of DCOption.
	TCPObfuscatedOnly bool
	// CDN field of DCOption.
	CDN bool
	// Static field of DCOption.
	Static bool
	// ID field of DCOption.
	ID int
	// IPAddress field of DCOption.
	IPAddress string
	// Port field of DCOption.
	Port int
	// Secret field of DCOption.
	//
	// Use SetSecret and GetSecret helpers.
	Secret []byte
}

// DCOptionTypeID is TL type id of DCOption.
const DCOptionTypeID = 0x18b7a10d

// Ensuring interfaces in compile-time for DCOption.
var (
	_ bin.Encoder     = &DCOption{}
	_ bin.Decoder     = &DCOption{}
	_ bin.BareEncoder = &DCOption{}
	_ bin.BareDecoder = &DCOption{}
)

func (d *DCOption) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Ipv6 == false) {
		return false
	}
	if !(d.MediaOnly == false) {
		return false
	}
	if !(d.TCPObfuscatedOnly == false) {
		return false
	}
	if !(d.CDN == false) {
		return false
	}
	if !(d.Static == false) {
		return false
	}
	if !(d.ID == 0) {
		return false
	}
	if !(d.IPAddress == "") {
		return false
	}
	if !(d.Port == 0) {
		return false
	}
	if !(d.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DCOption) String() string {
	if d == nil {
		return "DCOption(nil)"
	}
	type Alias DCOption
	return fmt.Sprintf("DCOption%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DCOption) TypeID() uint32 {
	return DCOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*DCOption) TypeName() string {
	return "dcOption"
}

// TypeInfo returns info about TL type.
func (d *DCOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dcOption",
		ID:   DCOptionTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Ipv6",
			SchemaName: "ipv6",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "MediaOnly",
			SchemaName: "media_only",
			Null:       !d.Flags.Has(1),
		},
		{
			Name:       "TCPObfuscatedOnly",
			SchemaName: "tcpo_only",
			Null:       !d.Flags.Has(2),
		},
		{
			Name:       "CDN",
			SchemaName: "cdn",
			Null:       !d.Flags.Has(3),
		},
		{
			Name:       "Static",
			SchemaName: "static",
			Null:       !d.Flags.Has(4),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "IPAddress",
			SchemaName: "ip_address",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
			Null:       !d.Flags.Has(10),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DCOption) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dcOption#18b7a10d as nil")
	}
	b.PutID(DCOptionTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DCOption) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dcOption#18b7a10d as nil")
	}
	if !(d.Ipv6 == false) {
		d.Flags.Set(0)
	}
	if !(d.MediaOnly == false) {
		d.Flags.Set(1)
	}
	if !(d.TCPObfuscatedOnly == false) {
		d.Flags.Set(2)
	}
	if !(d.CDN == false) {
		d.Flags.Set(3)
	}
	if !(d.Static == false) {
		d.Flags.Set(4)
	}
	if !(d.Secret == nil) {
		d.Flags.Set(10)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dcOption#18b7a10d: field flags: %w", err)
	}
	b.PutInt(d.ID)
	b.PutString(d.IPAddress)
	b.PutInt(d.Port)
	if d.Flags.Has(10) {
		b.PutBytes(d.Secret)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DCOption) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dcOption#18b7a10d to nil")
	}
	if err := b.ConsumeID(DCOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode dcOption#18b7a10d: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DCOption) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dcOption#18b7a10d to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field flags: %w", err)
		}
	}
	d.Ipv6 = d.Flags.Has(0)
	d.MediaOnly = d.Flags.Has(1)
	d.TCPObfuscatedOnly = d.Flags.Has(2)
	d.CDN = d.Flags.Has(3)
	d.Static = d.Flags.Has(4)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field ip_address: %w", err)
		}
		d.IPAddress = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field port: %w", err)
		}
		d.Port = value
	}
	if d.Flags.Has(10) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode dcOption#18b7a10d: field secret: %w", err)
		}
		d.Secret = value
	}
	return nil
}

// SetIpv6 sets value of Ipv6 conditional field.
func (d *DCOption) SetIpv6(value bool) {
	if value {
		d.Flags.Set(0)
		d.Ipv6 = true
	} else {
		d.Flags.Unset(0)
		d.Ipv6 = false
	}
}

// GetIpv6 returns value of Ipv6 conditional field.
func (d *DCOption) GetIpv6() (value bool) {
	return d.Flags.Has(0)
}

// SetMediaOnly sets value of MediaOnly conditional field.
func (d *DCOption) SetMediaOnly(value bool) {
	if value {
		d.Flags.Set(1)
		d.MediaOnly = true
	} else {
		d.Flags.Unset(1)
		d.MediaOnly = false
	}
}

// GetMediaOnly returns value of MediaOnly conditional field.
func (d *DCOption) GetMediaOnly() (value bool) {
	return d.Flags.Has(1)
}

// SetTCPObfuscatedOnly sets value of TCPObfuscatedOnly conditional field.
func (d *DCOption) SetTCPObfuscatedOnly(value bool) {
	if value {
		d.Flags.Set(2)
		d.TCPObfuscatedOnly = true
	} else {
		d.Flags.Unset(2)
		d.TCPObfuscatedOnly = false
	}
}

// GetTCPObfuscatedOnly returns value of TCPObfuscatedOnly conditional field.
func (d *DCOption) GetTCPObfuscatedOnly() (value bool) {
	return d.Flags.Has(2)
}

// SetCDN sets value of CDN conditional field.
func (d *DCOption) SetCDN(value bool) {
	if value {
		d.Flags.Set(3)
		d.CDN = true
	} else {
		d.Flags.Unset(3)
		d.CDN = false
	}
}

// GetCDN returns value of CDN conditional field.
func (d *DCOption) GetCDN() (value bool) {
	return d.Flags.Has(3)
}

// SetStatic sets value of Static conditional field.
func (d *DCOption) SetStatic(value bool) {
	if value {
		d.Flags.Set(4)
		d.Static = true
	} else {
		d.Flags.Unset(4)
		d.Static = false
	}
}

// GetStatic returns value of Static conditional field.
func (d *DCOption) GetStatic() (value bool) {
	return d.Flags.Has(4)
}

// GetID returns value of ID field.
func (d *DCOption) GetID() (value int) {
	return d.ID
}

// GetIPAddress returns value of IPAddress field.
func (d *DCOption) GetIPAddress() (value string) {
	return d.IPAddress
}

// GetPort returns value of Port field.
func (d *DCOption) GetPort() (value int) {
	return d.Port
}

// SetSecret sets value of Secret conditional field.
func (d *DCOption) SetSecret(value []byte) {
	d.Flags.Set(10)
	d.Secret = value
}

// GetSecret returns value of Secret conditional field and
// boolean which is true if field was set.
func (d *DCOption) GetSecret() (value []byte, ok bool) {
	if !d.Flags.Has(10) {
		return value, false
	}
	return d.Secret, true
}
