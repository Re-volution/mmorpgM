package main

const xyz = 5

type wMapCellManager struct {
	Data [xyz][xyz][1]map[string]*player //长宽高划分区域map[玩家唯一id]玩家
}

// TODO：可以考虑优化
func (mcm *wMapCellManager) Init() {
	for i := 0; i < xyz; i++ {
		for j := 0; j < xyz; j++ {
			for z := 0; z < 1; z++ {
				mcm.Data[i][j][z] = make(map[string]*player)
			}
		}
	}
	//TODO:加载已存在的玩家
}

// 检测进入
func (mcm *wMapCellManager) CheckInto(p *player, tox, toy int) {
	p.x = tox
	p.y = toy
	x := tox / (mapmaxx / xyz)
	y := toy / (mapmaxy / xyz)
	cellid := x*cell_temp + y //单格长度需要小于10000
	if cellid != p.xy {       //进入新的格子
		oldxy := p.xy
		p.xy = cellid
		//todo:如果是移动到地图边缘，应该通知进入新地图，通过发送方game去处理 或者直连
		mcm.broadcast(oldxy, cellid)
	}
	mcm.broadcast(p.xy)
}

func (mcm *wMapCellManager) broadcast(cellid ...int) {
	var tomap = make(map[int]bool, 9)
	for _, cid := range cellid {
		if cid != -1 { //刚进入地图为-1
			if _, ok := tomap[cid]; !ok {
				tomap[cid] = true
			}
			if (cid/cell_temp + 1) < xyz {
				if _, ok := tomap[cid+cell_temp]; !ok {
					tomap[cid+cell_temp] = true
				}
			}

			if (cid%cell_temp + 1) < xyz {
				if _, ok := tomap[cid+1]; !ok {
					tomap[cid+1] = true
				}
			}

			if (cid/cell_temp - 1) >= 0 {
				if _, ok := tomap[cid-cell_temp]; !ok {
					tomap[cid-cell_temp] = true
				}
			}

			if (cid%cell_temp - 1) >= 0 {
				if _, ok := tomap[cid-1]; !ok {
					tomap[cid-1] = true
				}
			}

			if (cid/cell_temp+1) < xyz && (cid%cell_temp+1) < xyz {
				if _, ok := tomap[cid+cell_temp+1]; !ok {
					tomap[cid+cell_temp+1] = true
				}
			}

			if (cid/cell_temp+1) < xyz && (cid%cell_temp-1) >= 0 {
				if _, ok := tomap[cid+cell_temp-1]; !ok {
					tomap[cid+cell_temp-1] = true
				}
			}

			if (cid/cell_temp-1) >= 0 && (cid%cell_temp-1) >= 0 {
				if _, ok := tomap[cid-cell_temp-1]; !ok {
					tomap[cid-cell_temp-1] = true
				}
			}

			if (cid/cell_temp-1) >= 0 && (cid%cell_temp+1) < xyz {
				if _, ok := tomap[cid-cell_temp+1]; !ok {
					tomap[cid-cell_temp+1] = true
				}
			}

		}
	}

	for cid, _ := range tomap {
		for _, p := range mcm.Data[cid/cell_temp][cid%cell_temp][0] {
			sendmsg(p)
		}

	}
}

func sendmsg(*player) {

}
