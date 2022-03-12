package main


import (
	"testing"
)

func TestProductInProducts(t *testing.T){
	product := &Product{Name: "test", Id: 1}
	products := make([]*Product, 0)
	products = append(products, product)
	
	got := ProductInProducts(product, products)
	want := true

	if got != want {
			t.Errorf("got %t, wanted %t", got, want)
	}
}

