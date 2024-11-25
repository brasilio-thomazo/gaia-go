package repository

import (
	"context"

	"gorm.io/gorm"
	"optimus.dev.br/gaia/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Count(ctx context.Context) int64 {
	var count int64
	r.db.WithContext(ctx).Model(&model.User{}).Count(&count)
	return count
}

func (r *UserRepository) FindAll(ctx context.Context) []model.User {
	var data []model.User
	r.db.WithContext(ctx).Preload("Group").Find(&data)
	return data
}

func (r *UserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var data model.User
	err := r.db.WithContext(ctx).Preload("Group").First(&data, id).Error
	return &data, err
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var data model.User
	err := r.db.WithContext(ctx).Preload("Group").Where("username = ?", username).First(&data).Error
	return &data, err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var data model.User
	err := r.db.WithContext(ctx).Preload("Group").Where("email = ?", email).First(&data).Error
	return &data, err
}

func (r *UserRepository) ExistsByUsername(ctx context.Context, username string) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (r *UserRepository) ExistsByUsernameAndIDNot(ctx context.Context, username string, id int64) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.User{}).Where("username = ? AND id != ?", username, id).Count(&count)
	return count > 0
}

func (r *UserRepository) ExistsByEmail(ctx context.Context, email string) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

func (r *UserRepository) ExistsByEmailAndIDNot(ctx context.Context, email string, id int64) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.User{}).Where("email = ? AND id != ?", email, id).Count(&count)
	return count > 0
}

func (r *UserRepository) Create(ctx context.Context, data *model.User) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *UserRepository) Update(ctx context.Context, data *model.User) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}
