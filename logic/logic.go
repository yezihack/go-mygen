package logic

import (
	"github.com/ThreeKing2018/k3log"
	"github.com/yezihack/m2m/mysql"
)

//生成原生的crud查询数据库
func (l *Logic) CreateCRUD() {
	//读取所有表列表
	dbTools := new(mysql.DbTools)
	tableList, err := dbTools.GetTableList()
	if err != nil {
		panic(err)
	}
	tableNameList := make([]*mysql.TableList, 0)
	//将表结构写入文件
	for tableName, tableComment := range tableList {
		tableNameList = append(tableNameList, &mysql.TableList{
			UpperTableName: l.T.ToUpper(tableName),
			TableName:      tableName,
			Comment:        tableComment,
		})
		//查询表结构信息
		tableDesc, err := l.db.GetTableDesc(tableName)
		if err != nil {
			panic(err)
		}
		//生成基础信息
		err = l.GenerateDBStructure(tableName, tableComment, tableDesc)
		if err != nil {
			panic(err)
		}
		//生成增,删,改,查文件
		err = l.GenerateCURDFile(tableName, tableDesc)
		if err != nil {
			panic(err)
		}
	}
	//生成所有表的文件
	err = l.GenerateTableList(tableNameList)
	if err != nil {
		k3log.Panic("生成所有表的文件", err)
	}
	k3log.Info("加载完成1")
}

//生成mysql markdown文档
func (l *Logic) CreateMarkdown() {
	//读取所有表列表
	dbTools := new(mysql.DbTools)
	tableList, err := dbTools.GetTableList()
	if err != nil {
		k3log.Panic("获取表列表", tableList)
	}
	data := new(mysql.MarkDownData)
	//将表结构写入文件
	i := 1
	for tableName, tableComment := range tableList {
		data.TableList = append(data.TableList, &mysql.TableList{
			Index:          i,
			UpperTableName: l.T.ToUpper(tableName),
			TableName:      tableName,
			Comment:        tableComment,
		})
		//查询表结构信息
		desc := new(mysql.MarkDownDataChild)
		desc.List, err = l.db.GetTableDesc(tableName)
		if err != nil {
			k3log.Panic("获取表详情", err)
		}
		desc.Index = i
		desc.TableName = tableName
		desc.Comment = tableComment
		data.DescList = append(data.DescList, desc)
		i++
	}
	//生成所有表的文件
	err = l.GenerateMarkdown(data)
	if err != nil {
		k3log.Panic("生成MARKDOWN文件", err)
	}
	k3log.Info("加载完成2")
}
