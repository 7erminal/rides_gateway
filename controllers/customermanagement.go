package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// CustomermanagementController operations for Customermanagement
type CustomermanagementController struct {
	beego.Controller
}

// URLMapping ...
func (c *CustomermanagementController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("UpdateCustomerImage", c.UpdateCustomerImage)
}

// Post ...
// @Title Create
// @Description create Customermanagement
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.AddCustomer	true		"body for Customermanagement content"
// @Success 200 {object} responses.CustomerGatewayResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *CustomermanagementController) Post() {
	v := c.Ctx.Input.GetData("user")
	userData, err := v.(*responses.UsersOri)
	userIdStr := strconv.FormatInt(userData.UserId, 10)

	fmt.Printf("Type of v: %T\n", v)
	fmt.Printf("Value of v: %+v\n", v)

	logs.Info("Error is ", err)

	var cust requests.AddCustomer
	json.Unmarshal(c.Ctx.Input.RequestBody, &cust)

	// Manipulate data if needed
	//
	//

	// customerImage := functions.UploadItemImage(&c.Controller, filePath)
	logs.Info("Customer image path is ", cust.ImagePath)

	isSuccess := false
	message := "An error occurred"
	customerResp := functions.AddCustomer(&c.Controller, cust, userIdStr, "Individual")

	if customerResp.StatusCode == 200 {
		isSuccess = true
		message = "Customer successfully added"

		var custGateway responses.CustomerGateway = responses.CustomerGateway{
			CustomerId:           customerResp.Customer.CustomerId,
			FullName:             customerResp.Customer.FullName,
			Email:                customerResp.Customer.Email,
			PhoneNumber:          customerResp.Customer.PhoneNumber,
			Location:             customerResp.Customer.Location,
			IdentificationType:   customerResp.Customer.IdentificationType,
			IdentificationNumber: customerResp.Customer.IdentificationNumber,
			DateCreated:          customerResp.Customer.DateCreated,
			Status:               customerResp.Customer.Active,
		}
		resp := responses.CustomerGatewayResponseDTO{Success: isSuccess, Result: &custGateway, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Error("An error occurred while saving customer details ", customerResp.StatusDesc)
		message = "An error occurred while saving customer details"
		var resp responses.CustomerGatewayResponseDTO = responses.CustomerGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Customermanagement by id
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customermanagement
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CustomermanagementController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	isSuccess := false
	message := "An error occurred"
	customerResp := functions.GetCustomerDetails(&c.Controller, id)

	if customerResp.StatusCode == 200 {
		isSuccess = true
		message = "Customer successfully added"

		var custGateway responses.CustomerGateway = responses.CustomerGateway{
			CustomerId:           customerResp.Customer.CustomerId,
			FullName:             customerResp.Customer.FullName,
			Email:                customerResp.Customer.Email,
			PhoneNumber:          customerResp.Customer.PhoneNumber,
			Location:             customerResp.Customer.Location,
			IdentificationType:   customerResp.Customer.IdentificationType,
			IdentificationNumber: customerResp.Customer.IdentificationNumber,
			DateCreated:          customerResp.Customer.DateCreated,
			Status:               customerResp.Customer.Active,
		}
		resp := responses.CustomerGatewayResponseDTO{Success: isSuccess, Result: &custGateway, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Error("An error occurred while fetching customer details ", customerResp.StatusDesc)
		message = "An error occurred while getting customer details. The customer may not exist"
		var resp responses.CustomerGatewayResponseDTO = responses.CustomerGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Customermanagement
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customermanagement
// @Failure 403
// @router / [get]
func (c *CustomermanagementController) GetAll() {
	v := c.Ctx.Input.GetData("user")
	userData, err := v.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", v)
	fmt.Printf("Value of v: %+v\n", v)

	logs.Info("Error is ", err)

	logs.Info("User received is ", v, " AND ", userData)

	var fields string
	var sortby string
	var order string
	var query string
	var limit string
	var offset string

	// limit: 10 (default is 10)
	if v := c.GetString("limit"); v != "" {
		limit = v
	}
	// offset: 0 (default is 0)
	if v := c.GetString("offset"); v != "" {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = v
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = v
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		query = v
	}

	isSuccess := false
	message := "An error occurred"
	customerResp := functions.GetCustomers(&c.Controller, query, fields, sortby, order, offset, limit)

	if customerResp.StatusCode == 200 {
		isSuccess = true
		message = "Customers fetched successfully"

		customers := []responses.CustomerGateway{}

		for _, customer := range *customerResp.Customers {
			var custGateway responses.CustomerGateway = responses.CustomerGateway{
				CustomerId:           customer.CustomerId,
				FullName:             customer.FullName,
				Email:                customer.Email,
				PhoneNumber:          customer.PhoneNumber,
				Location:             customer.Location,
				IdentificationType:   customer.IdentificationType,
				IdentificationNumber: customer.IdentificationNumber,
				DateCreated:          customer.DateCreated,
				Status:               customer.Active,
				LastDeal:             customer.LastTxnDate,
			}

			customers = append(customers, custGateway)
		}

		resp := responses.CustomersGatewayResponseDTO{Success: isSuccess, Result: &customers, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Error("An error occurred while fetching customer details ", customerResp.StatusDesc)
		message = "An error occurred while fetching customer details"
		var resp responses.CustomersGatewayResponseDTO = responses.CustomersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Customermanagement
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Customermanagement	true		"body for Customermanagement content"
// @Success 200 {object} models.Customermanagement
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CustomermanagementController) Put() {

}

// UpdateCustomerImage ...
// @Title UpdateCustomerImage
// @Description Update Customer's Image
// @Param	Authorization		header 	string true		"header for User"
// @Param	Image		formData 	file	true		"Item Image"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /upload-customer-image [post]
func (c *CustomermanagementController) UpdateCustomerImage() {
	var isSuccess bool = false

	image, header, err := c.GetFile("Image")

	if err != nil {
		var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
		c.Data["json"] = resp
	} else {
		logs.Info("Success response received")
		isSuccess = false
		respCode, filePath := functions.SaveImage(&c.Controller, "Image", image, *header)

		if respCode == 200 {
			isSuccess = true

			var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: &filePath, StatusDesc: "Successfully uploaded image"}
			c.Data["json"] = resp

		} else {
			var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to upload file. Tmp"}
			c.Data["json"] = resp
		}

	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customermanagement
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CustomermanagementController) Delete() {

}
