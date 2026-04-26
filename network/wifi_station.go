package network

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var (
	pathNameCleaner = regexp.MustCompile("[^a-zA-Z0-9]+")
)

type Station struct {
	*Endpoint
	Frequency      int               `json:"frequency"`
	Channel        int               `json:"channel"`
	RSSI           int8              `json:"rssi"`
	Sent           uint64            `json:"sent"`
	Received       uint64            `json:"received"`
	Encryption     string            `json:"encryption"`
	Cipher         string            `json:"cipher"`
	Authentication string            `json:"authentication"`
	WPS            map[string]string `json:"wps"`
	Handshake      *Handshake        `json:"-"`

	// guards Encryption/Cipher/Authentication and the WPS map.
	// Sent/Received use sync/atomic for lock-free counter updates.
	mu sync.RWMutex
}

// AddSent atomically increases the bytes-sent counter.
func (s *Station) AddSent(n uint64) {
	atomic.AddUint64(&s.Sent, n)
}

// AddReceived atomically increases the bytes-received counter.
func (s *Station) AddReceived(n uint64) {
	atomic.AddUint64(&s.Received, n)
}

// SentBytes returns the current bytes-sent counter.
func (s *Station) SentBytes() uint64 {
	return atomic.LoadUint64(&s.Sent)
}

// ReceivedBytes returns the current bytes-received counter.
func (s *Station) ReceivedBytes() uint64 {
	return atomic.LoadUint64(&s.Received)
}

// SetEncryption updates encryption-related fields under a write lock.
func (s *Station) SetEncryption(encryption, cipher, authentication string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Encryption = encryption
	s.Cipher = cipher
	s.Authentication = authentication
}

// EncryptionInfo returns a consistent snapshot of encryption/cipher/auth.
func (s *Station) EncryptionInfo() (encryption, cipher, authentication string) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Encryption, s.Cipher, s.Authentication
}

// SetWPS sets a WPS info entry under a write lock.
func (s *Station) SetWPS(name, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.WPS == nil {
		s.WPS = make(map[string]string)
	}
	s.WPS[name] = value
}

// WPSData returns a copy of the WPS info map under a read lock.
func (s *Station) WPSData() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make(map[string]string, len(s.WPS))
	for k, v := range s.WPS {
		out[k] = v
	}
	return out
}

func cleanESSID(essid string) string {
	res := ""
	for _, c := range essid {
		if strconv.IsPrint(c) {
			res += string(c)
		} else {
			break
		}
	}
	return res
}

func NewStation(essid, bssid string, frequency int, rssi int8) *Station {
	return &Station{
		Endpoint:  NewEndpointNoResolve(MonitorModeAddress, bssid, cleanESSID(essid), 0),
		Frequency: frequency,
		Channel:   Dot11Freq2Chan(frequency),
		RSSI:      rssi,
		WPS:       make(map[string]string),
		Handshake: NewHandshake(),
	}
}

func (s *Station) BSSID() string {
	return s.HwAddress
}

func (s *Station) ESSID() string {
	return s.Hostname
}

func (s *Station) HasWPS() bool {
	return len(s.WPS) > 0
}

func (s *Station) IsOpen() bool {
	return s.Encryption == "" || s.Encryption == "OPEN"
}

func (s *Station) PathFriendlyName() string {
	name := ""
	bssid := strings.Replace(s.HwAddress, ":", "", -1)
	if essid := pathNameCleaner.ReplaceAllString(s.Hostname, ""); essid != "" {
		name = fmt.Sprintf("%s_%s", essid, bssid)
	} else {
		name = bssid
	}
	return name
}
