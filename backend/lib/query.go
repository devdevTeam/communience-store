package lib

import "strconv"

func InsertNewUser(uid, mail, password, userName, hash string) error {
	query := "INSERT INTO users (uid, mail, password, name, hash) VALUES ($1, $2, $3, $4, $5)"
	err := Conn.Exec(query, uid, mail, password, userName, hash)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewDefaultCard(uid string) error {
	query := `INSERT INTO default_cards (name, hurigana, birthday, instagram, twitter, facebook, free, hobby, uid) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	err := Conn.Exec(query, "", "", "", "", "", "", "", "", uid)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewRoom(rid, roomName, password string, have_form bool, hash string) error {
	query := "INSERT INTO rooms (rid, name, password, have_form, hash) VALUES ($1, $2, $3, $4, $5)"
	err := Conn.Exec(query, rid, roomName, password, have_form, hash)
	if err != nil {
		return err
	}
	return nil
}

func InsertNewForm(rid string, colList, bool_str_list []string) error {
	for idx, colName := range colList {
		query := "INSERT INTO forms (rid, col_name, col_idx, display) VALUES ($1, $2, $3, $4)"
		Bbool, _ := strconv.ParseBool(bool_str_list[idx])
		err := Conn.Exec(query, rid, colName, idx, Bbool)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertNewCardValue(uid, rid string) error {
	query := "SELECT COUNT(*) FROM forms WHERE rid = $1"
	num, err := Conn.GetRow(query, rid)
	if err != nil {
		return err
	}
	for idx := int64(0); idx < num[0][0].(int64); idx++ {
		query := "INSERT INTO card_value (uid, rid, col_idx, value) VALUES ($1, $2, $3, $4)"
		err := Conn.Exec(query, uid, rid, idx, "")
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

func InsertUserFriendRelation(uid, fid string) error {
	query := "INSERT INTO user_friend_relation (uid, fid) VALUES ($1, $2)"
	err := Conn.Exec(query, uid, fid)
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

func SelectDefaultCard(uid string) ([]string, error) {
	query := "SELECT * FROM default_cards WHERE uid = $1"
	row, err := Conn.GetRow(query, uid)
	if err != nil {
		return nil, err
	}
	var result []string
	for idx, col := range row[0] {
		if idx == 2 {
			result = append(result, col.(string))
		} else {
			result = append(result, col.(string))
		}
	}
	return result, nil
}

func SelectCardValue(uid, rid string) ([]string, error) {
	query := "SELECT COUNT(*) FROM forms WHERE rid = $1"
	num, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []string
	for idx := int64(0); idx < num[0][0].(int64); idx++ {
		query := "SELECT value FROM card_value WHERE uid = $1 AND rid = $2 AND col_idx = $3"
		row, err := Conn.GetRow(query, uid, rid, idx)
		if err != nil {
			return nil, err
		}
		if row[0][0] == nil {
			result = append(result, "null")
		} else {
			result = append(result, row[0][0].(string))
		}
	}
	return result, nil
}

func SelectUserRoomRelation(uid, rid string) (map[string]interface{}, error) {
	query := `SELECT * FROM user_room_relation WHERE uid=$1 AND rid=$2`
	row, err := Conn.GetRow(query, uid, rid)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	if len(row) != 0 {
		result["admin"] = row[0][0].(bool)
		result["uid"] = row[0][1].(string)
		result["rid"] = row[0][2].(string)
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
		tmp["admin"] = row[3].(bool)
		result = append(result, tmp)
		tmp = map[string]interface{}{}
	}
	return result, nil
}

func SelectRoomUserDefaultCard(rid string) ([]interface{}, error) {
	query := `SELECT default_cards.*
				FROM default_cards 
				INNER JOIN user_room_relation ON default_cards.uid = user_room_relation.uid 
				AND user_room_relation.rid=$1`
	rows, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]interface{}{}
	for _, row := range rows {
		tmp["name"] = row[0].(string)
		tmp["hurigana"] = row[1].(string)
		tmp["birthday"] = row[2].(string)
		tmp["instagram"] = row[3].(string)
		tmp["twitter"] = row[4].(string)
		tmp["facebook"] = row[5].(string)
		tmp["free"] = row[6].(string)
		result = append(result, tmp)
		tmp = map[string]interface{}{}
	}
	return result, nil
}

func SelectRoomDisplayInfo_forms(rid string) ([]interface{}, error) {
	query := `SELECT user_room_relation.uid, card_value.value
				FROM user_room_relation
				INNER JOIN forms USING(rid) 
				INNER JOIN card_value USING(uid, col_idx)
				WHERE user_room_relation.rid = $1
				AND display = true`
	rows, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]string{}
	for _, row := range rows {
		tmp["uid"] = row[0].(string)
		tmp["displayValue"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]string{}
	}
	return result, nil
}

func SelectRoomDisplayInfo_default(rid string) ([]interface{}, error) {
	query := `SELECT user_room_relation.uid, default_cards.name
				FROM user_room_relation 
				INNER JOIN default_cards USING(uid)
				WHERE user_room_relation.rid = $1;`
	rows, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]string{}
	for _, row := range rows {
		tmp["uid"] = row[0].(string)
		tmp["displayValue"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]string{}
	}
	return result, nil
}

func SelectUserFriendRelation(uid, fid string) (map[string]interface{}, error) {
	query := `SELECT * FROM user_friend_relation WHERE uid=$1 AND fid=$2`
	row, err := Conn.GetRow(query, uid, fid)
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}
	if len(row) != 0 {
		result["uid"] = row[0][0].(string)
		result["fid"] = row[0][1].(string)
	}
	return result, nil
}

func SelectUserFriend(rid string) ([]interface{}, error) {
	query := `SELECT user_friend_relation.uid, default_cards.name
				FROM user_friend_relation 
				INNER JOIN default_cards ON user_friend_relation.fid = default_cards.uid
				WHERE user_friend_relation.uid = $1`
	rows, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]string{}
	for _, row := range rows {
		tmp["fid"] = row[0].(string)
		tmp["displayValue"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]string{}
	}
	return result, nil
}

func SelectRoom(rid string) ([]interface{}, error) {
	// |rid|name|password|have_form|hash|
	query := "SELECT * FROM rooms WHERE rid = $1"
	row, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	for idx, col := range row[0] {
		if idx == 3 {
			result = append(result, col.(bool))
		} else {
			result = append(result, col.(string))
		}
	}
	return result, nil
}

func SelectRoomWithHash(hash string) ([]string, error) {
	// |rid|name|password|have_form|hash|
	query := "SELECT rid, password FROM rooms WHERE hash = $1"
	row, err := Conn.GetRow(query, hash)
	if err != nil {
		return nil, err
	}
	var result []string
	if len(row) != 0 {
		for _, col := range row[0] {
			result = append(result, col.(string))
		}
	}
	return result, nil
}

func SelectUserWithHash(hash string) ([]string, error) {
	// |rid|name|password|have_form|hash|
	query := "SELECT uid, password FROM users WHERE hash = $1"
	row, err := Conn.GetRow(query, hash)
	if err != nil {
		return nil, err
	}
	var result []string
	if len(row) != 0 {
		for _, col := range row[0] {
			result = append(result, col.(string))
		}
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
	query := `SELECT events.eid, events.org_uid, rooms.name
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
		tmp["org_uid"] = row[1].(string)
		tmp["name"] = row[2].(string)
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

func SelectEventCol(eid string) ([]interface{}, error) {
	query := "SELECT col_name, hidden FROM event_col WHERE eid = $1 ORDER BY col_idx"
	rows, err := Conn.GetRow(query, eid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]interface{}{}
	for _, row := range rows {
		tmp["colName"] = row[0].(string)
		tmp["hidden"] = row[1].(bool)
		result = append(result, tmp)
		tmp = map[string]interface{}{}
	}
	return result, nil
}

func UpdateEventCol(eid string, hidden_list []bool) error {
	for idx, hidden := range hidden_list {
		query := "UPDATE event_col SET hidden=$1 WHERE eid=$2 AND col_idx=$3"
		err := Conn.Exec(query, hidden, eid, idx)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateDefaultCard(uid, name, hurigana, birthday, instagram, twitter, facebook, free, hobby string) error {
	query := `UPDATE default_cards 
				SET name=$1, hurigana=$2, birthday=$3, 
				instagram=$4, twitter=$5, facebook=$6, 
				free=$7, hobby=$8 
				WHERE uid=$9`
	err := Conn.Exec(query, name, hurigana, birthday, instagram, twitter, facebook, free, hobby, uid)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCardValue(uid, rid string, value_list []string) error {
	for idx, value := range value_list {
		query := `UPDATE card_value SET value=$1 WHERE uid=$2 AND rid=$3 AND col_idx=$4`
		err := Conn.Exec(query, value, uid, rid, idx)
		if err != nil {
			return err
		}
	}
	return nil
}

func SearchRoom(name, rid string) ([]interface{}, error) {
	query := ""
	var err error
	var rows [][]interface{}
	if name != "" {
		query = "SELECT name, rid FROM rooms WHERE name = $1"
		rows, err = Conn.GetRow(query, name)
	} else {
		query = "SELECT name, rid FROM rooms WHERE rid = $1"
		rows, err = Conn.GetRow(query, rid)
	}
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]string{}
	for _, row := range rows {
		tmp["name"] = row[0].(string)
		tmp["rid"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]string{}
	}
	return result, nil
}

func SearchRoomUser_form(rid, col_name, value string) ([]interface{}, error) {
	query := `SELECT user_room_relation.uid, card_value.value
				FROM user_room_relation
				INNER JOIN forms USING(rid) 
				INNER JOIN card_value USING(uid, col_idx)
				WHERE user_room_relation.rid = $1
				AND forms.display = true
				AND user_room_relation.uid in (
					SELECT user_room_relation.uid
					FROM user_room_relation
					INNER JOIN forms USING(rid) 
					INNER JOIN card_value USING(uid, col_idx)
					WHERE user_room_relation.rid = $1
					AND forms.col_name = $2
					AND card_value.value LIKE '%` + value + `%'
				)`
	rows, err := Conn.GetRow(query, rid, col_name)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]string{}
	for _, row := range rows {
		tmp["uid"] = row[0].(string)
		tmp["displayValue"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]string{}
	}
	return result, nil
}

func SearchRoomUser_default(rid, col_name, value string) ([]interface{}, error) {
	query := `SELECT user_room_relation.uid, default_cards.name
				FROM user_room_relation
				INNER JOIN default_cards USING(uid)
				WHERE user_room_relation.rid = $1
				AND user_room_relation.uid in (
					SELECT user_room_relation.uid
					FROM user_room_relation
					INNER JOIN default_cards USING(uid)
					WHERE user_room_relation.rid = $1
					AND default_cards.` + col_name + ` LIKE '%` + value + `%'
				)`
	rows, err := Conn.GetRow(query, rid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]string{}
	for _, row := range rows {
		tmp["uid"] = row[0].(string)
		tmp["displayValue"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]string{}
	}
	return result, nil
}

func SearchFriend(uid, col_name, value string) ([]interface{}, error) {
	query := `SELECT user_friend_relation.fid, default_cards.name
				FROM user_friend_relation
				INNER JOIN default_cards ON user_friend_relation.fid = default_cards.uid
				WHERE user_friend_relation.uid = $1
				AND user_friend_relation.fid in (
					SELECT user_friend_relation.fid
					FROM user_friend_relation
					INNER JOIN default_cards ON user_friend_relation.fid = default_cards.uid
					WHERE user_friend_relation.uid = $1
					AND default_cards.` + col_name + ` LIKE '%` + value + `%'
				)`
	rows, err := Conn.GetRow(query, uid)
	if err != nil {
		return nil, err
	}
	var result []interface{}
	tmp := map[string]string{}
	for _, row := range rows {
		tmp["fid"] = row[0].(string)
		tmp["displayValue"] = row[1].(string)
		result = append(result, tmp)
		tmp = map[string]string{}
	}
	return result, nil
}
