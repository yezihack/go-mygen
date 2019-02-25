package mysql

import (
	"database/sql"
	"encoding/json"
	"strings"

	"github.com/ThreeKing2018/k3log"
)

//打印错误日志
func (t *Book) errWrite(err error, sql string, addParams ...interface{}) {
	type ErrorType struct {
		Book
		AddParams []interface{} `json:"add_params"`
	}
	var infoList []interface{}
	for k, _ := range addParams {
		infoList = append(infoList, addParams[k])
		errType := ErrorType{
			Book:      *t,
			AddParams: infoList,
		}
		paramsJson, _ := json.Marshal(errType)
		k3log.Error("msg", "dbError", "tableName", BOOK, "error", err.Error(), "sql", sql, "params", string(paramsJson))
	}
}

//获取表的所有字段
func (t *Book) selectColumn() string {
	return " `id`,`book_name`,`book_author`,`book_province`,`created_at`,`updated_at` "
}

//查询数据,基础函数
func (t *Book) _selectBody(db *sql.DB, sqlText string, params []interface{}) (_bodyArr []*Book, err error) {
	_bodyArr = make([]*Book, 0)
	rows, err := db.Query(sqlText, params...)
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	defer rows.Close()
	for rows.Next() {
		_one := nullBook{}
		err = rows.Scan(
			&_one.Id,           //
			&_one.BookName,     //
			&_one.BookAuthor,   //
			&_one.BookProvince, //
			&_one.CreatedAt,    //
			&_one.UpdatedAt,    //
		)
		if nil != err {
			t.errWrite(err, sqlText, params)
			continue
		}
		_bodyArr = append(_bodyArr, &Book{
			Id:           int(_one.Id.Int64),       //
			BookName:     _one.BookName.String,     //
			BookAuthor:   _one.BookAuthor.String,   //
			BookProvince: _one.BookProvince.String, //
			CreatedAt:    _one.CreatedAt.String,    //
			UpdatedAt:    _one.UpdatedAt.String,    //
		})
	}
	return
}

//更新数据基础函数
func (t *Book) _updateBody(db *sql.DB, sqlText string, params []interface{}) (b bool, err error) {
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		params...,
	)
	_count, err := res.RowsAffected()
	if nil != err {
		t.errWrite(err, sqlText, params)
		return
	}
	return _count > 0, nil
}

//插入数据
func (t *Book) Insert(db *sql.DB) (b bool, err error) {
	const sqlText = "INSERT INTO " + BOOK + " (book_name,book_author,book_province) VALUES (?,?,?)"
	stmt, err := db.Prepare(sqlText)
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		t.BookName,     //
		t.BookAuthor,   //
		t.BookProvince, //
	)
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	n, err := res.RowsAffected()
	if nil != err {
		t.errWrite(err, sqlText)
		return
	}
	return n > 0, nil
}

//查询单列数据
func (t *Book) First(db *sql.DB, id int64) (result *Book, err error) {
	sqlText := "SELECT" + t.selectColumn() + "FROM " + BOOK + " WHERE id = ? limit 1"
	list, err := t._selectBody(db, sqlText, []interface{}{id})
	if err != nil {
		return
	}
	if len(list) > 0 {
		result = list[0]
	}
	return
}

//查询多条数据,传入id,id2,id3
func (t *Book) Find(db *sql.DB, ids []int) (resultList []*Book, err error) {
	resultList = make([]*Book, 0)
	sqlText := "SELECT" + t.selectColumn() + "FROM " + BOOK + " WHERE id in (" + strings.TrimRight(strings.Repeat("?,", len(ids)), ",") + ")"
	var args []interface{}
	for _, id := range ids {
		args = append(args, id)
	}
	resultList, err = t._selectBody(db, sqlText, args)
	if err != nil {
		return
	}
	return
}

//主键更新数据
func (t *Book) Update(db *sql.DB, id int) (err error) {
	sqlText := "UPDATE " + BOOK + " SET book_name=?,book_author=?,book_province=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, t.BookName)
	params = append(params, t.BookAuthor)
	params = append(params, t.BookProvince)

	//主键id
	params = append(params, id)
	_, err = t._updateBody(db, sqlText, params)
	if err != nil {
		return
	}
	return
}
