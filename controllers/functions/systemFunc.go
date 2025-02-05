package functions

import (
	"AMC_gateway/api"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"io"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func AddBranch(c *beego.Controller, req requests.BranchRequestDTO, addedBy int64) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Sending user name ", strconv.FormatInt(addedBy, 10))

	request := api.NewRequest(
		host,
		"/v1/branches/",
		api.POST)
	request.InterfaceParams["Branch"] = req.Branch
	request.InterfaceParams["CountryCode"] = req.CountryCode
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Location"] = req.Location
	request.InterfaceParams["AddedBy"] = strconv.FormatInt(addedBy, 10)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
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
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateBranch(c *beego.Controller, req requests.BranchRequestDTO, addedBy int64, branchId string) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Sending user name ", strconv.FormatInt(addedBy, 10))

	request := api.NewRequest(
		host,
		"/v1/branches/"+branchId,
		api.PUT)
	request.InterfaceParams["Branch"] = req.Branch
	request.InterfaceParams["CountryCode"] = req.CountryCode
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Location"] = req.Location
	request.InterfaceParams["AddedBy"] = strconv.FormatInt(addedBy, 10)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
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
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetBranch(c *beego.Controller, branchid int64) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Getting branch details for ", branchid)

	request := api.NewRequest(
		host,
		"/v1/branches/"+strconv.FormatInt(branchid, 10),
		api.GET)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
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
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func DeleteBranch(c *beego.Controller, branchid string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Getting branch details for ", branchid)

	request := api.NewRequest(
		host,
		"/v1/branches/"+branchid,
		api.DELETE)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetBranches(c *beego.Controller) (resp responses.BranchesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/branches/",
		api.GET)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
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
	var data responses.BranchesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateBranchBranchManger(c *beego.Controller, userid string, branchid string) (resp responses.BranchesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/branches/branch-manager/"+branchid,
		api.PUT)
	request.InterfaceParams["BranchManager"] = userid
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
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
	var data responses.BranchesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCountries(c *beego.Controller) (resp responses.CountriesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/countries/",
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
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
	var data responses.CountriesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
