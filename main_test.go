package main


import (
	"testing"
)

func TestproductInProducts(t *testing.T){
	product := &Product{Name: "test", Id: 1}
	products := make([]*Product, 0)
	products = append(products, product)
	
	got := productInProducts(product, products)
	want := true

	if got != want {
			t.Errorf("got %t, wanted %t", got, want)
	}
}

