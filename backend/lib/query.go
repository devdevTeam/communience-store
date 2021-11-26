package lib

func InsertTest(test_str string) error {
	query := `INSERT INTO test (test1) VALUES ($1)`
	err := Conn.Exec(query, test_str)
	if err != nil {
		return err
	}
	return nil
}
