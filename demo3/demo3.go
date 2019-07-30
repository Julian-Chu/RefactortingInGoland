package demo3

import "math"

type Order struct {
	Qty       int
	ItemPrice float64
}

func Price(o Order) float64 {
	// 基本價 - 數量折扣 + 運費
	return math.Min(float64(o.Qty)*o.ItemPrice*0.1, float64(100)) -
		math.Max(0, float64(o.Qty-500))*o.ItemPrice*0.05 +
		float64(o.Qty)*o.ItemPrice
}
