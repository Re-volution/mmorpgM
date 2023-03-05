package LocalMapInfos

import (
	"fmt"
	"net"
)

type worldMapManager struct {
	ws   []*WorldConns
	anti int // 玩家视野不宜超过地图大小太多，否则会造成大量消息传输  //这里不采用九宫格，而是实时以自己为中心200坐标以内的方形区域做视野裁剪，方便查询是否存在其他大地图视野
}

var wmm *worldMapManager

type WorldConns struct {
	ID  int   //大地图的id
	XY0 []int //大地图的左上角坐标
	XY1 []int //大地图的右下角坐标
	C   net.Conn
}

// 添加地图信息
func AddMapsInfo(w *WorldConns) {
	wmm.ws = append(wmm.ws, w)
}

// 初始化地图管理器 anti视野范围
func InitMapsInfo(anti int) {
	wmm = new(worldMapManager)
	wmm.anti = anti
	fmt.Println("初始化地图管理器完成")
}

type WorldMapMsg struct {
	ID   int    //消息类型 保留
	X    int    //坐标x  如果是特殊消息可以不需要坐标
	Y    int    //坐标y
	Data []byte //具体要发送的消息 理应包含当前坐标以及到达坐标
}

func ToWorldMapMsg(d *WorldMapMsg, currlx, currly int) {
	worlds := wmm.getWorldConn(currlx, currly)
	worlds = append(worlds, wmm.getWorldConn(d.X, d.Y)...)
	var rec = make(map[int]bool) //做个重复地图筛选
	for _, v := range worlds {   //判断是否离开当前大地图 可以超出大地图依然保留所在位置，因为要保留视野消息情况，相应的，在还未进入下一张地图前，视野也可能会提前进入
		if ok, _ := rec[v.ID]; !ok {
			rec[v.ID] = true
			v.C.Write(d.Data)
		}
	}
}

// 如果出现超过视野的技能投掷，则需要生成单个物体结构体，里面应当包含玩家信息，发生反馈之后找到该玩家作相应处理，如果玩家下线则广播并在玩家大地图上做操作，等待玩家上线后处理反馈
// 这里需要获取玩家视野所在服务器
func (w *worldMapManager) getWorldConn(x, y int) []*WorldConns {
	var wreturn []*WorldConns
	for _, v := range w.ws {
		if isInAnti(v.XY0, v.XY1, x, y) {
			wreturn = append(wreturn, v)
		}
	}
	return wreturn
}

func isInAnti(xy0, xy1 []int, x, y int) bool {
	if xy0[0] >= x-wmm.anti && xy0[1] >= y-wmm.anti && xy0[0] < x+wmm.anti && xy0[1] < y+wmm.anti {
		return true
	}
	if xy0[0] >= x-wmm.anti && xy1[1] >= y-wmm.anti && xy0[0] < x+wmm.anti && xy1[1] < y+wmm.anti {
		return true
	}
	if xy1[0] >= x-wmm.anti && xy0[1] >= y-wmm.anti && xy1[0] < x+wmm.anti && xy0[1] < y+wmm.anti {
		return true
	}
	if xy1[0] >= x-wmm.anti && xy1[1] >= y-wmm.anti && xy1[0] < x+wmm.anti && xy1[1] < y+wmm.anti {
		return true
	}
	return false

}

func (w *worldMapManager) getAllWConns() {

}
