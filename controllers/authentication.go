package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// AuthenticationController operations for Authentication
type AuthenticationController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthenticationController) URLMapping() {
	// c.Mapping("VerifyToken", c.VerifyToken)
	c.Mapping("SignIn", c.SignIn)
	c.Mapping("Register", c.Register)
	c.Mapping("ChangePassword", c.ChangePassword)
}

// Register User ...
// @Title Create New User
// @Description Register a User
// @Param	body		body 	requests.Registration	true		"body for Registration content."
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /register [post]
func (c *AuthenticationController) Register() {
	var v requests.Registration
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if v.Token == "" {
		var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: false, Result: nil, StatusDesc: "User role not specified"}

		c.Data["json"] = resp
	} else {
		verifyUserToken := functions.VerifyInviteToken(&c.Controller, v.Token)

		logs.Info("User token verification is ", verifyUserToken)

		if verifyUserToken.StatusCode == 200 {

			logs.Info("Checking role ", verifyUserToken.Value.RoleId)

			userRole := functions.GetRole(&c.Controller, strings.TrimSpace(verifyUserToken.Value.RoleId))

			if userRole.StatusCode == 200 && strings.TrimSpace(verifyUserToken.Value.Email) == v.Email {

				var req requests.RegisterUser = requests.RegisterUser{Email: v.Email, Name: v.FirstName + " | " + v.LastName, Gender: "", PhoneNumber: v.PhoneNumber, Password: v.Password, RoleId: verifyUserToken.Value.RoleId}

				regResp := functions.RegistrationRequest(&c.Controller, req)

				var isSuccess bool = false

				// var data models.UserGateway

				var tkn *string

				if regResp.StatusCode == 200 {
					// go functions.UpdateUserInvite(&c.Controller, idStr, v.Value)
					splitName := strings.Split(regResp.User.FullName, " | ")

					logs.Debug("Name is ", splitName[0])

					// role := models.Role{Role: serRole.Role.Role}

					// data = models.UserGateway{
					// 	// UserId:         regResp.User.UserId,
					// 	// UserType:    regResp.User.UserType,
					// 	FirstName:   splitName[0],
					// 	LastName:    splitName[1],
					// 	Username:    regResp.User.Username,
					// 	Email:       regResp.User.Email,
					// 	PhoneNumber: regResp.User.PhoneNumber,
					// 	Role:        &role,
					// 	// Gender:         regResp.User.Gender,
					// 	// Dob:            regResp.User.Dob,
					// 	// Address:        regResp.User.Address,
					// 	// IdType:         regResp.User.IdType,
					// 	// IdNumber:       regResp.User.IdNumber,
					// 	// Active:         regResp.User.Active,
					// 	// IsVerified:     regResp.User.IsVerified,
					// 	// DateRegistered: regResp.User.DateCreated,
					// }

					isSuccess = true

					var signInReq requests.SignIn = requests.SignIn{Email: v.Email, Password: v.Password}

					loginResp := functions.SignInRequest(&c.Controller, signInReq)

					if loginResp.StatusCode == 200 {
						isSuccess = true
						tkn = &loginResp.Value
					}
				}

				var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: tkn, StatusDesc: regResp.StatusDesc}

				c.Data["json"] = resp

			} else {
				var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: false, Result: nil, StatusDesc: "User verification failed"}

				c.Data["json"] = resp
			}
		} else {
			var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: false, Result: nil, StatusDesc: "User role not specified"}

			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// Sign In ...
// @Title SignIn
// @Description sign user in
// @Param	body		body 	requests.SignIn	true		"body for Authentication content"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /sign-in [post]
func (c *AuthenticationController) SignIn() {
	var v requests.SignIn
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Info("Received ", v.Password, v.Email)

	loginResp := functions.SignInRequest(&c.Controller, v)

	// var data models.UserGateway

	var isSuccess bool = false
	var tkn *string

	if loginResp.StatusCode == 200 {
		// splitName := strings.Split(loginResp.Result.FullName, " | ")

		// logs.Debug("Name is ", splitName[0])
		// data = models.UserGateway{
		// 	UserId:         loginResp.Result.UserId,
		// 	UserType:       loginResp.Result.UserType,
		// 	FirstName:      splitName[0],
		// 	LastName:       splitName[1],
		// 	Username:       loginResp.Result.Username,
		// 	Email:          loginResp.Result.Email,
		// 	PhoneNumber:    loginResp.Result.PhoneNumber,
		// 	Gender:         loginResp.Result.Gender,
		// 	Dob:            loginResp.Result.Dob,
		// 	Address:        loginResp.Result.Address,
		// 	IdType:         loginResp.Result.IdType,
		// 	IdNumber:       loginResp.Result.IdNumber,
		// 	Active:         loginResp.Result.Active,
		// 	IsVerified:     loginResp.Result.IsVerified,
		// 	DateRegistered: loginResp.Result.DateCreated,
		// }

		isSuccess = true
		tkn = &loginResp.Value
	}

	var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: tkn, StatusDesc: loginResp.StatusDesc}

	c.Data["json"] = resp

	c.ServeJSON()
}

// Change Password ...
// @Title Change Password
// @Description Change user password
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.ChangePassword	true		"body for Authentication content"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /change-password [post]
func (c *AuthenticationController) ChangePassword() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			var v requests.ChangePassword
			json.Unmarshal(c.Ctx.Input.RequestBody, &v)

			logs.Info("Received ", v.OldPassword, v.NewPassword)

			idStr := strconv.FormatInt(verifyToken.User.UserId, 10)

			loginResp := functions.ChangePassword(&c.Controller, idStr, v)

			// var data models.UserGateway

			var isSuccess bool = false
			var tkn *string

			if loginResp.StatusCode == 200 {

				isSuccess = true
				tkn = &loginResp.Value
			}

			var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: tkn, StatusDesc: loginResp.StatusDesc}

			c.Data["json"] = resp
		} else {
			logs.Error("Unable to verify user")
			var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Unable to verify user")
		var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// // VerifyToken ...
// // @Title VerifyToken
// // @Description verify user token
// // @Param	body		body 	requests.StringRequestDTO	true		"body for Authentication content"
// // @Success 200 {object} responses.StringResponseDTO
// // @Failure 403 body is empty
// // @router /verify-token [post]
func (c *AuthenticationController) VerifyToken() {
	var v requests.StringRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Info("Received ", v.Value)

	loginResp := functions.VerifyToken(&c.Controller, v.Value)

	var isSuccess bool = false
	var tkn *string

	if loginResp.StatusCode == 200 {
		isSuccess = true
		tkn = &v.Value
	}

	var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: tkn, StatusDesc: loginResp.StatusDesc}

	c.Data["json"] = resp

	c.ServeJSON()
}
