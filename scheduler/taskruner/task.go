package taskruner

import (
	"GOFFPmgSystem/scheduler/dbops"
	"errors"
	"log"
	"os"
	"sync"
)

func VideoClearDispatcher(dc dataChan) error {
	res, e := dbops.ReadVideoDeletionRecord(3)
	if e != nil {
		log.Printf("Video clear dispatcher error:%v",e)
		return e
	}

	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _,id := range res {
		dc <- id
	}
	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	forloop:
		for {
			select {
				case vid := <- dc:
					go func(id interface{}) {
						if err := deleteVideo(id.(string));err!=nil{
							errMap.Store(id,err)
							return
						}
						if err := dbops.DelVideoDeletionRecord(id.(string));err!=nil{
							errMap.Store(id,err)
							return
						}
					}(vid)
				default:
					break forloop
			}
		}
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})
	return err
}

func deleteVideo(vid string) error {
	err := os.Remove(VIDOE_DIR + vid)
	if err != nil && os.IsNotExist(err) {
		log.Printf("Deleting video error:%v",err.Error())
		return err
	}
	return nil
}