package main

type ConnLimiter struct {
	concurrentConn int //定义令牌个数
	bucket chan int // 定义桶
}

func NewConnLimiter(cc int) *ConnLimiter {

}