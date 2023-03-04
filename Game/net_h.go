package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
)

type tcpManager struct {
	sync.RWMutex
	TcpConns map[uint32]*userConn
}

type userConn struct {
	id uint32
	n  net.Conn
}

var CCtcpManager = new(tcpManager)

func (tm *tcpManager) Init() {
	tm.TcpConns = make(map[uint32]*userConn)
	fmt.Println("链接管理器完成.......")
}

func (tm *tcpManager) addTcp(c net.Conn) *userConn {
	id := getid()
	var u = new(userConn)
	u.n = c
	u.id = id
	tm.Lock()
	tm.TcpConns[id] = u
	tm.Unlock()
	return u
}

func (tm *tcpManager) getTcp(id uint32) *userConn {
	tm.RLock()
	tm.RUnlock()
	return tm.TcpConns[id]
}

func (tm *tcpManager) delTcp(id uint32) {
	tm.Lock()
	delete(tm.TcpConns, id)
	tm.Unlock()
}

func (*tcpManager) startListen() {
	addr, e := net.ResolveTCPAddr("tcp", ":7999")
	if e != nil {
		return
	}
	fmt.Println("开始监听.......")
	l, e := net.ListenTCP("tcp", addr)
	if e != nil {
		return
	}
	c, e := l.Accept()

	var nuser = new(userTcp)
	nuser.c = c
	nuser.ex = make(chan bool, 2)
	u := CCtcpManager.addTcp(c)
	go u.Read(nuser)
}

func (u *userConn) Read(ut *userTcp) {
	var b = make([]byte, 512)
	var data = make([]byte, 0, 512)
	var length = uint32(0)
	var protocolID = uint16(0)
	for {
		l, e := u.n.Read(b)
		if e != nil {
			fmt.Println("read err:", e)
			CCtcpManager.delTcp(u.id)
			ut.ex <- true
			return
		}
		data = append(data, b[:l]...)
		for {
			if len(data) >= int(length) && length >= 6 {
				protocolID = binary.BigEndian.Uint16(data[4:6])
				ut.handMsg(protocolID, data[6:length])
				length = 0
				data = data[length:]
			}

			if length == 0 && len(data) >= 6 { //长度 取协议号
				length = binary.BigEndian.Uint32(data[:4])
			}

			if int(length) > len(data) || length == 0 {
				break
			}
		}

	}

}
