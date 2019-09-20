package main

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

/**
	对象池的应用

 */
func main() {
	pool := NewObjPool(10)

	for i:= 0;i<11 ;i++  {
		if obj,err := pool.GetObj(time.Second*1);err != nil{
			fmt.Println(err.Error())
		}else{
			fmt.Printf("%T\n",obj)
			if err := pool.ReleasObj(obj);err != nil{
				fmt.Println(err.Error())
			}
		}
	}
}

type ReusableObj struct {
}

type ObjPool struct {
	bufchan chan *ReusableObj
}
//获取对象.该方法定义了一个参数，是超时时间的设置
func (obj *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case obj := <-obj.bufchan:
		return obj, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

//释放对象
func (obj *ObjPool) ReleasObj(resObj *ReusableObj) error {
	select {
	case obj.bufchan <- resObj:
		return nil
	default:
		return errors.New("chan full")
	}
}

//初始化对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufchan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufchan <- &ReusableObj{}
	}
	return &objPool
}
