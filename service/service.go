package service

import "login/storage"

type IServiceManager interface {
	Customer() customerService
}

type Service struct {
	customerService customerService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.customerService = NewCustomerService(storage)

	return services
}

func (s Service) Customer() customerService {
	return s.customerService
}
