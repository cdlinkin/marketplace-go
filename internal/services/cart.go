package services

type CartService struct {
}

func NewCartService() *CartService {
	return &CartService{}
}

func (s *CartService) AddProduct(userID, productID, quantity int) error {
	return nil
}
func (s *CartService) RemoveProduct(userID, productID int) error {
	return nil
}
func (s *CartService) GetCart(userID int) error {
	return nil
}
