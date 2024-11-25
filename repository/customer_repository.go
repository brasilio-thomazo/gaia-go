package repository

import (
	"context"

	"gorm.io/gorm"
	"optimus.dev.br/gaia/model"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Count(ctx context.Context) int64 {
	var count int64
	r.db.WithContext(ctx).Model(&model.Customer{}).Count(&count)
	return count
}

func (r *CustomerRepository) FindAll(ctx context.Context) []model.Customer {
	var data []model.Customer
	r.db.WithContext(ctx).Find(&data)
	return data
}

func (r *CustomerRepository) FindByID(ctx context.Context, id int64) (*model.Customer, error) {
	var data model.Customer
	err := r.db.WithContext(ctx).First(&data, id).Error
	return &data, err
}

func (r *CustomerRepository) FindByName(ctx context.Context, name string) (*model.Customer, error) {
	var data model.Customer
	err := r.db.WithContext(ctx).Where("name = ?", name).First(&data).Error
	return &data, err
}

func (r *CustomerRepository) ExistsByName(ctx context.Context, name string) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.Customer{}).Where("name = ?", name).Count(&count)
	return count > 0
}

func (r *CustomerRepository) ExistsByNameAndIDNot(ctx context.Context, name string, id int64) bool {
	var count int64
	r.db.WithContext(ctx).Model(&model.Customer{}).Where("name = ? AND id != ?", name, id).Count(&count)
	return count > 0
}

func (r *CustomerRepository) Create(ctx context.Context, data *model.Customer) error {
	return r.db.WithContext(ctx).Create(data).Error
}

func (r *CustomerRepository) Update(ctx context.Context, data *model.Customer) error {
	return r.db.WithContext(ctx).Save(data).Error
}

func (r *CustomerRepository) Delete(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Delete(&model.Customer{}, id).Error
}
