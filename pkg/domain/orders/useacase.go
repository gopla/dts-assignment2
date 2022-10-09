package orders

import "context"

type OrderUsecase interface {
  GetOrderSvc(ctx context.Context) (result []Order, err error)
  CreateOrderSvc(ctx context.Context, input Order) (result Order, err error)
}
