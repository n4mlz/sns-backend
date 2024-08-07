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

	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/model"
)

func newReply(db *gorm.DB, opts ...gen.DOOption) reply {
	_reply := reply{}

	_reply.replyDo.UseDB(db, opts...)
	_reply.replyDo.UseModel(&model.Reply{})

	tableName := _reply.replyDo.TableName()
	_reply.ALL = field.NewAsterisk(tableName)
	_reply.ID = field.NewString(tableName, "id")
	_reply.CommentID = field.NewString(tableName, "comment_id")
	_reply.UserID = field.NewString(tableName, "user_id")
	_reply.Content = field.NewString(tableName, "content")
	_reply.CreatedAt = field.NewTime(tableName, "created_at")

	_reply.fillFieldMap()

	return _reply
}

type reply struct {
	replyDo

	ALL       field.Asterisk
	ID        field.String
	CommentID field.String
	UserID    field.String
	Content   field.String
	CreatedAt field.Time

	fieldMap map[string]field.Expr
}

func (r reply) Table(newTableName string) *reply {
	r.replyDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reply) As(alias string) *reply {
	r.replyDo.DO = *(r.replyDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reply) updateTableName(table string) *reply {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewString(table, "id")
	r.CommentID = field.NewString(table, "comment_id")
	r.UserID = field.NewString(table, "user_id")
	r.Content = field.NewString(table, "content")
	r.CreatedAt = field.NewTime(table, "created_at")

	r.fillFieldMap()

	return r
}

func (r *reply) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reply) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 5)
	r.fieldMap["id"] = r.ID
	r.fieldMap["comment_id"] = r.CommentID
	r.fieldMap["user_id"] = r.UserID
	r.fieldMap["content"] = r.Content
	r.fieldMap["created_at"] = r.CreatedAt
}

func (r reply) clone(db *gorm.DB) reply {
	r.replyDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reply) replaceDB(db *gorm.DB) reply {
	r.replyDo.ReplaceDB(db)
	return r
}

type replyDo struct{ gen.DO }

type IReplyDo interface {
	gen.SubQuery
	Debug() IReplyDo
	WithContext(ctx context.Context) IReplyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReplyDo
	WriteDB() IReplyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReplyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReplyDo
	Not(conds ...gen.Condition) IReplyDo
	Or(conds ...gen.Condition) IReplyDo
	Select(conds ...field.Expr) IReplyDo
	Where(conds ...gen.Condition) IReplyDo
	Order(conds ...field.Expr) IReplyDo
	Distinct(cols ...field.Expr) IReplyDo
	Omit(cols ...field.Expr) IReplyDo
	Join(table schema.Tabler, on ...field.Expr) IReplyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReplyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReplyDo
	Group(cols ...field.Expr) IReplyDo
	Having(conds ...gen.Condition) IReplyDo
	Limit(limit int) IReplyDo
	Offset(offset int) IReplyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReplyDo
	Unscoped() IReplyDo
	Create(values ...*model.Reply) error
	CreateInBatches(values []*model.Reply, batchSize int) error
	Save(values ...*model.Reply) error
	First() (*model.Reply, error)
	Take() (*model.Reply, error)
	Last() (*model.Reply, error)
	Find() ([]*model.Reply, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Reply, err error)
	FindInBatches(result *[]*model.Reply, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Reply) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReplyDo
	Assign(attrs ...field.AssignExpr) IReplyDo
	Joins(fields ...field.RelationField) IReplyDo
	Preload(fields ...field.RelationField) IReplyDo
	FirstOrInit() (*model.Reply, error)
	FirstOrCreate() (*model.Reply, error)
	FindByPage(offset int, limit int) (result []*model.Reply, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReplyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r replyDo) Debug() IReplyDo {
	return r.withDO(r.DO.Debug())
}

func (r replyDo) WithContext(ctx context.Context) IReplyDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r replyDo) ReadDB() IReplyDo {
	return r.Clauses(dbresolver.Read)
}

func (r replyDo) WriteDB() IReplyDo {
	return r.Clauses(dbresolver.Write)
}

func (r replyDo) Session(config *gorm.Session) IReplyDo {
	return r.withDO(r.DO.Session(config))
}

func (r replyDo) Clauses(conds ...clause.Expression) IReplyDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r replyDo) Returning(value interface{}, columns ...string) IReplyDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r replyDo) Not(conds ...gen.Condition) IReplyDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r replyDo) Or(conds ...gen.Condition) IReplyDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r replyDo) Select(conds ...field.Expr) IReplyDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r replyDo) Where(conds ...gen.Condition) IReplyDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r replyDo) Order(conds ...field.Expr) IReplyDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r replyDo) Distinct(cols ...field.Expr) IReplyDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r replyDo) Omit(cols ...field.Expr) IReplyDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r replyDo) Join(table schema.Tabler, on ...field.Expr) IReplyDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r replyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReplyDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r replyDo) RightJoin(table schema.Tabler, on ...field.Expr) IReplyDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r replyDo) Group(cols ...field.Expr) IReplyDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r replyDo) Having(conds ...gen.Condition) IReplyDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r replyDo) Limit(limit int) IReplyDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r replyDo) Offset(offset int) IReplyDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r replyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReplyDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r replyDo) Unscoped() IReplyDo {
	return r.withDO(r.DO.Unscoped())
}

func (r replyDo) Create(values ...*model.Reply) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r replyDo) CreateInBatches(values []*model.Reply, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r replyDo) Save(values ...*model.Reply) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r replyDo) First() (*model.Reply, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Reply), nil
	}
}

func (r replyDo) Take() (*model.Reply, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Reply), nil
	}
}

func (r replyDo) Last() (*model.Reply, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Reply), nil
	}
}

func (r replyDo) Find() ([]*model.Reply, error) {
	result, err := r.DO.Find()
	return result.([]*model.Reply), err
}

func (r replyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Reply, err error) {
	buf := make([]*model.Reply, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r replyDo) FindInBatches(result *[]*model.Reply, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r replyDo) Attrs(attrs ...field.AssignExpr) IReplyDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r replyDo) Assign(attrs ...field.AssignExpr) IReplyDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r replyDo) Joins(fields ...field.RelationField) IReplyDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r replyDo) Preload(fields ...field.RelationField) IReplyDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r replyDo) FirstOrInit() (*model.Reply, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Reply), nil
	}
}

func (r replyDo) FirstOrCreate() (*model.Reply, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Reply), nil
	}
}

func (r replyDo) FindByPage(offset int, limit int) (result []*model.Reply, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r replyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r replyDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r replyDo) Delete(models ...*model.Reply) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *replyDo) withDO(do gen.Dao) *replyDo {
	r.DO = *do.(*gen.DO)
	return r
}
