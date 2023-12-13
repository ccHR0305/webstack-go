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

	"github.com/ccHR0305/webstack-go/internal/repository/mysql/model"
)

func newAuthorizedAPI(db *gorm.DB, opts ...gen.DOOption) authorizedAPI {
	_authorizedAPI := authorizedAPI{}

	_authorizedAPI.authorizedAPIDo.UseDB(db, opts...)
	_authorizedAPI.authorizedAPIDo.UseModel(&model.AuthorizedAPI{})

	tableName := _authorizedAPI.authorizedAPIDo.TableName()
	_authorizedAPI.ALL = field.NewAsterisk(tableName)
	_authorizedAPI.ID = field.NewInt64(tableName, "id")
	_authorizedAPI.BusinessKey = field.NewString(tableName, "business_key")
	_authorizedAPI.Method = field.NewString(tableName, "method")
	_authorizedAPI.API = field.NewString(tableName, "api")
	_authorizedAPI.IsDeleted = field.NewInt64(tableName, "is_deleted")
	_authorizedAPI.CreatedAt = field.NewTime(tableName, "created_at")
	_authorizedAPI.CreatedUser = field.NewString(tableName, "created_user")
	_authorizedAPI.UpdatedAt = field.NewTime(tableName, "updated_at")
	_authorizedAPI.UpdatedUser = field.NewString(tableName, "updated_user")

	_authorizedAPI.fillFieldMap()

	return _authorizedAPI
}

type authorizedAPI struct {
	authorizedAPIDo

	ALL         field.Asterisk
	ID          field.Int64  // 主键
	BusinessKey field.String // 调用方key
	Method      field.String // 请求方式
	API         field.String // 请求地址
	IsDeleted   field.Int64  // 是否删除 1:是  -1:否
	CreatedAt   field.Time   // 创建时间
	CreatedUser field.String // 创建人
	UpdatedAt   field.Time   // 更新时间
	UpdatedUser field.String // 更新人

	fieldMap map[string]field.Expr
}

func (a authorizedAPI) Table(newTableName string) *authorizedAPI {
	a.authorizedAPIDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authorizedAPI) As(alias string) *authorizedAPI {
	a.authorizedAPIDo.DO = *(a.authorizedAPIDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authorizedAPI) updateTableName(table string) *authorizedAPI {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.BusinessKey = field.NewString(table, "business_key")
	a.Method = field.NewString(table, "method")
	a.API = field.NewString(table, "api")
	a.IsDeleted = field.NewInt64(table, "is_deleted")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.CreatedUser = field.NewString(table, "created_user")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.UpdatedUser = field.NewString(table, "updated_user")

	a.fillFieldMap()

	return a
}

func (a *authorizedAPI) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authorizedAPI) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["business_key"] = a.BusinessKey
	a.fieldMap["method"] = a.Method
	a.fieldMap["api"] = a.API
	a.fieldMap["is_deleted"] = a.IsDeleted
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["created_user"] = a.CreatedUser
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["updated_user"] = a.UpdatedUser
}

func (a authorizedAPI) clone(db *gorm.DB) authorizedAPI {
	a.authorizedAPIDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a authorizedAPI) replaceDB(db *gorm.DB) authorizedAPI {
	a.authorizedAPIDo.ReplaceDB(db)
	return a
}

type authorizedAPIDo struct{ gen.DO }

type IAuthorizedAPIDo interface {
	gen.SubQuery
	Debug() IAuthorizedAPIDo
	WithContext(ctx context.Context) IAuthorizedAPIDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAuthorizedAPIDo
	WriteDB() IAuthorizedAPIDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAuthorizedAPIDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthorizedAPIDo
	Not(conds ...gen.Condition) IAuthorizedAPIDo
	Or(conds ...gen.Condition) IAuthorizedAPIDo
	Select(conds ...field.Expr) IAuthorizedAPIDo
	Where(conds ...gen.Condition) IAuthorizedAPIDo
	Order(conds ...field.Expr) IAuthorizedAPIDo
	Distinct(cols ...field.Expr) IAuthorizedAPIDo
	Omit(cols ...field.Expr) IAuthorizedAPIDo
	Join(table schema.Tabler, on ...field.Expr) IAuthorizedAPIDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthorizedAPIDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthorizedAPIDo
	Group(cols ...field.Expr) IAuthorizedAPIDo
	Having(conds ...gen.Condition) IAuthorizedAPIDo
	Limit(limit int) IAuthorizedAPIDo
	Offset(offset int) IAuthorizedAPIDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthorizedAPIDo
	Unscoped() IAuthorizedAPIDo
	Create(values ...*model.AuthorizedAPI) error
	CreateInBatches(values []*model.AuthorizedAPI, batchSize int) error
	Save(values ...*model.AuthorizedAPI) error
	First() (*model.AuthorizedAPI, error)
	Take() (*model.AuthorizedAPI, error)
	Last() (*model.AuthorizedAPI, error)
	Find() ([]*model.AuthorizedAPI, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuthorizedAPI, err error)
	FindInBatches(result *[]*model.AuthorizedAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AuthorizedAPI) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthorizedAPIDo
	Assign(attrs ...field.AssignExpr) IAuthorizedAPIDo
	Joins(fields ...field.RelationField) IAuthorizedAPIDo
	Preload(fields ...field.RelationField) IAuthorizedAPIDo
	FirstOrInit() (*model.AuthorizedAPI, error)
	FirstOrCreate() (*model.AuthorizedAPI, error)
	FindByPage(offset int, limit int) (result []*model.AuthorizedAPI, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthorizedAPIDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authorizedAPIDo) Debug() IAuthorizedAPIDo {
	return a.withDO(a.DO.Debug())
}

func (a authorizedAPIDo) WithContext(ctx context.Context) IAuthorizedAPIDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authorizedAPIDo) ReadDB() IAuthorizedAPIDo {
	return a.Clauses(dbresolver.Read)
}

