package gomygen

import (
	"bytes"
	"fmt"
	"github.com/yezihack/colorlog"
	"strings"
)

//获取表名的列表
func (m *ModelS) GetTableList(onTable ...string) (tableResult map[string]string, err error) {
	result, err := m.Find("show tables")
	if err != nil {
		return
	}
	tableList := make([]string, 0)
	//获取配置文件里的table_list设定
	for _, mapVal := range result {
		for _, tableName := range mapVal {
			tb := strings.TrimSpace(tableName.(string))
			if len(onTable) > 0 {
				for k := range onTable {
					if onTable[k] == tableName {
						tableList = append(tableList, fmt.Sprintf("'%s'", tb))
						break
					}
				}
			} else {
				tableList = append(tableList, fmt.Sprintf("'%s'", tb))
			}
		}
	}
	tableResult = make(map[string]string)
	var sqlBuff bytes.Buffer
	sqlBuff.WriteString("select distinct table_name, table_comment from information_schema.tables")
	//子查询
	sqlBuff.WriteString(fmt.Sprintf(" WHERE table_name in (%s)", strings.Join(tableList, ",")))
	//查找
	result, err = m.Find(sqlBuff.String())
	if err != nil {
		colorlog.Warn("查找表时,出错了, %v", err)
	}
	for _, info := range result {
		if err != nil {
			return
		}
		table_name := info["table_name"].(string)
		table_comment := info["table_comment"].(string)
		tableResult[table_name] = table_comment
	}
	return
}

//获取表结构详情
func (m *ModelS) GetTableDesc(tableName string) (reply []*TableDesc, err error) {
	result, err := m.Find("select column_name,data_type, column_key, is_nullable,column_default,column_type, column_comment from information_schema.columns where table_name = ?", tableName)
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
