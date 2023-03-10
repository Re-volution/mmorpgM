package main

import (
	"fmt"
	"mmorpg/LocalMapInfos"
)

const (
	id = 1
)
const mapminx, mapminy, mapmaxx, mapmaxy, cell_temp = 0, 0, 1000, 1000, 10000

// 世界地图系统仅加载自身地图
func InitMaps() {
	var m = new(LocalMapInfos.WorldConns)
	m.ID = 1
	m.XY0 = []int{0, 0}
	m.XY1 = []int{mapmaxx, mapmaxy}
	LocalMapInfos.AddMapsInfo(m)
	fmt.Println("加载地图服务器完成.......")
}
func main() {
	fmt.Println("世界地图", id)
	LocalMapInfos.InitMapsInfo(200)
	InitMaps()
	WManager.InitWorldManager()
	HandMsg()
	LocalMapInfos.InitServerMa(getNewUT, 6999)
	LocalMapInfos.StartTCPListen()
}
