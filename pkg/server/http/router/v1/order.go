package v1

import (
  "github.com/gin-gonic/gin"
  engine "github.com/gopla/assignment-2/config/gin"
  "github.com/gopla/assignment-2/pkg/domain/orders"
  "github.com/gopla/assignment-2/pkg/server/http/router"
)

type OrderRouterImpl struct {
  ginEngine    engine.HttpServer
  routerGroup  *gin.RouterGroup
  orderHandler orders.OrderHandler
}

func NewOrderRouter(ginEngine engine.HttpServer, orderHandler orders.OrderHandler) router.Router {
  routerGroup := ginEngine.GetGin().Group("/v1/orders")
  return &OrderRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, orderHandler: orderHandler}
}

func (o *OrderRouterImpl) get() {
  o.routerGroup.GET("", o.orderHandler.GetOrder)
}

func (o *OrderRouterImpl) post() {
  o.routerGroup.POST("", o.orderHandler.CreateOrder)
}

func (o *OrderRouterImpl) put() {
  o.routerGroup.PUT("/:id", o.orderHandler.UpdateOrder)
}

func (o *OrderRouterImpl) show() {
  o.routerGroup.GET("/:id", o.orderHandler.ShowOrder)
}

func (o *OrderRouterImpl) Routers() {
  o.get()
  o.post()
  o.put()
  o.show()
}
