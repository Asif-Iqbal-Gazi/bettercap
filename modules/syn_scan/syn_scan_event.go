package syn_scan

import (
	"github.com/Asif-Iqbal-Gazibettercap/network"
	"github.com/Asif-Iqbal-Gazibettercap/session"
)

type SynScanEvent struct {
	Address string
	Host    *network.Endpoint
	Port    int
}

func NewSynScanEvent(address string, h *network.Endpoint, port int) SynScanEvent {
	return SynScanEvent{
		Address: address,
		Host:    h,
		Port:    port,
	}
}

func (e SynScanEvent) Push() {
	session.I.Events.Add("syn.scan", e)
	session.I.Refresh()
}
