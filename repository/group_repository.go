package repository

import (
	"context"

	"gorm.io/gorm"
	"optimus.dev.br/gaia/model"
)

type GroupRepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) Count(ctx context.Context) int64 {
	var count int64
	r.db.WithContext(ctx).Model(&model.Group{}).Count(&count)
	return count
}

func (r *GroupRepository) FindAll(ctx context.Context) []model.Group {
	var data []model.Group
	r.db.WithContext(ctx).Find(&data)
	return data
}

func (r *GroupRepository) FindByID(ctx context.Context, id int) (*model.Group, error) {
	var data model.Group
	err := r.db.WithContext(ctx).First(&data, id).Error
	return &data, err
}

func (r *GroupRepository) FindByName(ctx context.Context, name string) (*model.Group, error) {
	var data model.Group
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&data).Error
	return &data, err
}

func (r *GroupRepository) ExistsByName(ctx context.Context, name string) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.Group{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func (r *GroupRepository) ExistsByNameAndIDNot(ctx context.Context, name string, id int) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.Group{}).Where("name = ? AND id != ?", name, id).Count(&count)
	return count > 0
}

func (r *GroupRepository) Create(ctx context.Context, data *model.Group) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *GroupRepository) Update(ctx context.Context, data *model.Group) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *GroupRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&model.Group{}, id).Error
}
