package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {

	if quantity <= 0 {
		return errors.New("invalid quantity")
	}
	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return errors.New("product not found")
	}

	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	found := false
	for i, item := range cartItems {
		if item.ProductName == productName {
			cartItems[i].Quantity += quantity
			found = true
			break
		}
	}

	if !found {
		newCartItem := entity.CartItem{
			ProductName: productName,
			Price:       product.Price,
			Quantity:    quantity,
		}
		cartItems = append(cartItems, newCartItem)
	}

	return s.database.SaveCartItems(cartItems)
}

func (s *Service) RemoveCart(productName string) error {
	cartItems, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	var updatedCartItems []entity.CartItem
	found := false
	for _, item := range cartItems {
		if item.ProductName == productName {
			found = true
			continue
		}
		updatedCartItems = append(updatedCartItems, item)
	}

	if !found {
		return errors.New("product not found")
	}

	return s.database.SaveCartItems(updatedCartItems)
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	return s.database.GetCartItems()
}

func (s *Service) ResetCart() error {
	return s.database.SaveCartItems([]entity.CartItem{})
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	return s.database.GetProductData(), nil
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	cartItems, err := s.ShowCart()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	totalPrice := 0
	for _, item := range cartItems {
		totalPrice += item.Price * item.Quantity
	}

	change := money - totalPrice
	if change < 0 {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	s.ResetCart()

	return entity.PaymentInformation{
		TotalPrice:  totalPrice,
		Change:      change,
		ProductList: cartItems,
		MoneyPaid:   money,
	}, nil

}
