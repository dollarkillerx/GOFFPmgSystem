package taskruner

const (
	READY_TO_DISPATH = "d" //生成通知
	READY_TO_EXECUTE = "e" // 消费通知
	CLOSE = "c" // 错误通知
	VIDOE_DIR = "./VIDEOS/" //视频地址
)

type controlChan chan string //通知chan

type dataChan chan interface{} //data chan

type fn func(dc dataChan) error //消费