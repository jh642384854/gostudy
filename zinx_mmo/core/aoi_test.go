package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	aoiManager := NewAOIManager(0,250,5,0,250,5)
	fmt.Println(aoiManager)
}

func TestFun1(t *testing.T)  {
	fmt.Println(10%3,10/3)
}

func TestGetAroundGridIDs(t *testing.T) {
	aoiManager := NewAOIManager(0,250,5,0,250,5)
	for i:=0;i<len(aoiManager.grids) ;i++  {
		//得到当前格子周围的其他所有格子对象信息
		gridObjets := aoiManager.GetAroundGridIDs(i)
		//下面从这些格子对象信息中只是获取格子的ID信息
		grids := make([]int,0,len(gridObjets))
		for j:=0;j<len(gridObjets) ;j++  {
			grids = append(grids,gridObjets[j].GID)
		}
		fmt.Printf("格子编号为%d的相邻格子id是：",i)
		fmt.Println(grids)
	}
}