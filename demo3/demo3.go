package demo3

import "math"

type Order struct {
	Qty       int
	ItemPrice float64
}

func Price(o Order) float64 {
	// 基本價 - 數量折扣 + 運費

	return float64(o.Qty)*o.ItemPrice - math.Max(0, float64(o.Qty-500))*o.ItemPrice*0.05 +
		math.Min(float64(o.Qty)*o.ItemPrice*0.1, float64(100))
}
