//收藏文章表
package models

import (
	"database/sql"
	"strings"
)

type CellectModel struct {
	DB *sql.DB
	Tx *sql.Tx
}

func NewCellect(db *sql.DB) *CellectModel {
	return &CellectModel{
		DB: db,
	}
}

func NewCellectTx(db *sql.Tx) *CellectModel {
	return &CellectModel{
		Tx: db,
	}
}

//获取所有的表字段
func (m *CellectModel) getColumns() string {
	return " `id`,`article_id`,`user_id`,`state`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *CellectModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*Cellect, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer query.Close()
	for query.Next() {
		row := CellectNull{}
		err = query.Scan(
			&row.Id,        //
			&row.ArticleId, //文章ID
			&row.UserId,    //用户ID
			&row.State,     //状态,1公开,0私用
			&row.CreatedAt, //创建时间
			&row.UpdatedAt, //更新时间
		)
		if nil != err && err != sql.ErrNoRows {
			continue
		}
		if err == sql.ErrNoRows {
			err = nil
		}
		rowsResult = append(rowsResult, &Cellect{
			Id:        row.Id.Int64,        //
			ArticleId: row.ArticleId.Int64, //文章ID
			UserId:    row.UserId.Int64,    //用户ID
			State:     row.State.Int64,     //状态,1公开,0私用
			CreatedAt: row.CreatedAt.Time,  //创建时间
			UpdatedAt: row.UpdatedAt.Time,  //更新时间
		})
	}
	return
}

//获取单行数据
func (m *CellectModel) getRow(sqlTxt string, params ...interface{}) (rowResult *Cellect, err error) {
	query := m.DB.QueryRow(sqlTxt, params...)
	row := CellectNull{}
	err = query.Scan(
		&row.Id,        //
		&row.ArticleId, //文章ID
		&row.UserId,    //用户ID
		&row.State,     //状态,1公开,0私用
		&row.CreatedAt, //创建时间
		&row.UpdatedAt, //更新时间
	)
	if nil != err && err != sql.ErrNoRows {
		return
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	rowResult = &Cellect{
		Id:        row.Id.Int64,        //
		ArticleId: row.ArticleId.Int64, //文章ID
		UserId:    row.UserId.Int64,    //用户ID
		State:     row.State.Int64,     //状态,1公开,0私用
		CreatedAt: row.CreatedAt.Time,  //创建时间
		UpdatedAt: row.UpdatedAt.Time,  //更新时间
	}

	return
}

//_更新数据
func (m *CellectModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *CellectModel) Create(value *Cellect) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_CELLECT + " (article_id,user_id,state) VALUES (?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.ArticleId, //文章ID
		value.UserId,    //用户ID
		value.State,     //状态,1公开,0私用
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
func (m *CellectModel) Update(value *Cellect) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_CELLECT + " SET article_id=?,user_id=?,state=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.ArticleId)
	params = append(params, value.UserId)
	params = append(params, value.State)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//_更新数据 支持事务
func (m *CellectModel) SaveTx(sqlTxt string, value ...interface{}) (b bool, err error) {
	stmt, err := m.Tx.Prepare(sqlTxt)
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

//新增信息 支持事务
func (m *CellectModel) CreateTx(value *Cellect) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_CELLECT + " (article_id,user_id,state) VALUES (?,?,?)"
	stmt, err := m.Tx.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.ArticleId, //文章ID
		value.UserId,    //用户ID
		value.State,     //状态,1公开,0私用
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

//更新数据 支持事务
func (m *CellectModel) UpdateTx(value *Cellect) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_CELLECT + " SET article_id=?,user_id=?,state=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.ArticleId)
	params = append(params, value.UserId)
	params = append(params, value.State)
	params = append(params, value.Id)

	return m.SaveTx(sqlText, params...)
}

//查询多行数据
func (m *CellectModel) Find(value *Cellect) (resList []*Cellect, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_CELLECT
	resList, err = m.getRows(sqlText)
	return
}

//In 查询多行数据
func (m *CellectModel) FindIn(ids []int) (resList []*Cellect, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_CELLECT + " WHERE id in (" + strings.TrimRight(strings.Repeat("?,", len(ids)), ",") + ")"
	param := make([]interface{}, 0)
	for _, id := range ids {
		param = append(param, id)
	}
	resList, err = m.getRows(sqlText, param...)
	return
}

//获取单行数据
func (m *CellectModel) First(value *Cellect) (result *Cellect, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_CELLECT + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *CellectModel) Last(value *Cellect) (result *Cellect, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_CELLECT + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *CellectModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_CELLECT
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
