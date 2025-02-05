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

func AddAuditTrail(c *beego.Controller, id string, req requests.UpdateUserRequestDTO) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("auditTrailBaseUrl")

	logs.Info("Sending first name ", req.FirstName)

	request := api.NewRequest(
		host,
		"/v1/transactions/"+id,
		api.PUT)
	request.InterfaceParams["FullName"] = req.FirstName + " | " + req.LastName
	request.InterfaceParams["Address"] = req.Address
	request.InterfaceParams["Dob"] = req.Dob
	request.InterfaceParams["Gender"] = req.Gender
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Username"] = req.Username
	request.InterfaceParams["MaritalStatus"] = ""
	request.InterfaceParams["BranchId"] = req.BranchId
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
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}
