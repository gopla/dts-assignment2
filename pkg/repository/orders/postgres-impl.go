package orders

import (
	"context"
	"errors"
	"log"

	"github.com/gopla/assignment-2/config/postgres"
	"github.com/gopla/assignment-2/pkg/domain/orders"
)

type OrderRepoImpl struct {
  pgCln postgres.PostgresClient
}

func NewOrderRepo(pgCln postgres.PostgresClient) orders.OrderRepo {
  return &OrderRepoImpl{pgCln: pgCln}
}

func (o *OrderRepoImpl) GetOrder(ctx context.Context) (result []orders.Order, err error) {
  db := o.pgCln.GetClient()

  db.Model(&orders.Order{}).Preload("Items").Find(&result)

  if err = db.Error; err != nil {
    return
  }

  return result, err
}

func (o *OrderRepoImpl) CreateOrder(ctx context.Context, inputOrder *orders.Order) (err error) {
  db := o.pgCln.GetClient()

  db.Model(&orders.Order{}).Create(&inputOrder)

  if err = db.Error; err != nil {
    log.Println("Error when creating new order")
  }

  return err
}

func (o *OrderRepoImpl) UpdateOrder(ctx context.Context, inputOrder *orders.Order, oldData *orders.Order) (err error) {
  db := o.pgCln.GetClient()
  db.Model(oldData).Updates(inputOrder)

  for _, v := range inputOrder.Items {
    db.Model(v).Updates(v)
  }

  if db.Error != nil {
    return db.Error
  }

  return err
}

func (o *OrderRepoImpl) ShowOrder(ctx context.Context, inputOrder *orders.Order, id int) (result orders.Order, err error) {
  db := o.pgCln.GetClient()

  db.Preload("Items").First(&result, id)

  if err = db.Error; err != nil {
    return
  }

  return result, err
}

func (o *OrderRepoImpl) DeleteOrder(ctx context.Context, id int) (err error) {
  db := o.pgCln.GetClient()

  result := db.Delete(&orders.Order{}, id)

  if result.RowsAffected < 1 {
    return errors.New("record not found")
  }

  if result.Error != nil {
    return err
  }

  return err
}
