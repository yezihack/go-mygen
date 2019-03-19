## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ tmp](#1tmp)       || |
| 2|[ activitiy_stat](#2activitiy_stat)       |活动统计表| |
| 3|[ book](#3book)       |书表| |
| 4|[ data_comment](#4data_comment)       |评论数据| |



### 1.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 2.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 3.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 4.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ tmp](#1tmp)       || |
| 2|[ activitiy_stat](#2activitiy_stat)       |活动统计表| |
| 3|[ book](#3book)       |书表| |
| 4|[ data_comment](#4data_comment)       |评论数据| |



### 1.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 2.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 3.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 4.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ tmp](#1tmp)       || |
| 2|[ activitiy_stat](#2activitiy_stat)       |活动统计表| |
| 3|[ book](#3book)       |书表| |
| 4|[ data_comment](#4data_comment)       |评论数据| |



### 1.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 2.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 3.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 4.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ tmp](#1tmp)       || |
| 2|[ activitiy_stat](#2activitiy_stat)       |活动统计表| |
| 3|[ book](#3book)       |书表| |
| 4|[ data_comment](#4data_comment)       |评论数据| |



### 1.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 2.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 3.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 4.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ data_comment](#1data_comment)       |评论数据| |
| 2|[ tmp](#2tmp)       || |
| 3|[ activitiy_stat](#3activitiy_stat)       |活动统计表| |
| 4|[ book](#4book)       |书表| |



### 1.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 2.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 3.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 4.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ book](#1book)       |书表| |
| 2|[ data_comment](#2data_comment)       |评论数据| |
| 3|[ tmp](#3tmp)       || |
| 4|[ activitiy_stat](#4activitiy_stat)       |活动统计表| |



### 1.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 2.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 3.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)

### 4.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
## 表总榄

|序列| 表名 | 说明 | 备注 |
|:---|:----|:----| :----|
| 1|[ activitiy_stat](#1activitiy_stat)       |活动统计表| |
| 2|[ book](#2book)       |书表| |
| 3|[ data_comment](#3data_comment)       |评论数据| |
| 4|[ tmp](#4tmp)       || |



### 1.activitiy_stat
> 活动统计表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|stat_type  | smallint(6) unsigned | NO| 0 | false | 统计类型, 1 打款| |
|stat_day  | datetime | NO| CURRENT_TIMESTAMP | false | 统计日期| |
|stat_data  | varchar(512) | NO|  | false | 统计数据| |
|stat_remark  | varchar(50) | NO|  | false | 备注| |
|identify  | int(11) unsigned | NO| 0 | false | 乐观锁| |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间| |

[TOP](#表总榄)

### 2.book
> 书表

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|book_name  | varchar(64) | NO|  | false | 名称| |
|book_author  | varchar(125) | NO|  | false | 作者| |
|book_province  | varchar(25) | NO|  | false | 省| |
|updated_at  | datetime | NO| CURRENT_TIMESTAMP | false | 更新时间 | |
|created_at  | datetime | NO| CURRENT_TIMESTAMP | false | 创建时间 | |

[TOP](#表总榄)

### 3.data_comment
> 评论数据

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned | NO|  | true | | |
|data  | text | NO|  | false | | |

[TOP](#表总榄)

### 4.tmp
> 

- 表字段说明

|字段 | 类型 | 是否为空| 默认值 | 是否是主键 | 说明 | 备注|
|:----|:---|:----|:----|:---|:----|:---|
|id  | int(11) unsigned zerofill | NO|  | true | | |
|t1  | tinyint(1) unsigned zerofill | NO|  | false | | |
|t2  | tinyint(2) unsigned zerofill | NO|  | false | | |
|t3  | tinyint(3) unsigned zerofill | NO|  | false | | |
|t4  | tinyint(4) unsigned zerofill | NO|  | false | | |
|t10  | tinyint(10) unsigned zerofill | NO|  | false | | |
|tt  | tinyint(4) | NO|  | false | | |

[TOP](#表总榄)
