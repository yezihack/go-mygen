//
package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type SgUsersModel struct {
	DB *sql.DB
}

func NewSgUsers(db *sql.DB) *SgUsersModel {
	return &SgUsersModel{
		DB: db,
	}
}

//获取所有的表字段
func (m *SgUsersModel) getColumns() string {
	return " `id`,`name`,`pass`,`salt`,`google2fa_key`,`login_count`,`token`,`login_ip`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *SgUsersModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*SgUsers, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	defer query.Close()
	if err != nil {
		return
	}
	for query.Next() {
		row := SgUsersNull{}
		err = query.Scan(
			&row.Id,           //
			&row.Name,         //用户名
			&row.Pass,         //密码
			&row.Salt,         //盐
			&row.Google2faKey, //谷歌两步验证密钥
			&row.LoginCount,   //登陆次数
			&row.Token,        //记住密码生成的随机数
			&row.LoginIp,      //登陆IP
			&row.CreatedAt,    //创建时间
			&row.UpdatedAt,    //更新时间
		)
		if nil != err {
			continue
		}
		rowsResult = append(rowsResult, &SgUsers{
			Id:           row.Id.Int64,            //
			Name:         row.Name.String,         //用户名
			Pass:         row.Pass.String,         //密码
			Salt:         row.Salt.String,         //盐
			Google2faKey: row.Google2faKey.String, //谷歌两步验证密钥
			LoginCount:   row.LoginCount.Int64,    //登陆次数
			Token:        row.Token.String,        //记住密码生成的随机数
			LoginIp:      row.LoginIp.String,      //登陆IP
			CreatedAt:    row.CreatedAt.Int64,     //创建时间
			UpdatedAt:    row.UpdatedAt.Int64,     //更新时间
		})
	}
	return
}

//获取单行数据
func (m *SgUsersModel) getRow(sql string, params ...interface{}) (rowResult *SgUsers, err error) {
	query := m.DB.QueryRow(sql, params...)
	if err != nil {
		return
	}
	row := SgUsersNull{}
	err = query.Scan(
		&row.Id,           //
		&row.Name,         //用户名
		&row.Pass,         //密码
		&row.Salt,         //盐
		&row.Google2faKey, //谷歌两步验证密钥
		&row.LoginCount,   //登陆次数
		&row.Token,        //记住密码生成的随机数
		&row.LoginIp,      //登陆IP
		&row.CreatedAt,    //创建时间
		&row.UpdatedAt,    //更新时间
	)
	if nil != err {
		return
	}
	rowResult = &SgUsers{
		Id:           row.Id.Int64,            //
		Name:         row.Name.String,         //用户名
		Pass:         row.Pass.String,         //密码
		Salt:         row.Salt.String,         //盐
		Google2faKey: row.Google2faKey.String, //谷歌两步验证密钥
		LoginCount:   row.LoginCount.Int64,    //登陆次数
		Token:        row.Token.String,        //记住密码生成的随机数
		LoginIp:      row.LoginIp.String,      //登陆IP
		CreatedAt:    row.CreatedAt.Int64,     //创建时间
		UpdatedAt:    row.UpdatedAt.Int64,     //更新时间
	}

	return
}

//_更新数据
func (m *SgUsersModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *SgUsersModel) Create(value *SgUsers) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_SG_USERS + " (name,pass,salt,google2fa_key,login_count,token,login_ip) VALUES (?,?,?,?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Name,         //用户名
		value.Pass,         //密码
		value.Salt,         //盐
		value.Google2faKey, //谷歌两步验证密钥
		value.LoginCount,   //登陆次数
		value.Token,        //记住密码生成的随机数
		value.LoginIp,      //登陆IP
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
func (m *SgUsersModel) Update(value *SgUsers) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_SG_USERS + " SET name=?,pass=?,salt=?,google2fa_key=?,login_count=?,token=?,login_ip=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Name)
	params = append(params, value.Pass)
	params = append(params, value.Salt)
	params = append(params, value.Google2faKey)
	params = append(params, value.LoginCount)
	params = append(params, value.Token)
	params = append(params, value.LoginIp)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//查询多行数据
func (m *SgUsersModel) Find(value *SgUsers) (resList []*SgUsers, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_USERS
	resList, err = m.getRows(sqlText)
	return
}

//获取单行数据
func (m *SgUsersModel) First(value *SgUsers) (result *SgUsers, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_USERS + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *SgUsersModel) Last(value *SgUsers) (result *SgUsers, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_USERS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *SgUsersModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_SG_USERS
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
