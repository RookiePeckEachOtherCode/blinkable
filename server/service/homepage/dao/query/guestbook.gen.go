// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"blinkable/server/service/homepage/model"
)

func newGuestbook(db *gorm.DB, opts ...gen.DOOption) guestbook {
	_guestbook := guestbook{}

	_guestbook.guestbookDo.UseDB(db, opts...)
	_guestbook.guestbookDo.UseModel(&model.Guestbook{})

	tableName := _guestbook.guestbookDo.TableName()
	_guestbook.ALL = field.NewAsterisk(tableName)
	_guestbook.ID = field.NewInt64(tableName, "id")
	_guestbook.CreateTime = field.NewTime(tableName, "create_time")
	_guestbook.Context = field.NewString(tableName, "context")
	_guestbook.FromuserID = field.NewInt64(tableName, "fromuser_id")
	_guestbook.UserID = field.NewInt64(tableName, "user_id")

	_guestbook.fillFieldMap()

	return _guestbook
}

type guestbook struct {
	guestbookDo

	ALL        field.Asterisk
	ID         field.Int64
	CreateTime field.Time
	Context    field.String
	FromuserID field.Int64
	UserID     field.Int64

	fieldMap map[string]field.Expr
}

func (g guestbook) Table(newTableName string) *guestbook {
	g.guestbookDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g guestbook) As(alias string) *guestbook {
	g.guestbookDo.DO = *(g.guestbookDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *guestbook) updateTableName(table string) *guestbook {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.CreateTime = field.NewTime(table, "create_time")
	g.Context = field.NewString(table, "context")
	g.FromuserID = field.NewInt64(table, "fromuser_id")
	g.UserID = field.NewInt64(table, "user_id")

	g.fillFieldMap()

	return g
}

func (g *guestbook) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *guestbook) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 5)
	g.fieldMap["id"] = g.ID
	g.fieldMap["create_time"] = g.CreateTime
	g.fieldMap["context"] = g.Context
	g.fieldMap["fromuser_id"] = g.FromuserID
	g.fieldMap["user_id"] = g.UserID
}

func (g guestbook) clone(db *gorm.DB) guestbook {
	g.guestbookDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g guestbook) replaceDB(db *gorm.DB) guestbook {
	g.guestbookDo.ReplaceDB(db)
	return g
}

type guestbookDo struct{ gen.DO }

type IGuestbookDo interface {
	gen.SubQuery
	Debug() IGuestbookDo
	WithContext(ctx context.Context) IGuestbookDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGuestbookDo
	WriteDB() IGuestbookDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGuestbookDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGuestbookDo
	Not(conds ...gen.Condition) IGuestbookDo
	Or(conds ...gen.Condition) IGuestbookDo
	Select(conds ...field.Expr) IGuestbookDo
	Where(conds ...gen.Condition) IGuestbookDo
	Order(conds ...field.Expr) IGuestbookDo
	Distinct(cols ...field.Expr) IGuestbookDo
	Omit(cols ...field.Expr) IGuestbookDo
	Join(table schema.Tabler, on ...field.Expr) IGuestbookDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGuestbookDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGuestbookDo
	Group(cols ...field.Expr) IGuestbookDo
	Having(conds ...gen.Condition) IGuestbookDo
	Limit(limit int) IGuestbookDo
	Offset(offset int) IGuestbookDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGuestbookDo
	Unscoped() IGuestbookDo
	Create(values ...*model.Guestbook) error
	CreateInBatches(values []*model.Guestbook, batchSize int) error
	Save(values ...*model.Guestbook) error
	First() (*model.Guestbook, error)
	Take() (*model.Guestbook, error)
	Last() (*model.Guestbook, error)
	Find() ([]*model.Guestbook, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Guestbook, err error)
	FindInBatches(result *[]*model.Guestbook, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Guestbook) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGuestbookDo
	Assign(attrs ...field.AssignExpr) IGuestbookDo
	Joins(fields ...field.RelationField) IGuestbookDo
	Preload(fields ...field.RelationField) IGuestbookDo
	FirstOrInit() (*model.Guestbook, error)
	FirstOrCreate() (*model.Guestbook, error)
	FindByPage(offset int, limit int) (result []*model.Guestbook, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGuestbookDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g guestbookDo) Debug() IGuestbookDo {
	return g.withDO(g.DO.Debug())
}

func (g guestbookDo) WithContext(ctx context.Context) IGuestbookDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g guestbookDo) ReadDB() IGuestbookDo {
	return g.Clauses(dbresolver.Read)
}

func (g guestbookDo) WriteDB() IGuestbookDo {
	return g.Clauses(dbresolver.Write)
}

func (g guestbookDo) Session(config *gorm.Session) IGuestbookDo {
	return g.withDO(g.DO.Session(config))
}

func (g guestbookDo) Clauses(conds ...clause.Expression) IGuestbookDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g guestbookDo) Returning(value interface{}, columns ...string) IGuestbookDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g guestbookDo) Not(conds ...gen.Condition) IGuestbookDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g guestbookDo) Or(conds ...gen.Condition) IGuestbookDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g guestbookDo) Select(conds ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g guestbookDo) Where(conds ...gen.Condition) IGuestbookDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g guestbookDo) Order(conds ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g guestbookDo) Distinct(cols ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g guestbookDo) Omit(cols ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g guestbookDo) Join(table schema.Tabler, on ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g guestbookDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g guestbookDo) RightJoin(table schema.Tabler, on ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g guestbookDo) Group(cols ...field.Expr) IGuestbookDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g guestbookDo) Having(conds ...gen.Condition) IGuestbookDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g guestbookDo) Limit(limit int) IGuestbookDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g guestbookDo) Offset(offset int) IGuestbookDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g guestbookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGuestbookDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g guestbookDo) Unscoped() IGuestbookDo {
	return g.withDO(g.DO.Unscoped())
}

func (g guestbookDo) Create(values ...*model.Guestbook) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g guestbookDo) CreateInBatches(values []*model.Guestbook, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g guestbookDo) Save(values ...*model.Guestbook) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g guestbookDo) First() (*model.Guestbook, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guestbook), nil
	}
}

func (g guestbookDo) Take() (*model.Guestbook, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guestbook), nil
	}
}

func (g guestbookDo) Last() (*model.Guestbook, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guestbook), nil
	}
}

func (g guestbookDo) Find() ([]*model.Guestbook, error) {
	result, err := g.DO.Find()
	return result.([]*model.Guestbook), err
}

func (g guestbookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Guestbook, err error) {
	buf := make([]*model.Guestbook, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g guestbookDo) FindInBatches(result *[]*model.Guestbook, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g guestbookDo) Attrs(attrs ...field.AssignExpr) IGuestbookDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g guestbookDo) Assign(attrs ...field.AssignExpr) IGuestbookDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g guestbookDo) Joins(fields ...field.RelationField) IGuestbookDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g guestbookDo) Preload(fields ...field.RelationField) IGuestbookDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g guestbookDo) FirstOrInit() (*model.Guestbook, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guestbook), nil
	}
}

func (g guestbookDo) FirstOrCreate() (*model.Guestbook, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Guestbook), nil
	}
}

func (g guestbookDo) FindByPage(offset int, limit int) (result []*model.Guestbook, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g guestbookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g guestbookDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g guestbookDo) Delete(models ...*model.Guestbook) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *guestbookDo) withDO(do gen.Dao) *guestbookDo {
	g.DO = *do.(*gen.DO)
	return g
}