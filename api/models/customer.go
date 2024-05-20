package models

type CustomerCreate struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	Phone     []string `json:"phone"`
	Mail      string   `json:"mail"`
	Birthday  string   `json:"birthday"`
	Sex       string   `json:"sex"`
}

type Customer struct {
	Id          string   `json:"id"`
	External_id string   `json:"external_id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Age         int      `json:"age"`
	Phone       []string `json:"phone"`
	Mail        string   `json:"mail"`
	Birthday    string   `json:"birthday"`
	Sex         string   `json:"sex"`
}

type GetCustomer struct {
	Id          string   `json:"id"`
	External_id string   `json:"external_id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Age         int      `json:"age"`
	Phone       []string `json:"phone"`
	Mail        string   `json:"mail"`
	Birthday    string   `json:"birthday"`
	Sex         string   `json:"sex"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type GetAllCustomersRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllCustomersResponse struct {
	Customers []GetCustomer `json:"customers"`
	Count     int64         `json:"count"`
}

type Birthday struct {
	Id       string `json:"id"`
	Birthday string `json:"birtday"`
}
