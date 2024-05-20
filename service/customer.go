package service

import (
	"context"
	"fmt"
	"login/api/models"
	"login/storage"
)

type customerService struct {
	storage storage.IStorage
}

func NewCustomerService(storage storage.IStorage) customerService {
	return customerService{storage: storage}
}

func (s customerService) Create(ctx context.Context, customer models.CustomerCreate) (string, error) {
	id, err := s.storage.CustomerStorage().Create(ctx, customer)
	if err != nil {
		fmt.Println("error while creating customer, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s customerService) Update(ctx context.Context, customer models.CustomerCreate, id string) (string, error) {
	// business logic
	id, err := s.storage.CustomerStorage().Update(ctx, customer, id)
	if err != nil {
		fmt.Println("error while updating customer, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s customerService) GetAllCustomers(ctx context.Context, customer models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	// business logic
	customers, err := s.storage.CustomerStorage().GetAll(ctx, models.GetAllCustomersRequest{})
	if err != nil {
		fmt.Println("error while getting all customer, err: ", err)
		return customers, err
	}
	// business logic
	return customers, nil
}
func (s customerService) GetCustomerById(ctx context.Context, id string) (models.GetCustomer, error) {
	// business logic
	resp, err := s.storage.CustomerStorage().GetCustomerById(ctx, id)
	if err != nil {
		fmt.Println("error while getting  customer, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}

func (s customerService) Delete(ctx context.Context, id string) error {
	// business logic
	_, err := s.storage.CustomerStorage().Delete(ctx, id)
	if err != nil {
		fmt.Println("error while deletting all customer, err: ", err)
		return err
	}
	// business logic
	return nil
}

func (s customerService) UpdateBirthday(ctx context.Context, customer models.Birthday) (string, error) {
	// business logic
	id, err := s.storage.CustomerStorage().UpdateBirthday(ctx, customer)
	if err != nil {
		fmt.Println("error while updating customer, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s customerService) GetAgeById(ctx context.Context, id string) (int, error) {
	// business logic
	resp, err := s.storage.CustomerStorage().GetAgeById(ctx, id)
	if err != nil {
		fmt.Println("error while getting  age, err: ", err)
		return resp, err
	}
	// business logic
	return resp, nil
}
