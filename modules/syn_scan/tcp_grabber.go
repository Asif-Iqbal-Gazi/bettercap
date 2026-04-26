package syn_scan

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

func cleanBanner(banner string) string {
	clean := ""
	for _, c := range banner {
		if strconv.IsPrint(c) {
			clean += string(c)
		}
	}
	return clean
}

func tcpGrabber(mod *SynScanner, ip string, port int) string {
	dialer := net.Dialer{
		Timeout: bannerGrabTimeout,
	}

	addr := net.JoinHostPort(ip, strconv.Itoa(port))
	if conn, err := dialer.Dial("tcp", addr); err == nil {
		defer conn.Close()
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		return cleanBanner(strings.Trim(msg, "\r\n\t "))
	} else {
		mod.Debug("%s : %v", addr, err)
	}
	return ""
}
