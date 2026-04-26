package wifi

import (
	"errors"
	"net"
	"sync/atomic"
	"time"

	"github.com/bettercap/bettercap/v2/network"
	"github.com/bettercap/bettercap/v2/packets"
	"github.com/bettercap/bettercap/v2/session"

	"github.com/evilsocket/islazy/tui"
)

var errNoRecon = errors.New("Module wifi.ap requires module wifi.recon to be activated.")

func (mod *WiFiModule) parseApConfig() (err error) {
	var bssid string
	if err, mod.apConfig.SSID = mod.StringParam("wifi.ap.ssid"); err != nil {
		return
	} else if err, bssid = mod.StringParam("wifi.ap.bssid"); err != nil {
		return
	} else if mod.apConfig.BSSID, err = net.ParseMAC(network.NormalizeMac(bssid)); err != nil {
		return
	} else if err, mod.apConfig.Channel = mod.IntParam("wifi.ap.channel"); err != nil {
		return
	} else if err, mod.apConfig.Encryption = mod.BoolParam("wifi.ap.encryption"); err != nil {
		return
	}
	return
}

func (mod *WiFiModule) startAp() error {
	// we need channel hopping and packet injection for this
	if !mod.Running() {
		return errNoRecon
	} else if !atomic.CompareAndSwapInt32(&mod.apRunning, 0, 1) {
		return session.ErrAlreadyStarted(mod.Name())
	}

	mod.writes.Add(1)
	go func() {
		defer mod.writes.Done()
		defer atomic.StoreInt32(&mod.apRunning, 0)

		enc := tui.Yellow("WPA2")
		if !mod.apConfig.Encryption {
			enc = tui.Green("Open")
		}
		mod.Info("sending beacons as SSID %s (%s) on channel %d (%s).",
			tui.Bold(mod.apConfig.SSID),
			mod.apConfig.BSSID.String(),
			mod.apConfig.Channel,
			enc)

		for seqn := uint16(0); mod.Running(); seqn++ {
			if err, pkt := packets.NewDot11Beacon(mod.apConfig, seqn); err != nil {
				mod.Error("could not create beacon packet: %s", err)
			} else {
				mod.injectPacket(pkt)
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	return nil
}
