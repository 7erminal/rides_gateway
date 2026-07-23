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

func SignInRequest(c *beego.Controller, req requests.SignIn) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending email ", req.Email)
	logs.Info("Sending password ", req.Password)

	request := api.NewRequest(
		host,
		"/v1/auth/login/token",
		api.POST)
	request.InterfaceParams["Username"] = req.Email
	request.InterfaceParams["Password"] = req.Password
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func SendOTP(c *beego.Controller, phoneNumber string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending OTP to ", phoneNumber)

	request := api.NewRequest(
		host,
		"/v1/auth/send-activation-code",
		api.POST)
	request.InterfaceParams["MobileNumber"] = phoneNumber
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyOTP(c *beego.Controller, req requests.SignIn) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending email ", req.Email)
	logs.Info("Sending password ", req.Password)

	request := api.NewRequest(
		host,
		"/v1/auth/verify-activation-code",
		api.POST)
	request.InterfaceParams["Username"] = req.Email
	request.InterfaceParams["Password"] = req.Password
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func SendActivationCode(c *beego.Controller, phoneNumber string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending OTP to ", phoneNumber)

	request := api.NewRequest(
		host,
		"/v1/auth/send-activation-code",
		api.POST)
	request.InterfaceParams["MobileNumber"] = phoneNumber
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyActivationCode(c *beego.Controller, number string, code string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending number ", number)
	logs.Info("Sending password ", code)

	request := api.NewRequest(
		host,
		"/v1/auth/verify-activation-code",
		api.POST)
	request.InterfaceParams["MobileNumber"] = number
	request.InterfaceParams["Password"] = code
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func ChangePassword(c *beego.Controller, userid string, req requests.ChangePassword) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("Sending old password ", req.OldPassword)
	logs.Info("Sending new password ", req.NewPassword)

	request := api.NewRequest(
		host,
		"/v1/auth/change-password/"+userid,
		api.PUT)
	request.InterfaceParams["OldPassword"] = req.OldPassword
	request.InterfaceParams["NewPassword"] = req.NewPassword
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
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyToken(c *beego.Controller, token string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("About to verify token ", token)

	request := api.NewRequest(
		host,
		"/v1/auth/token/check",
		api.POST)
	request.InterfaceParams["Value"] = token
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
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}

func VerifyTokenNew(token string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("About to verify token ", token)

	request := api.NewRequest(
		host,
		"/v1/auth/token/check",
		api.POST)
	request.InterfaceParams["Value"] = token
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		// c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		// c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)

	return data
}

func RegistrationRequest(c *beego.Controller, req requests.RegisterUser) (resp responses.UserResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending email ", req.Email)

	request := api.NewRequest(
		host,
		"/v1/users/sign-up",
		api.POST)
	request.InterfaceParams["Email"] = req.Email
	request.InterfaceParams["Password"] = req.Password
	request.InterfaceParams["Name"] = req.Name
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Role"] = req.RoleId
	request.InterfaceParams["RoleRequired"] = true

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
	var data responses.UserResponseDTO
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

func VerifyInviteToken(c *beego.Controller, token string) (resp responses.InviteDecodeResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("About to verify token ", token)

	request := api.NewRequest(
		host,
		"/v1/users/verify-invite",
		api.POST)
	request.InterfaceParams["Value"] = token
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
	var data responses.InviteDecodeResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	return data
}
