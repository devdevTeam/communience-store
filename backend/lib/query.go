package lib

func InsertNewUser(uid, mail, password, userName string) error {
	query := "INSERT INTO users (uid, mail, password, name) VALUES ($1, $2, $3, $4)"
	err := Conn.Exec(query, uid, mail, password, userName)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewRoom(rid, roomName, password string) error {
	query := "INSERT INTO rooms (rid, name, password) VALUES ($1, $2, $3)"
	err := Conn.Exec(query, rid, roomName, password)
	if err != nil {
		return err
	}
	return nil
}
