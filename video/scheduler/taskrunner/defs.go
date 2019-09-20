package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_PATH = "E:/GoProjects/src/dev/video/"
)
//定义当前数据的状态(是被分发的还是被执行的)
type controllerChan chan string
//要被处理的数据
type dataChan chan interface { }
//定义数据进行分发和执行的处理函数
type fn func(dc dataChan) error

