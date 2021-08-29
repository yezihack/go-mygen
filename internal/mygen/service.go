package mygen

import (
	"github.com/yezihack/go-mygen/internal/config"
	"github.com/yezihack/go-mygen/internal/entity"
	"sort"
	"strings"
)

//将表名赋值给结构对象, 供其它方法使用
func (m *entity.ModelS) GetTableNameAndComment() (err error) {
	//读取所有表列表
	if len(m.Tables) == 0 {
		m.Tables, err = m.findDbTables()
		if err != nil {
			return
		}
	}
	//如果正在处理表数据为空,则将所有的表赋值给它
	if len(m.DoTables) == 0 {
		m.DoTables = m.Tables
	}
	return
}

//向数据库里读取所有的表
func (m *entity.ModelS) findDbTables() (NameAndComment []entity.TableNameAndComment, err error) {
	m.l.Lock()
	defer m.l.Unlock()
	result, err := m.Find("SELECT `TABLE_NAME` AS 'table_name', `TABLE_COMMENT` AS 'table_comment' FROM "+
		"information_schema.tables WHERE `TABLE_SCHEMA` = ?", m.DBName)
	if err != nil {
		return
	}
	//获取表名 与 表注释
	NameAndComment = make([]entity.TableNameAndComment, 0)
	//获取库里所有的表名
	for idx, info := range result {
		idx++
		NameAndComment = append(NameAndComment, entity.TableNameAndComment{
			Index:   idx,
			Name:    info["table_name"].(string),
			Comment: info["table_comment"].(string),
		})
	}
	//排序, 采用升序
	sort.Slice(NameAndComment, func(i, j int) bool {
		return strings.ToLower(NameAndComment[i].Name) < strings.ToLower(NameAndComment[j].Name)
	})
	return
}

// 获取表结构详情
func (m *entity.ModelS) GetTableDesc(tableName string) (reply []*entity.TableDesc, err error) {
	m.l.Lock()
	defer m.l.Unlock()
	result, err := m.Find("select `COLUMN_NAME` AS column_name,`DATA_TYPE` AS data_type, `COLUMN_KEY` AS column_key, "+
		"`IS_NULLABLE` AS is_nullable, `COLUMN_DEFAULT` AS column_default,`COLUMN_TYPE` AS column_type, `COLUMN_COMMENT` "+
		"AS column_comment from information_schema.columns where table_name = ? and table_schema = ?",
		tableName, m.DBName)
	if err != nil {
		return
	}
	reply = make([]*entity.TableDesc, 0)
	i := 0
	for _, row := range result {
		var keyBool bool
		if strings.ToUpper(row["column_key"].(string)) == "PRI" {
			keyBool = true
		}
		oriType := row["data_type"].(string)
		var columnDefault string
		val, ok := row["column_default"].(string)
		if ok {
			columnDefault = val
		}
		reply = append(reply, &entity.TableDesc{
			Index:            i,
			ColumnName:       row["column_name"].(string),
			GoColumnName:     m.T.Capitalize(row["column_name"].(string)),
			OriMysqlType:     oriType,
			UpperMysqlType:   strings.ToUpper(oriType),
			GolangType:       config.MysqlTypeToGoType[oriType],
			MysqlNullType:    config.MysqlTypeToGoNullType[oriType],
			ColumnComment:    row["column_comment"].(string),
			IsNull:           row["is_nullable"].(string),
			DefaultValue:     columnDefault,
			ColumnTypeNumber: row["column_type"].(string),
			PrimaryKey:       keyBool,
		})
		i++
	}
	return
}
