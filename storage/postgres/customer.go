package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"login/api/models"
	"login/pkg"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomer(db *pgxpool.Pool) customerRepo {
	return customerRepo{
		db: db,
	}
}

func generateExternalID(db *pgxpool.Pool, ctx context.Context) (string, error) {
	var nextVal int
	err := db.QueryRow(ctx, "SELECT nextval('customers_external_id_seq')").Scan(&nextVal)
	if err != nil {
		return "", err
	}

	externalID := "num-" + fmt.Sprintf("%06d", nextVal)
	return externalID, nil
}

func (s *customerRepo) Create(ctx context.Context, customer models.CustomerCreate) (string, error) {

	id := uuid.New()
	externalID, err := generateExternalID(s.db, ctx)
	if err != nil {
		log.Fatal(err)
	}
	query := ` INSERT INTO customers (
		id,
		external_id,
		first_name,
		last_name,
		age,
		phone,
		mail,
		birthday,
		sex) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err = s.db.Exec(ctx,
		query,
		id,
		externalID,
		customer.FirstName,
		customer.LastName,
		customer.Age,
		customer.Phone,
		customer.Mail,
		customer.Birthday,
		customer.Sex)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *customerRepo) Update(ctx context.Context, customer models.CustomerCreate, id string) (string, error) {

	query := `UPDATE customers SET 
	first_name = $1, 
	last_name =$2, 
	age = $3, 
	phone = $4,
	mail = $5,
	birthday = $6, 
	sex = $7,
	updated_at = 'NOW()'
	WHERE id = $8`

	_, err := s.db.Exec(ctx, query,
		customer.FirstName,
		customer.LastName,
		customer.Age,
		pq.Array(customer.Phone),
		customer.Mail,
		customer.Birthday,
		customer.Sex,
		id)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (s *customerRepo) GetAll(ctx context.Context, req models.GetAllCustomersRequest) (models.GetAllCustomersResponse, error) {
	resp := models.GetAllCustomersResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit
	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}
	query := `SELECT id,
					external_id,
					first_name,
					last_name,
					age,
					phone,
					mail,
					birthday,
					sex
				FROM customers
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	fmt.Println(offest, req.Limit)

	for rows.Next() {
		var (
			customer    models.GetCustomer
			external_id sql.NullString
			firstName   sql.NullString
			lastName    sql.NullString
			phone       sql.NullString
			mail        sql.NullString
			birthday    sql.NullString
			sex         sql.NullString
			created_at  sql.NullString
			updated_at  sql.NullString
		)
		if err = rows.Scan(
			&customer.Id,
			&external_id,
			&firstName,
			&lastName,
			&customer.Age,
			&birthday,
			&phone,
			&mail,
			&sex,
			&created_at,
			&updated_at,
		); err != nil {
			return resp, err
		}
		customer.External_id = pkg.NullStringToString(external_id)
		customer.FirstName = pkg.NullStringToString(firstName)
		customer.LastName = pkg.NullStringToString(lastName)
		customer.Birthday = pkg.NullStringToString(birthday)
		customer.Phone[0] = pkg.NullStringToString(phone)
		customer.Mail = pkg.NullStringToString(mail)
		customer.Sex = pkg.NullStringToString(sex)
		customer.CreatedAt = pkg.NullStringToString(created_at)
		customer.UpdatedAt = pkg.NullStringToString(updated_at)
		resp.Customers = append(resp.Customers, customer)
	}

	err = s.db.QueryRow(ctx, `SELECT count(*) from customers WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *customerRepo) GetCustomerById(ctx context.Context, id string) (models.GetCustomer, error) {
	var (
		customer    models.GetCustomer
		external_id sql.NullString
		firstName   sql.NullString
		lastName    sql.NullString
		phone       []string
		mail        sql.NullString
		birthday    sql.NullString
		sex         sql.NullString
		created_at  sql.NullString
		updated_at  sql.NullString
	)
	query := `SELECT id,
					external_id,
					first_name,
					last_name,
					age,
					phone,
					mail,
					birthday,
					sex,
					created_at,
					updated_at
				FROM customers
				WHERE id = $1 LIMIT 1`
	rows := s.db.QueryRow(ctx, query, id)
	err := rows.Scan(
		&customer.Id,
		&external_id,
		&firstName,
		&lastName,
		&customer.Age,
		&birthday,
		pq.Array(&phone),
		&mail,
		&sex,
		&created_at,
		&updated_at)
	if err != nil {
		return customer, err
	}
	customer.External_id = pkg.NullStringToString(external_id)
	customer.FirstName = pkg.NullStringToString(firstName)
	customer.LastName = pkg.NullStringToString(lastName)
	customer.Birthday = pkg.NullStringToString(birthday)
	customer.Phone = phone
	customer.Mail = pkg.NullStringToString(mail)
	customer.Sex = pkg.NullStringToString(sex)
	customer.CreatedAt = pkg.NullStringToString(created_at)
	customer.UpdatedAt = pkg.NullStringToString(updated_at)

	return customer, nil
}

func (s *customerRepo) Delete(ctx context.Context, id string) (string, error) {
	query := `DELETE FROM customers WHERE id = $1`

	_, err := s.db.Exec(ctx, query, id)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *customerRepo) UpdateBirthday(ctx context.Context, customer models.Birthday) (string, error) {
	query := "UPDATE customers SET birthday = $1 WHERE id = $2"
	_, err := s.db.Exec(ctx, query, customer.Birthday, customer.Id)
	if err != nil {
		return "", err
	}

	return customer.Id, nil
}

func (s *customerRepo) GetAgeById(ctx context.Context, id string) (int, error) {
	var age int
	query := `SELECT age
				FROM customers
				WHERE id = $1 LIMIT 1`
	rows := s.db.QueryRow(ctx, query, id)
	err := rows.Scan(
		&age)
	if err != nil {
		return age, err
	}

	return age, nil
}