func (a authorizedAPIDo) WriteDB() IAuthorizedAPIDo {
	return a.Clauses(dbresolver.Write)
}

func (a authorizedAPIDo) Session(config *gorm.Session) IAuthorizedAPIDo {
	return a.withDO(a.DO.Session(config))
}

func (a authorizedAPIDo) Clauses(conds ...clause.Expression) IAuthorizedAPIDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authorizedAPIDo) Returning(value interface{}, columns ...string) IAuthorizedAPIDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authorizedAPIDo) Not(conds ...gen.Condition) IAuthorizedAPIDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authorizedAPIDo) Or(conds ...gen.Condition) IAuthorizedAPIDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authorizedAPIDo) Select(conds ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authorizedAPIDo) Where(conds ...gen.Condition) IAuthorizedAPIDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authorizedAPIDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAuthorizedAPIDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a authorizedAPIDo) Order(conds ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authorizedAPIDo) Distinct(cols ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authorizedAPIDo) Omit(cols ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authorizedAPIDo) Join(table schema.Tabler, on ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authorizedAPIDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authorizedAPIDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authorizedAPIDo) Group(cols ...field.Expr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authorizedAPIDo) Having(conds ...gen.Condition) IAuthorizedAPIDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authorizedAPIDo) Limit(limit int) IAuthorizedAPIDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authorizedAPIDo) Offset(offset int) IAuthorizedAPIDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authorizedAPIDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthorizedAPIDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authorizedAPIDo) Unscoped() IAuthorizedAPIDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authorizedAPIDo) Create(values ...*model.AuthorizedAPI) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authorizedAPIDo) CreateInBatches(values []*model.AuthorizedAPI, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authorizedAPIDo) Save(values ...*model.AuthorizedAPI) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authorizedAPIDo) First() (*model.AuthorizedAPI, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizedAPI), nil
	}
}

func (a authorizedAPIDo) Take() (*model.AuthorizedAPI, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizedAPI), nil
	}
}

func (a authorizedAPIDo) Last() (*model.AuthorizedAPI, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizedAPI), nil
	}
}

func (a authorizedAPIDo) Find() ([]*model.AuthorizedAPI, error) {
	result, err := a.DO.Find()
	return result.([]*model.AuthorizedAPI), err
}

func (a authorizedAPIDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuthorizedAPI, err error) {
	buf := make([]*model.AuthorizedAPI, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authorizedAPIDo) FindInBatches(result *[]*model.AuthorizedAPI, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authorizedAPIDo) Attrs(attrs ...field.AssignExpr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authorizedAPIDo) Assign(attrs ...field.AssignExpr) IAuthorizedAPIDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authorizedAPIDo) Joins(fields ...field.RelationField) IAuthorizedAPIDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authorizedAPIDo) Preload(fields ...field.RelationField) IAuthorizedAPIDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authorizedAPIDo) FirstOrInit() (*model.AuthorizedAPI, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizedAPI), nil
	}
}

func (a authorizedAPIDo) FirstOrCreate() (*model.AuthorizedAPI, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthorizedAPI), nil
	}
}

func (a authorizedAPIDo) FindByPage(offset int, limit int) (result []*model.AuthorizedAPI, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a authorizedAPIDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authorizedAPIDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a authorizedAPIDo) Delete(models ...*model.AuthorizedAPI) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *authorizedAPIDo) withDO(do gen.Dao) *authorizedAPIDo {
	a.DO = *do.(*gen.DO)
	return a
}
