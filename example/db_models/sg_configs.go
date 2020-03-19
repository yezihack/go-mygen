//
package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SgConfigsModel struct {
	DB *sql.DB
}

func NewSgConfigs(db *sql.DB) *SgConfigsModel {
	return &SgConfigsModel{
		DB: db,
	}
}

//获取所有的表字段
func (m *SgConfigsModel) getColumns() string {
	return " `id`,`name`,`key`,`value`,`config_type`,`remark`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *SgConfigsModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*SgConfigs, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	defer query.Close()
	if err != nil {
		return
	}
	for query.Next() {
		row := SgConfigsNull{}
		err = query.Scan(
			&row.Id,         //
			&row.Name,       //名称
			&row.Key,        //键
			&row.Value,      //值
			&row.ConfigType, //变量类型：system 系统内置/user 用户定义
			&row.Remark,     //
			&row.CreatedAt,  //创建时间
			&row.UpdatedAt,  //更新时间
		)
		if nil != err {
			continue
		}
		rowsResult = append(rowsResult, &SgConfigs{
			Id:         row.Id.Int64,          //
			Name:       row.Name.String,       //名称
			Key:        row.Key.String,        //键
			Value:      row.Value.String,      //值
			ConfigType: row.ConfigType.String, //变量类型：system 系统内置/user 用户定义
			Remark:     row.Remark.String,     //
			CreatedAt:  row.CreatedAt.Int64,   //创建时间
			UpdatedAt:  row.UpdatedAt.Int64,   //更新时间
		})
	}
	return
}

//获取单行数据
func (m *SgConfigsModel) getRow(sql string, params ...interface{}) (rowResult *SgConfigs, err error) {
	query := m.DB.QueryRow(sql, params...)
	if err != nil {
		return
	}
	row := SgConfigsNull{}
	err = query.Scan(
		&row.Id,         //
		&row.Name,       //名称
		&row.Key,        //键
		&row.Value,      //值
		&row.ConfigType, //变量类型：system 系统内置/user 用户定义
		&row.Remark,     //
		&row.CreatedAt,  //创建时间
		&row.UpdatedAt,  //更新时间
	)
	if nil != err {
		return
	}
	rowResult = &SgConfigs{
		Id:         row.Id.Int64,          //
		Name:       row.Name.String,       //名称
		Key:        row.Key.String,        //键
		Value:      row.Value.String,      //值
		ConfigType: row.ConfigType.String, //变量类型：system 系统内置/user 用户定义
		Remark:     row.Remark.String,     //
		CreatedAt:  row.CreatedAt.Int64,   //创建时间
		UpdatedAt:  row.UpdatedAt.Int64,   //更新时间
	}

	return
}

//_更新数据
func (m *SgConfigsModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *SgConfigsModel) Create(value *SgConfigs) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_SG_CONFIGS + " (name,key,value,config_type,remark) VALUES (?,?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Name,       //名称
		value.Key,        //键
		value.Value,      //值
		value.ConfigType, //变量类型：system 系统内置/user 用户定义
		value.Remark,     //
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
func (m *SgConfigsModel) Update(value *SgConfigs) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_SG_CONFIGS + " SET name=?,key=?,value=?,config_type=?,remark=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Name)
	params = append(params, value.Key)
	params = append(params, value.Value)
	params = append(params, value.ConfigType)
	params = append(params, value.Remark)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//查询多行数据
func (m *SgConfigsModel) Find(value *SgConfigs) (resList []*SgConfigs, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_CONFIGS
	resList, err = m.getRows(sqlText)
	return
}

//获取单行数据
func (m *SgConfigsModel) First(value *SgConfigs) (result *SgConfigs, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_CONFIGS + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *SgConfigsModel) Last(value *SgConfigs) (result *SgConfigs, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_CONFIGS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *SgConfigsModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_SG_CONFIGS
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
