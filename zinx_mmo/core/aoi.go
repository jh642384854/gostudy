package core

import "fmt"

/**
	AOI区域管理模块
 */
type AOIManager struct {
	//区域的左边界坐标
	MinX int
	//区域的右边界坐标
	MaxX int
	//X方向格子的数量
	CntsX int
	//区域的上边界坐标
	MinY int
	//区域的下边界坐标
	MaxY int
	//Y轴方向格子的数量
	CntsY int
	//当前区域中有哪些格子  map-key=格子ID，value=格子对象
	grids map[int]*Grid
}

//初始化一个AOI区域管理模块
func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiManger := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}
	//给AOI初始化区域的格子所有的格子进行编号和初始化
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			//计算格子ID，根据x,y编号。
			//格子编号计算公式：gid = y*cntsX+x
			gid := y*cntsY + x

			//初始化gid格子
			girdMinX := aoiManger.MinX + x*aoiManger.gridWidth()
			girdMaxX := aoiManger.MinX + (x+1)*aoiManger.gridWidth()
			girdMinY := aoiManger.MinY + y*aoiManger.gridHeight()
			girdMaxY := aoiManger.MinY + (y+1)*aoiManger.gridHeight()
			aoiManger.grids[gid] = NewGrid(gid, girdMinX, girdMaxX, girdMinY, girdMaxY)
		}
	}
	return aoiManger
}

//得到每个格子在X轴方向的宽度。就是得到每个格子实际的宽度是多少
func (this *AOIManager) gridWidth() int {
	return (this.MaxX - this.MinX) / this.CntsX
}

//得到每个格子在Y轴方向的长度。就是得到每个格子实际的高度是多少
func (this *AOIManager) gridHeight() int {
	return (this.MaxY - this.MinY) / this.CntsY
}

//打印调试消息
func (this *AOIManager) String() string {
	//打印AOIManager信息
	s := fmt.Sprintf("AOIManager:\n MinX:%d, MaxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\n Grids in AOIManager:\n",
		this.MinX, this.MaxX, this.CntsX, this.MinY, this.MaxY, this.CntsY)

	//打印全部格子信息
	for _, grid := range this.grids {
		s += fmt.Sprintln(grid)
	}

	return s
}

/**
	@param gid int 这个是格子的编号ID
	@return 得到指定格子相邻的其他所有格子对象的切片信息

	通过地图上面的某一个格子信息来获取这个格子周边的所有格子信息
	实现思路如下：
	以格子A为说明
	第一步：我们会首先找到与A横向相连的左右两个格子。这样就有可能得到2个(靠边的就只有两个)或3个格子
	第二步：得到了A格子横向相连的格子后，我们在依次处理这几个相连格子的纵向相连的格子。这样我们就得到了A格子的所有相邻的格子节点。
	难点：

	格子示例：
		0	1	2	3	4
    0	0	1	2	3	4
	1	5	6	7	8	9
	2	10	11	12	13	14
	3	15	16	17	18	19
	4	20	21	22	23	24
	需要注意的是，我们这里的25个格子是平分的250*250面积的地图，也就说，每个格子其实是各自占有50的宽和高的。为什么要说这个呢？因为我们给出一个地图的坐标点的时候，
	比如(130,90)这个坐标点，我们能根据这个坐标点来定位到是那个格子。

 */
