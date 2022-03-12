package main


import (
    "fmt"
	"math/rand"
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

		uniqueProducts := make([]*Product, 0)
		for i := 0; i < len(products) - 1; i++ {
			product := products[i]

			if !ProductInProducts(product, uniqueProducts) {
				uniqueProducts = append(uniqueProducts, product)
			}
		}

		fmt.Printf("Initial solution \n\n")
		fmt.Printf("products: %+v", uniqueProducts)

    /// Task: given products - construct a list of unique products where all the duplicates have been removed
    /// print the list of unique products.
    /// vvvv YOUR CODE GOES HERE vvvv ///
    

    /// ^^^^ YOUR CODE GOES HERE ^^^^ /// 
}

