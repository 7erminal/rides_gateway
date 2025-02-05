package functions

import (
	"AMC_gateway/api"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"io"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func PlaceOrder(c *beego.Controller, req requests.PostTransactionRequest) (resp responses.OrderOriResponseDTO) {
	host, _ := beego.AppConfig.String("transactionsBaseUrl")

	logs.Info("Sending currency ", req.CurrencyId)

	request := api.NewRequest(
		host,
		"/v1/orders/",
		api.POST)
	request.InterfaceParams["Currency"] = req.CurrencyId
	request.InterfaceParams["Items"] = req.Items
	request.InterfaceParams["RequestType"] = req.RequestType
	request.InterfaceParams["Comment"] = req.Comment
	request.InterfaceParams["CreatedBy"] = req.OrderBy
	request.InterfaceParams["OrderEndDate"] = req.OrderEndDate
	request.InterfaceParams["OrderStartDate"] = req.OrderStartDate
	request.InterfaceParams["OrderLocation"] = req.OrderLocation
	request.InterfaceParams["Customer"] = req.CustomerId
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.OrderOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func UploadPaymentProof(c *beego.Controller, paymentProofImage string) (resp responses.ItemImageOriResponseDTO) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	logs.Info("Sending file ", paymentProofImage)

	request := api.NewRequest(
		host,
		"/v1/payments/upload-payment-proof",
		api.POST)

	request.FileField["Image"] = paymentProofImage
	// request.HeaderField["content-type"] = "multipart/form-data"
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}

	// client.Request.HeaderField["content-type"] = "multipart/form-data"
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.ItemImageOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
