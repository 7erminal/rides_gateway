package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// RegistrationController operations for Registration
type RegistrationController struct {
	beego.Controller
}

// URLMapping ...
func (c *RegistrationController) URLMapping() {
	c.Mapping("RegisterAlt", c.RegisterAlt)
}

// Post ...
// @Title Create
// @Description Register a User
// @Param	body		body 	requests.Registration	true		"body for Registration content. Role should be the role ID"
// @Success 200 {object} responses.UserGatewayResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *RegistrationController) RegisterAlt() {
	var v requests.Registration
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if v.Token == "" {
		var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: false, Result: nil, StatusDesc: "User role not specified"}

		c.Data["json"] = resp
	} else {
		userRole := functions.GetRole(&c.Controller, v.Token)

		if userRole.StatusCode == 200 {

			var req requests.RegisterUser = requests.RegisterUser{Email: v.Email, Name: v.FirstName + " | " + v.LastName, Gender: "", PhoneNumber: v.PhoneNumber, Password: v.Password, RoleId: v.Token}

			regResp := functions.RegistrationRequest(&c.Controller, req)

			var isSuccess bool = false

			var data responses.UserGateway

			if regResp.StatusCode == 200 {
				splitName := strings.Split(regResp.User.FullName, " | ")

				logs.Debug("Name is ", splitName[0])

				role := responses.Role{Role: userRole.Role.Role}

				data = responses.UserGateway{
					// UserId:         regResp.User.UserId,
					// UserType:    regResp.User.UserType,
					FirstName:   splitName[0],
					LastName:    splitName[1],
					Username:    regResp.User.Username,
					Email:       regResp.User.Email,
					PhoneNumber: regResp.User.PhoneNumber,
					Role:        &role,
					// Gender:         regResp.User.Gender,
					// Dob:            regResp.User.Dob,
					// Address:        regResp.User.Address,
					// IdType:         regResp.User.IdType,
					// IdNumber:       regResp.User.IdNumber,
					// Active:         regResp.User.Active,
					// IsVerified:     regResp.User.IsVerified,
					// DateRegistered: regResp.User.DateCreated,
				}

				isSuccess = true
			}

			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: &data, StatusDesc: regResp.StatusDesc}

			c.Data["json"] = resp
		} else {
			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: false, Result: nil, StatusDesc: "User role not specified"}

			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}
