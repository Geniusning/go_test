package dpops

import "testing"

func clearTables() {
	dbconn.Exec("truncate users")
	dbconn.Exec("truncate video_info")
	dbconn.Exec("truncate comment")
	dbconn.Exec("truncate sessions")
}
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDelUser)
	t.Run("ReGet", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("liuning1", "123")
	if err != nil {
		t.Errorf("Error of Adduser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("liuning1")
	if pwd != "123" && err != nil {
		t.Errorf("error of GetUser:%v", err)
	}
}

func testDelUser(t *testing.T) {
	err := DeleteUser("liuning1", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("liuning1")
	if err != nil {
		t.Errorf("error of ReGetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("delete user failed")
	}
}
