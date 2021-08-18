package ufgutil

import (
	"log"
	"net"
	"runtime/pprof"
)

func startPprof(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println("accept error", err)
				return
			}

			pprof.StartCPUProfile(conn)
			var buf [1]byte
			conn.Read(buf[:])
			pprof.StopCPUProfile()
			conn.Close()
		}
	}()

	return nil
}

func init() {
	if err := startPprof(":46157"); err != nil {
		log.Println("start pprof error", err)
	}
}
