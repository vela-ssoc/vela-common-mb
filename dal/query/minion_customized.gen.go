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

func newMinionCustomized(db *gorm.DB, opts ...gen.DOOption) minionCustomized {
	_minionCustomized := minionCustomized{}

	_minionCustomized.minionCustomizedDo.UseDB(db, opts...)
	_minionCustomized.minionCustomizedDo.UseModel(&model.MinionCustomized{})

	tableName := _minionCustomized.minionCustomizedDo.TableName()
	_minionCustomized.ALL = field.NewAsterisk(tableName)
	_minionCustomized.ID = field.NewInt64(tableName, "id")
	_minionCustomized.Name = field.NewString(tableName, "name")
	_minionCustomized.Icon = field.NewString(tableName, "icon")
	_minionCustomized.CreatedAt = field.NewTime(tableName, "created_at")
	_minionCustomized.UpdatedAt = field.NewTime(tableName, "updated_at")

	_minionCustomized.fillFieldMap()

	return _minionCustomized
}

type minionCustomized struct {
	minionCustomizedDo minionCustomizedDo

	ALL       field.Asterisk
	ID        field.Int64
	Name      field.String
	Icon      field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (m minionCustomized) Table(newTableName string) *minionCustomized {
	m.minionCustomizedDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m minionCustomized) As(alias string) *minionCustomized {
	m.minionCustomizedDo.DO = *(m.minionCustomizedDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *minionCustomized) updateTableName(table string) *minionCustomized {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.Name = field.NewString(table, "name")
	m.Icon = field.NewString(table, "icon")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")

	m.fillFieldMap()

	return m
}

func (m *minionCustomized) WithContext(ctx context.Context) *minionCustomizedDo {
	return m.minionCustomizedDo.WithContext(ctx)
}

func (m minionCustomized) TableName() string { return m.minionCustomizedDo.TableName() }

func (m minionCustomized) Alias() string { return m.minionCustomizedDo.Alias() }

func (m minionCustomized) Columns(cols ...field.Expr) gen.Columns {
	return m.minionCustomizedDo.Columns(cols...)
}

func (m *minionCustomized) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *minionCustomized) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 5)
	m.fieldMap["id"] = m.ID
	m.fieldMap["name"] = m.Name
	m.fieldMap["icon"] = m.Icon
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
}

func (m minionCustomized) clone(db *gorm.DB) minionCustomized {
	m.minionCustomizedDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m minionCustomized) replaceDB(db *gorm.DB) minionCustomized {
	m.minionCustomizedDo.ReplaceDB(db)
	return m
}

type minionCustomizedDo struct{ gen.DO }

func (m minionCustomizedDo) Debug() *minionCustomizedDo {
	return m.withDO(m.DO.Debug())
}

func (m minionCustomizedDo) WithContext(ctx context.Context) *minionCustomizedDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m minionCustomizedDo) ReadDB() *minionCustomizedDo {
	return m.Clauses(dbresolver.Read)
}

func (m minionCustomizedDo) WriteDB() *minionCustomizedDo {
	return m.Clauses(dbresolver.Write)
}

func (m minionCustomizedDo) Session(config *gorm.Session) *minionCustomizedDo {
	return m.withDO(m.DO.Session(config))
}

func (m minionCustomizedDo) Clauses(conds ...clause.Expression) *minionCustomizedDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m minionCustomizedDo) Returning(value interface{}, columns ...string) *minionCustomizedDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m minionCustomizedDo) Not(conds ...gen.Condition) *minionCustomizedDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m minionCustomizedDo) Or(conds ...gen.Condition) *minionCustomizedDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m minionCustomizedDo) Select(conds ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m minionCustomizedDo) Where(conds ...gen.Condition) *minionCustomizedDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m minionCustomizedDo) Order(conds ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m minionCustomizedDo) Distinct(cols ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m minionCustomizedDo) Omit(cols ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m minionCustomizedDo) Join(table schema.Tabler, on ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m minionCustomizedDo) LeftJoin(table schema.Tabler, on ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m minionCustomizedDo) RightJoin(table schema.Tabler, on ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m minionCustomizedDo) Group(cols ...field.Expr) *minionCustomizedDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m minionCustomizedDo) Having(conds ...gen.Condition) *minionCustomizedDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m minionCustomizedDo) Limit(limit int) *minionCustomizedDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m minionCustomizedDo) Offset(offset int) *minionCustomizedDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m minionCustomizedDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *minionCustomizedDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m minionCustomizedDo) Unscoped() *minionCustomizedDo {
	return m.withDO(m.DO.Unscoped())
}

func (m minionCustomizedDo) Create(values ...*model.MinionCustomized) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m minionCustomizedDo) CreateInBatches(values []*model.MinionCustomized, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m minionCustomizedDo) Save(values ...*model.MinionCustomized) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m minionCustomizedDo) First() (*model.MinionCustomized, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionCustomized), nil
	}
}

func (m minionCustomizedDo) Take() (*model.MinionCustomized, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionCustomized), nil
	}
}

func (m minionCustomizedDo) Last() (*model.MinionCustomized, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionCustomized), nil
	}
}

func (m minionCustomizedDo) Find() ([]*model.MinionCustomized, error) {
	result, err := m.DO.Find()
	return result.([]*model.MinionCustomized), err
}

func (m minionCustomizedDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MinionCustomized, err error) {
	buf := make([]*model.MinionCustomized, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m minionCustomizedDo) FindInBatches(result *[]*model.MinionCustomized, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m minionCustomizedDo) Attrs(attrs ...field.AssignExpr) *minionCustomizedDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m minionCustomizedDo) Assign(attrs ...field.AssignExpr) *minionCustomizedDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m minionCustomizedDo) Joins(fields ...field.RelationField) *minionCustomizedDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m minionCustomizedDo) Preload(fields ...field.RelationField) *minionCustomizedDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m minionCustomizedDo) FirstOrInit() (*model.MinionCustomized, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionCustomized), nil
	}
}

func (m minionCustomizedDo) FirstOrCreate() (*model.MinionCustomized, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionCustomized), nil
	}
}

func (m minionCustomizedDo) FindByPage(offset int, limit int) (result []*model.MinionCustomized, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m minionCustomizedDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m minionCustomizedDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m minionCustomizedDo) Delete(models ...*model.MinionCustomized) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *minionCustomizedDo) withDO(do gen.Dao) *minionCustomizedDo {
	m.DO = *do.(*gen.DO)
	return m
}
