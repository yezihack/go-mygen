//
package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SgTagsModel struct {
	DB *sql.DB
}

func NewSgTags(db *sql.DB) *SgTagsModel {
	return &SgTagsModel{
		DB: db,
	}
}

//获取所有的表字段
func (m *SgTagsModel) getColumns() string {
	return " `id`,`name`,`remark`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *SgTagsModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*SgTags, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	defer query.Close()
	if err != nil {
		return
	}
	for query.Next() {
		row := SgTagsNull{}
		err = query.Scan(
			&row.Id,        //
			&row.Name,      //名称
			&row.Remark,    //备注
			&row.CreatedAt, //创建时间
			&row.UpdatedAt, //更新时间
		)
		if nil != err {
			continue
		}
		rowsResult = append(rowsResult, &SgTags{
			Id:        row.Id.Int64,        //
			Name:      row.Name.String,     //名称
			Remark:    row.Remark.String,   //备注
			CreatedAt: row.CreatedAt.Int64, //创建时间
			UpdatedAt: row.UpdatedAt.Int64, //更新时间
		})
	}
	return
}

//获取单行数据
func (m *SgTagsModel) getRow(sql string, params ...interface{}) (rowResult *SgTags, err error) {
	query := m.DB.QueryRow(sql, params...)
	if err != nil {
		return
	}
	row := SgTagsNull{}
	err = query.Scan(
		&row.Id,        //
		&row.Name,      //名称
		&row.Remark,    //备注
		&row.CreatedAt, //创建时间
		&row.UpdatedAt, //更新时间
	)
	if nil != err {
		return
	}
	rowResult = &SgTags{
		Id:        row.Id.Int64,        //
		Name:      row.Name.String,     //名称
		Remark:    row.Remark.String,   //备注
		CreatedAt: row.CreatedAt.Int64, //创建时间
		UpdatedAt: row.UpdatedAt.Int64, //更新时间
	}

	return
}

//_更新数据
func (m *SgTagsModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *SgTagsModel) Create(value *SgTags) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_SG_TAGS + " (name,remark) VALUES (?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Name,   //名称
		value.Remark, //备注
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
func (m *SgTagsModel) Update(value *SgTags) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_SG_TAGS + " SET name=?,remark=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Name)
	params = append(params, value.Remark)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//查询多行数据
func (m *SgTagsModel) Find(value *SgTags) (resList []*SgTags, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_TAGS
	resList, err = m.getRows(sqlText)
	return
}

//获取单行数据
func (m *SgTagsModel) First(value *SgTags) (result *SgTags, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_TAGS + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *SgTagsModel) Last(value *SgTags) (result *SgTags, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_TAGS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *SgTagsModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_SG_TAGS
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
