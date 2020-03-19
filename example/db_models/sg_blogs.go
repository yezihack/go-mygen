//
package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type SgBlogsModel struct {
	DB *sql.DB
}

func NewSgBlogs(db *sql.DB) *SgBlogsModel {
	return &SgBlogsModel{
		DB: db,
	}
}

//获取所有的表字段
func (m *SgBlogsModel) getColumns() string {
	return " `id`,`user_id`,`title`,`content`,`type`,`is_push`,`status`,`first_tag_id`,`read_count`,`star_count`,`created_at`,`updated_at` "
}

//获取多行数据.
func (m *SgBlogsModel) getRows(sqlTxt string, params ...interface{}) (rowsResult []*SgBlogs, err error) {
	query, err := m.DB.Query(sqlTxt, params...)
	defer query.Close()
	if err != nil {
		return
	}
	for query.Next() {
		row := SgBlogsNull{}
		err = query.Scan(
			&row.Id,         //
			&row.UserId,     //用户ID,与users表相关
			&row.Title,      //标题
			&row.Content,    //
			&row.Type,       //文章类型,1为原创,2为转载
			&row.IsPush,     //是否推送，0未，1是
			&row.Status,     //文章状态:1发布,2草稿
			&row.FirstTagId, //首标签ID
			&row.ReadCount,  //阅读次数
			&row.StarCount,  //点赞次数
			&row.CreatedAt,  //创建时间
			&row.UpdatedAt,  //更新时间
		)
		if nil != err {
			continue
		}
		rowsResult = append(rowsResult, &SgBlogs{
			Id:         row.Id.Int64,         //
			UserId:     row.UserId.Int64,     //用户ID,与users表相关
			Title:      row.Title.String,     //标题
			Content:    row.Content.String,   //
			Type:       row.Type.Int64,       //文章类型,1为原创,2为转载
			IsPush:     row.IsPush.Int64,     //是否推送，0未，1是
			Status:     row.Status.Int64,     //文章状态:1发布,2草稿
			FirstTagId: row.FirstTagId.Int64, //首标签ID
			ReadCount:  row.ReadCount.Int64,  //阅读次数
			StarCount:  row.StarCount.Int64,  //点赞次数
			CreatedAt:  row.CreatedAt.Int64,  //创建时间
			UpdatedAt:  row.UpdatedAt.Int64,  //更新时间
		})
	}
	return
}

//获取单行数据
func (m *SgBlogsModel) getRow(sql string, params ...interface{}) (rowResult *SgBlogs, err error) {
	query := m.DB.QueryRow(sql, params...)
	if err != nil {
		return
	}
	row := SgBlogsNull{}
	err = query.Scan(
		&row.Id,         //
		&row.UserId,     //用户ID,与users表相关
		&row.Title,      //标题
		&row.Content,    //
		&row.Type,       //文章类型,1为原创,2为转载
		&row.IsPush,     //是否推送，0未，1是
		&row.Status,     //文章状态:1发布,2草稿
		&row.FirstTagId, //首标签ID
		&row.ReadCount,  //阅读次数
		&row.StarCount,  //点赞次数
		&row.CreatedAt,  //创建时间
		&row.UpdatedAt,  //更新时间
	)
	if nil != err {
		return
	}
	rowResult = &SgBlogs{
		Id:         row.Id.Int64,         //
		UserId:     row.UserId.Int64,     //用户ID,与users表相关
		Title:      row.Title.String,     //标题
		Content:    row.Content.String,   //
		Type:       row.Type.Int64,       //文章类型,1为原创,2为转载
		IsPush:     row.IsPush.Int64,     //是否推送，0未，1是
		Status:     row.Status.Int64,     //文章状态:1发布,2草稿
		FirstTagId: row.FirstTagId.Int64, //首标签ID
		ReadCount:  row.ReadCount.Int64,  //阅读次数
		StarCount:  row.StarCount.Int64,  //点赞次数
		CreatedAt:  row.CreatedAt.Int64,  //创建时间
		UpdatedAt:  row.UpdatedAt.Int64,  //更新时间
	}

	return
}

//_更新数据
func (m *SgBlogsModel) Save(sqlTxt string, value ...interface{}) (b bool, err error) {
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
func (m *SgBlogsModel) Create(value *SgBlogs) (lastId int64, err error) {
	sqlText := "INSERT INTO " + TABLE_SG_BLOGS + " (user_id,title,content,type,is_push,status,first_tag_id,read_count,star_count) VALUES (?,?,?,?,?,?,?,?,?)"
	stmt, err := m.DB.Prepare(sqlText)
	if err != nil {
		return
	}
	defer stmt.Close()
	result, err := stmt.Exec(
		value.UserId,     //用户ID,与users表相关
		value.Title,      //标题
		value.Content,    //
		value.Type,       //文章类型,1为原创,2为转载
		value.IsPush,     //是否推送，0未，1是
		value.Status,     //文章状态:1发布,2草稿
		value.FirstTagId, //首标签ID
		value.ReadCount,  //阅读次数
		value.StarCount,  //点赞次数
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
func (m *SgBlogsModel) Update(value *SgBlogs) (b bool, err error) {
	sqlText := "UPDATE " + TABLE_SG_BLOGS + " SET user_id=?,title=?,content=?,type=?,is_push=?,status=?,first_tag_id=?,read_count=?,star_count=? WHERE id = ?"
	params := make([]interface{}, 0)
	params = append(params, value.UserId)
	params = append(params, value.Title)
	params = append(params, value.Content)
	params = append(params, value.Type)
	params = append(params, value.IsPush)
	params = append(params, value.Status)
	params = append(params, value.FirstTagId)
	params = append(params, value.ReadCount)
	params = append(params, value.StarCount)
	params = append(params, value.Id)

	return m.Save(sqlText, params...)
}

//查询多行数据
func (m *SgBlogsModel) Find(value *SgBlogs) (resList []*SgBlogs, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_BLOGS
	resList, err = m.getRows(sqlText)
	return
}

//获取单行数据
func (m *SgBlogsModel) First(value *SgBlogs) (result *SgBlogs, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_BLOGS + " LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取单行数据
func (m *SgBlogsModel) Last(value *SgBlogs) (result *SgBlogs, err error) {
	sqlText := "SELECT" + m.getColumns() + "FROM " + TABLE_SG_BLOGS + " ORDER BY ID DESC LIMIT 1"
	result, err = m.getRow(sqlText)
	if err != nil {
		return
	}
	return
}

//获取行数
func (m *SgBlogsModel) Count() (count int64, err error) {
	sqlText := "SELECT COUNT(*) FROM " + TABLE_SG_BLOGS
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
