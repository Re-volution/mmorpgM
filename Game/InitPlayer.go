package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type playerManager struct {
	players map[string]*player
	sync.Locker
}

// 玩家视野事件
type seeEvent struct {
	id     int //事件类型
	x      int
	y      int
	param1 int   //参数
	param2 []int //参数
}

// 广播消息
type boradcastMsg struct {
	events []seeEvent
}

type player struct {
	uid string
	c   net.Conn
	ch  chan *userMsg
	ch2 chan boradcastMsg
	x   int //玩家所在坐标轴
	y   int //同上
}

var mps = new(playerManager)

func (pm *playerManager) Init() {
	pm.players = make(map[string]*player)
	fmt.Println("玩家管理器初始化完成.......")
}

func (pm *playerManager) Add(p *player) {
	mps.Lock()
	mps.players[p.uid] = p
	mps.Unlock()
}

func (pm *playerManager) Del(uid string) {
	mps.Lock()
	delete(mps.players, uid)
	mps.Unlock()
}

func LogoutPlayer(uid string) {
	//TODO：保存数据库等
	mps.Del(uid)
}

// todo:玩家的实际数据传输
func LoginPlayer(wd chanUid /*,p playerDB*/) {
	//TODO:检测链接是否断开 select <-wd.exchan
	var p = new(player)
	p.c = wd.conn
	p.ch = wd.ch
	p.uid = wd.uid
	mps.Add(p)
	defer mps.Del(wd.uid)

	for {
		var count = 0
		for {
			select {
			case msg := <-p.ch:
				p.handleMsg(msg.protocolID, msg.msg)
			case <-wd.exchan: //TODO:断线处理操作，如一次性处理完剩余的消息，看需求
				return
			case _ = <-p.ch2:
				//TODO:处理广播消息
			default:
				count = 50
				time.Sleep(5 * time.Millisecond)
			}
			count++
			if count >= 50 {
				break
			}
		}

	}
}

// 处理玩家消息
func (p *player) handleMsg(prorocolID uint16, msg []byte) {
	switch prorocolID {
	case 9:

	default:
	}
}
