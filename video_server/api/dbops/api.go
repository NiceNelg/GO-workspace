package dbops

func AddUserCredential(userName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO user(username, pwd) VALUES(?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(userName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(userName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM user WHERE username=?")
	if err != nil {
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(userName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUserCredential(userName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM user WHERE username=? AND pwd=?")
	if err != nil {
		return err
	}
	stmtDel.Exec(userName, pwd)
	stmtDel.Close()
	return nil
}
