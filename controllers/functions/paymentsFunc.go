package functions

import (
	"AMC_gateway/api"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"io"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func GetPaymentMethods(c *beego.Controller) (resp responses.PaymentMethodsOriResponseDTO) {
	host, _ := beego.AppConfig.String("paymentBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/payment-methods/",
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
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
	var data responses.PaymentMethodsOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}
