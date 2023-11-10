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

func newMinionBin(db *gorm.DB, opts ...gen.DOOption) minionBin {
	_minionBin := minionBin{}

	_minionBin.minionBinDo.UseDB(db, opts...)
	_minionBin.minionBinDo.UseModel(&model.MinionBin{})

	tableName := _minionBin.minionBinDo.TableName()
	_minionBin.ALL = field.NewAsterisk(tableName)
	_minionBin.ID = field.NewInt64(tableName, "id")
	_minionBin.FileID = field.NewInt64(tableName, "file_id")
	_minionBin.Goos = field.NewString(tableName, "goos")
	_minionBin.Arch = field.NewString(tableName, "arch")
	_minionBin.Name = field.NewString(tableName, "name")
	_minionBin.Customized = field.NewString(tableName, "customized")
	_minionBin.Unstable = field.NewBool(tableName, "unstable")
	_minionBin.Caution = field.NewString(tableName, "caution")
	_minionBin.Ability = field.NewString(tableName, "ability")
	_minionBin.Size = field.NewInt64(tableName, "size")
	_minionBin.Hash = field.NewString(tableName, "hash")
	_minionBin.Semver = field.NewString(tableName, "semver")
	_minionBin.Changelog = field.NewString(tableName, "changelog")
	_minionBin.Weight = field.NewInt64(tableName, "weight")
	_minionBin.Deprecated = field.NewBool(tableName, "deprecated")
	_minionBin.CreatedAt = field.NewTime(tableName, "created_at")
	_minionBin.UpdatedAt = field.NewTime(tableName, "updated_at")

	_minionBin.fillFieldMap()

	return _minionBin
}

type minionBin struct {
	minionBinDo minionBinDo

	ALL        field.Asterisk
	ID         field.Int64
	FileID     field.Int64
	Goos       field.String
	Arch       field.String
	Name       field.String
	Customized field.String
	Unstable   field.Bool
	Caution    field.String
	Ability    field.String
	Size       field.Int64
	Hash       field.String
	Semver     field.String
	Changelog  field.String
	Weight     field.Int64
	Deprecated field.Bool
	CreatedAt  field.Time
	UpdatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (m minionBin) Table(newTableName string) *minionBin {
	m.minionBinDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m minionBin) As(alias string) *minionBin {
	m.minionBinDo.DO = *(m.minionBinDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *minionBin) updateTableName(table string) *minionBin {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.FileID = field.NewInt64(table, "file_id")
	m.Goos = field.NewString(table, "goos")
	m.Arch = field.NewString(table, "arch")
	m.Name = field.NewString(table, "name")
	m.Customized = field.NewString(table, "customized")
	m.Unstable = field.NewBool(table, "unstable")
	m.Caution = field.NewString(table, "caution")
	m.Ability = field.NewString(table, "ability")
	m.Size = field.NewInt64(table, "size")
	m.Hash = field.NewString(table, "hash")
	m.Semver = field.NewString(table, "semver")
	m.Changelog = field.NewString(table, "changelog")
	m.Weight = field.NewInt64(table, "weight")
	m.Deprecated = field.NewBool(table, "deprecated")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")

	m.fillFieldMap()

	return m
}

func (m *minionBin) WithContext(ctx context.Context) *minionBinDo {
	return m.minionBinDo.WithContext(ctx)
}

func (m minionBin) TableName() string { return m.minionBinDo.TableName() }

func (m minionBin) Alias() string { return m.minionBinDo.Alias() }

func (m minionBin) Columns(cols ...field.Expr) gen.Columns { return m.minionBinDo.Columns(cols...) }

func (m *minionBin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *minionBin) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 17)
	m.fieldMap["id"] = m.ID
	m.fieldMap["file_id"] = m.FileID
	m.fieldMap["goos"] = m.Goos
	m.fieldMap["arch"] = m.Arch
	m.fieldMap["name"] = m.Name
	m.fieldMap["customized"] = m.Customized
	m.fieldMap["unstable"] = m.Unstable
	m.fieldMap["caution"] = m.Caution
	m.fieldMap["ability"] = m.Ability
	m.fieldMap["size"] = m.Size
	m.fieldMap["hash"] = m.Hash
	m.fieldMap["semver"] = m.Semver
	m.fieldMap["changelog"] = m.Changelog
	m.fieldMap["weight"] = m.Weight
	m.fieldMap["deprecated"] = m.Deprecated
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
}

