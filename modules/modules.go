package modules

import (
	"github.com/Asif-Iqbal-Gazibettercap/modules/any_proxy"
	"github.com/Asif-Iqbal-Gazibettercap/modules/api_rest"
	"github.com/Asif-Iqbal-Gazibettercap/modules/arp_spoof"
	"github.com/Asif-Iqbal-Gazibettercap/modules/ble"
	"github.com/Asif-Iqbal-Gazibettercap/modules/c2"
	"github.com/Asif-Iqbal-Gazibettercap/modules/caplets"
	"github.com/Asif-Iqbal-Gazibettercap/modules/dhcp6_spoof"
	"github.com/Asif-Iqbal-Gazibettercap/modules/dns_spoof"
	"github.com/Asif-Iqbal-Gazibettercap/modules/events_stream"
	"github.com/Asif-Iqbal-Gazibettercap/modules/gps"
	"github.com/Asif-Iqbal-Gazibettercap/modules/hid"
	"github.com/Asif-Iqbal-Gazibettercap/modules/http_proxy"
	"github.com/Asif-Iqbal-Gazibettercap/modules/http_server"
	"github.com/Asif-Iqbal-Gazibettercap/modules/https_proxy"
	"github.com/Asif-Iqbal-Gazibettercap/modules/https_server"
	"github.com/Asif-Iqbal-Gazibettercap/modules/mac_changer"
	"github.com/Asif-Iqbal-Gazibettercap/modules/mdns_server"
	"github.com/Asif-Iqbal-Gazibettercap/modules/mysql_server"
	"github.com/Asif-Iqbal-Gazibettercap/modules/ndp_spoof"
	"github.com/Asif-Iqbal-Gazibettercap/modules/net_probe"
	"github.com/Asif-Iqbal-Gazibettercap/modules/net_recon"
	"github.com/Asif-Iqbal-Gazibettercap/modules/net_sniff"
	"github.com/Asif-Iqbal-Gazibettercap/modules/packet_proxy"
	"github.com/Asif-Iqbal-Gazibettercap/modules/syn_scan"
	"github.com/Asif-Iqbal-Gazibettercap/modules/tcp_proxy"
	"github.com/Asif-Iqbal-Gazibettercap/modules/ticker"
	"github.com/Asif-Iqbal-Gazibettercap/modules/ui"
	"github.com/Asif-Iqbal-Gazibettercap/modules/update"
	"github.com/Asif-Iqbal-Gazibettercap/modules/wifi"
	"github.com/Asif-Iqbal-Gazibettercap/modules/wol"

	"github.com/Asif-Iqbal-Gazibettercap/session"
)

func LoadModules(sess *session.Session) {
	sess.Register(any_proxy.NewAnyProxy(sess))
	sess.Register(arp_spoof.NewArpSpoofer(sess))
	sess.Register(api_rest.NewRestAPI(sess))
	sess.Register(ble.NewBLERecon(sess))
	sess.Register(dhcp6_spoof.NewDHCP6Spoofer(sess))
	sess.Register(net_recon.NewDiscovery(sess))
	sess.Register(dns_spoof.NewDNSSpoofer(sess))
	sess.Register(events_stream.NewEventsStream(sess))
	sess.Register(gps.NewGPS(sess))
	sess.Register(http_proxy.NewHttpProxy(sess))
	sess.Register(http_server.NewHttpServer(sess))
	sess.Register(https_proxy.NewHttpsProxy(sess))
	sess.Register(https_server.NewHttpsServer(sess))
	sess.Register(mac_changer.NewMacChanger(sess))
	sess.Register(mysql_server.NewMySQLServer(sess))
	sess.Register(mdns_server.NewMDNSServer(sess))
	sess.Register(net_sniff.NewSniffer(sess))
	sess.Register(packet_proxy.NewPacketProxy(sess))
	sess.Register(net_probe.NewProber(sess))
	sess.Register(syn_scan.NewSynScanner(sess))
	sess.Register(tcp_proxy.NewTcpProxy(sess))
	sess.Register(ticker.NewTicker(sess))
	sess.Register(wifi.NewWiFiModule(sess))
	sess.Register(wol.NewWOL(sess))
	sess.Register(hid.NewHIDRecon(sess))
	sess.Register(c2.NewC2(sess))
	sess.Register(ndp_spoof.NewNDPSpoofer(sess))

	sess.Register(caplets.NewCapletsModule(sess))
	sess.Register(update.NewUpdateModule(sess))
	sess.Register(ui.NewUIModule(sess))
}
