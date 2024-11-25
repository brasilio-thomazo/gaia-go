package repository

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"optimus.dev.br/gaia/model"
)

type AppRepository struct {
	db *gorm.DB
}

func NewAppRepository(db *gorm.DB) *AppRepository {
	return &AppRepository{db: db}
}

func (r *AppRepository) FindAll(ctx context.Context) []model.App {
	var data []model.App
	r.db.WithContext(ctx).Find(&data)
	return data
}

func (r *AppRepository) FindByID(ctx context.Context, id uuid.UUID) (*model.App, error) {
	var data model.App
	err := r.db.WithContext(ctx).First(&data, id).Error
	return &data, err
}

func (r *AppRepository) ExistsByName(ctx context.Context, name string) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.App{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func (r *AppRepository) ExistsByNameAndIDNot(ctx context.Context, name string, id uuid.UUID) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.App{}).Where("name = ? AND id != ?", name, id).Count(&count)
	return count > 0
}

func (r *AppRepository) Create(ctx context.Context, data *model.App) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *AppRepository) Update(ctx context.Context, data *model.App) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *AppRepository) Delete(ctx context.Context, id uuid.UUID) error {
	data, err := r.FindByID(ctx, id)
	if err != nil {
		return err
	}
	return r.db.WithContext(ctx).Delete(data).Error
}
