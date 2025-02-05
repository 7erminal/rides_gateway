package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// TransactionsController operations for Transactions
type TransactionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *TransactionsController) URLMapping() {
	c.Mapping("PlaceRentalRequest", c.PlaceRentalRequest)
	c.Mapping("PlaceRentalRequest", c.PlaceRentalRequest)
}

// PlaceRentalRequest ...
// @Title Place Rental Request
// @Description Place an order to rent a product
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.RentalRequestDTO	true		"body for Transactions content"
// @Success 201 {object} models.Transactions
// @Failure 403 body is empty
// @router /place-rental-request [post]
func (c *TransactionsController) PlaceRentalRequest() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", u)
	fmt.Printf("Value of v: %+v\n", u)

	if err != false {
		logs.Error("An error occurred ", err)
	}
	var v requests.RentalRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	requestType := "RENTAL"

	isSuccess := false

	itemImage := functions.UploadPaymentProof(&c.Controller, v.PaymentProofImageUrl)

	if itemImage.StatusCode == 200 {
		logs.Info("Payment proof uploaded successfully")
	}

	products := []requests.Item{}

	for _, i := range v.Products {
		tProduct := requests.Item{ItemId: i.ProductId, Quantity: i.Quantity}
		products = append(products, tProduct)
	}

	q := requests.PostTransactionRequest{
		Items:                products,
		RequestType:          requestType,
		PaymentProofImageUrl: v.PaymentProofImageUrl,
		PaymentMethodId:      v.PaymentMethodId,
		Comment:              "",
		OrderLocation:        v.OrderLocation,
		OrderBy:              userData.UserId,
		OrderStartDate:       v.OrderStartDate,
		OrderEndDate:         v.OrderEndDate,
		CustomerId:           v.CustomerId,
		CurrencyId:           1,
	}

	fmt.Printf("Item request of v: %+v\n", q)

	postRequestResponse := functions.PlaceOrder(&c.Controller, q)

	message := "Unable to complete order"

	if postRequestResponse.StatusCode == 200 {
		isSuccess = true

		message = "Order completed successfully"
		resp := responses.OrderResponseDTO{Success: isSuccess, Result: postRequestResponse.Transaction, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		resp := responses.OrderResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// PlaceSalesRequest ...
// @Title Place Sales Request
// @Description Place an order to buy a product
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.SalesRequestDTO	true		"body for Transactions content"
// @Success 201 {object} models.Transactions
// @Failure 403 body is empty
// @router /place-sales-request [post]
func (c *TransactionsController) PlaceSalesRequest() {
	u := c.Ctx.Input.GetData("user")
	userData, err := u.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", u)
	fmt.Printf("Value of v: %+v\n", u)

	if err != false {
		logs.Error("An error occurred ", err)
	}
	var v requests.SalesRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	requestType := "SALES"

	isSuccess := false

	products := []requests.Item{}

	for _, i := range v.Products {
		tProduct := requests.Item{ItemId: i.ProductId, Quantity: i.Quantity}
		products = append(products, tProduct)
	}

	q := requests.PostTransactionRequest{
		Items:                products,
		RequestType:          requestType,
		PaymentProofImageUrl: v.PaymentProofImageUrl,
		PaymentMethodId:      v.PaymentMethodId,
		Comment:              "",
		OrderLocation:        "",
		OrderBy:              userData.UserId,
		OrderStartDate:       "",
		OrderEndDate:         "",
		CustomerId:           v.CustomerId,
		CurrencyId:           1,
	}

	postRequestResponse := functions.PlaceOrder(&c.Controller, q)

	message := "Unable to complete order"

	if postRequestResponse.StatusCode == 200 {
		isSuccess = true

		message = "Order completed successfully"
		resp := responses.OrderResponseDTO{Success: isSuccess, Result: postRequestResponse.Transaction, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		resp := responses.OrderResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
