package model

import (
	"chai-hotel/common/core"
	"context"
	"gorm.io/gorm"
)

type (
	NoticeUserLinkModel interface {
		WithTrans(ctx context.Context) NoticeUserLinkModel
		Insert(ctx context.Context, arg *NoticeUserLink) error
		BatchInsert(ctx context.Context, args []*NoticeUserLink) error
		Delete(ctx context.Context, id string) error
		Update(ctx context.Context, id string, v map[string]interface{}) error
		FindOne(ctx context.Context, id string) (*NoticeUserLink, error)
		FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*NoticeUserLink, error)
		Enable(ctx context.Context, id string) error
		IsExist(ctx context.Context, id *string, code string) (bool, error)
	}

	NoticeUserLink struct {
		ID        string         `json:"id" gorm:"column:id;type:varchar(32);primaryKey"`          //ID
		IsEnable  bool           `json:"is_enable" gorm:"column:is_enable;not null"`               //是否启用
		CreatedAt int64          `json:"created_at" gorm:"column:created_at,autoCreateTime:milli"` //创建时间
		UpdatedAt int64          `json:"updated_at" gorm:"column:updated_at,autoUpdateTime:milli"` //更新时间
		DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`                      //删除时间
	}

	defaultNoticeUserLinkModel struct {
		db *gorm.DB
	}
)

func NewNoticeUserLinkModel(isMigration bool, db *gorm.DB) NoticeUserLinkModel {
	if isMigration {
		err := db.AutoMigrate(&NoticeUserLink{})
		if err != nil {
			panic(err)
		}
	}
	return &defaultNoticeUserLinkModel{db: db}
}

func (m *defaultNoticeUserLinkModel) WithTrans(ctx context.Context) NoticeUserLinkModel {
	return &defaultNoticeUserLinkModel{db: core.GetDB(ctx, m.db)}
}

func (m *defaultNoticeUserLinkModel) Insert(ctx context.Context, arg *NoticeUserLink) error {
	return m.db.Create(&arg).Error
}

func (m *defaultNoticeUserLinkModel) BatchInsert(ctx context.Context, args []*NoticeUserLink) error {
	return m.db.CreateInBatches(&args, 100).Error
}

func (m *defaultNoticeUserLinkModel) Delete(ctx context.Context, id string) error {
	return m.db.Delete(&NoticeUserLink{}, id).Error
}

func (m *defaultNoticeUserLinkModel) Update(ctx context.Context, id string, v map[string]interface{}) error {
	return m.db.Model(&NoticeUserLink{}).Where("id = ?", id).Updates(v).Error
}

func (m *defaultNoticeUserLinkModel) FindOne(ctx context.Context, id string) (*NoticeUserLink, error) {
	var result *NoticeUserLink

	err := m.db.Where("id = ?", id).First(&result).Error

	return result, err
}

func (m *defaultNoticeUserLinkModel) FindPage(ctx context.Context, page, pageSize int, filter *string) (int64, []*NoticeUserLink, error) {
	var (
		total int64
		list  []*NoticeUserLink
		err   error
	)

	p := core.PageHandle(page, pageSize, filter)
	query := m.db.Model(&NoticeUserLink{})
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

func (m *defaultNoticeUserLinkModel) Enable(ctx context.Context, id string) error {
	var (
		result *NoticeUserLink
		err    error
	)

	result, err = m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	return m.db.Model(&NoticeUserLink{}).Where("id = ?", id).Updates(map[string]interface{}{"is_enable": !result.IsEnable}).Error
}

func (m *defaultNoticeUserLinkModel) IsExist(ctx context.Context, id *string, code string) (bool, error) {
	var count int64

	query := m.db.Model(&NoticeUserLink{})
	if id != nil {
		query = query.Where("id != ?", id)
	}
	err := query.Where("code = ?", code).Count(&count).Error

	return count > 0, err
}