func (m minionBin) clone(db *gorm.DB) minionBin {
	m.minionBinDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m minionBin) replaceDB(db *gorm.DB) minionBin {
	m.minionBinDo.ReplaceDB(db)
	return m
}

type minionBinDo struct{ gen.DO }

func (m minionBinDo) Debug() *minionBinDo {
	return m.withDO(m.DO.Debug())
}

func (m minionBinDo) WithContext(ctx context.Context) *minionBinDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m minionBinDo) ReadDB() *minionBinDo {
	return m.Clauses(dbresolver.Read)
}

func (m minionBinDo) WriteDB() *minionBinDo {
	return m.Clauses(dbresolver.Write)
}

func (m minionBinDo) Session(config *gorm.Session) *minionBinDo {
	return m.withDO(m.DO.Session(config))
}

func (m minionBinDo) Clauses(conds ...clause.Expression) *minionBinDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m minionBinDo) Returning(value interface{}, columns ...string) *minionBinDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m minionBinDo) Not(conds ...gen.Condition) *minionBinDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m minionBinDo) Or(conds ...gen.Condition) *minionBinDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m minionBinDo) Select(conds ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m minionBinDo) Where(conds ...gen.Condition) *minionBinDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m minionBinDo) Order(conds ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m minionBinDo) Distinct(cols ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m minionBinDo) Omit(cols ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m minionBinDo) Join(table schema.Tabler, on ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m minionBinDo) LeftJoin(table schema.Tabler, on ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m minionBinDo) RightJoin(table schema.Tabler, on ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m minionBinDo) Group(cols ...field.Expr) *minionBinDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m minionBinDo) Having(conds ...gen.Condition) *minionBinDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m minionBinDo) Limit(limit int) *minionBinDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m minionBinDo) Offset(offset int) *minionBinDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m minionBinDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *minionBinDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m minionBinDo) Unscoped() *minionBinDo {
	return m.withDO(m.DO.Unscoped())
}

func (m minionBinDo) Create(values ...*model.MinionBin) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m minionBinDo) CreateInBatches(values []*model.MinionBin, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m minionBinDo) Save(values ...*model.MinionBin) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m minionBinDo) First() (*model.MinionBin, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionBin), nil
	}
}

func (m minionBinDo) Take() (*model.MinionBin, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionBin), nil
	}
}

func (m minionBinDo) Last() (*model.MinionBin, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionBin), nil
	}
}

func (m minionBinDo) Find() ([]*model.MinionBin, error) {
	result, err := m.DO.Find()
	return result.([]*model.MinionBin), err
}

func (m minionBinDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MinionBin, err error) {
	buf := make([]*model.MinionBin, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m minionBinDo) FindInBatches(result *[]*model.MinionBin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m minionBinDo) Attrs(attrs ...field.AssignExpr) *minionBinDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m minionBinDo) Assign(attrs ...field.AssignExpr) *minionBinDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m minionBinDo) Joins(fields ...field.RelationField) *minionBinDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m minionBinDo) Preload(fields ...field.RelationField) *minionBinDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m minionBinDo) FirstOrInit() (*model.MinionBin, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionBin), nil
	}
}

func (m minionBinDo) FirstOrCreate() (*model.MinionBin, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinionBin), nil
	}
}

func (m minionBinDo) FindByPage(offset int, limit int) (result []*model.MinionBin, count int64, err error) {
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

func (m minionBinDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m minionBinDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m minionBinDo) Delete(models ...*model.MinionBin) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *minionBinDo) withDO(do gen.Dao) *minionBinDo {
	m.DO = *do.(*gen.DO)
	return m
}
