package main

import (
	"sort"
	"strings"
)

//将表名赋值给结构对象, 供其它方法使用
func (m *ModelS) GetTableNameAndComment() (err error) {
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
func (m *ModelS) findDbTables() (NameAndComment []TableNameAndComment, err error) {
	m.l.Lock()
	defer m.l.Unlock()
	result, err := m.Find("SELECT `table_name`, table_comment FROM information_schema.tables WHERE table_schema = ?", m.DBName)
	if err != nil {
		return
	}
	//获取表名 与 表注释
	NameAndComment = make([]TableNameAndComment, 0)
	//获取库里所有的表名
	for idx, info := range result {
		idx++
		NameAndComment = append(NameAndComment, TableNameAndComment{
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
func (m *ModelS) GetTableDesc(tableName string) (reply []*TableDesc, err error) {
	m.l.Lock()
	defer m.l.Unlock()
	result, err := m.Find("select column_name,data_type, column_key, is_nullable,column_default,column_type, column_comment from information_schema.columns "+
		"where table_name = ? and table_schema = ?", tableName, m.DBName)
	if err != nil {
		return
	}
	reply = make([]*TableDesc, 0)
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
		reply = append(reply, &TableDesc{
			Index:            i,
			ColumnName:       row["column_name"].(string),
			GoColumnName:     m.T.Capitalize(row["column_name"].(string)),
			OriMysqlType:     oriType,
			UpperMysqlType:   strings.ToUpper(oriType),
			GolangType:       MysqlTypeToGoType[oriType],
			MysqlNullType:    MysqlTypeToGoNullType[oriType],
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
