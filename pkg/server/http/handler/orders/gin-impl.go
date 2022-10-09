package orders

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/gopla/assignment-2/pkg/domain/orders"
)

type OrderHdlImpl struct {
  orderUsecase orders.OrderUsecase
}

func NewOrderHandler(ordersUsecase orders.OrderUsecase) orders.OrderHandler {
  return &OrderHdlImpl{orderUsecase: ordersUsecase}
}

// Get Order
func (o *OrderHdlImpl) GetOrder(ctx *gin.Context) {
  result, err := o.orderUsecase.GetOrderSvc(ctx)
  if err != nil {
    ctx.JSON(http.StatusBadGateway, gin.H{
      "Error": err.Error(),
    })
  }

  ctx.JSON(http.StatusOK, gin.H{
    "Data": result,
  })
}

// Create Order
func (o *OrderHdlImpl) CreateOrder(ctx *gin.Context) {
  var order orders.Order

  if err := ctx.ShouldBind(&order); err != nil {
    ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "Error": "failed to bind payload",
    })
    return
  }

  result, err := o.orderUsecase.CreateOrderSvc(ctx, order)
  if err != nil {
    ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
      "Error": err,
    })
    return
  }

  ctx.JSON(http.StatusOK, gin.H{
    "Data": result,
  })

}