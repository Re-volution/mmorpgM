package LocalMapInfos

import (
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"sync"
)

type tcpManager struct {
	sync.RWMutex                                //读写锁
	TcpConns     map[uint32]*userConn           //链接池
	getHandle    func(conn net.Conn) HandleMsgI //处理消息函数
	Port         int                            //端口
}

type userConn struct {
	id uint32
	n  net.Conn
}

type HandleMsgI interface {
	HandMsg(uint16, []byte)
	Exit()
}

var CCtcpManager = new(tcpManager)

func InitServerMa(f func(conn net.Conn) HandleMsgI, p int) {
	CCtcpManager.TcpConns = make(map[uint32]*userConn)
	CCtcpManager.getHandle = f
	CCtcpManager.Port = p
	fmt.Println("tcp链接管理器完成.......")
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

func StartTCPListen() {
	addr, e := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(CCtcpManager.Port))
	if e != nil {
		return
	}
	fmt.Println("开始监听.......")
	l, e := net.ListenTCP("tcp", addr)
	if e != nil {
		return
	}
	c, e := l.Accept()

	u := CCtcpManager.addTcp(c)
	nuser := CCtcpManager.getHandle(c)
	go u.Read(nuser)
}

func (u *userConn) Read(ut HandleMsgI) {
	var b = make([]byte, 512)
	var data = make([]byte, 0, 512)
	var length = uint32(0)
	var protocolID = uint16(0)
	for {
		l, e := u.n.Read(b)
		if e != nil {
			fmt.Println("read err:", e)
			CCtcpManager.delTcp(u.id)
			ut.Exit()
			return
		}
		data = append(data, b[:l]...)
		for {
			if length > 0 && len(data) >= int(length) && len(data) >= 6 { //前四位是长度，后两位是协议号
				protocolID = binary.BigEndian.Uint16(data[4:6])
				ut.HandMsg(protocolID, data[6:length])
				length = 0
				data = data[length:]
			}
			if length == 0 && len(data) >= 4 { //长度
				length = binary.BigEndian.Uint32(data[:4])
			}
			if int(length) > len(data) || length == 0 {
				break
			}
		}

	}

}
