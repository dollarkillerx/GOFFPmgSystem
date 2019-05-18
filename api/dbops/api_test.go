package dbops

import (
	"GOFFPmgSystem/api/utils"
	"testing"
)

// init(dblogin,truncate tables)->run tests->clear data(truncate table)

var (
	tempvid string // 测试用视频uuid
)

func clearTables() {
	dbConn.Exec("TRUNCATE `users`")
	dbConn.Exec("TRUNCATE `video_info`")
	dbConn.Exec("TRUNCATE `comments`")
	dbConn.Exec("TRUNCATE `sessions`")
}

func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

// 用户相关测试
func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add",testAddUserCredential)
	t.Run("Get",testGetUserCredential)
	t.Run("Del",testDeleteUser)
	t.Run("Reget",testRegetUser)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("testUser", "123")
	if err != nil {
		t.Errorf("Error of AddUser:%v",err)
	}
}

func testGetUserCredential(t *testing.T) {
	s, e := GetUserCredential("testUser")
	pwd := utils.Md5String("123")
	if e!=nil || s!=pwd {
		t.Errorf("Error of GetUser:%v",e)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("testUser", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v",err)
	}
}

func testRegetUser(t *testing.T)  {
	s, e := GetUserCredential("testUser")
	if e!=nil || s!="" {
		t.Errorf("Error of RegetUser:%v",e)
	}
}


// video相关的测试
func TestVideoWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("AddVideo",testAddNewVideo)
	t.Run("GetVideo",testGetVideoInfo)
	t.Run("DelVideo",testDeleteVideoInfo)
}

func testAddNewVideo(t *testing.T) {
	info, e := AddNewVideo(1, "my-video")
	if e != nil {
		t.Error(e.Error())
	}
	tempvid = info.VideoId
	t.Log(tempvid)
}

func testGetVideoInfo(t *testing.T) {
	info, e := GetVideoInfo(tempvid)
	if e != nil {
		t.Error(e.Error())
	}
	t.Log(info)
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid, 1)
	if err != nil {
		t.Error(err.Error())
	}
}

// 评论相关的测试
func TestCommentWorkFlow(t *testing.T)  {
	clearTables()
	t.Run("AddUser",testAddUserCredential)
	t.Run("AddComment",testAddNewComment)
	t.Run("ListComments",testListComments)
}

func testAddNewComment(t *testing.T) {
	error := AddNewComment("123", 1, "tset 1")
	if error != nil {
		t.Error(error.Error())
	}
}


func testListComments(t *testing.T) {
	data, e := ListComments("123", 1558057715, 0)
	if e!=nil {
		t.Error(e.Error())
	}

	if data == nil {
		t.Error("未知")
	}

	for k,v := range data {
		t.Logf("comment:%d,%v \n",k,v)
	}

}

