package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/gopla/assignment-2/config/postgres"
  "github.com/gopla/assignment-2/pkg/domain/items"
  "github.com/gopla/assignment-2/pkg/domain/orders"

  engine "github.com/gopla/assignment-2/config/gin"

  orderrepo "github.com/gopla/assignment-2/pkg/repository/orders"
  orderhandler "github.com/gopla/assignment-2/pkg/server/http/handler/orders"
  orderrouter "github.com/gopla/assignment-2/pkg/server/http/router/v1"
  orderusecase "github.com/gopla/assignment-2/pkg/usecase/orders"
)

func main() {
  postgreCln := postgres.NewPostgresConnection(postgres.Config{
    Host:         "localhost",
    Port:         "5432",
    User:         "postgres",
    Password:     "root",
    DatabaseName: "orders_by",
  })
  postgreCln.MigrateDB(orders.Order{}, items.Item{})

  ginEngine := engine.NewGinHttp(engine.Config{
    Port: ":8080",
  })

  ginEngine.GetGin().Use(
    gin.Recovery(),
    gin.Logger(),
  )

  ginEngine.GetGin().GET("/", func(ctx *gin.Context) {
    respMap := map[string]any{
      "message": "server up and running",
    }

    ctx.JSON(http.StatusOK, respMap)
  })

  orderRepo := orderrepo.NewOrderRepo(postgreCln)
  orderUseCase := orderusecase.NewOrderUsecase(orderRepo)
  orderHandler := orderhandler.NewOrderHandler(orderUseCase)
  orderrouter.NewOrderRouter(ginEngine, orderHandler).Routers()

  ginEngine.Serve()
}
