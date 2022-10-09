package orders

import (
  "context"
)

type OrderRepo interface {
  GetOrder(ctx context.Context) (order []Order, err error)
  CreateOrder(ctx context.Context, order *Order) (err error)
}
