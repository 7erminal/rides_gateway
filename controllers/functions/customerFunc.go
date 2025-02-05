package functions

import (
	"AMC_gateway/api"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"io"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func AddCustomer(c *beego.Controller, req requests.AddCustomer, addedBy string, custType string) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending email ", req.Email)

	// Get date
	now := time.Now()
	y, m, d := now.Date()
	d_str := strconv.Itoa(d)
	m_str := strconv.Itoa(int(m))
	if len(d_str) < 2 {
		d_str = "0" + d_str
	}
	if len(m_str) < 2 {
		m_str = "0" + m_str
	}
	dob := strconv.Itoa(y) + "/" + m_str + "/" + d_str

	request := api.NewRequest(
		host,
		"/v1/customers/add-customer",
		api.POST)
	request.InterfaceParams["Name"] = req.Email
	request.InterfaceParams["Email"] = req.Email
	request.InterfaceParams["IdType"] = req.IdType
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["IdNumber"] = req.IdNumber
	request.InterfaceParams["Dob"] = dob
	request.InterfaceParams["AddedBy"] = addedBy
	request.InterfaceParams["Location"] = req.Location
	request.FileField["CustomerImage"] = req.ImagePath
	request.InterfaceParams["Category"] = custType
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
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data["access_token"])
	// logs.Info("Expires in ", data["expires_in"])
	// logs.Info("Scope is ", data["scope"])
	// logs.Info("Token Type is ", data["token_type"])
	// logs.Info("Response received ", c.Data["json"])
	// logs.Info("Access token ", data.Access_token)
	// logs.Info("Expires in ", data.Expires_in)
	// logs.Info("Scope is ", data.Scope)
	// logs.Info("Token Type is ", data.Token_type)

	return data
}

func GetCustomerDetails(c *beego.Controller, userid int64) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting user details ", userid)

	request := api.NewRequest(
		host,
		"/v1/customers/"+strconv.FormatInt(userid, 10),
		api.GET)
	// request.Params["username"] = username
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
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func GetCustomers(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.CustomersResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/customers",
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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
	var data responses.CustomersResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Customers)

	return data
}

func UpdateCustomer(c *beego.Controller, id string, req requests.UpdateCustomer) (resp responses.CustomerResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.Name)

	request := api.NewRequest(
		host,
		"/v1/customers/"+id,
		api.PUT)
	request.InterfaceParams["Name"] = req.Email
	request.InterfaceParams["Email"] = req.Email
	request.InterfaceParams["IdType"] = req.IdType
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["IdNumber"] = req.IdNumber
	// request.InterfaceParams["Dob"] = req.IdNumber
	// request.InterfaceParams["AddedBy"] = req.IdNumber
	request.InterfaceParams["Location"] = req.Location
	request.FileField["ImagePath"] = req.ImagePath
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
	var data responses.CustomerResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func DeleteCustomer(c *beego.Controller, id string, req requests.UpdateCustomer) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.Name)

	request := api.NewRequest(
		host,
		"/v1/customers/"+id,
		api.DELETE)
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data
}

func GetIdTypes(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.IDTypeResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting Id Types ")

	request := api.NewRequest(
		host,
		"/v1/id-types",
		api.GET)
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	request.Params["query"] = query
	request.Params["fields"] = fields
	request.Params["sortby"] = sortby
	request.Params["order"] = order
	request.Params["offset"] = offset
	request.Params["limit"] = limit
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
	var data responses.IDTypeResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
