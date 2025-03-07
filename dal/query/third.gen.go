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

func newThird(db *gorm.DB, opts ...gen.DOOption) third {
	_third := third{}

	_third.thirdDo.UseDB(db, opts...)
	_third.thirdDo.UseModel(&model.Third{})

	tableName := _third.thirdDo.TableName()
	_third.ALL = field.NewAsterisk(tableName)
	_third.ID = field.NewInt64(tableName, "id")
	_third.FileID = field.NewInt64(tableName, "file_id")
	_third.Name = field.NewString(tableName, "name")
	_third.Hash = field.NewString(tableName, "hash")
	_third.Desc = field.NewString(tableName, "desc")
	_third.Size = field.NewInt64(tableName, "size")
	_third.Customized = field.NewString(tableName, "customized")
	_third.Extension = field.NewString(tableName, "extension")
	_third.CreatedID = field.NewInt64(tableName, "created_id")
	_third.UpdatedID = field.NewInt64(tableName, "updated_id")
	_third.CreatedAt = field.NewTime(tableName, "created_at")
	_third.UpdatedAt = field.NewTime(tableName, "updated_at")

	_third.fillFieldMap()

	return _third
}

type third struct {
	thirdDo thirdDo

	ALL        field.Asterisk
	ID         field.Int64
	FileID     field.Int64
	Name       field.String
	Hash       field.String
	Desc       field.String
	Size       field.Int64
	Customized field.String
	Extension  field.String
	CreatedID  field.Int64
	UpdatedID  field.Int64
	CreatedAt  field.Time
	UpdatedAt  field.Time

	fieldMap map[string]field.Expr
}

func (t third) Table(newTableName string) *third {
	t.thirdDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t third) As(alias string) *third {
	t.thirdDo.DO = *(t.thirdDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *third) updateTableName(table string) *third {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.FileID = field.NewInt64(table, "file_id")
	t.Name = field.NewString(table, "name")
	t.Hash = field.NewString(table, "hash")
	t.Desc = field.NewString(table, "desc")
	t.Size = field.NewInt64(table, "size")
	t.Customized = field.NewString(table, "customized")
	t.Extension = field.NewString(table, "extension")
	t.CreatedID = field.NewInt64(table, "created_id")
	t.UpdatedID = field.NewInt64(table, "updated_id")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")

	t.fillFieldMap()

	return t
}

func (t *third) WithContext(ctx context.Context) *thirdDo { return t.thirdDo.WithContext(ctx) }

func (t third) TableName() string { return t.thirdDo.TableName() }

func (t third) Alias() string { return t.thirdDo.Alias() }

func (t third) Columns(cols ...field.Expr) gen.Columns { return t.thirdDo.Columns(cols...) }

func (t *third) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *third) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.ID
	t.fieldMap["file_id"] = t.FileID
	t.fieldMap["name"] = t.Name
	t.fieldMap["hash"] = t.Hash
	t.fieldMap["desc"] = t.Desc
	t.fieldMap["size"] = t.Size
	t.fieldMap["customized"] = t.Customized
	t.fieldMap["extension"] = t.Extension
	t.fieldMap["created_id"] = t.CreatedID
	t.fieldMap["updated_id"] = t.UpdatedID
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
}

func (t third) clone(db *gorm.DB) third {
	t.thirdDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t third) replaceDB(db *gorm.DB) third {
	t.thirdDo.ReplaceDB(db)
	return t
}

type thirdDo struct{ gen.DO }

func (t thirdDo) Debug() *thirdDo {
	return t.withDO(t.DO.Debug())
}

func (t thirdDo) WithContext(ctx context.Context) *thirdDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t thirdDo) ReadDB() *thirdDo {
	return t.Clauses(dbresolver.Read)
}

func (t thirdDo) WriteDB() *thirdDo {
	return t.Clauses(dbresolver.Write)
}

func (t thirdDo) Session(config *gorm.Session) *thirdDo {
	return t.withDO(t.DO.Session(config))
}

func (t thirdDo) Clauses(conds ...clause.Expression) *thirdDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t thirdDo) Returning(value interface{}, columns ...string) *thirdDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t thirdDo) Not(conds ...gen.Condition) *thirdDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t thirdDo) Or(conds ...gen.Condition) *thirdDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t thirdDo) Select(conds ...field.Expr) *thirdDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t thirdDo) Where(conds ...gen.Condition) *thirdDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t thirdDo) Order(conds ...field.Expr) *thirdDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t thirdDo) Distinct(cols ...field.Expr) *thirdDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t thirdDo) Omit(cols ...field.Expr) *thirdDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t thirdDo) Join(table schema.Tabler, on ...field.Expr) *thirdDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t thirdDo) LeftJoin(table schema.Tabler, on ...field.Expr) *thirdDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t thirdDo) RightJoin(table schema.Tabler, on ...field.Expr) *thirdDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t thirdDo) Group(cols ...field.Expr) *thirdDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t thirdDo) Having(conds ...gen.Condition) *thirdDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t thirdDo) Limit(limit int) *thirdDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t thirdDo) Offset(offset int) *thirdDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t thirdDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *thirdDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t thirdDo) Unscoped() *thirdDo {
	return t.withDO(t.DO.Unscoped())
}

func (t thirdDo) Create(values ...*model.Third) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t thirdDo) CreateInBatches(values []*model.Third, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t thirdDo) Save(values ...*model.Third) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t thirdDo) First() (*model.Third, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Third), nil
	}
}

func (t thirdDo) Take() (*model.Third, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Third), nil
	}
}

func (t thirdDo) Last() (*model.Third, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Third), nil
	}
}

func (t thirdDo) Find() ([]*model.Third, error) {
	result, err := t.DO.Find()
	return result.([]*model.Third), err
}

func (t thirdDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Third, err error) {
	buf := make([]*model.Third, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t thirdDo) FindInBatches(result *[]*model.Third, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t thirdDo) Attrs(attrs ...field.AssignExpr) *thirdDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t thirdDo) Assign(attrs ...field.AssignExpr) *thirdDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t thirdDo) Joins(fields ...field.RelationField) *thirdDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t thirdDo) Preload(fields ...field.RelationField) *thirdDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t thirdDo) FirstOrInit() (*model.Third, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Third), nil
	}
}

func (t thirdDo) FirstOrCreate() (*model.Third, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Third), nil
	}
}

func (t thirdDo) FindByPage(offset int, limit int) (result []*model.Third, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t thirdDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t thirdDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t thirdDo) Delete(models ...*model.Third) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *thirdDo) withDO(do gen.Dao) *thirdDo {
	t.DO = *do.(*gen.DO)
	return t
}
