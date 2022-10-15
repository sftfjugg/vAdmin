package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
	orm "vAdmin/database"
	"vAdmin/tools"
)

type BaseModel struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

type WhereParam struct {
	Field   string
	Tag     string
	Prepare interface{}
}

type QueryParam struct {
	Fields     string
	Offset     int
	Limit      int
	Order      string
	Where      []WhereParam
}

func Create(model interface{}) bool {
	db := orm.Eloquent.Create(model)
	if err := db.Error; err != nil {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

func GetByPk(model interface{}, id interface{}) bool {
	db := orm.Eloquent.Model(model)
	db.First(model, id)
	if err := db.Error; err != nil && !db.RecordNotFound() {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

func UpdateByPk(model interface{}) bool {
	db := orm.Eloquent.Model(model)
	db = db.Updates(model)
	if err := db.Error; err != nil {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

func GetOne(model interface{}, query QueryParam) bool {
	db := orm.Eloquent.Model(model)
	if query.Fields != "" {
		db = db.Select(query.Fields)
	}
	db = parseWhereParam(db, query.Where)
	db = db.First(model)
	if err := db.Error; err != nil && !db.RecordNotFound() {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

func GetMulti(model interface{}, query QueryParam) bool {
	db := orm.Eloquent.Offset(query.Offset)
	if query.Limit > 0 {
		db = db.Limit(query.Limit)
	}
	if query.Fields != "" {
		db = db.Select(query.Fields)
	}
	if query.Order != "" {
		db = db.Order(query.Order)
	}
	db = parseWhereParam(db, query.Where)
	db.Find(model)
	if err := db.Error; err != nil {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

func Count(model interface{}, count *int, query QueryParam) bool {
	db := orm.Eloquent.Model(model)
	db = parseWhereParam(db, query.Where)
	db = db.Count(count)
	if err := db.Error; err != nil {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

func Update(model interface{}, data interface{}, query QueryParam) bool {
	db := orm.Eloquent.Model(model)
	db = parseWhereParam(db, query.Where)
	db = db.Updates(data)
	if err := db.Error; err != nil {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

func Delete(model interface{}, query QueryParam) bool {
	if len(query.Where) == 0 {
		//s := fmt.Sprintf("mysql query error: delete failed, where conditions cannot be empty")
		return false
	}
	db := orm.Eloquent.Model(model)
	db = parseWhereParam(db, query.Where)
	db = db.Delete(model)
	if err := db.Error; err != nil {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

// 表清空Delete
func DeleteByModel(model interface{}) (count int64 ,err error){
	db := orm.Eloquent.Delete(model)
	err=db.Error
	if err!=nil {
		return
	}
	count=db.RowsAffected
	return
}

// 根据where条件Delete
func DeleteByWhere(model,where interface{}) (count int64 ,err error){
	db := orm.Eloquent.Where(where).Delete(model)
	err=db.Error
	if err!=nil {
		return
	}
	count=db.RowsAffected
	return
}

// 单项Delete
func DeleteByID(model interface{},id uint64) (count int64 ,err error){
	db := orm.Eloquent.Where("id=?", id).Delete(model)
	err=db.Error
	if err!=nil {
		return
	}
	count=db.RowsAffected
	return
}

// 批量Delete
func DeleteByIDS(model interface{},ids []uint64) (count int64 ,err error){
	db:=orm.Eloquent.Where("id in (?)", ids).Delete(model)
	err=db.Error
	if err!=nil {
		return
	}
	count=db.RowsAffected
	return
}

func parseWhereParam(db *gorm.DB, where []WhereParam) *gorm.DB {
	if len(where) == 0 {
		return db
	}
	var (
		plain []string
		prepare []interface{}
	)
	for _, w := range where {
		tag := w.Tag
		if tag == "" {
			tag = "="
		}
		var plainFmt string
		switch tag {
		case "IN":
			plainFmt = fmt.Sprintf("%s IN (?)", w.Field)
		default:
			plainFmt = fmt.Sprintf("%s %s ?", w.Field, tag)
		}
		plain = append(plain, plainFmt)
		prepare = append(prepare, w.Prepare)
	}
	return db.Where(strings.Join(plain, " AND "), prepare...)
}

func DeleteByPk(model interface{}) bool {
	db := orm.Eloquent.Model(model)
	db.Delete(model)
	if err := db.Error; err != nil {
		s := fmt.Sprintf("mysql query error: %s sql[%v]", err.Error(), db.QueryExpr())
		tools.HasError(err, s, -1)
		return false
	}
	return true
}

// 通过ID查找First
func FirstByID(out interface{},id int)(notFound bool,err error){
	err= orm.Eloquent.First(out, id).Error
	if err!=nil {
		notFound=gorm.IsRecordNotFoundError(err)
	}
	return
}

// 通过条件查找First
func First(where interface{},out interface{})(notFound bool,err error){
	err= orm.Eloquent.Where(where).First(out).Error
	if err!=nil {
		notFound=gorm.IsRecordNotFoundError(err)
	}
	return
}

// 条件查询Find
func Find(where interface{},out interface{},orders ...string)error{
	db:=orm.Eloquent.Where(where)
	if len(orders)>0 {
		for _,order:=range orders {
			db=db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan
func Scan(model,where interface{},out interface{})(notFound bool,err error){
	err= orm.Eloquent.Model(model).Where(where).Scan(out).Error
	if err!=nil {
		notFound=gorm.IsRecordNotFoundError(err)
	}
	return
}

// ScanList
func ScanList(model,where interface{},out interface{},orders ...string)error{
	db:=orm.Eloquent.Model(model).Where(where)
	if len(orders)>0 {
		for _,order:=range orders {
			db=db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// 分页条件
type PageWhereOrder struct {
	Order string
	Where string
	Value []interface{}
}
// GetPage
func GetPageCom(model,where interface{},out interface{},pageIndex, pageSize uint64,totalCount *uint64,whereOrder ...PageWhereOrder) error{
	db:=orm.Eloquent.Model(model).Where(where)
	if len(whereOrder)>0 {
		for _,wo:=range whereOrder {
			if wo.Order !="" {
				db=db.Order(wo.Order)
			}
			if wo.Where !="" {
				db=db.Where(wo.Where,wo.Value...)
			}
		}
	}
	err:=db.Count(totalCount).Error
	if err!=nil{
		return err
	}
	if *totalCount==0{
		return nil
	}
	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// PluckList
func PluckList(model,where interface{},out interface{},fieldName string)error{
	return orm.Eloquent.Model(model).Where(where).Pluck(fieldName, out).Error
}