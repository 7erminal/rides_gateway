package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// PaymentController operations for Payment
type PaymentController struct {
	beego.Controller
}

// URLMapping ...
func (c *PaymentController) URLMapping() {
	c.Mapping("GetAll", c.GetPaymentMethods)
	c.Mapping("UploadPaymentProof", c.UploadPaymentProof)
}

// GetAll ...
// @Title GetAll
// @Description get Payment
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.PaymentMethodsResponseDTO
// @Failure 403
// @router /payment-methods [get]
func (c *PaymentController) GetPaymentMethods() {
	paymentMethodsResp := functions.GetPaymentMethods(&c.Controller)

	isSuccess := false
	message := "An error occurred"
	if paymentMethodsResp.StatusCode == 200 {
		isSuccess = true
		message = "Payment methods fetched successfully"
		resp := responses.PaymentMethodsResponseDTO{Success: isSuccess, Result: paymentMethodsResp.PaymentMethods, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		message = "Failed to fetch payment methods"
		resp := responses.PaymentMethodsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Upload Payment Proof ...
// @Title Update Payment Proof
// @Description Update User's Image
// @Param	Authorization		header 	string true		"header for User"
// @Param	Image		formData 	file	true		"Item Image"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /upload-payment-proof [post]
func (c *PaymentController) UploadPaymentProof() {

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
