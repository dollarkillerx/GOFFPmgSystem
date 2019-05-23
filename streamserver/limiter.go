package main

import "log"

type ConnLimiter struct {
	concurrentConn int //定义令牌个数
	bucket chan int // 定义桶
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		cc,
		make(chan int,cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Println("Reached the rate limitation.")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn()  {
	i := <-cl.bucket
	log.Printf("New connction coming:%d\n",i)
}