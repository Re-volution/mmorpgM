package main

import "mmorpg/LocalMapInfos"

type wMoveManger struct {
	msg      chan interface{}
	movelist map[string]*MoveInfo
}

type MoveInfo struct {
	Uid   string
	Dir   int //方向
	Speed int //速度
	X     int
	Y     int
	Z     int
}

func (t *wMoveManger) Run() {
	t.msg = make(chan interface{})
	t.movelist = make(map[string]*MoveInfo)
	t.Move()
}

// 移动数据处理
func (t *wMoveManger) Move() {
	go func() {
		defer LocalMapInfos.RecoErr()
		for {
			i := 100 //防止数据处理过多
			for {
				select {
				case _ = <-t.msg:
					//TODO:加入行走队列
					i--
				default:
					i = -1
				}
				if i <= 0 {
					break
				}
			}
			//TODO:障碍物检测等

			//TODO:处理所有玩家行走，更新视野信息，移动玩家进行区域归类（建议自身检测是否超出区域之后，再进行重新划分区域
		}

	}()
}
