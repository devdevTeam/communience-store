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

func InsertNewForm(rid string, colList []string) error {
	for idx, colName := range colList {
		query := "INSERT INTO forms (rid, col_name, col_idx) VALUES ($1, $2, $3)"
		err := Conn.Exec(query, rid, colName, idx)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertUserRoomRelation(uid, rid string, admin bool) error {
	query := "INSERT INTO user_room_relation (uid, rid, admin) VALUES ($1, $2, $3)"
	err := Conn.Exec(query, uid, rid, admin)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewEvent(eid, password, org_uid, rid string) error {
	query := "INSERT INTO events (eid, password, org_uid, rid) VALUES ($1, $2, $3, $4)"
	err := Conn.Exec(query, eid, password, org_uid, rid)
	if err != nil {
		return err
	}
	return nil
}

func InsertEventCol(eid string, colList []string) error {
	for idx, colName := range colList {
		query := "INSERT INTO event_col (eid, col_name, col_idx, hidden) VALUES ($1, $2, $3, $4)"
		err := Conn.Exec(query, eid, colName, idx, false)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertParticipant(eid, uid string) error {
	query := "INSERT INTO participants (eid, uid) VALUES ($1, $2)"
	err := Conn.Exec(query, eid, uid)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUserRoomRelation(uid, rid string) error {
	query := "DELETE FROM user_room_relation WHERE uid = $1 AND rid = $2"
	err := Conn.Exec(query, uid, rid)
	if err != nil {
		return err
	}
	return nil
}

func SelectUser(uid string) ([]string, error) {
	query := "SELECT * FROM users WHERE uid = $1"
	row, err := Conn.GetRow(query, uid)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, col := range row[0] {
		result = append(result, col.(string))
	}
	return result, nil
}

func SelectUserRoomList(uid string) ([]interface{}, error) {
	query := `SELECT rooms.*, user_room_relation.admin
				FROM user_room_relation 
				INNER JOIN rooms ON user_room_relation.rid = rooms.rid AND user_room_relation.uid=$1`
	rows, err := Conn.GetRow(query, uid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]interface{}{}
	for _, row := range rows {
		tmp["rid"] = row[0].(string)
		tmp["name"] = row[1].(string)
		tmp["password"] = row[2].(string)
		tmp["admin"] = row[3].(bool)
		result = append(result, tmp)
		tmp = map[string]interface{}{}
	}
	return result, nil
}

func SelectRoom(rid string) ([]string, error) {
	query := "SELECT * FROM rooms WHERE rid = $1"
	row, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, col := range row[0] {
		result = append(result, col.(string))
	}
	return result, nil
}

func SelectForm(rid string) ([]string, error) {
	query := "SELECT col_name, col_idx FROM forms WHERE rid = $1 ORDER BY col_idx"
	rows, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, row := range rows {
		result = append(result, row[0].(string))
	}
	return result, nil
}

func SelectUserEventList(uid string) ([]interface{}, error) {
	query := `SELECT events.eid, rooms.name
				FROM events 
				INNER JOIN user_room_relation
				ON events.rid = user_room_relation.rid
				AND user_room_relation.uid=$1
				INNER JOIN rooms
				ON events.rid = rooms.rid`
	rows, err := Conn.GetRow(query, uid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]interface{}{}
	for _, row := range rows {
		tmp["eid"] = row[0].(string)
		tmp["name"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]interface{}{}
	}
	return result, nil
}

func SelectEvent(eid string) ([]string, error) {
	// eid | org_uid | passwprd | rid
	query := "SELECT * FROM events WHERE eid = $1"
	row, err := Conn.GetRow(query, eid)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, col := range row[0] {
		result = append(result, col.(string))
	}
	return result, nil
}

func SelectEventCol(eid string) ([]string, error) {
	query := "SELECT col_name, col_idx FROM event_col WHERE eid = $1 ORDER BY col_idx"
	rows, err := Conn.GetRow(query, eid)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, row := range rows {
		result = append(result, row[0].(string))
	}
	return result, nil
}

func UpdateEventCol(eid string, hidden_list []bool) error {
	for idx, hidden := range hidden_list {
		query := "UPDATE event_col SET col_idx=$1, hidden=$2 WHERE eid=$3"
		err := Conn.Exec(query, idx, hidden, eid)
		if err != nil {
			return err
		}
	}
	return nil
}
