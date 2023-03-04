package main

import (
	"fmt"
	"mmorpg/LocalMapInfos"
)

type DBManager struct {
	LoginChan chan chanUid
}

var DB = new(DBManager)

func (D *DBManager) Init() {
	D.LoginChan = make(chan chanUid, 512)
	go func() {
		defer LocalMapInfos.RecoErr()
		for {
			select {
			case wd := <-D.LoginChan: //TODO:此处目前是单线处理数据库读取，可改为并发执行
				//Todo:处理上线
				if false { //登陆失败
					wd.ch <- nil
				}
				//登陆成功返回链接
				wd.ch <- &userMsg{}

				//todo:处理玩家上线
				go LoginPlayer(wd)

			}
		}
	}()
	fmt.Println("玩家登陆数据库系统初始化完成.......")
}
