package storage

import (
	"context"
	"login/api/models"
)

type IStorage interface {
	CloseDB()
	CustomerStorage() CustomerStorage
}

type CustomerStorage interface {
	Create(ctx context.Context, customer models.CustomerCreate) (string, error)
	Update(ctx context.Context, customer models.CustomerCreate, id string) (string, error)
	GetAll(ctx context.Context, customer models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error)
	GetCustomerById(ctx context.Context, id string) (models.GetCustomer, error)
	Delete(ctx context.Context, id string) (string, error)
	UpdateBirthday(ctx context.Context, customer models.Birthday) (string, error)
	GetAgeById(ctx context.Context, id string) (int, error)
}
