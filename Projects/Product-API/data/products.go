package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required,description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	validate.RegisterValidation("description", ValidateDescription)
	return validate.Struct(p)

}

func validateSKU(fl validator.FieldLevel) bool  {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	match := re.FindAllString(fl.Field().String(), -1)
	fmt.Println(match)
	fmt.Println(len(match))
	if len(match) != 1 {
		return false
	}
	return true
}

func ValidateDescription(fl validator.FieldLevel) bool  {
	re := regexp.MustCompile(`[a-z]+\s[a-z]+\s[a-z]+\s[a-z]+\s[a-z]+`)
	match := re.FindAllString(fl.Field().String(), -1)
	if len(match) != 1 {
		return false
	}
	return true
}

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProducts (id int, p *Product) error {
	_, pos, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func FindProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

func DeleteProducts(id int, p *Product) error{
	_, pos, err := FindProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func getNextID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}

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
		UpdatedOn:   time.Now().UTC().String()},
}
