// This file was generated by go-dao-code-gen,
// you can modify it to be more suitable.

package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/huandu/go-sqlbuilder"
)

const ( // svModelTableName specifies the table name.
	SvModelTableName = "sv_model"
	// MaxSvModelLimit specifies limit of insert and select operations.
	MaxSvModelLimit int = 1000
)

var (
	svModelAlias           SvModelAlias
	svModelFields          []string
	svModelUpdateFields    map[string]struct{} // Fields are allowed to be modified
	svModelOmitEmptyFields map[string]struct{} // Fields skipped when modifying with empty value
)

// SvModelDao specifies the dao object.
type SvModelDao struct {
	db          *sql.DB
	forceMaster bool
	*SvModelAlias
}

// SvModelAlias alias of fields in table sv_model.
type SvModelAlias struct {
	ID           string // id
	Name         string // name
	ThumbURL     string // thumb_url
	ModelURL     string // model_url
	Price        string // price
	Brand        string // brand
	UserEditable string
	CreatedAt    string // created_at
	UpdatedAt    string // updated_at
	DeletedAt    string // deleted_at
}

// SvModelEntity sv_model table mapping,
// Please manually remove the update tag in the field that is not allowed to be modified.
type SvModelEntity struct {
	ID           int64  `db:"id" `                         // 模型id
	Name         string `db:"name" fieldtag:"update"`      // 模型名称
	ThumbURL     string `db:"thumb_url" fieldtag:"update"` // 模型缩略图
	ModelURL     string `db:"model_url" fieldtag:"update"` // 模型源文件
	Price        int64  `db:"price" fieldtag:"update"`     // 价格
	Brand        string `db:"brand" fieldtag:"update"`     // 手机品牌 eg:{"iphone":["iphone7"]}
	UserEditable int64  `db:"user_editable" fieldtag:"update"`
	CreatedAt    int64  `db:"created_at" fieldtag:"update"`
	UpdatedAt    int64  `db:"updated_at" fieldtag:"update"`
	DeletedAt    int64  `db:"deleted_at" fieldtag:"update"`
}

func init() {
	InitTableAlias(SvModelEntity{}, &svModelAlias)
	InitTableFields(SvModelEntity{}, &svModelFields)
	svModelUpdateFields = make(map[string]struct{}, len(svModelFields))
	InitTableUpdateFields(SvModelEntity{}, svModelUpdateFields)
	svModelOmitEmptyFields = make(map[string]struct{})
	InitTableOmitEmptyFields(SvModelEntity{}, svModelOmitEmptyFields)
}

// NewSvModelDao create a new table object.
func NewSvModelDao() *SvModelDao {
	return &SvModelDao{
		db:           globalDB,
		SvModelAlias: &svModelAlias,
	}
}

// Insert insert one data record.
func (d *SvModelDao) Insert(ctx context.Context, values map[string]interface{}) (lastInsertID int64, err error) {
	if len(values) == 0 {
		return lastInsertID, errors.New("param values cannot be empty")
	}
	cols := make([]string, 0, len(values))
	vals := make([]interface{}, 0, len(values))
	for _, field := range svModelFields {
		if val, ok := values[field]; ok {
			cols = append(cols, field)
			vals = append(vals, val)
		}
	}
	if len(cols) == 0 {
		return lastInsertID, errors.New("no valid field data found")
	}

	// Here you can add some default values for the fields.
	// example:
	// curTime := time.Now().Unix()
	// if _, ok := values["create_time"]; !ok {
	// cols = append(cols, "create_time")
	// vals = append(vals, curTime)
	// }
	// if _, ok := values["update_time"]; !ok {
	// cols = append(cols, "update_time")
	// vals = append(vals, curTime)
	// }
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto(SvModelTableName)
	ib.Cols(cols...)
	ib.Values(vals...)
	sql, args := ib.Build()
	result, err := d.db.Exec(sql, args...)
	if err != nil {
		// log error
		return lastInsertID, err
	}
	return result.LastInsertId()
}

