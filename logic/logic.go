package logic

import (
	"errors"

	"github.com/yezihack/gm2m/common"
	"github.com/yezihack/gm2m/conf"
	"github.com/yezihack/gm2m/log"
	"github.com/yezihack/gm2m/mysql"
)

//生成结构体文件
func (l *Logic) CreateStructure() error {
	//读取所有表列表
	tableList, err := l.DB.GetTableList()
	if err != nil {
		return err
	}
	//表结构文件路径
	path := common.GetExeRootDir() + "structure/"
	if l.T.IsDirOrFileExist(path) == false {
		if !l.T.CreateDir(path) {
			return errors.New("创建目录失败, path: " + path)
		}
	}
	path += conf.GOFILE_STRUCTURE
	if l.T.IsDirOrFileExist(path) == false {
		if !l.T.CreateFile(path) {
			err = errors.New("创建文件失败, path: " + path)
			return err
		}
	}
	//将表结构写入文件
	for tableName, tableComment := range tableList {
		//查询表结构信息
		tableDesc, err := l.DB.GetTableDesc(tableName)
		if err != nil {
			return err
		}
		//生成基础信息
		err = l.GenerateDBStructure(tableName, tableComment, path, tableDesc)
		if err != nil {
			return err
		}
	}
	return nil
}

//生成原生的crud查询数据库
func (l *Logic) CreateCRUD() error {
	//读取所有表列表
	tableList, err := l.DB.GetTableList()
	if err != nil {
		return err
	}
	tableNameList := make([]*mysql.TableList, 0)
	//表结构文件路径
	structPath := l.GetMysqlDir() + conf.GOFILE_STRUCTURE
	//将表结构写入文件
	for tableName, tableComment := range tableList {
		tableNameList = append(tableNameList, &mysql.TableList{
			UpperTableName: l.T.ToUpper(tableName),
			TableName:      tableName,
			Comment:        tableComment,
		})
		//查询表结构信息
		tableDesc, err := l.DB.GetTableDesc(tableName)
		if err != nil {
			return err
		}
		//生成基础信息
		err = l.GenerateDBStructure(tableName, tableComment, structPath, tableDesc)
		if err != nil {
			return err
		}
		//生成增,删,改,查文件
		err = l.GenerateCURDFile(tableName, tableDesc)
		if err != nil {
			return err
		}
	}
	//生成所有表的文件
	err = l.GenerateTableList(tableNameList)
	if err != nil {
		return err
	}
	log.Info("生成CRUD文件 完成")
	return nil
}

//生成mysql markdown文档
func (l *Logic) CreateMarkdown() error {
	//读取所有表列表
	tableList, err := l.DB.GetTableList()
	if err != nil {
		return err
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
		desc.List, err = l.DB.GetTableDesc(tableName)
		if err != nil {
			return err
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
		return err
	}
	return nil
}
