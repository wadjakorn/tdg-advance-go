package products

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (ProductService) products() []Product {
	r := []Product{}
	// no DB, so mock products
	r = append(r, Product{Title: "Nuphy Air75", Price: 3290})
	r = append(r, Product{Title: "Logitech MX Keys", Price: 3990})
	return r
}
