package dbops

import "log"

//api->video_id->mysql
//dispatcher->mysql->video_id->datachannel
//executor->datachannel->video_id->delete videos


// 读数据库需要删除的
func ReadVideoDeletionRecord(count int) ([]string,error) {
	stmt, e := dbConn.Prepare("SELECT `video_id` FROM `video_del_rec` LIMIT ?")
	defer stmt.Close()

	var ids []string
	if e != nil {
		return ids,e
	}

	rows, e := stmt.Query(count)
	if e != nil {
		log.Printf("Query VideoDeletionRecord error:%v",e)
		return ids,e
	}

	for rows.Next() {
		var id string
		if e := rows.Scan(&id);e != nil {
			return ids,err
		}
		ids = append(ids,id)
	}
	return ids,nil
}

// 删除已经删除的数据
func DelVideoDeletionRecord(vid string) error {
	stmt, e := dbConn.Prepare("DELETE FROM `video_del_rec` WHERE `video_id` = ?")
	defer stmt.Close()
	if e != nil {
		return e
	}
	_, e = stmt.Exec(vid)
	return e
}