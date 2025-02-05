package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/requests"
	"AMC_gateway/structs/responses"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// UserManagementController operations for UserManagement
type UserManagementController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserManagementController) URLMapping() {
	c.Mapping("GetUser", c.GetUser)
	c.Mapping("GetUserWithId", c.GetUserWithId)
	c.Mapping("InviteUserReg", c.InviteUserReg)
	c.Mapping("GetRoles", c.GetRoles)
	c.Mapping("UpdateUserImage", c.UpdateUserImage)
	c.Mapping("GetUsers", c.GetUsers)
	c.Mapping("GetBranchManagers", c.GetBranchManagers)
	c.Mapping("GetUsersUnderBranch", c.GetUsersUnderBranch)
	c.Mapping("VerifyInvite", c.VerifyInvite)
	c.Mapping("UpdateUser", c.UpdateUser)
	c.Mapping("UpdateUserRole", c.UpdateUserRole)
	c.Mapping("UpdateUserBranch", c.UpdateUserBranch)
}

// GetUserSession ...
// @Title Get User Session
// @Description Get user
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.UserGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-user-session [post]
func (c *UserManagementController) GetUser() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")
	var isSuccess bool = false
	logs.Info("Token name is ", token[0])
	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			userResp := functions.GetUserDetails(&c.Controller, verifyToken.User.UserId)

			var data responses.UserGateway

			if userResp.StatusCode == 200 {
				logs.Info("Name returned: ", userResp.User.FullName)
				splitName := strings.Split(userResp.User.FullName, " | ")

				// logs.Debug("Name is ", splitName[0])
				firstname := ""
				lastname := ""
				if len(splitName) > 1 {
					firstname = splitName[0]
					lastname = splitName[1]
				} else {
					firstname = splitName[0]
				}

				status := "ACTIVE"

				if userResp.User.Active == 6 {
					status = "DELETED"
				}
				if userResp.User.Active == 2 {
					status = "PENDING"
				}
				if userResp.User.Active == 4 {
					status = "INACTIVE"
				}
				// branch := &responses.BranchResp{}

				// if userResp.User.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: userResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: userResp.User.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: userResp.User.Branch.Country.Country, CountryCode: userResp.User.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: userResp.User.Branch.BranchId, Branch: userResp.User.Branch.Branch, Country: &country, Location: userResp.User.Branch.Location, PhoneNumber: userResp.User.Branch.PhoneNumber}
				// 	// userResp.User.Customer.Branch.Country = country
				// 	} else {
				// 	branch = nil
				// }
				data = responses.UserGateway{
					UserId: userResp.User.UserId,
					// UserType:    userResp.User.UserType,
					FirstName:   firstname,
					LastName:    lastname,
					Username:    userResp.User.Username,
					Email:       userResp.User.Email,
					PhoneNumber: userResp.User.PhoneNumber,
					Role:        userResp.User.Role,
					Customer:    userResp.User.UserDetails,
					ImagePath:   userResp.User.ImagePath,
					Status:      status,
					// Gender:         userResp.User.Gender,
					// Dob:            userResp.User.Dob,
					// Address:        userResp.User.Address,
					// IdType:         userResp.User.IdType,
					// IdNumber:       userResp.User.IdNumber,
					// Active:         userResp.User.Active,
					// IsVerified:     userResp.User.IsVerified,
					// DateRegistered: userResp.User.DateCreated,
				}

				isSuccess = true

				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: &data, StatusDesc: userResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetUserWithID ...
// @Title Get User with ID
// @Description Get user using ID
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.UserGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-user-with-id/:id [get]
func (c *UserManagementController) GetUserWithId() {
	authorization := c.Ctx.Input.Header("Authorization")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	token := strings.Split(authorization, " ")
	var isSuccess bool = false
	logs.Info("Token name is ", token[0])
	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			userResp := functions.GetUserDetails(&c.Controller, id)

			var data responses.UserGateway

			if userResp.StatusCode == 200 {
				logs.Info("Name returned: ", userResp.User.FullName)
				splitName := strings.Split(userResp.User.FullName, " | ")

				// logs.Debug("Name is ", splitName[0])
				firstname := ""
				lastname := ""
				if len(splitName) > 1 {
					firstname = splitName[0]
					lastname = splitName[1]
				} else {
					firstname = splitName[0]
				}
				// branch := &responses.BranchResp{}

				// if userResp.User.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: userResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: userResp.User.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: userResp.User.Branch.Country.Country, CountryCode: userResp.User.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: userResp.User.Branch.BranchId, Branch: userResp.User.Branch.Branch, Country: &country, Location: userResp.User.Branch.Location, PhoneNumber: userResp.User.Branch.PhoneNumber}
				// 	// userResp.User.Customer.Branch.Country = country
				// 	} else {
				// 	branch = nil
				// }
				status := "ACTIVE"

				if userResp.User.Active == 6 {
					status = "DELETED"
				}
				if userResp.User.Active == 2 {
					status = "PENDING"
				}
				if userResp.User.Active == 4 {
					status = "INACTIVE"
				}

				data = responses.UserGateway{
					UserId: userResp.User.UserId,
					// UserType:    userResp.User.UserType,
					FirstName:   firstname,
					LastName:    lastname,
					Username:    userResp.User.Username,
					Email:       userResp.User.Email,
					PhoneNumber: userResp.User.PhoneNumber,
					Role:        userResp.User.Role,
					Customer:    userResp.User.UserDetails,
					ImagePath:   userResp.User.ImagePath,
					Status:      status,
					// Gender:         userResp.User.Gender,
					// Dob:            userResp.User.Dob,
					// Address:        userResp.User.Address,
					// IdType:         userResp.User.IdType,
					// IdNumber:       userResp.User.IdNumber,
					// Active:         userResp.User.Active,
					// IsVerified:     userResp.User.IsVerified,
					// DateRegistered: userResp.User.DateCreated,
				}

				isSuccess = true

				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: &data, StatusDesc: userResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Error getting user details"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to access this resource"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Invalid authorization token"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetUsers ...
// @Title Get Users
// @Description Get all users
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.UsersGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-users [get]
func (c *UserManagementController) GetUsers() {
	v := c.Ctx.Input.GetData("user")
	userData, err := v.(*responses.UsersOri)

	fmt.Printf("Type of v: %T\n", v)
	fmt.Printf("Value of v: %+v\n", v)

	logs.Info("Error is ", err)

	logs.Info("User received is ", v, " AND ", userData)

	var isSuccess bool = false

	var fields string
	var sortby string
	var order string
	var query string
	var limit string
	var offset string

	// limit: 10 (default is 10)
	if v := c.GetString("limit"); v != "" {
		limit = v
	}
	// offset: 0 (default is 0)
	if v := c.GetString("offset"); v != "" {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = v
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = v
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		query = v
	}

	managingDirectorUserRole, _ := beego.AppConfig.String("managingDirectorRoleName")
	superAdminRole, _ := beego.AppConfig.String("superAdminRoleName")
	usersResp := responses.UsersOriResponseDTO{}
	hasRole := false
	if userData.Role != nil {
		if userData.Role.Role == superAdminRole || userData.Role.Role == managingDirectorUserRole {
			hasRole = true
			usersResp = functions.GetUsers(&c.Controller, query, fields, sortby, order, offset, limit)
		} else {
			logs.Info("User ", userData.UserId, " branch ID is ", userData.UserDetails)
			branchStr := strconv.FormatInt(userData.UserDetails.Branch.BranchId, 10)
			usersResp = functions.GetUsersWithBranch(&c.Controller, branchStr, query, fields, sortby, order, offset, limit)
		}
	}

	if !hasRole {
		var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "User does not have a role"}
		c.Data["json"] = resp
		c.ServeJSON()
	}

	if usersResp.StatusCode == 200 {

		// branch := &responses.BranchResp{}

		// if userResp.User.Branch != nil {
		// 	currency := responses.CurrencyResp{Symbol: userResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: userResp.User.Branch.Country.DefaultCurrency.Currency}
		// 	country := responses.CountryResp{Country: userResp.User.Branch.Country.Country, CountryCode: userResp.User.Branch.Country.CountryCode, Currency: &currency}
		// 	branch = &responses.BranchResp{BranchId: userResp.User.Branch.BranchId, Branch: userResp.User.Branch.Branch, Country: &country, Location: userResp.User.Branch.Location, PhoneNumber: userResp.User.Branch.PhoneNumber}
		// 	// userResp.User.Customer.Branch.Country = country
		// 	} else {
		// 	branch = nil
		// }
		users := []responses.UserGateway{}
		for _, user := range *usersResp.Users {
			logs.Info("Fullname is ", user.FullName)
			splitName := strings.Split(user.FullName, " | ")
			firstname := ""
			lastname := ""
			if len(splitName) > 1 {
				firstname = splitName[0]
				lastname = splitName[1]
			} else {
				firstname = splitName[0]
			}

			logs.Info("User Role is ", user.Role)

			status := "ACTIVE"

			if user.Active == 1 {
				status = "ACTIVE"
			}

			if user.Active == 6 {
				status = "INACTIVE"
			}
			if user.Active == 2 {
				status = "PENDING"
			}
			if user.Active == 4 {
				status = "INACTIVE"
			}

			data := responses.UserGateway{
				UserId: user.UserId,
				// UserType:    userResp.User.UserType,
				FirstName:   firstname,
				LastName:    lastname,
				Username:    user.Username,
				Email:       user.Email,
				PhoneNumber: user.PhoneNumber,
				Role:        user.Role,
				Customer:    user.UserDetails,
				ImagePath:   user.ImagePath,
				Status:      status,
				// Gender:         userResp.User.Gender,
				// Dob:            userResp.User.Dob,
				// Address:        userResp.User.Address,
				// IdType:         userResp.User.IdType,
				// IdNumber:       userResp.User.IdNumber,
				// Active:         userResp.User.Active,
				// IsVerified:     userResp.User.IsVerified,
				DateRegistered: user.DateCreated,
			}

			users = append(users, data)
		}

		isSuccess = true

		var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: &users, StatusDesc: usersResp.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetBranchManagers ...
// @Title Get Branch Managers
// @Description Get all branch managers users
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.UsersGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-branch-managers [get]
func (c *UserManagementController) GetBranchManagers() {

	// Retrieve the user data from the context
	userData := c.Ctx.Input.GetData("user")

	logs.Debug("User data is ", userData)

	var isSuccess bool = false
	logs.Info("about to fetch Managing director")
	// role_name := "Customer Service Maker"
	// role_name := "Branch Manager"
	role_name, _ := beego.AppConfig.String("branchManagerRoleName")
	role := functions.GetRoleWithRoleName(&c.Controller, role_name)
	if role.StatusCode == 200 {
		var fields string
		var sortby string
		var order string
		var query string
		var limit string
		var offset string

		// limit: 10 (default is 10)
		if v := c.GetString("limit"); v != "" {
			limit = v
		}
		// offset: 0 (default is 0)
		if v := c.GetString("offset"); v != "" {
			offset = v
		}
		// sortby: col1,col2
		if v := c.GetString("sortby"); v != "" {
			sortby = v
		}
		// order: desc,asc
		if v := c.GetString("order"); v != "" {
			order = v
		}
		// query: k:v,k:v
		if v := c.GetString("query"); v != "" {
			query = v
		}
		role_idStr := strconv.FormatInt(role.Role.RoleId, 10)
		usersResp := functions.GetUsersWithRole(&c.Controller, role_idStr, query, fields, sortby, order, offset, limit)

		if usersResp.StatusCode == 200 {

			// branch := &responses.BranchResp{}

			// if userResp.User.Branch != nil {
			// 	currency := responses.CurrencyResp{Symbol: userResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: userResp.User.Branch.Country.DefaultCurrency.Currency}
			// 	country := responses.CountryResp{Country: userResp.User.Branch.Country.Country, CountryCode: userResp.User.Branch.Country.CountryCode, Currency: &currency}
			// 	branch = &responses.BranchResp{BranchId: userResp.User.Branch.BranchId, Branch: userResp.User.Branch.Branch, Country: &country, Location: userResp.User.Branch.Location, PhoneNumber: userResp.User.Branch.PhoneNumber}
			// 	// userResp.User.Customer.Branch.Country = country
			// 	} else {
			// 	branch = nil
			// }
			users := []responses.UserGateway{}
			for _, user := range *usersResp.Users {
				logs.Info("Fullname is ", user.FullName)
				splitName := strings.Split(user.FullName, " | ")
				firstname := ""
				lastname := ""
				if len(splitName) > 1 {
					firstname = splitName[0]
					lastname = splitName[1]
				} else {
					firstname = splitName[0]
				}

				logs.Info("User Role is ", user.Role)

				status := "ACTIVE"

				if user.Active == 6 {
					status = "DELETED"
				}
				if user.Active == 2 {
					status = "PENDING"
				}
				if user.Active == 4 {
					status = "INACTIVE"
				}

				data := responses.UserGateway{
					UserId: user.UserId,
					// UserType:    userResp.User.UserType,
					FirstName:   firstname,
					LastName:    lastname,
					Username:    user.Username,
					Email:       user.Email,
					PhoneNumber: user.PhoneNumber,
					Role:        user.Role,
					Customer:    user.UserDetails,
					ImagePath:   user.ImagePath,
					Status:      status,
					// Gender:         userResp.User.Gender,
					// Dob:            userResp.User.Dob,
					// Address:        userResp.User.Address,
					// IdType:         userResp.User.IdType,
					// IdNumber:       userResp.User.IdNumber,
					// Active:         userResp.User.Active,
					// IsVerified:     userResp.User.IsVerified,
					// DateRegistered: userResp.User.DateCreated,
				}

				users = append(users, data)
			}

			isSuccess = true

			var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: &users, StatusDesc: usersResp.StatusDesc}
			c.Data["json"] = resp
		} else {
			logs.Error("Unable to get users using role")
			var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred."}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Unable to get role ")
		var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Given Role not found."}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetUsersUnderBranch ...
// @Title Get Users Under Branch
// @Description Get all users under specified branch
// @Param	Authorization		header 	string true		"header for User"
// @Param	branch_id		path 	string	true		"The key for staticblock"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.UsersGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-users-under-branch/:branch_id [get]
func (c *UserManagementController) GetUsersUnderBranch() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")
	var isSuccess bool = false
	logs.Info("Token name is ", token[0])
	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			var fields string
			var sortby string
			var order string
			var query string
			var limit string
			var offset string

			// limit: 10 (default is 10)
			if v := c.GetString("limit"); v != "" {
				limit = v
			}
			// offset: 0 (default is 0)
			if v := c.GetString("offset"); v != "" {
				offset = v
			}
			// sortby: col1,col2
			if v := c.GetString("sortby"); v != "" {
				sortby = v
			}
			// order: desc,asc
			if v := c.GetString("order"); v != "" {
				order = v
			}
			// query: k:v,k:v
			if v := c.GetString("query"); v != "" {
				query = v
			}
			branch_idStr := c.Ctx.Input.Param(":branch_id")
			logs.Info("about to fetch Users under specified branch")
			usersResp := functions.GetUsersWithBranch(&c.Controller, branch_idStr, query, fields, sortby, order, offset, limit)
			if usersResp.StatusCode == 200 {

				// branch := &responses.BranchResp{}

				// if userResp.User.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: userResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: userResp.User.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: userResp.User.Branch.Country.Country, CountryCode: userResp.User.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: userResp.User.Branch.BranchId, Branch: userResp.User.Branch.Branch, Country: &country, Location: userResp.User.Branch.Location, PhoneNumber: userResp.User.Branch.PhoneNumber}
				// 	// userResp.User.Customer.Branch.Country = country
				// 	} else {
				// 	branch = nil
				// }
				users := []responses.UserGateway{}
				if usersResp.Users != nil {
					for _, user := range *usersResp.Users {
						logs.Info("Fullname is ", user.FullName)
						splitName := strings.Split(user.FullName, " | ")
						firstname := ""
						lastname := ""
						if len(splitName) > 1 {
							firstname = splitName[0]
							lastname = splitName[1]
						} else {
							firstname = splitName[0]
						}

						logs.Info("User Role is ", user.Role)

						status := "ACTIVE"

						if user.Active == 6 {
							status = "DELETED"
						}
						if user.Active == 2 {
							status = "PENDING"
						}
						if user.Active == 4 {
							status = "INACTIVE"
						}

						data := responses.UserGateway{
							UserId: user.UserId,
							// UserType:    userResp.User.UserType,
							FirstName:   firstname,
							LastName:    lastname,
							Username:    user.Username,
							Email:       user.Email,
							PhoneNumber: user.PhoneNumber,
							Role:        user.Role,
							Customer:    user.UserDetails,
							ImagePath:   user.ImagePath,
							Status:      status,
							// Gender:         userResp.User.Gender,
							// Dob:            userResp.User.Dob,
							// Address:        userResp.User.Address,
							// IdType:         userResp.User.IdType,
							// IdNumber:       userResp.User.IdNumber,
							// Active:         userResp.User.Active,
							// IsVerified:     userResp.User.IsVerified,
							// DateRegistered: userResp.User.DateCreated,
						}

						users = append(users, data)
					}
				} else {
					users = []responses.UserGateway{}
				}

				isSuccess = true

				var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: &users, StatusDesc: usersResp.StatusDesc}
				c.Data["json"] = resp

			} else {
				logs.Error("Unable to get role ")
				var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Given Role not found."}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Unable to verify token ")
		var resp responses.UsersGatewayResponseDTO = responses.UsersGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Post ...
// @Title Invite User
// @Description Invite user to register using email or username or phone number
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.InviteRequestDTO	true		"body for Registration content"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /invite-user [post]
func (c *UserManagementController) InviteUserReg() {
	var v requests.InviteRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")
	var isSuccess bool = false
	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			link := "https://amc-flowpos.com/auth/user/invite/"
			inviteResp := functions.InviteUser(&c.Controller, v.Email, v.Role, link, verifyToken.User.UserId)

			var message string

			if inviteResp.StatusCode == 200 {
				logs.Info("Success response received")
				isSuccess = true
				message = "Email sent"

				var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: &message, StatusDesc: inviteResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				logs.Error("Failed generating invite token")
				var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: inviteResp.StatusDesc}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Post ...
// @Title Verify Invite
// @Description Verify invite
// @Param	body		body 	requests.StringRequestDTO	true		"body for Registration content"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /verify-invite [post]
func (c *UserManagementController) VerifyInvite() {
	var v requests.StringRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	logs.Info("Success response")

	verifyUserToken := functions.VerifyInviteToken(&c.Controller, v.Value)

	logs.Info("User token verification is ", verifyUserToken)

	var isSuccess bool = false

	if verifyUserToken.StatusCode == 200 {
		logs.Info("Success response received")
		isSuccess = true
		message := "Verification Successful"

		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: &message, StatusDesc: verifyUserToken.StatusDesc}
		c.Data["json"] = resp
	} else {
		var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: verifyUserToken.StatusDesc}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetRoles ...
// @Title Get Roles
// @Description Get all roles
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.RolesAllGatewayResponseDTO
// @Failure 403 body is empty
// @router /get-roles [get]
func (c *UserManagementController) GetRoles() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")
	var isSuccess bool = false
	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			rolesResp := functions.GetRoles(&c.Controller)

			// var message string

			if rolesResp.StatusCode == 200 {
				logs.Info("Name returned: ", verifyToken.User.FullName)

				isSuccess = true
				// message = "Email sent"

				var resp responses.RolesAllGatewayResponseDTO = responses.RolesAllGatewayResponseDTO{Success: isSuccess, Result: rolesResp.Roles, StatusDesc: rolesResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.RolesAllGatewayResponseDTO = responses.RolesAllGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.RolesAllGatewayResponseDTO = responses.RolesAllGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.RolesAllGatewayResponseDTO = responses.RolesAllGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetInvites ...
// @Title Get Invites
// @Description Get all invites
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.UserInvitesResponse
// @Failure 403 body is empty
// @router /get-invites [get]
func (c *UserManagementController) GetUserInvites() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")
	var isSuccess bool = false
	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			invitesResp := functions.GetInvites(&c.Controller)

			// var message string

			if invitesResp.StatusCode == 200 {
				userInvites := []responses.UserInvites{}

				logs.Info("Name returned: ", verifyToken.User.FullName)

				if invitesResp.UserInvites != nil {
					for _, userInvite := range *invitesResp.UserInvites {
						// splitName := strings.Split(userInvite.InvitedBy.FullName, " | ")

						// firstname := ""
						// lastname := ""
						// if len(splitName) > 1 {
						// 	firstname = splitName[0]
						// 	lastname = splitName[1]
						// } else {
						// 	firstname = splitName[0]
						// }

						// logs.Debug("Name is ", splitName[0])
						logs.Debug("Email is ", userInvite.Email)
						// branch := &responses.BranchResp{}

						// if userInvite.InvitedBy.Branch != nil {
						// 	currency := responses.CurrencyResp{Symbol: userInvite.InvitedBy.Branch.Country.DefaultCurrency.Symbol, Currency: userInvite.InvitedBy.Branch.Country.DefaultCurrency.Currency}
						// 	country := responses.CountryResp{Country: userInvite.InvitedBy.Branch.Country.Country, CountryCode: userInvite.InvitedBy.Branch.Country.CountryCode, Currency: &currency}
						// 	branch = &responses.BranchResp{BranchId: userInvite.InvitedBy.Branch.BranchId, Branch: userInvite.InvitedBy.Branch.Branch, Country: &country, Location: userInvite.InvitedBy.Branch.Location, PhoneNumber: userInvite.InvitedBy.Branch.PhoneNumber}
						// } else {
						// 	branch = nil
						// }
						// user := responses.UserGateway{
						// 	UserId: userInvite.InvitedBy.UserId,
						// 	// UserType:    userResp.User.UserType,
						// 	FirstName:   firstname,
						// 	LastName:    lastname,
						// 	Username:    userInvite.InvitedBy.Username,
						// 	Email:       userInvite.InvitedBy.Email,
						// 	PhoneNumber: userInvite.InvitedBy.PhoneNumber,
						// 	Role:        userInvite.InvitedBy.Role,
						// 	Customer:    userInvite.InvitedBy.Customer,
						// 	ImagePath:   verifyToken.User.ImagePath,
						// 	// Gender:         userResp.User.Gender,
						// 	// Dob:            userResp.User.Dob,
						// 	// Address:        userResp.User.Address,
						// 	// IdType:         userResp.User.IdType,
						// 	// IdNumber:       userResp.User.IdNumber,
						// 	// Active:         userResp.User.Active,
						// 	// IsVerified:     userResp.User.IsVerified,
						// 	// DateRegistered: userResp.User.DateCreated,
						// }

						// userInvite := responses.UserInvites{UserInviteId: userInvite.UserInviteId, InvitedBy: &user, InvitationToken: userInvite.InvitationToken, Status: userInvite.Status}
						userInvite := responses.UserInvites{UserInviteId: userInvite.UserInviteId, Status: userInvite.Status, Email: userInvite.Email, Role: userInvite.Role.Role}

						userInvites = append(userInvites, userInvite)
					}
				} else {
					userInvites = []responses.UserInvites{}
				}

				isSuccess = true

				var resp responses.UserInvitesResponse = responses.UserInvitesResponse{Success: isSuccess, Result: &userInvites, StatusDesc: invitesResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.UserInvitesResponse = responses.UserInvitesResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UserInvitesResponse = responses.UserInvitesResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserInvitesResponse = responses.UserInvitesResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateImage ...
// @Title UpdateImage
// @Description Update User's Image
// @Param	Authorization		header 	string true		"header for User"
// @Param	UserImage		formData 	file	true		"User Image"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /update-user-image [post]
func (c *UserManagementController) UpdateUserImage() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			image, header, err := c.GetFile("UserImage")

			if err != nil {
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
				c.Data["json"] = resp
			} else {
				logs.Info("Success response received")
				isSuccess = false

				respCode, filePath := functions.SaveImage(&c.Controller, "UserImage", image, *header)

				if respCode == 200 {
					var data responses.UserGateway

					userResp := functions.UpdateUserImage(&c.Controller, filePath, verifyToken.User.UserId)

					if userResp.StatusCode == 200 {
						logs.Info("Name returned: ", userResp.User.FullName)
						splitName := strings.Split(userResp.User.FullName, " | ")

						// logs.Debug("Name is ", splitName[0])
						firstname := ""
						lastname := ""
						if len(splitName) > 1 {
							firstname = splitName[0]
							lastname = splitName[1]
						} else {
							firstname = splitName[0]
						}

						// branch := &responses.BranchResp{}
						// if userResp.User.Branch != nil {
						// 	currency := responses.CurrencyResp{Symbol: userResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: userResp.User.Branch.Country.DefaultCurrency.Currency}
						// 	country := responses.CountryResp{Country: userResp.User.Branch.Country.Country, CountryCode: userResp.User.Branch.Country.CountryCode, Currency: &currency}
						// 	branch = &responses.BranchResp{BranchId: userResp.User.Branch.BranchId, Branch: userResp.User.Branch.Branch, Country: &country, Location: userResp.User.Branch.Location, PhoneNumber: userResp.User.Branch.PhoneNumber}
						// } else {
						// 	branch = nil
						// }
						status := "ACTIVE"

						if userResp.User.Active == 6 {
							status = "INACTIVE"
						}
						if userResp.User.Active == 2 {
							status = "PENDING"
						}
						if userResp.User.Active == 4 {
							status = "INACTIVE"
						}

						data = responses.UserGateway{
							UserId: userResp.User.UserId,
							// UserType:    userResp.User.UserType,
							FirstName:   firstname,
							LastName:    lastname,
							Username:    userResp.User.Username,
							Email:       userResp.User.Email,
							PhoneNumber: userResp.User.PhoneNumber,
							Role:        userResp.User.Role,
							ImagePath:   userResp.User.ImagePath,
							Customer:    userResp.User.UserDetails,
							Status:      status,
							// Gender:         userResp.User.Gender,
							// Dob:            userResp.User.Dob,
							// Address:        userResp.User.Address,
							// IdType:         userResp.User.IdType,
							// IdNumber:       userResp.User.IdNumber,
							// Active:         userResp.User.Active,
							// IsVerified:     userResp.User.IsVerified,
							// DateRegistered: userResp.User.DateCreated,
						}

						isSuccess = true

						var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: &data, StatusDesc: userResp.StatusDesc}
						c.Data["json"] = resp
					} else {
						var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
						c.Data["json"] = resp
					}
				} else {
					var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Error updating image"}
					c.Data["json"] = resp
				}
			}
		} else {
			var resp responses.StringResponseDTO = responses.StringResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.RolesAllGatewayResponseDTO = responses.RolesAllGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateUserDetails ...
// @Title UpdateUserDetails
// @Description Update User's details
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateUserRequestDTO	true		"body to update"
// @Success 200 {object} responses.UserGatewayResponseDTO
// @Failure 403 body is empty
// @router /update-user/:id [put]
func (c *UserManagementController) UpdateUser() {
	var v requests.UpdateUserRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	authorization := c.Ctx.Input.Header("Authorization")
	logs.Info("Updating user ...")

	idStr := c.Ctx.Input.Param(":id")
	// id, _ := strconv.ParseInt(idStr, 0, 64)

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			updateUserResp := functions.UpdateUser(&c.Controller, idStr, v)

			if updateUserResp.StatusCode == 200 {
				var data responses.UserGateway

				splitName := strings.Split(updateUserResp.User.FullName, " | ")

				firstname := ""
				lastname := ""
				if len(splitName) > 1 {
					firstname = splitName[0]
					lastname = splitName[1]
				} else {
					firstname = splitName[0]
				}

				// branch := &responses.BranchResp{}
				// if updateUserResp.User.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: updateUserResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: updateUserResp.User.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: updateUserResp.User.Branch.Country.Country, CountryCode: updateUserResp.User.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: updateUserResp.User.Branch.BranchId, Branch: updateUserResp.User.Branch.Branch, Country: &country, Location: updateUserResp.User.Branch.Location, PhoneNumber: updateUserResp.User.Branch.PhoneNumber}
				// } else {
				// 	branch = nil
				// }
				status := "ACTIVE"

				if updateUserResp.User.Active == 6 {
					status = "DELETED"
				}
				if updateUserResp.User.Active == 2 {
					status = "PENDING"
				}
				if updateUserResp.User.Active == 4 {
					status = "INACTIVE"
				}

				data = responses.UserGateway{
					UserId: updateUserResp.User.UserId,
					// UserType:    regResp.User.UserType,
					FirstName:   firstname,
					LastName:    lastname,
					Username:    updateUserResp.User.Username,
					Email:       updateUserResp.User.Email,
					PhoneNumber: updateUserResp.User.PhoneNumber,
					Role:        updateUserResp.User.Role,
					Customer:    updateUserResp.User.UserDetails,
					ImagePath:   updateUserResp.User.ImagePath,
					Status:      status,
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
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: &data, StatusDesc: "User updated"}
				c.Data["json"] = resp

			} else {
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Update User Role ...
// @Title Update User Role
// @Description Update User's role
// @Param	Authorization		header 	string true		"header for User"
// @Param	userid		path 	string	true		"The userid id you want to update"
// @Param	body		body 	requests.UpdateUserRoleRequestDTO	true		"body to update"
// @Success 200 {object} responses.UserGatewayResponseDTO
// @Failure 403 body is empty
// @router /update-user-role/:userid [put]
func (c *UserManagementController) UpdateUserRole() {
	var v requests.UpdateUserRoleRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	authorization := c.Ctx.Input.Header("Authorization")
	logs.Info("Updating user ...")

	idStr := c.Ctx.Input.Param(":userid")
	// id, _ := strconv.ParseInt(idStr, 0, 64)

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			updateUserResp := functions.UpdateUserRole(&c.Controller, idStr, v)

			if updateUserResp.StatusCode == 200 {
				var data responses.UserGateway

				splitName := strings.Split(updateUserResp.User.FullName, " | ")

				firstname := ""
				lastname := ""
				if len(splitName) > 1 {
					firstname = splitName[0]
					lastname = splitName[1]
				} else {
					firstname = splitName[0]
				}

				// branch := &responses.BranchResp{}
				// if updateUserResp.User.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: updateUserResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: updateUserResp.User.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: updateUserResp.User.Branch.Country.Country, CountryCode: updateUserResp.User.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: updateUserResp.User.Branch.BranchId, Branch: updateUserResp.User.Branch.Branch, Country: &country, Location: updateUserResp.User.Branch.Location, PhoneNumber: updateUserResp.User.Branch.PhoneNumber}
				// } else {
				// 	branch = nil
				// }
				status := "ACTIVE"

				if updateUserResp.User.Active == 6 {
					status = "DELETED"
				}
				if updateUserResp.User.Active == 2 {
					status = "PENDING"
				}
				if updateUserResp.User.Active == 4 {
					status = "INACTIVE"
				}

				data = responses.UserGateway{
					UserId: updateUserResp.User.UserId,
					// UserType:    regResp.User.UserType,
					FirstName:   firstname,
					LastName:    lastname,
					Username:    updateUserResp.User.Username,
					Email:       updateUserResp.User.Email,
					PhoneNumber: updateUserResp.User.PhoneNumber,
					Role:        updateUserResp.User.Role,
					Customer:    updateUserResp.User.UserDetails,
					ImagePath:   updateUserResp.User.ImagePath,
					Status:      status,
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
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: &data, StatusDesc: "User updated"}
				c.Data["json"] = resp

			} else {
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred " + updateUserResp.StatusDesc}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to access this resource"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Invalid authorization token"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Update User Branch ...
// @Title Update User Branch
// @Description Update User's Branch
// @Param	Authorization		header 	string true		"header for User"
// @Param	userid		path 	string	true		"The userid id you want to update"
// @Param	body		body 	requests.UpdateUserBranchRequestDTO	true		"body to update"
// @Success 200 {object} responses.UserGatewayResponseDTO
// @Failure 403 body is empty
// @router /update-user-branch/:userid [put]
func (c *UserManagementController) UpdateUserBranch() {
	var v requests.UpdateUserBranchRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	authorization := c.Ctx.Input.Header("Authorization")
	logs.Info("Updating user ...")

	idStr := c.Ctx.Input.Param(":userid")
	// id, _ := strconv.ParseInt(idStr, 0, 64)

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			updateUserResp := functions.UpdateUserBranch(&c.Controller, idStr, v)

			if updateUserResp.StatusCode == 200 {
				var data responses.UserGateway

				splitName := strings.Split(updateUserResp.User.FullName, " | ")

				firstname := ""
				lastname := ""
				if len(splitName) > 1 {
					firstname = splitName[0]
					lastname = splitName[1]
				} else {
					firstname = splitName[0]
				}

				// branch := &responses.BranchResp{}
				// if updateUserResp.User.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: updateUserResp.User.Branch.Country.DefaultCurrency.Symbol, Currency: updateUserResp.User.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: updateUserResp.User.Branch.Country.Country, CountryCode: updateUserResp.User.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: updateUserResp.User.Branch.BranchId, Branch: updateUserResp.User.Branch.Branch, Country: &country, Location: updateUserResp.User.Branch.Location, PhoneNumber: updateUserResp.User.Branch.PhoneNumber}
				// } else {
				// 	branch = nil
				// }
				status := "ACTIVE"

				if updateUserResp.User.Active == 6 {
					status = "INACTIVE"
				}
				if updateUserResp.User.Active == 2 {
					status = "PENDING"
				}
				if updateUserResp.User.Active == 4 {
					status = "INACTIVE"
				}

				data = responses.UserGateway{
					UserId: updateUserResp.User.UserId,
					// UserType:    regResp.User.UserType,
					FirstName:   firstname,
					LastName:    lastname,
					Username:    updateUserResp.User.Username,
					Email:       updateUserResp.User.Email,
					PhoneNumber: updateUserResp.User.PhoneNumber,
					Role:        updateUserResp.User.Role,
					Customer:    updateUserResp.User.UserDetails,
					ImagePath:   updateUserResp.User.ImagePath,
					Status:      status,
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
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: &data, StatusDesc: "User updated"}
				c.Data["json"] = resp

			} else {
				var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred " + updateUserResp.StatusDesc}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to access this resource"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Invalid authorization token"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateUserInviteToken ...
// @Title UpdateUserInvite Token
// @Description Update User's Invite token
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Success 200 {object} responses.UserInviteResponse
// @Failure 403 body is empty
// @router /revoke-invite/:id [put]
func (c *UserManagementController) UpdateInviteToken() {
	idStr := c.Ctx.Input.Param(":id")

	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			updateInvite := functions.UpdateUserInvite(&c.Controller, idStr, "CANCELLED")

			if updateInvite.StatusCode == 200 {
				// splitName := strings.Split(updateInvite.UserInvite.InvitedBy.FullName, " | ")

				// firstname := ""
				// lastname := ""
				// if len(splitName) > 1 {
				// 	firstname = splitName[0]
				// 	lastname = splitName[1]
				// } else {
				// 	firstname = splitName[0]
				// }
				// branch := &responses.BranchResp{}

				// if updateInvite.UserInvite.InvitedBy.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: updateInvite.UserInvite.InvitedBy.Branch.Country.DefaultCurrency.Symbol, Currency: updateInvite.UserInvite.InvitedBy.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: updateInvite.UserInvite.InvitedBy.Branch.Country.Country, CountryCode: updateInvite.UserInvite.InvitedBy.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: updateInvite.UserInvite.InvitedBy.Branch.BranchId, Branch: updateInvite.UserInvite.InvitedBy.Branch.Branch, Country: &country, Location: updateInvite.UserInvite.InvitedBy.Branch.Location, PhoneNumber: updateInvite.UserInvite.InvitedBy.Branch.PhoneNumber}
				// } else {
				// 	branch = nil
				// }
				// user := responses.UserGateway{
				// 	UserId: updateInvite.UserInvite.InvitedBy.UserId,
				// 	// UserType:    userResp.User.UserType,
				// 	FirstName:   firstname,
				// 	LastName:    lastname,
				// 	Username:    updateInvite.UserInvite.InvitedBy.Username,
				// 	Email:       updateInvite.UserInvite.InvitedBy.Email,
				// 	PhoneNumber: updateInvite.UserInvite.InvitedBy.PhoneNumber,
				// 	Role:        updateInvite.UserInvite.InvitedBy.Role,
				// 	Customer:    updateInvite.UserInvite.InvitedBy.Customer,
				// 	ImagePath:   updateInvite.UserInvite.InvitedBy.ImagePath,
				// 	// Gender:         userResp.User.Gender,
				// 	// Dob:            userResp.User.Dob,
				// 	// Address:        userResp.User.Address,
				// 	// IdType:         userResp.User.IdType,
				// 	// IdNumber:       userResp.User.IdNumber,
				// 	// Active:         userResp.User.Active,
				// 	// IsVerified:     userResp.User.IsVerified,
				// 	// DateRegistered: userResp.User.DateCreated,
				// }

				// userInvite := responses.UserInvites{UserInviteId: updateInvite.UserInvite.UserInviteId, InvitedBy: &user, InvitationToken: updateInvite.UserInvite.InvitationToken, Status: updateInvite.UserInvite.Status}
				userInvite := responses.UserInvites{UserInviteId: updateInvite.UserInvite.UserInviteId, Status: updateInvite.UserInvite.Status}

				logs.Error("Success")
				isSuccess = true
				var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: &userInvite, StatusDesc: "User Invite updated"}
				c.Data["json"] = resp
			} else {
				message := "An Error occurred. Unable to update invite"
				if updateInvite.StatusCode == 503 {
					message = updateInvite.StatusDesc
				}
				logs.Error("Unable to update token")
				var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: message}
				c.Data["json"] = resp
			}
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

// GetUserInvite ...
// @Title Get User Invite
// @Description Get user invite
// @Param	Authorization		header 	string true		"header for User"
// @Param	token		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.UserInviteResponse
// @Failure 403 body is empty
// @router /get-user-invite/:token [get]
func (c *UserManagementController) GetUserInvite() {
	inviteToken := c.Ctx.Input.Param(":token")
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")
	var isSuccess bool = false
	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			inviteResp := functions.GetInvite(&c.Controller, inviteToken)

			// var message string

			if inviteResp.StatusCode == 200 {
				logs.Info("Name returned: ", verifyToken.User.FullName)

				// splitName := strings.Split(inviteResp.UserInvite.InvitedBy.FullName, " | ")

				// firstname := ""
				// lastname := ""
				// if len(splitName) > 1 {
				// 	firstname = splitName[0]
				// 	lastname = splitName[1]
				// } else {
				// 	firstname = splitName[0]
				// }
				// branch := &responses.BranchResp{}

				// if inviteResp.UserInvite.InvitedBy.Branch != nil {
				// 	currency := responses.CurrencyResp{Symbol: inviteResp.UserInvite.InvitedBy.Branch.Country.DefaultCurrency.Symbol, Currency: inviteResp.UserInvite.InvitedBy.Branch.Country.DefaultCurrency.Currency}
				// 	country := responses.CountryResp{Country: inviteResp.UserInvite.InvitedBy.Branch.Country.Country, CountryCode: inviteResp.UserInvite.InvitedBy.Branch.Country.CountryCode, Currency: &currency}
				// 	branch = &responses.BranchResp{BranchId: inviteResp.UserInvite.InvitedBy.Branch.BranchId, Branch: inviteResp.UserInvite.InvitedBy.Branch.Branch, Country: &country, Location: inviteResp.UserInvite.InvitedBy.Branch.Location, PhoneNumber: inviteResp.UserInvite.InvitedBy.Branch.PhoneNumber}
				// } else {
				// 	branch = nil
				// }
				// user := responses.UserGateway{
				// 	UserId: inviteResp.UserInvite.InvitedBy.UserId,
				// 	// UserType:    userResp.User.UserType,
				// 	FirstName:   firstname,
				// 	LastName:    lastname,
				// 	Username:    inviteResp.UserInvite.InvitedBy.Username,
				// 	Email:       inviteResp.UserInvite.InvitedBy.Email,
				// 	PhoneNumber: inviteResp.UserInvite.InvitedBy.PhoneNumber,
				// 	Role:        inviteResp.UserInvite.InvitedBy.Role,
				// 	Customer:    inviteResp.UserInvite.InvitedBy.Customer,
				// 	ImagePath:   inviteResp.UserInvite.InvitedBy.ImagePath,
				// 	// Gender:         userResp.User.Gender,
				// 	// Dob:            userResp.User.Dob,
				// 	// Address:        userResp.User.Address,
				// 	// IdType:         userResp.User.IdType,
				// 	// IdNumber:       userResp.User.IdNumber,
				// 	// Active:         userResp.User.Active,
				// 	// IsVerified:     userResp.User.IsVerified,
				// 	// DateRegistered: userResp.User.DateCreated,
				// }

				// userInvite := responses.UserInvites{UserInviteId: inviteResp.UserInvite.UserInviteId, InvitedBy: &user, InvitationToken: inviteResp.UserInvite.InvitationToken, Status: inviteResp.UserInvite.Status}

				userInvite := responses.UserInvites{UserInviteId: inviteResp.UserInvite.UserInviteId, Role: inviteResp.UserInvite.Role.Role, Status: inviteResp.UserInvite.Status}
				logs.Error("Success")
				isSuccess = true
				var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: &userInvite, StatusDesc: "User Invite fetched"}
				c.Data["json"] = resp
			} else {
				var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Log Out ...
// @Title Log Out
// @Description Log User Out
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.UserGatewayResponseDTO
// @Failure 403 body is empty
// @router /log-out [get]
func (c *UserManagementController) LogOut() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {

			logout := functions.LogOut(&c.Controller, token[1])

			if logout.StatusCode == 200 {
				isSuccess = true
			}

			var resp responses.UserGatewayResponseDTO = responses.UserGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "User logout complete"}
			c.Data["json"] = resp
		} else {
			var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.UserInviteResponse = responses.UserInviteResponse{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