// InsertMany insert multiple data records.
func (d *SvModelDao) InsertMany(ctx context.Context, valueList []map[string]interface{}) (err error) {
	if len(valueList) == 0 {
		return
	}
	if len(valueList) > MaxSvModelLimit {
		return fmt.Errorf("received %d data, exceeding the maximum %d limit", len(valueList), MaxSvModelLimit)
	}
	var cols []string
	var valsList [][]interface{}
	for index, values := range valueList {
		if len(values) == 0 {
			return errors.New("param values cannot be empty")
		}
		vals := make([]interface{}, 0, len(values))
		for _, field := range svModelFields {
			if val, ok := values[field]; ok {
				if index == 0 {
					cols = append(cols, field)
				}
				vals = append(vals, val)
			}
		}
		valsList = append(valsList, vals)
	}
	if len(cols) == 0 {
		return errors.New("no valid field data found")
	}
	// Add default columns.
	// var hasAddCreate, hasAddUpdate bool
	// if _, ok := valueList[0]["create_time"]; !ok {
	//     cols = append(cols, "create_time")
	//     hasAddCreate = true
	// }
	// if _, ok := valueList[0]["update_time"]; !ok {
	//     cols = append(cols, "update_time")
	//     hasAddUpdate = true
	// }
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto(SvModelTableName)
	ib.Cols(cols...)
	// curTime := time.Now().Unix()
	for _, vals := range valsList {
		// Add default values.
		// if hasAddCreate {
		//     vals = append(vals, curTime)
		// }
		// if hasAddUpdate {
		//     vals = append(vals, curTime)
		// }
		ib.Values(vals...)
	}
	sql, args := ib.Build()
	_, err = d.db.Exec(sql, args...)
	if err != nil {
		return err
	}
	return nil
}

// Get one data record that meets the query criteria
func (d *SvModelDao) Get(ctx context.Context, conds ...SvModelCond) (svModelEntity *SvModelEntity, err error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(svModelFields...)
	sb.From(SvModelTableName)
	o := NewSvModelConds(conds...)
	sqlArgs := BuildSvModelConds(&sb.Cond, &o)
	sb.Where(sqlArgs...)
	sb.OrderBy("id").Desc()
	sb.Limit(1)
	sql, args := sb.Build()
	if d.forceMaster {
		sql = ForceMasterIdenti + sql
	}
	rows, err := d.db.Query(sql, args...)
	if err != nil {
		return
	}
	defer rows.Close()
	if rows.Next() {
		svModelEntity = &SvModelEntity{}
		svModelStruct := sqlbuilder.NewStruct(new(SvModelEntity))
		err = rows.Scan(svModelStruct.Addr(svModelEntity)...)
		if err != nil {
			return
		}
	}
	return
}

// Count total of data record that meets the query criteria.
func (d *SvModelDao) Count(ctx context.Context, conds ...SvModelCond) (total int, err error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(*)")
	sb.From(SvModelTableName)
	o := NewSvModelConds(conds...)
	sqlArgs := BuildSvModelConds(&sb.Cond, &o)
	sb.Where(sqlArgs...)
	sql, args := sb.Build()
	if d.forceMaster {
		sql = ForceMasterIdenti + sql
	}
	rows, err := d.db.Query(sql, args...)
	if err != nil {
		// log error
		return
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&total)
		if err != nil {
			return
		}
	}
	return
}

