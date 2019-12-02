//
package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SgMigrationsModel struct {
	DB *sql.DB
}

func NewSgMigrations(db *sql.DB) *SgMigrationsModel {
	return &SgMigrationsModel{
		DB: db,
	}
}

//获取所有的表字段
func (m *SgMigrationsModel) getColumns() string {
	return " `id`,`migration`,`batch` "
}

//获取多行数据.
func (m *SgMigrationsModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*SgMigrations, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	defer query.Close()
	if err != nil {
		return
	}
	for query.Next() {
		row := SgMigrationsNull{}
		err = query.Scan(
			&row.Id,        //
			&row.Migration, //
			&row.Batch,     //
		)
		if nil != err {
			continue
		}
		rowsResult = append(rowsResult, &SgMigrations{
			Id:        row.Id.Int64,         //
			Migration: row.Migration.String, //
			Batch:     row.Batch.Int64,      //
		})
	}
	return
}

//获取单行数据
func (m *SgMigrationsModel) getRow(sql string, params ...interface{}) (rowResult *SgMigrations, err error) {
	query := m.DB.QueryRow(sql, params...)
	if err != nil {
		return
	}
	row := SgMigrationsNull{}
	err = query.Scan(
		&row.Id,        //
		&row.Migration, //
		&row.Batch,     //
	)
	if nil != err {
		return
	}
	rowResult = &SgMigrations{
		Id:        row.Id.Int64,         //
		Migration: row.Migration.String, //
		Batch:     row.Batch.Int64,      //
	}

	return
}

//_更新数据
func (m *SgMigrationsModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(value...)
	if err != nil {
		return
	}
	var affectCount int64
	affectCount, err = result.RowsAffected()
	if err != nil {
		return
	}
	b = affectCount > 0
	return
}

//新增信息
func (m *SgMigrationsModel) Create(value *SgMigrations) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_SG_MIGRATIONS + " (migration,batch) VALUES (?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Migration, //
		value.Batch,     //
	)
	if err != nil {
		return
	}
	lastId, err = result.LastInsertId()
	if err != nil {
		return
	}
	return
}

//更新数据
func (m *SgMigrationsModel) Update(value *SgMigrations) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_SG_MIGRATIONS + " SET migration=?,batch=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Migration)
	params = append(params, value.Batch)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//查询多行数据
func (m *SgMigrationsModel) Find(value *SgMigrations) (resList []*SgMigrations, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_MIGRATIONS
	resList, err = m.getRows(sqlText)
	return
}

//获取单行数据
func (m *SgMigrationsModel) First(value *SgMigrations) (result *SgMigrations, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_MIGRATIONS + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *SgMigrationsModel) Last(value *SgMigrations) (result *SgMigrations, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_MIGRATIONS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *SgMigrationsModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_SG_MIGRATIONS
	query := m.DB.QueryRow(sqlText)
	if err != nil {
		return
	}
	err = query.Scan(&count)
	if err != nil {
		return
	}
	return
}
