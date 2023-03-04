package main

import (
	"fmt"
	"mmorpg/LocalMapInfos"
	"net"
)

const (
	LoginMsgID = 1
)

type userTcp struct {
	c  net.Conn
	u  chan *userMsg
	ex chan bool
}

type userMsg struct {
	protocolID uint16
	msg        []byte
}

func (ut *userTcp) Exit() {
	ut.ex <- true
}
func getNewUT(c net.Conn) LocalMapInfos.HandleMsgI {
	var nuser = new(userTcp)
	nuser.c = c
	nuser.ex = make(chan bool, 2)
	return nuser
}
func (ut *userTcp) HandMsg(protocolID uint16, msg []byte) {
	switch protocolID {
	case LoginMsgID:
		if ut.u != nil {
			fmt.Println("ut.u != nil:", string(msg))
			return
		}

		//todo:等待检测其他服务器是否上线该玩家
		b := DBLogin(string(msg), ut)
		if !b { //TODO:处理登陆失败返回
			fmt.Println("DBLogin fail:", string(msg))
			ut.u = nil
			return
		}
		ut.u = make(chan *userMsg, 512)
	default:
		if ut.u == nil {
			fmt.Println("user not login:", string(msg))
			return
		}
		ut.u <- &userMsg{protocolID, msg}
	}
}

type chanUid struct {
	uid    string
	conn   net.Conn
	ch     chan *userMsg
	exchan chan bool
}

func DBLogin(uid string, ut *userTcp) bool {
	DB.LoginChan <- chanUid{uid, ut.c, ut.u, ut.ex}
	//注意必须返回数据，或者这里写成超时处理
	if <-ut.u == nil {
		return false
	}
	return true
}
