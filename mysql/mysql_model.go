package mysql

import "database/sql"

type ModelS struct {
	DB *sql.DB
}

func NewDB(db *sql.DB) *ModelS {
	return &ModelS{
		DB:db,
	}
}

func (m *ModelS) Find(sql string, args ...interface{}) ([]string, error) {
	resultMap := make(map[string]string)
	query, err := m.DB.Query(sql, args...)
	defer query.Close()
	if err != nil {
		return nil, err
	}
	for query.Next()  {
		query.Columns()
	}
	return
}
