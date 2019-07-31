//用户表
package models

import (
	"database/sql"
	"strings"
)

type UsersModel struct {
	DB *sql.DB
	Tx *sql.Tx
}

func NewUsers(db *sql.DB) *UsersModel {
	return &UsersModel{
		DB: db,
	}
}

func NewUsersTx(db *sql.Tx) *UsersModel {
	return &UsersModel{
		Tx: db,
	}
}

//获取所有的表字段
func (m *UsersModel) getColumns() string {
	return " `id`,`username`,`email`,`password`,`openid`,`state`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *UsersModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*Users, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer query.Close()
	for query.Next() {
		row := UsersNull{}
		err = query.Scan(
			&row.Id,        //
			&row.Username,  //用户名称
			&row.Email,     //用户邮箱
			&row.Password,  //用户密码
			&row.Openid,    //微信OPENID
			&row.State,     //用户状态,1使用中,0禁用
			&row.CreatedAt, //创建时间
			&row.UpdatedAt, //更新时间
		)
		if nil != err && err != sql.ErrNoRows {
			continue
		}
		//if err == sql.ErrNoRows {
		//	err = nil
		//}
		rowsResult = append(rowsResult, &Users{
			Id:        row.Id.Int64,        //
			Username:  row.Username.String, //用户名称
			Email:     row.Email.String,    //用户邮箱
			Password:  row.Password.String, //用户密码
			Openid:    row.Openid.String,   //微信OPENID
			State:     row.State.Int64,     //用户状态,1使用中,0禁用
			CreatedAt: row.CreatedAt.Time,  //创建时间
			UpdatedAt: row.UpdatedAt.Time,  //更新时间
		})
	}
	return
}

//获取单行数据
func (m *UsersModel) getRow(sqlTxt string, params ...interface{}) (rowResult *Users, err error) {
	query := m.DB.QueryRow(sqlTxt, params...)
	row := UsersNull{}
	err = query.Scan(
		&row.Id,        //
		&row.Username,  //用户名称
		&row.Email,     //用户邮箱
		&row.Password,  //用户密码
		&row.Openid,    //微信OPENID
		&row.State,     //用户状态,1使用中,0禁用
		&row.CreatedAt, //创建时间
		&row.UpdatedAt, //更新时间
	)
	if nil != err && err != sql.ErrNoRows {
		return
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	rowResult = &Users{
		Id:        row.Id.Int64,        //
		Username:  row.Username.String, //用户名称
		Email:     row.Email.String,    //用户邮箱
		Password:  row.Password.String, //用户密码
		Openid:    row.Openid.String,   //微信OPENID
		State:     row.State.Int64,     //用户状态,1使用中,0禁用
		CreatedAt: row.CreatedAt.Time,  //创建时间
		UpdatedAt: row.UpdatedAt.Time,  //更新时间
	}

	return
}

//_更新数据
func (m *UsersModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *UsersModel) Create(value *Users) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_USERS + " (username,email,password,openid,state) VALUES (?,?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Username, //用户名称
		value.Email,    //用户邮箱
		value.Password, //用户密码
		value.Openid,   //微信OPENID
		value.State,    //用户状态,1使用中,0禁用
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
func (m *UsersModel) Update(value *Users) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_USERS + " SET username=?,email=?,password=?,openid=?,state=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Username)
	params = append(params, value.Email)
	params = append(params, value.Password)
	params = append(params, value.Openid)
	params = append(params, value.State)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//_更新数据 支持事务
func (m *UsersModel) SaveTx(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *UsersModel) CreateTx(value *Users) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_USERS + " (username,email,password,openid,state) VALUES (?,?,?,?,?)"
	stmt, err := m.Tx.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Username, //用户名称
		value.Email,    //用户邮箱
		value.Password, //用户密码
		value.Openid,   //微信OPENID
		value.State,    //用户状态,1使用中,0禁用
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
func (m *UsersModel) UpdateTx(value *Users) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_USERS + " SET username=?,email=?,password=?,openid=?,state=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Username)
	params = append(params, value.Email)
	params = append(params, value.Password)
	params = append(params, value.Openid)
	params = append(params, value.State)
	params = append(params, value.Id)

	return m.SaveTx(sqlText, params...)
}

//查询多行数据
func (m *UsersModel) Find(value *Users) (resList []*Users, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_USERS
	resList, err = m.getRows(sqlText)
	return
}

//In 查询多行数据
func (m *UsersModel) FindIn(ids []int) (resList []*Users, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_USERS + " WHERE id in (" + strings.TrimRight(strings.Repeat("?,", len(ids)), ",") + ")"
	param := make([]interface{}, 0)
	for _, id := range ids {
		param = append(param, id)
	}
	resList, err = m.getRows(sqlText, param...)
	return
}

//获取单行数据
func (m *UsersModel) First(value *Users) (result *Users, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_USERS + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *UsersModel) Last(value *Users) (result *Users, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_USERS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *UsersModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_USERS
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
