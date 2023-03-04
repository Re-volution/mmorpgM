package main

import "sync"

type wPlyerManager struct {
	lock    sync.RWMutex
	players map[string]*player
}

type player struct {
	uid string //玩家唯一id
	x   int
	y   int
}

// 初始化玩家管理器
func (wpm *wPlyerManager) InitPlayerManager() {
	//TODO:加载玩家数据等

	wpm.players = make(map[string]*player)
}

//处理玩家分格，不同玩家在不同区域内，进行广播优先查找视野内拥有全部格子视野的区域，根据实际情况划分
//这里有两种简化方案：
//1.视野内全中的区域进行广播，要求视野距离远大于格子区域（3倍以上）；
//2.视野内存在的格子全部选取发送，要求视野距离至少是格子区域的1/2以上。
