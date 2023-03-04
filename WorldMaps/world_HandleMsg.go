package main

import (
	"fmt"
	"mmorpg/LocalMapInfos"
	"net"
)

const protoID_Move = 1

type userTcp struct {
}

func (ut *userTcp) Exit() {
}

func getNewUT(_ net.Conn) LocalMapInfos.HandleMsgI {
	var nuser = new(userTcp)
	return nuser
}
func (ut *userTcp) HandMsg(protocolID uint16, msg []byte) {
	//TODO: 解析处理消息，塞入数据处理通道  如果处理慢，考虑分开处理,包括不限于更改消息传输方式，如 json ->  protobuf -> byte
	WManager.MsgChan <- &wHandleMsgInfo{protocolID, msg}
}

//============================================================================================

// TODO:目前是同一个地方处理，可以考虑不交叉的地方分开处理
func HandMsg() {
	go func() {
		defer LocalMapInfos.RecoErr()
		for {
			i := 100 //处理消息设立长度，防止一直处理陷入无限循环
			for {
				select {
				case d := <-WManager.MsgChan: //TODO:将数据处理做成队列形式，塞到相应处理队列，在下面去检查处理
					handleRecvMsg(d)
					i--
				default:
					i = -1
				}
				if i <= 0 {
					break
				}
			}
			//TODO:处理各系统循环队列消息，如行走，释放技能等出现不同系统化交叉时的情况
			//TODO:处理其他消息内容
		}
	}()
	fmt.Println("世界大地图消息处理启动..........")
}

func handleRecvMsg(d *wHandleMsgInfo) {
	switch d.protocolID {
	case protoID_Move:
		WManager.MoveManager.msg <- d.msg
	default:

	}
}
