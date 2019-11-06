package lepai_net

import (
	"net"
	"time"
)

func PingIPPort(ip string) (result bool) {
	var (
		conn net.Conn
		err  error
	)
	if conn, err = net.DialTimeout("tcp", ip, 3*time.Second); err != nil {
		//log.Println(err)
		return
	}
	conn.Close()
	result = true
	return
}
