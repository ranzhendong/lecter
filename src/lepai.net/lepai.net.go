package lepai_net

import (
	"net"
	"time"
)

func PingIPPort(ip string, pingTimeout int) (result bool) {
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.DialTimeout("tcp", ip, time.Duration(pingTimeout)*time.Millisecond); err != nil {
		//log.Println(err)
		return
	}
	conn.Close()
	result = true
	return
}
