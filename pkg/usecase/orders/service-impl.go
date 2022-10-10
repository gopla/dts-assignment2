package orders

import (
  "context"
  "errors"
  "log"

  "github.com/gopla/assignment-2/pkg/domain/orders"
)

type OrderUsecaseImpl struct {
  orderRepo orders.OrderRepo
}

func NewOrderUsecase(orderRepo orders.OrderRepo) orders.OrderUsecase {
  return &OrderUsecaseImpl{orderRepo: orderRepo}
}

func (u *OrderUsecaseImpl) GetOrderSvc(ctx context.Context) (result []orders.Order, err error) {
  result, err = u.orderRepo.GetOrder(ctx)

  if err != nil {
    log.Println("error when fetching data from database: " + err.Error())
    err = errors.New("INTERNAL_SERVER_ERROR")
    return result, err
  }

  return result, err
}

func (o *OrderUsecaseImpl) CreateOrderSvc(ctx context.Context, input orders.Order) (result orders.Order, err error) {
  if err = o.orderRepo.CreateOrder(ctx, &input); err != nil {
    log.Println("error while creating new order")
    err = errors.New("ERROR")
  }

  return input, err
}

func (o *OrderUsecaseImpl) UpdateOrderSvc(ctx context.Context, input orders.Order, id int) (result orders.Order, err error) {
  result, err = o.orderRepo.ShowOrder(ctx, &orders.Order{}, id)

  if err != nil {
    log.Println("error when fetching data from database: " + err.Error())
    err = errors.New("INTERNAL_SERVER_ERROR")
    return result, err
  }

  if err = o.orderRepo.UpdateOrder(ctx, &input, &result); err != nil {
    log.Println("error while updating old order")
    err = errors.New("ERROR")
  }

  return input, err
}

func (o *OrderUsecaseImpl) ShowOrderSvc(ctx context.Context, id int) (result orders.Order, err error) {
  result, err = o.orderRepo.ShowOrder(ctx, &orders.Order{}, id)

  if err != nil {
    log.Println("error when fetching data from database: " + err.Error())
    err = errors.New("INTERNAL_SERVER_ERROR")
    return result, err
  }

  return result, err
}

func (o *OrderUsecaseImpl) DeleteOrderSvc(ctx context.Context, id int) (result orders.Order, err error) {
  err = o.orderRepo.DeleteOrder(ctx, id)

  if err != nil {
    return result, err
  }

  return result, err
}