// List multiple data records that meets the query criteria.
func (d *SvModelDao) List(ctx context.Context, limit, offset int, conds ...SvModelCond) (svModelList []*SvModelEntity, err error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(svModelFields...)
	sb.From(SvModelTableName)
	o := NewSvModelConds(conds...)
	sqlArgs := BuildSvModelConds(&sb.Cond, &o)
	sb.Where(sqlArgs...)
	sb.OrderBy("id").Desc()
	if limit <= 0 || limit > MaxSvModelLimit {
		sb.Limit(MaxSvModelLimit)
	} else {
		sb.Limit(limit)
	}
	if offset >= 0 {
		sb.Offset(offset)
	}
	sql, args := sb.Build()
	if d.forceMaster {
		sql = ForceMasterIdenti + sql
	}
	rows, err := d.db.Query(sql, args...)
	if err != nil {
		// log error
		return
	}
	defer rows.Close()
	svModelStruct := sqlbuilder.NewStruct(new(SvModelEntity))
	for rows.Next() {
		svModelEntity := &SvModelEntity{}
		err = rows.Scan(svModelStruct.Addr(svModelEntity)...)
		if err != nil {
			return
		}
		svModelList = append(svModelList, svModelEntity)
	}
	return
}

// All multiple data records that meets the query criteria.
func (d *SvModelDao) All(ctx context.Context, limit int, conds ...SvModelCond) (svModelList []*SvModelEntity, err error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(svModelFields...)
	sb.From(SvModelTableName)
	o := NewSvModelConds(conds...)
	sqlArgs := BuildSvModelConds(&sb.Cond, &o)
	sb.Where(sqlArgs...)
	sb.OrderBy("id").Desc()
	if limit > 0 {
		sb.Limit(limit)
	}
	sql, args := sb.Build()
	if d.forceMaster {
		sql = ForceMasterIdenti + sql
	}
	rows, err := d.db.Query(sql, args...)
	if err != nil {
		// log error
		return
	}
	defer rows.Close()
	svModelStruct := sqlbuilder.NewStruct(new(SvModelEntity))
	for rows.Next() {
		svModelEntity := &SvModelEntity{}
		err = rows.Scan(svModelStruct.Addr(svModelEntity)...)
		if err != nil {
			return
		}
		svModelList = append(svModelList, svModelEntity)
	}
	return
}

// Update is modify the records that meet the query criteria.
func (d *SvModelDao) Update(ctx context.Context, values map[string]interface{}, conds ...SvModelCond) (total int64, err error) {
	if len(conds) == 0 {
		return
	}
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update(SvModelTableName)
	fieldList := make([]string, 0, len(values))
	for index, val := range values {
		if _, ok := svModelUpdateFields[index]; !ok {
			continue
		}
		if _, ok := svModelOmitEmptyFields[index]; ok {
			if reflect.ValueOf(val).IsZero() {
				continue
			}
		}
		fieldList = append(fieldList, ub.Assign(index, val))
	}
	values = nil
	if len(fieldList) == 0 {
		return
	}
	// deal update time
	// fieldList = append(fieldList, ub.Assign("modified", time.Now()))
	ub.Set(fieldList...)
	o := NewSvModelConds(conds...)
	sqlArgs := BuildSvModelConds(&ub.Cond, &o)
	ub.Where(sqlArgs...)
	sql, args := ub.Build()
	result, err := d.db.Exec(sql, args...)
	if err != nil {
		// log error
		return
	}
	return result.RowsAffected()
}

// Delete the records that meet the query criteria.
func (d *SvModelDao) Delete(ctx context.Context, conds ...SvModelCond) (total int64, err error) {
	if len(conds) == 0 {
		return
	}
	db := sqlbuilder.NewDeleteBuilder()
	db.DeleteFrom(SvModelTableName)
	o := NewSvModelConds(conds...)
	sqlArgs := BuildSvModelConds(&db.Cond, &o)
	db.Where(sqlArgs...)
	sql, args := db.Build()
	result, err := d.db.Exec(sql, args...)
	if err != nil {
		// log error
		return
	}
	return result.RowsAffected()
}

// ForceMaster add the master identity for operations of the current object.
func (d *SvModelDao) ForceMaster() {
	d.forceMaster = true
}

// DisableForceMaster remove the master identity for operations of the current object.
func (d *SvModelDao) DisableForceMaster() {
	d.forceMaster = false
}

func (d *SvModelDao) UseConn(db *sql.DB) {
	d.db = db
}

func (d *SvModelDao) CloneConn() (db *sql.DB) {
	return d.db
}
