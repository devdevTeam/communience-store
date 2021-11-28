package lib

import (
	"time"
)

func InsertNewUser(uid, mail, password, userName string) error {
	query := "INSERT INTO users (uid, mail, password, name) VALUES ($1, $2, $3, $4)"
	err := Conn.Exec(query, uid, mail, password, userName)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewDefaultCard(uid string) error {
	query := `INSERT INTO default_cards (name, hurigana, birthday, instagram, twitter, facebook, free, uid) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	err := Conn.Exec(query, nil, nil, nil, nil, nil, nil, nil, uid)
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

func SelectDefaultCard(uid string) ([]string, error) {
	query := "SELECT * FROM default_cards WHERE uid = $1"
	row, err := Conn.GetRow(query, uid)
	if err != nil {
		return nil, err
	}
	var result []string
	var layout = "2006-01-02 15:04:05"
	for idx, col := range row[0] {
		if idx == 2 {
			result = append(result, col.(time.Time).Format(layout))
		} else {
			result = append(result, col.(string))
		}
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

func UpdateDefaultCard(uid, name, hurigana, birthday, instagram, twitter, facebook, free string) error {
	query := `UPDATE default_cards 
				SET name=$1, hurigana=$2, birthday=$3, 
				instagram=$4, twitter=$5, facebook=$6, free=$7 
				WHERE uid=$8`
	err := Conn.Exec(query, name, hurigana, birthday, instagram, twitter, facebook, free, uid)
	if err != nil {
		return err
	}
	return nil
}
