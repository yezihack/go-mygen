## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#{Index}book)       |书表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 备注|
|:---- |:---- |:----|:---- |:--- |:---- |
|id  | int(11) unsigned | NO|  | true | |
|book_name  | varchar(64) | NO|  | false | 名称|
|book_author  | varchar(125) | NO|  | false | 作者|
|book_province  | varchar(25) | NO|  | false | 省|
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 |

