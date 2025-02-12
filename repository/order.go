package repository

import (
	"errors"
	"orderfoodonline/entity"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) GetUserID(orderID uint) (uint, error) {
	var order entity.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, errors.New("order not found")
		}
		return 0, err
	}
	return order.UserID, nil
}

func (r *OrderRepository) GetManyByUser(userID, page, limit int) ([]entity.Order, error) {
	var orders []entity.Order
	offset := (page - 1) * limit

	err := r.db.Where("user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&orders).Error

	return orders, err
}

func (r *OrderRepository) GetManyAdmin(page, limit int) ([]entity.Order, int64, error) {
	var orders []entity.Order
	var total int64

	err := r.db.Model(&entity.Order{}).
		Count(&total).
		Limit(limit).Offset((page - 1) * limit).
		Find(&orders).Error

	return orders, total, err
}

func (r *OrderRepository) GetManyByStatus(status string) ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Where("status = ?", status).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) Create(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) UpdatePayment(orderID uint) error {
	order := &entity.Order{}
	if err := r.db.First(order, orderID).Error; err != nil {
		return err
	}

	order.Payment = true
	order.Status = "success"

	return r.db.Save(order).Error
}
