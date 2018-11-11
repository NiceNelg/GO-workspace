package dbops

import (
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate user")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate session")
	dbConn.Exec("truncate video_info")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("GET", testGetUser)
	t.Run("DELETE", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("nelg", "123456")
	if err != nil {
		t.Errorf("Error or AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("nelg")
	if pwd != "123456" || err != nil {
		t.Errorf("Error or Get User: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUserCredential("nelg", "123456")
	if err != nil {
		t.Errorf("Error or AddUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	_, err := GetUserCredential("nelg")
	if err != nil {
		t.Errorf("Error or Delete User: %v", err)
	}
}
