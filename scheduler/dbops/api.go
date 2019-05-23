package dbops
// 1.任务下发 user->api service->delete video
// 2. api service->scheduler-> write video
// 3. timer
// 4. timer->runner->read wvdr ->exec ->delete video from folder

import (
	_ "github.com/go-sql-driver/mysql"
)

func AddVideoDeletionRecord(vid string) error {
	stmt, e := dbConn.Prepare("INSERT INTO `video_del_rec`(`video_id`) VALUE (?)")
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(vid)
	return e
}

