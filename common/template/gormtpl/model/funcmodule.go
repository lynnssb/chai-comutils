package model

import (
	"chai-hotel/common/core"
	"context"
	"gorm.io/gorm"
)

type (
	FuncModuleModel interface {
		WithTrans(ctx context.Context) FuncModuleModel
		Insert(ctx context.Context, arg *FuncModule) error
		BatchInsert(ctx context.Context, args []*FuncModule) error
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		FindOne(ctx context.Context, id string) (*FuncModule, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*FuncModule, error)
		Enable(ctx context.Context, id string) error
		IsExist(ctx context.Context, id *string, code string) (bool, error)
	}

	FuncModule struct {
		ID        string         `json:"id" gorm:"column:id;type:varchar(32);primaryKey"`          //ID
		IsEnable  bool           `json:"is_enable" gorm:"column:is_enable;not null"`               //是否启用
		CreatedAt int64          `json:"created_at" gorm:"column:created_at,autoCreateTime:milli"` //创建时间
		UpdatedAt int64          `json:"updated_at" gorm:"column:updated_at,autoUpdateTime:milli"` //更新时间
		DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`                      //删除时间
	}

	defaultFuncModuleModel struct {
		db *gorm.DB
	}
)

func NewFuncModuleModel(isMigration bool, db *gorm.DB) FuncModuleModel {
	if isMigration {
		err := db.AutoMigrate(&FuncModule{})
		if err != nil {
			panic(err)
		}
	}
	return &defaultFuncModuleModel{db: db}
}

func (m *defaultFuncModuleModel) WithTrans(ctx context.Context) FuncModuleModel {
	return &defaultFuncModuleModel{db: core.GetDB(ctx, m.db)}
}

func (m *defaultFuncModuleModel) Insert(ctx context.Context, arg *FuncModule) error {
	return m.db.Create(&arg).Error
}

func (m *defaultFuncModuleModel) BatchInsert(ctx context.Context, args []*FuncModule) error {
	return m.db.CreateInBatches(&args, 100).Error
}

func (m *defaultFuncModuleModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&FuncModule{}, id).Error
}

func (m *defaultFuncModuleModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&FuncModule{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultFuncModuleModel) FindOne(ctx context.Context, id string) (*FuncModule, error) {
	var result *FuncModule

	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultFuncModuleModel) FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*FuncModule, error) {
	var (
		total int64
		list  []*FuncModule
		err   error
	)

	p := core.PageHandle(page, pageSize, filter)
	query := m.db.Model(&FuncModule{})
	if filter != nil {
		//todo......
	}
	err = query.Count(&total).Error
	if err != nil {
		return 0, nil, err
	}

	err = query.Offset(p.Page).Limit(p.PageSize).Order("created_at desc").Find(&list).Error

	return total, list, err
}

func (m *defaultFuncModuleModel) Enable(ctx context.Context, id string) error {
	var (
		result *FuncModule
		err    error
	)

	result, err = m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	return m.db.Model(&FuncModule{}).Where("id = ?", id).Updates(map[string]interface{}{"is_enable": !result.IsEnable}).Error
}

func (m *defaultFuncModuleModel) IsExist(ctx context.Context, id *string, code string) (bool, error) {
	var count int64

	query := m.db.Model(&FuncModule{})
	if id != nil {
		query = query.Where("id != ?", id)
	}
	err := query.Where("code = ?", code).Count(&count).Error

	return count > 0, err
}
