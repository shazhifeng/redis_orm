package test

import (
	"testing"
	"time"

	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/mkideal/log"

	"github.com/weikaishio/redis_orm"
	"github.com/weikaishio/redis_orm/test/models"
)

var (
	engine *redis_orm.Engine
)

//test and log?
func init() {
	options := redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       1,
	}

	redisClient := redis.NewClient(&options)

	engine = redis_orm.NewEngine(redisClient)
	engine.IsShowLog(true)

	log.SetLevelFromString("TRACE")
}

//func TestEngine_GetTable(t *testing.T) {
//	faq := &models.Faq{
//		Title:   "为啥",
//		Content: "我也不知道",
//		Hearts:  20,
//	}
//	val := reflect.ValueOf(faq)
//	table, err := engine.GetTable(val, reflect.Indirect(val))
//	t.Logf("table:%v,err:%v", table, err)
//}

//func TestEngine_GetDefaultValue(t *testing.T) {
//	faq := &models.Faq{}
//	err := engine.GetDefaultValue(faq)
//	bys, _ := json.Marshal(faq)
//	t.Logf("GetDefaultValue faq:%v,err:%v", string(bys), err)
//}

func TestEngine_Insert(t *testing.T) {
	faq := &models.Faq{
		Title:  "index3",
		Unique: time.Now().Unix(),
	}
	err := engine.Insert(faq)
	bys, _ := json.Marshal(faq)
	t.Logf("Insert faq:%v,err:%v", string(bys), err)
}

func TestEngine_Get(t *testing.T) {
	faq := &models.Faq{
		Id: 2,
	}
	has, err := engine.Get(faq)
	bys, _ := json.Marshal(faq)
	t.Logf("faq:%v,has:%v,err:%v", string(bys), has, err)
}

//func TestEngine_GetByCombinedIndex(t *testing.T) {
//	faq := &models.Faq{}
//	has, err := engine.GetByCondition(faq, redis_orm.NewSearchCondition(
//		redis_orm.IndexType_IdScore,
//		"1&index",
//		"1&index",
//		"Type",
//		"Title",
//	))
//
//	bys, _ := json.Marshal(faq)
//	t.Logf("faq:%v,has:%v,err:%v", string(bys), has, err)
//}

//func TestEngine_Find(t *testing.T) {
//	var faqAry []*models.Faq
//	count, err := engine.Find(0, 3, redis_orm.NewSearchCondition(
//		redis_orm.IndexType_IdScore,
//		redis_orm.ScoreMin,
//		redis_orm.ScoreMax,
//		"Id",
//	), &faqAry)
//	bys, _ := json.Marshal(faqAry)
//	t.Logf("faqAry:%v,count:%v,err:%v", string(bys), count, err)
//}

//func TestEngine_Update(t *testing.T) {
//	faq := &models.Faq{
//		Id:    1,
//		Title: "test51",
//	}
//	err := engine.Update(faq, "Title")
//	t.Logf("TestEngine_Update err:%v", err)
//}

//func TestEngine_UpdateMulti(t *testing.T) {
//	faq := &models.Faq{
//		Title: "test51",
//	}
//	affectedRow, err := engine.UpdateMulti(faq, redis_orm.NewSearchCondition(
//		redis_orm.IndexType_IdScore,
//		"1",
//		"1",
//		"Id",
//	),
//		"Title")
//	t.Logf("TestEneine_UpdateMulti affectedRow:%d,err:%v", affectedRow, err)
//}

//func TestEngine_Delete(t *testing.T) {
//	faq := &models.Faq{
//		Id: 1,
//	}
//	err := engine.Delete(faq)
//	t.Logf("Delete faq:%v, err:%v", faq, err)
//}
//
//func TestEngine_TableDrop(t *testing.T){
//	faq:=&models.Faq{}
//	err:=engine.TableDrop(faq)
//	t.Logf("TestEngine_TableDrop err:%v", err)
//}
func TestEngine_DeleteMulti(t *testing.T){
	faq := &models.Faq{
		Title: "test51",
	}
	affectedRow, err := engine.DeleteMulti(faq, redis_orm.NewSearchCondition(
		redis_orm.IndexType_IdScore,
		"1",
		"4",
		"Id",
	),
		"Title")
	t.Logf("TestEngine_DeleteMulti affectedRow:%d,err:%v", affectedRow, err)
}