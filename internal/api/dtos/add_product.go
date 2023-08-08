package dtos

type AddProduct struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
}