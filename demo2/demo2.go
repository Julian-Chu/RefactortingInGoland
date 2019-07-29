package demo2

import "math"

type Order struct {
	Qty       int
	ItemPrice float64
}

func Price(o Order) float64 {
	// 基本價 - 數量折扣 + 運費
	const minOrderQtyForDiscount = 500
	return float64(o.Qty)*o.ItemPrice - math.Max(0, float64(o.Qty-minOrderQtyForDiscount))*o.ItemPrice*0.05 +
		math.Min(float64(o.Qty)*o.ItemPrice*0.1, 100)
}
