## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ article](#1article)       |文章表| |
| 2|[ cellect](#2cellect)       |收藏文章表| |
| 3|[ users](#3users)       |用户表| |



### 1.article
> 文章表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|title  | varchar(50) | NO|  | false | 文章标题| |
|content  | text | YES|  | false | 文章内容| |
|read_count  | int(11) | NO| 0 | false | 阅读数量| |
|state  | tinyint(4) | NO| 1 | false | 状态,1使用中,0禁用| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.cellect
> 收藏文章表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|article_id  | int(11) | NO| 0 | false | 文章ID| |
|user_id  | int(11) | NO| 0 | false | 用户ID| |
|state  | tinyint(4) | NO| 1 | false | 状态,1公开,0私用| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 3.users
> 用户表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|username  | varchar(50) | NO|  | false | 用户名称| |
|email  | varchar(50) | NO|  | false | 用户邮箱| |
|password  | varchar(128) | NO|  | false | 用户密码| |
|openid  | char(32) | NO|  | false | 微信OPENID| |
|state  | tinyint(4) | NO| 1 | false | 用户状态,1使用中,0禁用| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
