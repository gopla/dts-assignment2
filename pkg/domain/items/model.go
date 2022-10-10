package items

type Item struct {
  Id          uint64 `json:"lineItemId,omitempty" gorm:"column:item_id;primaryKey"`
  ItemCode    string `json:"itemCode"`
  Description string `json:"description"`
  Quantity    uint64 `json:"quantity"`
  OrderID     uint64 `json:"-" gorm:"column:order_id"`
}
