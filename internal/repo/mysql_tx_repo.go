package repo

import (
	"database/sql"
	"gorm.io/gorm"
	"tiktok-server/internal/repo/op"
	"tiktok-server/internal/sources"
)

//提交 Commit
//回滚 Rollback

type WithTxRepo[T any] struct {
	baseRepo *MySQLRepo[T]
	*gorm.DB
}

func Begin(level ...sql.IsolationLevel) (tx *gorm.DB) {
	if level == nil || len(level) == 0 || len(level) >= 2 {
		level = []sql.IsolationLevel{DefaultIsolationLevel}
	}
	return sources.MysqlSource.Db.Begin(&sql.TxOptions{Isolation: level[0]})
}

func CommitOrRollback(tx *gorm.DB, err error) {
	if r := recover(); r != nil {
		tx.Rollback()
		return
	} else {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}
}

func (rep *MySQLRepo[T]) WithTx(tx *gorm.DB) *WithTxRepo[T] {
	return &WithTxRepo[T]{rep, tx}
}

func (rep *WithTxRepo[T]) Save(o *T) error {
	return rep.baseRepo.Save(o, rep.DB)
}

// Create 插入某个记录
func (rep *WithTxRepo[T]) Create(o *T) error {
	return rep.baseRepo.Save(o, rep.DB)
}

func (rep *WithTxRepo[T]) FindByID(pk string) (*T, error) {
	return rep.baseRepo.FindByID(pk, rep.DB)
}

func (rep *WithTxRepo[T]) FindOne(conditions ...*op.Condition) (*T, error) {
	return rep.baseRepo.FindOneWithTx(rep.DB, conditions...)
}

func (rep *WithTxRepo[T]) Find(conditions ...*op.Condition) ([]*T, error) {
	return rep.baseRepo.FindWithTx(rep.DB, conditions...)
}

func (rep *WithTxRepo[T]) UpdateFiled(pk string, newValue map[string]any) error {
	return rep.baseRepo.UpdateFiledWithTx(rep.DB, pk, newValue)
}
