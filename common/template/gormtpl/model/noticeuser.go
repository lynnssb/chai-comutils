package model

import (
	"chai-hotel/common/core"
	"context"
	"gorm.io/gorm"
)

type (
	NoticeUserModel interface {
		WithTrans(ctx context.Context) NoticeUserModel
		Insert(ctx context.Context, arg *NoticeUser) error
		BatchInsert(ctx context.Context, args []*NoticeUser) error
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		FindOne(ctx context.Context, id string) (*NoticeUser, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*NoticeUser, error)
		Enable(ctx context.Context, id string) error
		IsExist(ctx context.Context, id *string, code string) (bool, error)
	}

	NoticeUser struct {
		ID        string         `json:"id" gorm:"column:id;type:varchar(32);primaryKey"`          //ID
		IsEnable  bool           `json:"is_enable" gorm:"column:is_enable;not null"`               //是否启用
		CreatedAt int64          `json:"created_at" gorm:"column:created_at,autoCreateTime:milli"` //创建时间
		UpdatedAt int64          `json:"updated_at" gorm:"column:updated_at,autoUpdateTime:milli"` //更新时间
		DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`                      //删除时间
	}

	defaultNoticeUserModel struct {
		db *gorm.DB
	}
)

func NewNoticeUserModel(isMigration bool, db *gorm.DB) NoticeUserModel {
	if isMigration {
		err := db.AutoMigrate(&NoticeUser{})
		if err != nil {
			panic(err)
		}
	}
	return &defaultNoticeUserModel{db: db}
}

func (m *defaultNoticeUserModel) WithTrans(ctx context.Context) NoticeUserModel {
	return &defaultNoticeUserModel{db: core.GetDB(ctx, m.db)}
}

func (m *defaultNoticeUserModel) Insert(ctx context.Context, arg *NoticeUser) error {
	return m.db.Create(&arg).Error
}

func (m *defaultNoticeUserModel) BatchInsert(ctx context.Context, args []*NoticeUser) error {
	return m.db.CreateInBatches(&args, 100).Error
}

func (m *defaultNoticeUserModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&NoticeUser{}, id).Error
}

func (m *defaultNoticeUserModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&NoticeUser{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultNoticeUserModel) FindOne(ctx context.Context, id string) (*NoticeUser, error) {
	var result *NoticeUser

	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultNoticeUserModel) FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*NoticeUser, error) {
	var (
		total int64
		list  []*NoticeUser
		err   error
	)

	p := core.PageHandle(page, pageSize, filter)
	query := m.db.Model(&NoticeUser{})
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

func (m *defaultNoticeUserModel) Enable(ctx context.Context, id string) error {
	var (
		result *NoticeUser
		err    error
	)

	result, err = m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	return m.db.Model(&NoticeUser{}).Where("id = ?", id).Updates(map[string]interface{}{"is_enable": !result.IsEnable}).Error
}

func (m *defaultNoticeUserModel) IsExist(ctx context.Context, id *string, code string) (bool, error) {
	var count int64

	query := m.db.Model(&NoticeUser{})
	if id != nil {
		query = query.Where("id != ?", id)
	}
	err := query.Where("code = ?", code).Count(&count).Error

	return count > 0, err
}
