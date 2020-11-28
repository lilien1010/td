// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// AccountUploadWallPaperRequest represents TL type `account.uploadWallPaper#dd853661`.
type AccountUploadWallPaperRequest struct {
	// File field of AccountUploadWallPaperRequest.
	File InputFileClass
	// MimeType field of AccountUploadWallPaperRequest.
	MimeType string
	// Settings field of AccountUploadWallPaperRequest.
	Settings WallPaperSettings
}

// AccountUploadWallPaperRequestTypeID is TL type id of AccountUploadWallPaperRequest.
const AccountUploadWallPaperRequestTypeID = 0xdd853661

// Encode implements bin.Encoder.
func (u *AccountUploadWallPaperRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadWallPaper#dd853661 as nil")
	}
	b.PutID(AccountUploadWallPaperRequestTypeID)
	if u.File == nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#dd853661: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#dd853661: field file: %w", err)
	}
	b.PutString(u.MimeType)
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#dd853661: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUploadWallPaperRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadWallPaper#dd853661 to nil")
	}
	if err := b.ConsumeID(AccountUploadWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: %w", err)
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: field file: %w", err)
		}
		u.File = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: field mime_type: %w", err)
		}
		u.MimeType = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#dd853661: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUploadWallPaperRequest.
var (
	_ bin.Encoder = &AccountUploadWallPaperRequest{}
	_ bin.Decoder = &AccountUploadWallPaperRequest{}
)