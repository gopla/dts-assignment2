package orders

import "github.com/gopla/assignment-2/pkg/domain/items"

type Order struct {
  OrderID      uint64       `json:"-" gorm:"column:order_id;primaryKey"`
  OrderedAt    string       `json:"orderedAt" gorm:"column:ordered_at"`
  CustomerName string       `json:"customerName" gorm:"column:customer_name"`
  Items        []items.Item `json:"items" gorm:"foreignKey:order_id"`
}
