package core

import (
	"fmt"
	"sync"
)

/**
	一个AOI地图中的格子类型
 */

type Grid struct {
	//格子ID
	GID int
	//格子的左边边界坐标(x轴最小值)
	MinX int
	//格子的右边边界坐标(x轴最大值)
	MaxX int
	//格子的上边边界坐标(y轴最小值)
	MinY int
	//格子的下边边界坐标(y轴最大值)
	MaxY int
	//当前格子内玩家或是物体成员的ID集合
	playerIDs map[int]bool
	//保护当前集合的锁
	pIDLock sync.RWMutex
}

//初始化当前格子的方法
func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

//向格子中添加玩家
func (this *Grid) Add(playid int) {
	this.pIDLock.Lock()
	defer this.pIDLock.Unlock()
	this.playerIDs[playid] = true
}

//从格子中删除玩家
func (this *Grid) Remove(playid int) {
	this.pIDLock.Lock()
	defer this.pIDLock.Unlock()
	delete(this.playerIDs, playid)
}

//得到当前格子中所有的玩家ID
func (this *Grid) GetPlayerIDs() (playerIDs []int) {
	this.pIDLock.RLock()
	defer this.pIDLock.RUnlock()
	for k, _ := range this.playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return
}

//调试信息，用来打印格子的基本信息
func (this *Grid) String() string {
	return fmt.Sprintf("Grid id :%d,MinX:%d，MaxX:%d，MinY:%d，MaxY:%d，palyerids：%v",
		this.GID, this.MinX, this.MaxX, this.MinY, this.MaxY, this.playerIDs);
}
