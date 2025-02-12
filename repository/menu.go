package repository

import (
	"fmt"
	"orderfoodonline/entity"

	"gorm.io/gorm"
)

type MenuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) GetMany(page, limit int) ([]entity.Menu, error) {
	var menus []entity.Menu
	offset := (page - 1) * limit

	err := r.db.Limit(limit).Offset(offset).Find(&menus).Error
	return menus, err
}

func (r *MenuRepository) GetOne(id uint) (*entity.Menu, error) {
	menu := &entity.Menu{}
	err := r.db.First(menu, id).Error
	return menu, err
}

func (r *MenuRepository) Create(menu *entity.Menu) error {
	return r.db.Create(menu).Error
}

func (r *MenuRepository) Update(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entity.Menu{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *MenuRepository) UpdateQuantity(id uint, operation string, qty int) error {
	menu := &entity.Menu{}
	if err := r.db.First(menu, id).Error; err != nil {
		return err
	}

	if operation == "increase" {
		menu.Quantity += qty
	} else if operation == "decrease" && menu.Quantity >= qty {
		menu.Quantity -= qty
	} else {
		return fmt.Errorf("invalid operation or insufficient quantity")
	}

	return r.db.Save(menu).Error
}

func (r *MenuRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Menu{}, id).Error
}
