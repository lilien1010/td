package session

import (
	"github.com/ogen-go/errors"

	"github.com/gotd/td/session/tdesktop"
	"github.com/gotd/td/telegram/dcs"
	"github.com/gotd/td/tg"
)

func findDCAddr(list []tg.DCOption, dcID int) string {
	for _, opt := range list {
		if opt.ID != dcID {
			continue
		}
		if opt.TCPObfuscatedOnly ||
			opt.CDN ||
			opt.MediaOnly {
			continue
		}

		return opt.IPAddress
	}

	return ""
}

func mapConfig(mainDC int, cfg tdesktop.MTPConfig) tg.Config {
	var opts []tg.DCOption
	for _, dc := range cfg.DCOptions.Options {
		opts = append(opts, tg.DCOption{
			Flags:             dc.Flags,
			Ipv6:              dc.IPv6(),
			MediaOnly:         dc.MediaOnly(),
			TCPObfuscatedOnly: dc.TCPOOnly(),
			CDN:               dc.CDN(),
			Static:            dc.Static(),
			ID:                int(dc.ID),
			IPAddress:         dc.IP,
			Port:              int(dc.Port),
			Secret:            dc.Secret,
		})
	}
	return tg.Config{
		DCOptions:               opts,
		ThisDC:                  mainDC,
		ChatSizeMax:             int(cfg.ChatSizeMax),
		MegagroupSizeMax:        int(cfg.MegagroupSizeMax),
		ForwardedCountMax:       int(cfg.ForwardedCountMax),
		OnlineUpdatePeriodMs:    int(cfg.OnlineUpdatePeriod),
		OfflineBlurTimeoutMs:    int(cfg.OfflineBlurTimeout),
		OfflineIdleTimeoutMs:    int(cfg.OfflineIdleTimeout),
		OnlineCloudTimeoutMs:    int(cfg.OnlineCloudTimeout),
		NotifyCloudDelayMs:      int(cfg.NotifyCloudDelay),
		NotifyDefaultDelayMs:    int(cfg.NotifyDefaultDelay),
		SavedGifsLimit:          int(cfg.SavedGifsLimit),
		EditTimeLimit:           int(cfg.EditTimeLimit),
		RevokeTimeLimit:         int(cfg.RevokeTimeLimit),
		RevokePmTimeLimit:       int(cfg.RevokePrivateTimeLimit),
		RevokePmInbox:           cfg.RevokePrivateInbox,
		StickersRecentLimit:     int(cfg.StickersRecentLimit),
		StickersFavedLimit:      int(cfg.StickersFavedLimit),
		PinnedDialogsCountMax:   int(cfg.PinnedDialogsCountMax),
		PinnedInfolderCountMax:  int(cfg.PinnedDialogsInFolderMax),
		MeURLPrefix:             cfg.InternalLinksDomain,
		ChannelsReadMediaPeriod: int(cfg.ChannelsReadMediaPeriod),
		CallReceiveTimeoutMs:    int(cfg.CallReceiveTimeoutMs),
		CallRingTimeoutMs:       int(cfg.CallRingTimeoutMs),
		CallConnectTimeoutMs:    int(cfg.CallConnectTimeoutMs),
		CallPacketTimeoutMs:     int(cfg.CallPacketTimeoutMs),
		WebfileDCID:             int(cfg.WebFileDCID),
		DCTxtDomainName:         cfg.TxtDomainString,
		PhonecallsEnabled:       cfg.PhoneCallsEnabled,
		BlockedMode:             cfg.BlockedMode,
		CaptionLengthMax:        int(cfg.CaptionLengthMax),
	}
}

// TDesktopSession converts TDesktop's Account to Data.
func TDesktopSession(account tdesktop.Account) (*Data, error) {
	auth := account.Authorization
	cfg := account.Config
	dc := auth.MainDC

	key, ok := auth.Keys[dc]
	if !ok {
		return nil, errors.Errorf("key for main DC (%d) not found", dc)
	}
	keyID := key.ID()

	mappedCfg := mapConfig(dc, cfg)
	addr := findDCAddr(mappedCfg.DCOptions, dc)
	if addr == "" {
		// Fallback to built-in addrs.
		var list dcs.List
		if !cfg.Environment.Test() {
			list = dcs.Prod()
		} else {
			list = dcs.Test()
		}
		addr = findDCAddr(list.Options, dc)
		if addr == "" {
			return nil, errors.Errorf("can't find address for DC %d", dc)
		}
	}

	return &Data{
		DC:        dc,
		Addr:      addr,
		Config:    mappedCfg,
		AuthKey:   key[:],
		AuthKeyID: keyID[:],
	}, nil
}
