package orders

import (
  "context"
)

type OrderRepo interface {
  GetOrder(ctx context.Context) (order []Order, err error)
  CreateOrder(ctx context.Context, order *Order) (err error)
  UpdateOrder(ctx context.Context, order *Order, oldData *Order) (err error)
  ShowOrder(ctx context.Context, order *Order, id int) (result Order, err error)
  DeleteOrder(ctx context.Context, id int) (err error)
}
