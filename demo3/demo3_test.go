package demo3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrice(t *testing.T) {
	type args struct {
		o Order
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Qty:1, ItemPrice:1", args{Order{Qty: 1, ItemPrice: 1}}, 1.1},
		{"Qty:2, ItemPrice:1", args{Order{Qty: 2, ItemPrice: 1}}, 2.2},
		{"Qty:1000, ItemPrice:10", args{Order{Qty: 1000, ItemPrice: 10}}, 9850},
		{"Qty:500, ItemPrice:10", args{Order{Qty: 500, ItemPrice: 10}}, 5100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Price(tt.args.o); got != tt.want {
				t.Errorf("Price() = %v, want %v", got, tt.want)
			}
		})
	}
}