func (this *AOIManager) GetAroundGridIDs(gid int) (grids []*Grid) {
	//判断当前格子ID是否属于这个地图中
	if _, ok := this.grids[gid]; !ok {
		return
	}
	tempGids := []int{} //定义的这个变量就是用来专门记录格子的编号ID的，为后面的根据这些格子ID来获取纵向相邻的格子ID
	tempGids = append(tempGids, gid)
	grids = append(grids, this.grids[gid])
	//①、根据给定的格子ID来计算出横轴与之相邻的左右格子ID。如果
	//那怎么判断一个格子的左边或右边是否还有值呢？这个就需要利用x轴的最大值和最小值来判断了。如果当前格子的左边格子的x轴大于0，就说明当前格子的左边还有值，反之则没有。
	//这里为什么是x轴大于0呢？这是根据下面根据格子ID从而计算出格子的X轴编号来的。X轴和Y轴的编号开始都是0
	//同理，如果当前格子的右边格子的x轴小于x轴的最大值，也就说明当前格子的右边也还有值，反之则没有
	//那现在就要根据格子的编号来求出当前格子的X轴的编号。计算公式如下：XPosition = gid % XaxisNums（其中gid就是当前格子的编号。XaxisNums就是当前地图的X轴的总数，这里就是5）
	XPosition := gid % this.CntsX

	if XPosition > 0 {
		grids = append(grids, this.grids[gid-1]) //将当前格子的左边相邻的格子加入进来
		tempGids = append(tempGids, gid-1)
	}

	if XPosition < this.CntsX-1 {
		grids = append(grids, this.grids[gid+1]) //将当前格子的右边相邻的格子加入进来
		tempGids = append(tempGids, gid+1)
	}
	//经过上面的处理，我们就拿到了一个格子相邻的横向坐标的几个格子编号了。下面就要依次来对这几个格子编号做纵向相邻格子id的处理。
	//首先就是要根据格子编号ID来获取当前这个编号的Y轴信息，计算公式如下：YPosition = gid / XaxisNums  注意这里是除法，上面的是取余，这两个算法不一样
	for i := 0; i < len(tempGids); i++ {
		YPosition := tempGids[i] / this.CntsX
		if YPosition > 0 {
			grids = append(grids, this.grids[tempGids[i]-this.CntsX]) //将当前格子的上边相邻的格子加入进来
		}
		if YPosition < this.CntsY-1 {
			grids = append(grids, this.grids[tempGids[i]+this.CntsX]) //将当前格子的下边相邻的格子加入进来
		}
	}
	return
}
/**
	通过一个坐标点来得到是那个格子编号
	一个格子的编号可以由下面的公式来确定：
	Y轴编号*X轴总长度+当前X轴编号。
	比如上面地图中的23这个编号，我们依据这个公式，就可以得到4*5+3=23
	那现在的问题又变成了如果根据一个坐标点来确定这个坐标点对应的X轴和Y轴的编号呢？
	计算公式如下：
	x轴编号：当前X轴坐标 / 每个格子的宽度
	Y轴编号：当前Y轴坐标 / 每个格子的高度
	需要注意的是，我们给的地图如果起始是0开始的话，以上公式就没有什么问题，如果给出的地图并不是从0开始的，那就需要用X轴或Y轴减去起始的x轴和Y轴的开始编号，所以比较稳妥的公式如下：
	x轴编号：(当前X轴坐标-当前地图最小X轴编号) / 每个格子的宽度
	Y轴编号：(当前Y轴坐标-当前地图最小X轴编号) / 每个格子的高度
 */
func (this *AOIManager) GetGridNoByCoordinate(x, y float32) int {
	//得到当前坐标的X轴编号
	xPositionNo := (int(x) - this.MinX) / this.gridWidth()
	//得到当前坐标的Y轴编号
	yPositionNo := (int(y) - this.MinY) / this.gridHeight()

	return xPositionNo * this.CntsX + yPositionNo
}

//通过横纵坐标得到周边九宫格内全部的PlayerIDs
func (this *AOIManager) GetAllPlayerIdsByCoordinate(x,y float32) []int {
	//首先根据坐标来获取地图的格子编号
	gridNo := this.GetGridNoByCoordinate(x,y)
	//在根据格子编号获取当前格子的相邻的其他所有格子信息
	gridObjects := this.GetAroundGridIDs(gridNo)
	//然后循环遍历这些格子对象，从中提取玩家ID
	playIds := make([]int,0)
	for i:=0;i<len(gridObjects) ;i++  {
		playIds = append(playIds,gridObjects[i].GetPlayerIDs()...)
	}
	return playIds
}

//添加一个PlayerID到一个格子中
func (this *AOIManager) AddPlayerIdToGrid(playid ,grid int)  {
	this.grids[grid].Add(playid)
}

//移除一个格子中的PlayerID
func (this *AOIManager) RemovePlayerIdFromGrid(playid ,grid int)  {
	this.grids[grid].Remove(playid)
}

//通过GID获取全部的PlayerID
func (this *AOIManager) GetAllPlayerIdByGrid(grid int) []int {
	return this.grids[grid].GetPlayerIDs()
}

//通过坐标将Player添加到一个格子中
func (this *AOIManager) AddPlayerIdToGridByCoordinate(x,y float32,playid int)  {
	//①、根据坐标获取格子的编号
	gridNo := this.GetGridNoByCoordinate(x,y)
	//②、向这个格子中添加一个player
	this.grids[gridNo].Add(playid)
}

//通过坐标把一个Player从一个格子中删除
func (this *AOIManager) RemovePlayerIdFromGridByCoordinate(x,y float32,playid int)  {
	//①、根据坐标获取格子的编号
	gridNo := this.GetGridNoByCoordinate(x,y)
	//②、从格子中删除palyer
	this.grids[gridNo].Remove(playid)
}