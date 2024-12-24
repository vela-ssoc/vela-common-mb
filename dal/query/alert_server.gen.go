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

	"github.com/vela-ssoc/vela-common-mb/dal/model"
)

func newAlertServer(db *gorm.DB, opts ...gen.DOOption) alertServer {
	_alertServer := alertServer{}

	_alertServer.alertServerDo.UseDB(db, opts...)
	_alertServer.alertServerDo.UseModel(&model.AlertServer{})

	tableName := _alertServer.alertServerDo.TableName()
	_alertServer.ALL = field.NewAsterisk(tableName)
	_alertServer.ID = field.NewInt64(tableName, "id")
	_alertServer.Enabled = field.NewBool(tableName, "enabled")
	_alertServer.Mode = field.NewString(tableName, "mode")
	_alertServer.Name = field.NewString(tableName, "name")
	_alertServer.URL = field.NewString(tableName, "url")
	_alertServer.Token = field.NewString(tableName, "token")
	_alertServer.Account = field.NewString(tableName, "account")
	_alertServer.CreatedAt = field.NewTime(tableName, "updated_at")
	_alertServer.UpdatedAt = field.NewTime(tableName, "created_at")

	_alertServer.fillFieldMap()

	return _alertServer
}

type alertServer struct {
	alertServerDo alertServerDo

	ALL       field.Asterisk
	ID        field.Int64
	Enabled   field.Bool
	Mode      field.String
	Name      field.String
	URL       field.String
	Token     field.String
	Account   field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (a alertServer) Table(newTableName string) *alertServer {
	a.alertServerDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a alertServer) As(alias string) *alertServer {
	a.alertServerDo.DO = *(a.alertServerDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *alertServer) updateTableName(table string) *alertServer {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Enabled = field.NewBool(table, "enabled")
	a.Mode = field.NewString(table, "mode")
	a.Name = field.NewString(table, "name")
	a.URL = field.NewString(table, "url")
	a.Token = field.NewString(table, "token")
	a.Account = field.NewString(table, "account")
	a.CreatedAt = field.NewTime(table, "updated_at")
	a.UpdatedAt = field.NewTime(table, "created_at")

	a.fillFieldMap()

	return a
}

func (a *alertServer) WithContext(ctx context.Context) *alertServerDo {
	return a.alertServerDo.WithContext(ctx)
}

func (a alertServer) TableName() string { return a.alertServerDo.TableName() }

func (a alertServer) Alias() string { return a.alertServerDo.Alias() }

func (a alertServer) Columns(cols ...field.Expr) gen.Columns { return a.alertServerDo.Columns(cols...) }

func (a *alertServer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *alertServer) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["enabled"] = a.Enabled
	a.fieldMap["mode"] = a.Mode
	a.fieldMap["name"] = a.Name
	a.fieldMap["url"] = a.URL
	a.fieldMap["token"] = a.Token
	a.fieldMap["account"] = a.Account
	a.fieldMap["updated_at"] = a.CreatedAt
	a.fieldMap["created_at"] = a.UpdatedAt
}

func (a alertServer) clone(db *gorm.DB) alertServer {
	a.alertServerDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a alertServer) replaceDB(db *gorm.DB) alertServer {
	a.alertServerDo.ReplaceDB(db)
	return a
}

type alertServerDo struct{ gen.DO }

func (a alertServerDo) Debug() *alertServerDo {
	return a.withDO(a.DO.Debug())
}

func (a alertServerDo) WithContext(ctx context.Context) *alertServerDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a alertServerDo) ReadDB() *alertServerDo {
	return a.Clauses(dbresolver.Read)
}

func (a alertServerDo) WriteDB() *alertServerDo {
	return a.Clauses(dbresolver.Write)
}

func (a alertServerDo) Session(config *gorm.Session) *alertServerDo {
	return a.withDO(a.DO.Session(config))
}

func (a alertServerDo) Clauses(conds ...clause.Expression) *alertServerDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a alertServerDo) Returning(value interface{}, columns ...string) *alertServerDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a alertServerDo) Not(conds ...gen.Condition) *alertServerDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a alertServerDo) Or(conds ...gen.Condition) *alertServerDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a alertServerDo) Select(conds ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a alertServerDo) Where(conds ...gen.Condition) *alertServerDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a alertServerDo) Order(conds ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a alertServerDo) Distinct(cols ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a alertServerDo) Omit(cols ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a alertServerDo) Join(table schema.Tabler, on ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a alertServerDo) LeftJoin(table schema.Tabler, on ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a alertServerDo) RightJoin(table schema.Tabler, on ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a alertServerDo) Group(cols ...field.Expr) *alertServerDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a alertServerDo) Having(conds ...gen.Condition) *alertServerDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a alertServerDo) Limit(limit int) *alertServerDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a alertServerDo) Offset(offset int) *alertServerDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a alertServerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *alertServerDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a alertServerDo) Unscoped() *alertServerDo {
	return a.withDO(a.DO.Unscoped())
}

func (a alertServerDo) Create(values ...*model.AlertServer) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a alertServerDo) CreateInBatches(values []*model.AlertServer, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a alertServerDo) Save(values ...*model.AlertServer) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a alertServerDo) First() (*model.AlertServer, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertServer), nil
	}
}

func (a alertServerDo) Take() (*model.AlertServer, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertServer), nil
	}
}

func (a alertServerDo) Last() (*model.AlertServer, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertServer), nil
	}
}

func (a alertServerDo) Find() ([]*model.AlertServer, error) {
	result, err := a.DO.Find()
	return result.([]*model.AlertServer), err
}

func (a alertServerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AlertServer, err error) {
	buf := make([]*model.AlertServer, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a alertServerDo) FindInBatches(result *[]*model.AlertServer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a alertServerDo) Attrs(attrs ...field.AssignExpr) *alertServerDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a alertServerDo) Assign(attrs ...field.AssignExpr) *alertServerDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a alertServerDo) Joins(fields ...field.RelationField) *alertServerDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a alertServerDo) Preload(fields ...field.RelationField) *alertServerDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a alertServerDo) FirstOrInit() (*model.AlertServer, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertServer), nil
	}
}

func (a alertServerDo) FirstOrCreate() (*model.AlertServer, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AlertServer), nil
	}
}

func (a alertServerDo) FindByPage(offset int, limit int) (result []*model.AlertServer, count int64, err error) {
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

func (a alertServerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a alertServerDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a alertServerDo) Delete(models ...*model.AlertServer) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *alertServerDo) withDO(do gen.Dao) *alertServerDo {
	a.DO = *do.(*gen.DO)
	return a
}
