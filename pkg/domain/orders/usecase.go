package orders

import "context"

type OrderUsecase interface {
  GetOrderSvc(ctx context.Context) (result []Order, err error)
  CreateOrderSvc(ctx context.Context, input Order) (result Order, err error)
  UpdateOrderSvc(ctx context.Context, input Order, id int) (result Order, err error)
  ShowOrderSvc(ctx context.Context, id int) (result Order, err error)
}
