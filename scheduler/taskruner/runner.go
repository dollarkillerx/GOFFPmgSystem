package taskruner

type Runner struct {
	Controller controlChan // 控制信息
	Error controlChan // 错误信息
	Data dataChan // data info
	dataSize int
	longlived bool //是否长期存活
	Dispatcher fn // 生产
	Executor fn // 消费
}

func NewRunner (size int,longlived bool,d fn,e fn) *Runner {
	return &Runner{
		Controller:make(chan string,1),
		Error:make(chan string,1),
		Data:make(chan interface{},size),
		longlived:longlived,
		Dispatcher:d,
		Executor:e,
		dataSize:size,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		//如果不的话 就回收资源
		if !r.longlived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()
	for  {
		select {
			case c := <-r.Controller:
				if c == READY_TO_DISPATH {
					if err := r.Dispatcher(r.Data);err != nil {
						r.Error <- CLOSE
					}else{
						r.Controller <- READY_TO_EXECUTE
					}
				}
				if c == READY_TO_EXECUTE {
					if err := r.Executor(r.Data);err != nil {
						if err != nil {
							r.Error <- CLOSE
						}else{
							r.Controller <- READY_TO_DISPATH
						}
					}
				}
			case e := <- r.Error:
				if e == CLOSE{
					return
				}
			default:

			}
	}
}

func (r *Runner) StartAll()  {
	r.Controller <- READY_TO_DISPATH
	r.startDispatch()
}




// startDispatcher
// control chanel 任务判断
// data channel  数据