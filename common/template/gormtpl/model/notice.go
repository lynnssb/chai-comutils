package model

import (
	"chai-hotel/common/core"
	"context"
	"gorm.io/gorm"
)

type (
	NoticeModel interface {
		WithTrans(ctx context.Context) NoticeModel
		Insert(ctx context.Context, arg *Notice) error
		BatchInsert(ctx context.Context, args []*Notice) error
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		FindOne(ctx context.Context, id string) (*Notice, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*Notice, error)
		Enable(ctx context.Context, id string) error
		IsExist(ctx context.Context, id *string, code string) (bool, error)
	}

	Notice struct {
		ID        string         `json:"id" gorm:"column:id;type:varchar(32);primaryKey"`          //ID
		IsEnable  bool           `json:"is_enable" gorm:"column:is_enable;not null"`               //是否启用
		CreatedAt int64          `json:"created_at" gorm:"column:created_at,autoCreateTime:milli"` //创建时间
		UpdatedAt int64          `json:"updated_at" gorm:"column:updated_at,autoUpdateTime:milli"` //更新时间
		DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`                      //删除时间
	}

	defaultNoticeModel struct {
		db *gorm.DB
	}
)

func NewNoticeModel(isMigration bool, db *gorm.DB) NoticeModel {
	if isMigration {
		err := db.AutoMigrate(&Notice{})
		if err != nil {
			panic(err)
		}
	}
	return &defaultNoticeModel{db: db}
}

func (m *defaultNoticeModel) WithTrans(ctx context.Context) NoticeModel {
	return &defaultNoticeModel{db: core.GetDB(ctx, m.db)}
}

func (m *defaultNoticeModel) Insert(ctx context.Context, arg *Notice) error {
	return m.db.Create(&arg).Error
}

func (m *defaultNoticeModel) BatchInsert(ctx context.Context, args []*Notice) error {
	return m.db.CreateInBatches(&args, 100).Error
}

func (m *defaultNoticeModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&Notice{}, id).Error
}

func (m *defaultNoticeModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&Notice{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultNoticeModel) FindOne(ctx context.Context, id string) (*Notice, error) {
	var result *Notice

	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultNoticeModel) FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*Notice, error) {
	var (
		total int64
		list  []*Notice
		err   error
	)

	p := core.PageHandle(page, pageSize, filter)
	query := m.db.Model(&Notice{})
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

func (m *defaultNoticeModel) Enable(ctx context.Context, id string) error {
	var (
		result *Notice
		err    error
	)

	result, err = m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	return m.db.Model(&Notice{}).Where("id = ?", id).Updates(map[string]interface{}{"is_enable": !result.IsEnable}).Error
}

func (m *defaultNoticeModel) IsExist(ctx context.Context, id *string, code string) (bool, error) {
	var count int64

	query := m.db.Model(&Notice{})
	if id != nil {
		query = query.Where("id != ?", id)
	}
	err := query.Where("code = ?", code).Count(&count).Error

	return count > 0, err
}
