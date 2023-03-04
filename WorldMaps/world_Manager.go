package main

type WorldManager struct {
	MsgChan       chan *wHandleMsgInfo
	MoveManager   *wMoveManger
	PlayerManager *wPlyerManager
}

type wHandleMsgInfo struct {
	protocolID uint16
	msg        interface{}
}

var WManager = new(WorldManager)

func (wm *WorldManager) InitWorldManager() {
	wm.MsgChan = make(chan *wHandleMsgInfo)
	wm.MoveManager = new(wMoveManger)
	wm.MoveManager.Run()
	wm.PlayerManager = new(wPlyerManager)
	wm.PlayerManager.InitPlayerManager()
}
