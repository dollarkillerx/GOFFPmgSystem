package taskruner

import (
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(do dataChan) error {
		for i:= 0;i<30;i++ {
			do <- i
			log.Printf("Dispatcher send:%d\n",i)
		}
		return nil
	}

	e := func(dc dataChan) error {
		forloop:
			for {
				select {
					case d :=<-dc:
						log.Printf("Executor received:%v",d)
					default:
						break forloop

				}
			}
		return nil
	}
	runner := NewRunner(30,false,d,e)
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
