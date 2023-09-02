/**
 * @author:       wangxuebing
 * @fileName:     transaction.go
 * @date:         2023/1/14 16:45
 * @description:
 */

package core

import (
	"context"
	"gorm.io/gorm"
)

type (
	Transaction interface {
		ExecTx(context.Context, func(ctx context.Context) error) error
	}
	txDB struct {
		db *gorm.DB
	}

	// 用来承载事务上下文
	contextTxKey struct{}
)

func NewTransaction(db *gorm.DB) Transaction {
	return &txDB{db: db}
}

func (t *txDB) ExecTx(ctx context.Context, f func(ctx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return f(ctx)
	})
}

// GetDB 判断当前DB是不是使用事务DB
func GetDB(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB); ok {
		return tx
	}
	return db.Session(&gorm.Session{})
}
