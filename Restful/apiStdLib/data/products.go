package data

import (
	"encoding/json"
	"io"
	"time"
	"fmt"
)

// Product defines the structure for an API product with struct tags
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"-"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}
//to add a product
func (p *Product) FromJSON(r io.Reader) error{
	e := json.NewDecoder(r)
	return e.Decode(p)
}



// Products is a collection of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}
func PostProducts(p *Product){
	p.ID = getNextID()
	productList = append(productList,p)
}

func getNextID() int{
	lp := productList[len(productList) -1]
	return lp.ID +1
}

func UpdateProduct(id int, p *Product) error{
	_, pos, err := findProduct(id)
	if err != nil{
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProdNotFound = fmt.Errorf("Product Not Found")
func findProduct(id int) (*Product, int, error){
	for i, p := range productList{
		if p.ID ==id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProdNotFound

}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}