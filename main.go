package main


import (
  "fmt"
	"math/rand"
	"time"

)


type Product struct {
    Id int
    Name string
}


func RandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}


func RandomId() int {
	return rand.Int() % 5000
}


func NewProduct() *Product {
	return &Product{
			Id: RandomId(),
			Name: RandomString(50),
	}
}

func (p *Product) String() string  {
  return fmt.Sprintf("Product{Id: %d, Name: %s}\n", p.Id, p.Name)
}

func ProductInProducts(product *Product, products []*Product) bool {
	for _, element := range products {
		if element == product {
			return true
		}
	}
	return false
}

func LinearDedupProducts(products []*Product) []*Product{
	uniqueProducts := make([]*Product, 0)
	for i := 0; i < len(products); i++ {
		product := products[i]
		
		if !ProductInProducts(product, uniqueProducts) {
			uniqueProducts = append(uniqueProducts, product)
		}
	}
	return uniqueProducts
}

func LogarithmicDedupProducts(products []*Product) []*Product{
	productsMap := map[Product]struct{}{}
	uniqueProducts := []*Product{}
	for _, entry := range products {
			if _, value := productsMap[*entry]; !value {
				uniqueProducts = append(uniqueProducts, entry)
				productsMap[*entry] = struct{}{}
			}
	}

	return uniqueProducts
}

func main() {
    rand.Seed(234234)
    
    products := make([]*Product, 0)
    for i := 0; i < 5; i++ {
        product := NewProduct()
        products = append(products, product)

        if rand.Float32() > 0.5 {
            products = append(products, product)
        }
    }

    fmt.Printf("products: %+v", products)
		fmt.Printf("\n\n")

    /// Task: given products - construct a list of unique products where all the duplicates have been removed
    /// print the list of unique products.
    /// vvvv YOUR CODE GOES HERE vvvv ///

		// This solution takes linear time to run as it
		// loops through the entire slice to check
		// if the product is included
		fmt.Printf("O(n) solution\n")
		start := time.Now()
		linearProducts := LinearDedupProducts(products)
		fmt.Println("Time Taken: ", time.Since(start))
		fmt.Printf("products: %+v", linearProducts)
		fmt.Printf("\n\n")

		// This solution takes a logarithmic time to run as it
		// reduces the number of operations on the
		// next run for the same product.
		fmt.Printf("O(log n) solution \n")
		onStart := time.Now()
		logarithmicProducts := LogarithmicDedupProducts(products)
		fmt.Println(time.Since(onStart))
		fmt.Printf("products: %+v", logarithmicProducts)
    /// ^^^^ YOUR CODE GOES HERE ^^^^ /// 
}

