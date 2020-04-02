//
package models

import (
	"database/sql"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"
)

type SgBlogTagsModel struct {
	DB *sql.DB
}

func NewSgBlogTags(db *sql.DB) *SgBlogTagsModel {
	return &SgBlogTagsModel{
		DB: db,
	}
}

//获取所有的表字段
func (m *SgBlogTagsModel) getColumns() string {
	return " `id`,`blog_id`,`tag_id`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *SgBlogTagsModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*SgBlogTags, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	if err != nil {
		return
	}
	defer query.Close()
	for query.Next() {
		row := SgBlogTagsNull{}
		err = query.Scan(
			&row.Id,        //
			&row.BlogId,    //文章ID,与blogs表相关
			&row.TagId,     //标签ID,与tags表相关
			&row.CreatedAt, //创建时间
			&row.UpdatedAt, //更新时间
		)
		if err != nil && err != sql.ErrNoRows {
			err = errors.WithStack(err)
			return
		}
		rowsResult = append(rowsResult, &SgBlogTags{
			Id:        row.Id.Int64,        //
			BlogId:    row.BlogId.Int64,    //文章ID,与blogs表相关
			TagId:     row.TagId.Int64,     //标签ID,与tags表相关
			CreatedAt: row.CreatedAt.Int64, //创建时间
			UpdatedAt: row.UpdatedAt.Int64, //更新时间
		})
	}
	return
}

//获取单行数据
func (m *SgBlogTagsModel) getRow(sql string, params ...interface{}) (rowResult *SgBlogTags, err error) {
	query := m.DB.QueryRow(sql, params...)
	row := SgBlogTagsNull{}
	err = query.Scan(
		&row.Id,        //
		&row.BlogId,    //文章ID,与blogs表相关
		&row.TagId,     //标签ID,与tags表相关
		&row.CreatedAt, //创建时间
		&row.UpdatedAt, //更新时间
	)
	if nil != err {
		err = errors.WithStack(err)
		return
	}
	rowResult = &SgBlogTags{
		Id:        row.Id.Int64,        //
		BlogId:    row.BlogId.Int64,    //文章ID,与blogs表相关
		TagId:     row.TagId.Int64,     //标签ID,与tags表相关
		CreatedAt: row.CreatedAt.Int64, //创建时间
		UpdatedAt: row.UpdatedAt.Int64, //更新时间
	}

	return
}

//_更新数据
func (m *SgBlogTagsModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
	stmt, err := m.DB.Prepare(sqlTxt)
	if err != nil {
		err = errors.WithStack(err)
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
func (m *SgBlogTagsModel) Create(value *SgBlogTags) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_SG_BLOG_TAGS + " (blog_id,tag_id) VALUES (?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.BlogId, //文章ID,与blogs表相关
		value.TagId,  //标签ID,与tags表相关
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
func (m *SgBlogTagsModel) Update(value *SgBlogTags) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_SG_BLOG_TAGS + " SET blog_id=?,tag_id=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.BlogId)
	params = append(params, value.TagId)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//查询多行数据
func (m *SgBlogTagsModel) Find(value *SgBlogTags) (resList []*SgBlogTags, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_BLOG_TAGS
	resList, err = m.getRows(sqlText)
	return
}

//获取单行数据
func (m *SgBlogTagsModel) First(value *SgBlogTags) (result *SgBlogTags, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_BLOG_TAGS + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *SgBlogTagsModel) Last(value *SgBlogTags) (result *SgBlogTags, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_BLOG_TAGS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *SgBlogTagsModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_SG_BLOG_TAGS
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
