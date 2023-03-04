package main

import (
	"encoding/json"
	"fmt"
	"mmorpg/LocalMapInfos"
)

func (p *player) Move(x, y int) {
	var wdmsg = new(LocalMapInfos.WorldMapMsg)
	wdmsg.X = x
	wdmsg.Y = y
	wdmsg.ID = 1
	wdmsg.Data, _ = json.Marshal(struct {
		id   int
		oldx int
		oldy int
		x    int
		y    int
	}{wdmsg.ID, p.x, p.y, x, y})
	LocalMapInfos.ToWorldMapMsg(wdmsg, p.x, p.y)
}

func (p *player) boradcastMove(x, y int) {
	fmt.Println("收到移动消息")
	p.x = x
	p.y = y
	z, _ := json.Marshal(struct {
		id int
		x  int
		y  int
	}{1, x, y})
	p.c.Write(z)
}
