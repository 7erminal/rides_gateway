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

func GetUserDetails(c *beego.Controller, userid int64) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting user details ", userid)

	request := api.NewRequest(
		host,
		"/v1/users/"+strconv.FormatInt(userid, 10),
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
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User)

	return data
}

func InviteUser(c *beego.Controller, email string, role int64, link string, inviteby int64) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending email ", email)

	request := api.NewRequest(
		host,
		"/v1/users/invite-user/",
		api.POST)
	request.InterfaceParams["Email"] = email
	request.InterfaceParams["Link"] = link
	request.InterfaceParams["Role"] = strconv.FormatInt(role, 10)
	request.InterfaceParams["InviteBy"] = strconv.FormatInt(inviteby, 10)
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

	logs.Info("Resp is ", data)

	return data
}

func GetUsers(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.UsersOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/users",
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
	var data responses.UsersOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Users)

	return data
}

func GetUsersWithRole(c *beego.Controller, role_id string, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.UsersOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/users/role/"+role_id,
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
	var data responses.UsersOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Users)

	return data
}

func GetUsersWithBranch(c *beego.Controller, branch_id string, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.UsersOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting users ")

	request := api.NewRequest(
		host,
		"/v1/users/branch/"+branch_id,
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
	var data responses.UsersOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data.Users)

	return data
}

func GetRoles(c *beego.Controller) (resp responses.RolesAllResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting roles ")

	request := api.NewRequest(
		host,
		"/v1/roles",
		api.GET)
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
	var data responses.RolesAllResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetInvites(c *beego.Controller) (resp responses.UserInvitesResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting roles ")

	request := api.NewRequest(
		host,
		"/v1/users/get-user-invites",
		api.GET)
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
	var data responses.UserInvitesResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetRole(c *beego.Controller, role string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting roles ")

	request := api.NewRequest(
		host,
		"/v1/roles/"+role,
		api.GET)
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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetRoleWithRoleName(c *beego.Controller, role string) (resp responses.RoleResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting roles ")

	request := api.NewRequest(
		host,
		"/v1/roles/role/"+role,
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
	var data responses.RoleResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetInvite(c *beego.Controller, token string) (resp responses.UserInviteResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Getting user token ", token)

	request := api.NewRequest(
		host,
		"/v1/users/user-invite/"+token,
		api.GET)
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
	var data responses.UserInviteResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateUser(c *beego.Controller, id string, req requests.UpdateUserRequestDTO) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.FirstName)

	request := api.NewRequest(
		host,
		"/v1/users/"+id,
		api.PUT)
	request.InterfaceParams["FullName"] = req.FirstName + " | " + req.LastName
	request.InterfaceParams["Address"] = req.Address
	request.InterfaceParams["Dob"] = req.Dob
	request.InterfaceParams["Gender"] = req.Gender
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Username"] = req.Username
	request.InterfaceParams["MaritalStatus"] = ""
	request.InterfaceParams["BranchId"] = req.BranchId
	request.InterfaceParams["ImagePath"] = ""
	request.InterfaceParams["RoleId"] = req.RoleId
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

func UpdateUserRole(c *beego.Controller, id string, req requests.UpdateUserRoleRequestDTO) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending first name ", req.RoleId)

	request := api.NewRequest(
		host,
		"/v1/users/role/"+id,
		api.PUT)
	request.InterfaceParams["RoleId"] = req.RoleId
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

func UpdateUserBranch(c *beego.Controller, id string, req requests.UpdateUserBranchRequestDTO) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending branch ID ", req.BranchId)

	request := api.NewRequest(
		host,
		"/v1/users/branch/"+id,
		api.PUT)
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

func UpdateUserImage(c *beego.Controller, userImage string, userId int64) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Sending file ", userImage)

	request := api.NewRequest(
		host,
		"/v1/users/update-user-image/",
		api.POST)

	request.FileField["UserImage"] = userImage
	request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateUserInvite(c *beego.Controller, id string, status string) (resp responses.UserInviteResponseDTO) {
	host, _ := beego.AppConfig.String("customerBaseUrl")

	logs.Info("Updating invite ", status)

	request := api.NewRequest(
		host,
		"/v1/users/update-user-invite/"+id,
		api.PUT)
	request.InterfaceParams["Status"] = status

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
	var data responses.UserInviteResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	logs.Info("Resp is ", data)

	return data
}

func LogOut(c *beego.Controller, userToken string) (resp responses.UserOriResponseDTO) {
	host, _ := beego.AppConfig.String("authenticationBaseUrl")

	logs.Info("token", userToken)

	request := api.NewRequest(
		host,
		"/v1/auth/log-out/",
		api.POST)

	request.InterfaceParams["Token"] = userToken
	// request.HeaderField["content-type"] = "multipart/form-data"
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
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
	var data responses.UserOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
