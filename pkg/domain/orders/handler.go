package orders

import "github.com/gin-gonic/gin"

type OrderHandler interface {
  GetOrder(ctx *gin.Context)
  ShowOrder(ctx *gin.Context)
  CreateOrder(ctx *gin.Context)
  UpdateOrder(ctx *gin.Context)
  // DeleteOrder(ctx *gin.Context)
}
