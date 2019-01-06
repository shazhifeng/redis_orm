package redis_orm

import (
	"errors"
)

//todo: 多个独立的tag辩识和书写更方便些~ 需要加统一前缀，免得跟其他功能定义的tag冲突，然后又太长了，先不改
const (
	TagIdentifier = "redis_orm"
	//定义是否索引，索引名自动生成 e.g.fmt.Sprintf("%s%s_%s", KeyIndexPrefix, strings.ToLower(table.Name), strings.ToLower(columnName)),
	TagIndex = "index"
	//唯一索引 针对IndexType_IdMember有效，IndexType_IdScore的索引本来就是唯一的~
	//todo:唯一索引走hash更快
	TagUniqueIndex = "unique"
	/*
		要支持一种查询条件就得增加一个索引，用&连接联合索引中的字段
		组合索引 字符串则唯一！集合数据结构决定; 除非用int64,前4个字节和后4个字节，放Score
		1、id作为score, 可以组合但是member唯一，唯一查询可用
		此情况下的组合索引，直接按顺序拼接即可

		2、该字段类型必须是长整型，id作为member，同一个member只能有一个score，数值索引可用，可以查询范围
		此情况下的组合索引，仅仅支持两整型字段，左边32位 右边32位，支持范围查询的放左边
	*/
	TagCombinedindex = "combinedindex"
	//默认值
	TagDefaultValue = "dft"
	//是否主键
	TagPrimaryKey = "pk"
	//自增~ 暂只支持主键
	TagAutoIncrement = "autoincr"
	//备注名
	TagComment = "comment"
	//是否支持自动写创建时间
	TagCreatedAt = "created_at"
	//是否支持自动写更新时间
	TagUpdatedAt = "updated_at"

	//表key前缀+表名
	KeyTbPrefix = "tb:"
	//索引key前缀+表名+字段名
	KeyIndexPrefix = "ix:"
	//自增前缀+自增字段名
	KeyAutoIncrPrefix = "autoincr_last_"

	ScoreMax = "+inf"
	ScoreMin = "-inf"
)

var (
	ERR_UnKnowField           = errors.New("redis-orm-error:unknow field")
	ERR_UnKnowError           = errors.New("redis-orm-error:unknow error")
	ERR_NotSupportIndexField  = errors.New("redis-orm-error:not support this filed's index")
	Err_UnSupportedType       = errors.New("redis-orm-error:unsupported type")
	Err_UnSupportedTableModel = errors.New("redis-orm-error:unsupported table model")
	Err_FieldValueInvalid     = errors.New("redis-orm-error:field value invalid")
	Err_PrimaryKeyNotFound    = errors.New("redis-orm-error:primarykey not found")
	Err_PrimaryKeyTypeInvalid = errors.New("redis-orm-error:primarykey type invalid")

	Err_MoreThanOneIncrementColumn = errors.New("redis-orm-error:more than one increment column")
	Err_DataNotAvailable           = errors.New("redis-orm-error:data not exist")
	Err_DataHadAvailable           = errors.New("redis-orm-error:data had exist")
	Err_CombinedIndexColCountOver  = errors.New("redis-orm-error:combined index not support more than 2 column")
	Err_CombinedIndexTypeError     = errors.New("redis-orm-error:combined index not support this type of column")
	Err_NeedPointer                = errors.New("redis-orm-error:needs a pointer to a value")
	Err_NeedSlice                  = errors.New("redis-orm-error:value needs to be a slice")
	Err_NotSupportPointer2Pointer  = errors.New("redis-orm-error:pointer to pointer is not supported")
)
