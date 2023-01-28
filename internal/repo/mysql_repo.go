package repo

import (
	"database/sql"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"tiktok-server/internal/repo/op"
	"tiktok-server/internal/sources"
)

const DefaultIsolationLevel = sql.LevelReadCommitted

type MySQLRepo[T any] struct{}

func (rep *MySQLRepo[T]) SetupRepo() {

}

// Save 从数据库中查询出来，修改某个（些）字段后重新存储
func (rep *MySQLRepo[T]) Save(o *T, tx ...*gorm.DB) error {
	err := applyingTx(tx...).Save(o).Error
	if err != nil {
		return errors.Wrap(err, "create record")
	}
	return nil
}

// Create 插入某个记录
func (rep *MySQLRepo[T]) Create(o *T, tx ...*gorm.DB) error {
	err := applyingTx(tx...).Create(o).Error
	if err != nil {
		return errors.Wrap(err, "create record")
	}
	return nil
}

func (rep *MySQLRepo[T]) FindByID(pk string, tx ...*gorm.DB) (*T, error) {
	t := new(T)
	err := applyingTx(tx...).First(t, "id = ?", pk).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return t, err
}

func (rep *MySQLRepo[T]) FindOneWithTx(tx *gorm.DB, conditions ...*op.Condition) (*T, error) {
	t := new(T)
	db := applyingTx(tx)
	for _, condition := range conditions {
		db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
	}
	err := db.First(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return t, err
}

func (rep *MySQLRepo[T]) FindOne(conditions ...*op.Condition) (*T, error) {
	t := new(T)
	db := applyingTx()
	for _, condition := range conditions {
		db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
	}
	err := db.First(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return t, err
}

func (rep *MySQLRepo[T]) Find(conditions ...*op.Condition) ([]*T, error) {
	t := new([]*T)
	db := sources.MysqlSource.Db
	if conditions != nil && len(conditions) >= 1 {
		for _, condition := range conditions {
			db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
		}
	}
	err := db.Find(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return *t, err
}

func (rep *MySQLRepo[T]) FindWithTx(tx *gorm.DB, conditions ...*op.Condition) ([]*T, error) {
	t := new([]*T)
	db := applyingTx(tx)
	if conditions != nil && len(conditions) >= 1 {
		for _, condition := range conditions {
			db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
		}
	}
	err := db.Find(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return *t, err
}

type FindTransaction[M any] struct {
	*gorm.DB
}

func (tx FindTransaction[M]) Exec() ([]*M, error) {
	t := new([]*M)
	err := tx.DB.Find(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return *t, err
}

func (rep *MySQLRepo[T]) FindTx(filter map[string]interface{}, conditions ...*op.Condition) *FindTransaction[T] {
	db := sources.MysqlSource.Db
	t := new(T)
	db = db.Model(t)
	for _, condition := range conditions {
		db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
	}
	if filter != nil {
		db = db.Where(filter)
	}
	return &FindTransaction[T]{db}
}

func (rep *MySQLRepo[T]) Count(conditions ...*op.Condition) (int64, error) {
	t := new(T)
	db := sources.MysqlSource.Db
	for _, condition := range conditions {
		db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
	}
	var ans int64
	err := db.Model(t).Count(&ans).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return 0, nil
	}
	return ans, err
}

func (rep *MySQLRepo[T]) FindFields(conditions []*op.Condition, fields ...string) ([]*T, error) {
	t := new([]*T)
	db := sources.MysqlSource.Db
	if conditions != nil && len(conditions) >= 1 {
		for _, condition := range conditions {
			db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
		}
	}
	err := db.Select(fields).Find(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return *t, err
}

func (rep *MySQLRepo[T]) FindPKArray(ids []int, conditions ...*op.Condition) ([]*T, error) {
	t := new([]*T)
	db := sources.MysqlSource.Db
	for _, condition := range conditions {
		db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
	}
	err := db.Where(ids).Find(t).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { //没有找到返回空
		return nil, nil
	}
	return *t, err
}

func (rep *MySQLRepo[T]) UpdateFiledByPKArray(ids []int, newValue map[string]interface{}) error {
	t := new(T)
	return sources.MysqlSource.Db.Model(t).Where(ids).Updates(newValue).Error
}

func (rep *MySQLRepo[T]) UpdateFiled(pk string, newValue map[string]any) error {
	t := new(T)
	return sources.MysqlSource.Db.Model(t).Where("id = ?", pk).Updates(newValue).Error
}

func (rep *MySQLRepo[T]) UpdateFiledWithTx(tx *gorm.DB, pk string, newValue map[string]any) error {
	t := new(T)
	return applyingTx(tx).Model(t).Where("id = ?", pk).Updates(newValue).Error
}

func (rep *MySQLRepo[T]) DeleteByPK(pk int) error {
	t := new(T)
	return sources.MysqlSource.Db.Model(t).Delete(t, pk).Error
}

func (rep *MySQLRepo[T]) Delete(conditions ...*op.Condition) error {
	t := new(T)
	db := sources.MysqlSource.Db
	if conditions != nil {
		for _, condition := range conditions {
			db = db.Where(condition.Filed+" "+condition.Opr+" ?", condition.Value)
		}
	}
	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(t).Error
	return err
}

func applyingTx(tx ...*gorm.DB) *gorm.DB {
	if len(tx) == 0 {
		return sources.MysqlSource.Db
	}
	return tx[0]
}