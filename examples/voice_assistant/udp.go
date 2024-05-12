package main

import (
	"fmt"
	"net"
)

type Listener interface {
	GetPort() int
	Stop() error
	Read(num int) ([]byte, error)
}

type UDPListerner struct {
	port int
	conn *net.UDPConn
	addr *net.UDPAddr
}

func NewUDPListerner(port int) (*UDPListerner, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return &UDPListerner{port: port, conn: conn}, nil
}

// func (ul *UDPListerner) Start() (port int, err error) {
// 	return ul.port, err
// }

func (ul *UDPListerner) Stop() error {
	return ul.conn.Close()
}

func (ul UDPListerner) GetPort() int {
	return ul.port
}

func (ul *UDPListerner) Read(maxRead int) ([]byte, error) {
	read := 0
	var data []byte
	for read < maxRead {
		buf := make([]byte, 1024)
		_, _, err := ul.conn.ReadFromUDP(buf)
		if err != nil {
			return nil, err
		}
		read++
		data = append(data, buf...)
	}
	return data, nil
}
