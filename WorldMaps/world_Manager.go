package main

type WorldManager struct {
	MsgChan       chan *wHandleMsgInfo
	MoveManager   *wMoveManger
	PlayerManager *wPlyerManager
	MapCell       *wMapCellManager
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
	wm.MapCell = new(wMapCellManager)
	wm.MapCell.Init()
}
