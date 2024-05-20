package handler

import (
	"database/sql"
	"fmt"
	_ "login/api/docs"
	"login/api/models"
	"login/pkg/check"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/customer [POST]
// @Summary		creates a customer
// @Description	This api creates a customer and returns its id
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		customer body models.CustomerCreate true "customer"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateCustomer(c *gin.Context) {
	customer := models.CustomerCreate{}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePhone(customer.Phone[0]); err != nil {
		handleResponse(c, h.Log, "error while validating customer phone, phone: "+customer.Phone[0], http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateMail(customer.Mail); err != nil {
		handleResponse(c, h.Log, "error while validating customer mail, mail: "+customer.Mail, http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateBitrthday(customer.Birthday, customer.Age); err != nil {
		handleResponse(c, h.Log, "error while validating customer birthday, date: "+customer.Birthday, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Customer().Create(c.Request.Context(), customer)
	if err != nil {
		handleResponse(c, h.Log, "error while creating customer", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Router		/customer/update/{id} [PUT]
// @Summary		updates a customer
// @Description	This api updates a customer and returns its id
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		customer body models.CustomerCreate true "customer"
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateCustomer(c *gin.Context) {

	customer := models.CustomerCreate{}

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating customerId", http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidatePhone(customer.Phone[0]); err != nil {
		handleResponse(c, h.Log, "error while validating customer phone, phone: "+customer.Phone[0], http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateMail(customer.Mail); err != nil {
		handleResponse(c, h.Log, "error while validating customer mail, mail: "+customer.Mail, http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateBitrthday(customer.Birthday, customer.Age); err != nil {
		handleResponse(c, h.Log, "error while validating customer birthday, date: "+customer.Birthday, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Customer().Update(c.Request.Context(), customer, id)
	if err != nil {
		handleResponse(c, h.Log, "error while updating customer", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Router		/customer [GET]
// @Summary		get all customers
// @Description	This api get all customers
// @Tags		customer
// @Accept		json
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllCustomers(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(page, limit, search, "hand")

	resp, err := h.Service.Customer().GetAllCustomers(c.Request.Context(), models.GetAllCustomersRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all customers", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}

// @Router		/customer/{id} [GET]
// @Summary		get one customer
// @Description	This api get customer and returns its
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetCustomer(c *gin.Context) {

	id := c.Param("id")

	resp, err := h.Service.Customer().GetCustomerById(c.Request.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			handleResponse(c, h.Log, "customer not found", http.StatusNotFound, err.Error())
			return
		}
		handleResponse(c, h.Log, "error while getting customer", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "request successful", http.StatusOK, resp)
}

// @Router		/customer/{id} [DELETE]
// @Summary		delete one customer
// @Description	This api for delete customer
// @Tags		customer
// @Accept		json
// @Param 		id path string true "id"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) DeleteCustomer(c *gin.Context) {

	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating customerId", http.StatusBadRequest, err.Error())
		return
	}

	err := h.Service.Customer().Delete(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting customer", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Deleted successfully", http.StatusOK, nil)
}

// UpdateCustomerBirthday обновляет активность студента
// @Router		/customer/update_birthday/{id} [patch]
// @Summary		updates customer birthday
// @Description	This api updates customer birthday and returns its id
// @Tags		customer
// @Accept		json
// @Produce		json
// @Param		id path string true "customer id"
// @Param		birthday body models.Birthday true "customer birthday"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) UpdateBirthday(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, "error while validating customerId", http.StatusBadRequest, err.Error())
		return
	}

	var birthday models.Birthday
	if err := c.ShouldBindJSON(&birthday); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	age, err := h.Service.Customer().GetAgeById(c.Copy().Request.Context(), birthday.Id)
	if err != nil {
		fmt.Println("error while getting age for validating")
	}
	if err := check.ValidateBitrthday(birthday.Birthday, age); err != nil {
		handleResponse(c, h.Log, "error while validating customer birthday, date: "+birthday.Birthday, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.Service.Customer().UpdateBirthday(c.Request.Context(), birthday); err != nil {
		handleResponse(c, h.Log, "error while updating customer birthday", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated customer birthday successfully", http.StatusOK, id)
}
