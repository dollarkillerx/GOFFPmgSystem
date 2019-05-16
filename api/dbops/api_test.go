package dbops

import (
	"GOFFPmgSystem/api/tootl"
	"testing"
)

// init(dblogin,truncate tables)->run tests->clear data(truncate table)

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
	pwd := tootl.Md5String("123")
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