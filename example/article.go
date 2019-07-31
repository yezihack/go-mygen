//文章表
package models

import (
	"database/sql"
	"strings"
)

type ArticleModel struct {
	DB *sql.DB
	Tx *sql.Tx
}

func NewArticle(db *sql.DB) *ArticleModel {
	return &ArticleModel{
		DB: db,
	}
}

func NewArticleTx(db *sql.Tx) *ArticleModel {
	return &ArticleModel{
		Tx: db,
	}
}

//获取所有的表字段
func (m *ArticleModel) getColumns() string {
	return " `id`,`title`,`content`,`read_count`,`state`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *ArticleModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*Article, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil && err != sql.ErrNoRows {
		return
	}
	defer query.Close()
	for query.Next() {
		row := ArticleNull{}
		err = query.Scan(
			&row.Id,        //
			&row.Title,     //文章标题
			&row.Content,   //文章内容
			&row.ReadCount, //阅读数量
			&row.State,     //状态,1使用中,0禁用
			&row.CreatedAt, //创建时间
			&row.UpdatedAt, //更新时间
		)
		if nil != err && err != sql.ErrNoRows {
			continue
		}
		if err == sql.ErrNoRows {
			err = nil
		}
		rowsResult = append(rowsResult, &Article{
			Id:        row.Id.Int64,        //
			Title:     row.Title.String,    //文章标题
			Content:   row.Content.String,  //文章内容
			ReadCount: row.ReadCount.Int64, //阅读数量
			State:     row.State.Int64,     //状态,1使用中,0禁用
			CreatedAt: row.CreatedAt.Time,  //创建时间
			UpdatedAt: row.UpdatedAt.Time,  //更新时间
		})
	}
	return
}

//获取单行数据
func (m *ArticleModel) getRow(sqlTxt string, params ...interface{}) (rowResult *Article, err error) {
	query := m.DB.QueryRow(sqlTxt, params...)
	row := ArticleNull{}
	err = query.Scan(
		&row.Id,        //
		&row.Title,     //文章标题
		&row.Content,   //文章内容
		&row.ReadCount, //阅读数量
		&row.State,     //状态,1使用中,0禁用
		&row.CreatedAt, //创建时间
		&row.UpdatedAt, //更新时间
	)
	if nil != err && err != sql.ErrNoRows {
		return
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	rowResult = &Article{
		Id:        row.Id.Int64,        //
		Title:     row.Title.String,    //文章标题
		Content:   row.Content.String,  //文章内容
		ReadCount: row.ReadCount.Int64, //阅读数量
		State:     row.State.Int64,     //状态,1使用中,0禁用
		CreatedAt: row.CreatedAt.Time,  //创建时间
		UpdatedAt: row.UpdatedAt.Time,  //更新时间
	}

	return
}

//_更新数据
func (m *ArticleModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *ArticleModel) Create(value *Article) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_ARTICLE + " (title,content,read_count,state) VALUES (?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Title,     //文章标题
		value.Content,   //文章内容
		value.ReadCount, //阅读数量
		value.State,     //状态,1使用中,0禁用
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
func (m *ArticleModel) Update(value *Article) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_ARTICLE + " SET title=?,content=?,read_count=?,state=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Title)
	params = append(params, value.Content)
	params = append(params, value.ReadCount)
	params = append(params, value.State)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//_更新数据 支持事务
func (m *ArticleModel) SaveTx(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *ArticleModel) CreateTx(value *Article) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_ARTICLE + " (title,content,read_count,state) VALUES (?,?,?,?)"
	stmt, err := m.Tx.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.Title,     //文章标题
		value.Content,   //文章内容
		value.ReadCount, //阅读数量
		value.State,     //状态,1使用中,0禁用
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
func (m *ArticleModel) UpdateTx(value *Article) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_ARTICLE + " SET title=?,content=?,read_count=?,state=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.Title)
	params = append(params, value.Content)
	params = append(params, value.ReadCount)
	params = append(params, value.State)
	params = append(params, value.Id)

	return m.SaveTx(sqlText, params...)
}

//查询多行数据
func (m *ArticleModel) Find(value *Article) (resList []*Article, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_ARTICLE
	resList, err = m.getRows(sqlText)
	return
}

//In 查询多行数据
func (m *ArticleModel) FindIn(ids []int) (resList []*Article, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_ARTICLE + " WHERE id in (" + strings.TrimRight(strings.Repeat("?,", len(ids)), ",") + ")"
	param := make([]interface{}, 0)
	for _, id := range ids {
		param = append(param, id)
	}
	resList, err = m.getRows(sqlText, param...)
	return
}

//获取单行数据
func (m *ArticleModel) First(value *Article) (result *Article, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_ARTICLE + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *ArticleModel) Last(value *Article) (result *Article, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_ARTICLE + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *ArticleModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_ARTICLE
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
