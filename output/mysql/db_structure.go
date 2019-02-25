//数据库表内结构体信息
package mysql

import "database/sql" //
type Book struct {
	Id           int    `json:"id"`            //
	BookName     string `json:"book_name"`     //
	BookAuthor   string `json:"book_author"`   //
	BookProvince string `json:"book_province"` //
	CreatedAt    string `json:"created_at"`    //
	UpdatedAt    string `json:"updated_at"`    //

}
type nullBook struct {
	Id           sql.NullInt64  //
	BookName     sql.NullString //
	BookAuthor   sql.NullString //
	BookProvince sql.NullString //
	CreatedAt    sql.NullString //
	UpdatedAt    sql.NullString //

}
