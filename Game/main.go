package main

import (
	"fmt"
	"mmorpg/LocalMapInfos"
)

func InitMaps() {
	var m = new(LocalMapInfos.WorldConns)
	m.ID = 1
	m.XY0 = []int{0, 0}
	m.XY1 = []int{1000, 1000}
	LocalMapInfos.AddMapsInfo(m)
	fmt.Println("加载地图服务器完成.......")
}

func main() {
	fmt.Println("开始mmorpg框架构建.......")
	LocalMapInfos.InitMapsInfo(200)
	InitMaps()
	mps.Init()
	DB.Init()
	LocalMapInfos.InitServerMa(getNewUT, 7999)
	LocalMapInfos.StartTCPListen()
}
