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

// SystemController operations for System
type SystemController struct {
	beego.Controller
}

// URLMapping ...
func (c *SystemController) URLMapping() {
	c.Mapping("AddBranch", c.AddBranch)
	c.Mapping("GetOneBranch", c.GetOneBranch)
	c.Mapping("GetAllBranches", c.GetAllBranches)
	c.Mapping("Delete", c.Delete)
	c.Mapping("UpdateBranch", c.UpdateBranch)
}

// Post ...
// @Title Create
// @Description create System
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.BranchRequestDTO	true		"body for System content"
// @Success 200 {object} responses.BranchResponseDTO
// @Failure 403 body is empty
// @router /add-branch [post]
func (c *SystemController) AddBranch() {
	var v requests.BranchRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			userDetailsResp := functions.GetUserDetails(&c.Controller, v.BranchManager)
			if userDetailsResp.StatusCode == 200 {
				addBranchResp := functions.AddBranch(&c.Controller, v, verifyToken.User.UserId)

				if addBranchResp.StatusCode == 200 {
					// Assign branch manager to added branch
					splitName := strings.Split(userDetailsResp.User.FullName, " | ")
					firstname := ""
					lastname := ""
					if len(splitName) > 1 {
						firstname = splitName[0]
						lastname = splitName[1]
					} else {
						firstname = splitName[0]
					}
					userDetails := requests.UpdateUserRequestDTO{BranchId: addBranchResp.Branch.BranchId, FirstName: firstname, LastName: lastname, Username: userDetailsResp.User.Username, PhoneNumber: userDetailsResp.User.PhoneNumber, Gender: userDetailsResp.User.Gender, Dob: userDetailsResp.User.Dob.GoString(), Address: userDetailsResp.User.Address}
					userId := strconv.FormatInt(userDetailsResp.User.UserId, 10)
					updateUserResp := functions.UpdateUser(&c.Controller, userId, userDetails)
					branchIdStr := strconv.FormatInt(addBranchResp.Branch.BranchId, 10)
					updateBranchResp := functions.UpdateBranchBranchManger(&c.Controller, userId, branchIdStr)
					message := "Branch Added Successfully"
					if updateUserResp.StatusCode != 200 {
						message = "Branch added but failed to assign manager"
					}
					if updateBranchResp.StatusCode != 200 {
						message = "Branch added but failed to assign branch manager"
					}
					// var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: addBranchResp.Branch.Country.DefaultCurrency.Symbol, Currency: addBranchResp.Branch.Country.DefaultCurrency.Currency}
					// var country responses.CountryResp = responses.CountryResp{Country: addBranchResp.Branch.Country.Country, CountryCode: addBranchResp.Branch.Country.CountryCode, Currency: &curr}
					var data responses.BranchResp = responses.BranchResp{
						BranchId: addBranchResp.Branch.BranchId,
						Branch:   addBranchResp.Branch.Branch,
						// Country:     &country,
						Location:    addBranchResp.Branch.Location,
						PhoneNumber: addBranchResp.Branch.PhoneNumber,
					}

					isSuccess = true
					var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: &data, StatusDesc: message}
					c.Data["json"] = resp
				} else {
					var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: addBranchResp.StatusDesc}
					c.Data["json"] = resp
				}
			} else {
				var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Branch manager does not exist"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error verifying token")
			var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred." + verifyToken.StatusDesc}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to perform this request"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOneBranch ...
// @Title GetOneBranch
// @Description get Branch
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.BranchOriResponseDTO
// @Failure 403 :id is empty
// @router /get-branch/:id [get]
func (c *SystemController) GetOneBranch() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			getBranchResp := functions.GetBranch(&c.Controller, id)

			if getBranchResp.StatusCode == 200 {
				// var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: getBranchResp.Branch.Country.DefaultCurrency.Symbol, Currency: getBranchResp.Branch.Country.DefaultCurrency.Currency}
				// var country responses.CountryResp = responses.CountryResp{Country: getBranchResp.Branch.Country.Country, CountryCode: getBranchResp.Branch.Country.CountryCode, Currency: &curr}
				var data responses.BranchResp = responses.BranchResp{
					BranchId: getBranchResp.Branch.BranchId,
					Branch:   getBranchResp.Branch.Branch,
					// Country:     &country,
					Location:    getBranchResp.Branch.Location,
					PhoneNumber: getBranchResp.Branch.PhoneNumber,
					DateCreated: getBranchResp.Branch.DateCreated,
				}

				isSuccess = true
				var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: &data, StatusDesc: "Branch details fetched Successfully"}
				c.Data["json"] = resp
			} else {
				logs.Error("An error occurred fetching branches from api")
				var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error verifying token")
			var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllCountries ...
// @Title GetAllCountries
// @Description get Countries
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.CountriesResponseDTO
// @Failure 403
// @router /get-countries [get]
func (c *SystemController) GetAllCountries() {
	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			getCountriesResp := functions.GetCountries(&c.Controller)

			// json.Unmarshal(getBranchResp.Branches, &v)

			if getCountriesResp.StatusCode == 200 {
				logs.Info("Response is 200")
				var countries []responses.CountryResp
				if getCountriesResp.Countries != nil && len(*getCountriesResp.Countries) > 0 {
					for _, country := range *getCountriesResp.Countries {

						var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: country.DefaultCurrency.Symbol, Currency: country.DefaultCurrency.Currency}
						var country responses.CountryResp = responses.CountryResp{Country: country.Country, CountryCode: country.CountryCode, Currency: &curr}

						countries = append(countries, country)
					}
				} else {
					countries = []responses.CountryResp{}
				}

				isSuccess = true
				var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: &countries, StatusDesc: "Countries fetched Successfully"}
				c.Data["json"] = resp
			} else {
				logs.Error("An error occurred fetching countries from api")
				var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error verifying token")
			var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.CountriesResponseDTO = responses.CountriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetAllBranches ...
// @Title GetAllBranches
// @Description get branches
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.BranchesResponseDTO
// @Failure 403
// @router /get-branches [get]
func (c *SystemController) GetAllBranches() {
	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			getBranchResp := functions.GetBranches(&c.Controller)

			// json.Unmarshal(getBranchResp.Branches, &v)

			if getBranchResp.StatusCode == 200 {
				logs.Info("Response is 200")
				var branches []responses.BranchResp
				if getBranchResp.Branches != nil && len(*getBranchResp.Branches) > 0 {
					for _, branch := range *getBranchResp.Branches {

						var branchManager *responses.UserGateway

						logs.Info("Branch value is ", branch.BranchManager)
						if branch.BranchManager != nil {
							splitName := strings.Split(branch.BranchManager.FullName, " | ")
							firstname := ""
							lastname := ""
							if len(splitName) > 1 {
								firstname = splitName[0]
								lastname = splitName[1]
							} else {
								firstname = splitName[0]
							}

							// var curr responses.CurrencyResp = responses.CurrencyResp{Symbol: branch.Country.DefaultCurrency.Symbol, Currency: branch.Country.DefaultCurrency.Currency}
							// var country responses.CountryResp = responses.CountryResp{Country: branch.Country.Country, CountryCode: branch.Country.CountryCode, Currency: &curr}
							branchManager = &responses.UserGateway{UserId: branch.BranchManager.UserId, FirstName: firstname, LastName: lastname, Username: branch.BranchManager.Username, Email: branch.BranchManager.Email, PhoneNumber: branch.BranchManager.PhoneNumber, ImagePath: branch.BranchManager.ImagePath}
						} else {
							branchManager = nil
						}

						var data responses.BranchResp = responses.BranchResp{
							BranchId: branch.BranchId,
							Branch:   branch.Branch,
							// Country:     &country,
							Location:      branch.Location,
							PhoneNumber:   branch.PhoneNumber,
							DateCreated:   branch.DateCreated,
							BranchManager: branchManager,
						}

						branches = append(branches, data)
					}
				} else {
					branches = []responses.BranchResp{}
				}

				branchesData := responses.BranchesData{}
				branchesData.Data = &branches
				branchesData.Count = len(branches)

				isSuccess = true
				var resp responses.BranchesResponseDTO = responses.BranchesResponseDTO{Success: isSuccess, Result: &branchesData, StatusDesc: "Branches fetched Successfully"}
				c.Data["json"] = resp
			} else {
				var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}
		} else {
			logs.Error("Error verifying token")
			var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}

	} else {
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateBranch ...
// @Title Update Branch
// @Description update a Branch
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.BranchRequestDTO	true		"body for Branches content"
// @Success 200 {object} responses.BranchResponseDTO
// @Failure 403 :id is not int
// @router /update-branch/:id [put]
func (c *SystemController) UpdateBranch() {
	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			idStr := c.Ctx.Input.Param(":id")
			var r requests.BranchRequestDTO
			json.Unmarshal(c.Ctx.Input.RequestBody, &r)
			message := "Branch updated successfully"
			userDetailsResp := functions.GetUserDetails(&c.Controller, r.BranchManager)
			branchResp := &responses.BranchResp{}

			if userDetailsResp.StatusCode == 200 {
				updateBranch := functions.UpdateBranch(&c.Controller, r, verifyToken.User.UserId, idStr)

				if updateBranch.StatusCode == 200 {
					splitName := strings.Split(userDetailsResp.User.FullName, " | ")
					firstname := ""
					lastname := ""
					if len(splitName) > 1 {
						firstname = splitName[0]
						lastname = splitName[1]
					} else {
						firstname = splitName[0]
					}
					role_name, _ := beego.AppConfig.String("branchManagerRoleName")
					logs.Info("About to get data for role ", role_name)
					role := functions.GetRoleWithRoleName(&c.Controller, role_name)
					logs.Info("Get role response is ", role.Role.RoleId)
					var roleId int64 = 0
					if role.StatusCode == 200 {
						roleId = role.Role.RoleId
					}
					logs.Info("Sending role ", roleId)
					userDetails := requests.UpdateUserRequestDTO{RoleId: roleId, BranchId: updateBranch.Branch.BranchId, FirstName: firstname, LastName: lastname, Username: userDetailsResp.User.Username, PhoneNumber: userDetailsResp.User.PhoneNumber, Gender: userDetailsResp.User.Gender, Dob: userDetailsResp.User.Dob.GoString(), Address: userDetailsResp.User.Address}
					userId := strconv.FormatInt(userDetailsResp.User.UserId, 10)
					updateUserResp := functions.UpdateUser(&c.Controller, userId, userDetails)
					if updateUserResp.StatusCode == 200 {
						logs.Info("Update user response is ", updateUserResp.StatusDesc)
						branchIdStr := strconv.FormatInt(updateBranch.Branch.BranchId, 10)
						updateBranchResp := functions.UpdateBranchBranchManger(&c.Controller, userId, branchIdStr)
						splitName := strings.Split(updateUserResp.User.FullName, " | ")
						firstname := ""
						lastname := ""
						if len(splitName) > 1 {
							firstname = splitName[0]
							lastname = splitName[1]
						} else {
							firstname = splitName[0]
						}

						branchManager := responses.UserGateway{
							UserId:      updateUserResp.User.UserId,
							FirstName:   firstname,
							LastName:    lastname,
							Username:    updateUserResp.User.Username,
							Email:       updateUserResp.User.Email,
							PhoneNumber: updateUserResp.User.PhoneNumber,
							ImagePath:   updateUserResp.User.ImagePath,
							Customer:    updateUserResp.User.UserDetails,
							// Gender:
							// Dob:
							// Address:
							// IdType:
							// IdNumber:
							// Active:
							IsVerified: updateUserResp.User.IsVerified,
							Role:       updateUserResp.User.Role,
						}

						branchResp = &responses.BranchResp{
							BranchId:      updateBranch.Branch.BranchId,
							Branch:        updateBranch.Branch.Branch,
							Location:      updateBranch.Branch.Location,
							PhoneNumber:   updateBranch.Branch.PhoneNumber,
							DateCreated:   updateBranch.Branch.DateCreated,
							BranchManager: &branchManager,
						}

						if updateBranchResp.StatusCode != 200 {
							message = "Branch updated. Failed to update branch's branch manager"
						}
					} else {
						message = "Branch updated. Failed to update branch manager"
						logs.Error("Failed to update user", updateBranch.StatusDesc)
						resp := responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Branch update failed. " + updateUserResp.StatusDesc}
						c.Data["json"] = resp
					}

					isSuccess = true
				} else {
					message = "Failed to update branch"
					branchResp = nil
				}
			} else {
				logs.Info("Failed to get branch manager")
				message = "Failed to get specified branch manager"
			}
			resp := responses.BranchResponseDTO{Success: isSuccess, Result: branchResp, StatusDesc: message}
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = resp
		} else {
			logs.Error("Failed to verify token")
			resp := responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to verify token"}
			c.Data["json"] = resp
		}
	} else {
		c.Ctx.Output.SetStatus(200)
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// Delete Branch ...
// @Title Delete Branch
// @Description delete the Branches
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete-branch/:id [delete]
func (c *SystemController) Delete() {
	idStr := c.Ctx.Input.Param(":id")

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			message := "Deleted"
			deleteResp := functions.DeleteBranch(&c.Controller, idStr)
			if deleteResp.StatusCode == 200 {
				resp := responses.StringResponseDTO{Success: true, Result: &message, StatusDesc: "Branch deleted successfully"}
				c.Ctx.Output.SetStatus(200)
				c.Data["json"] = resp
			} else {
				message = "Deletion Failed:: "
				resp := responses.StringResponseDTO{Success: false, Result: &message, StatusDesc: deleteResp.StatusDesc}
				c.Ctx.Output.SetStatus(301)
				c.Data["json"] = resp
			}
		} else {
			c.Ctx.Output.SetStatus(200)
			logs.Info("Failed to verify token")
			resp := responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to verify token"}
			c.Data["json"] = resp
		}
	} else {
		c.Ctx.Output.SetStatus(200)
		var resp responses.BranchResponseDTO = responses.BranchResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Get All ID Types ...
// @Title Get All ID Types
// @Description get all ID types
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.IDTypesGatewayResponseDTO
// @Failure 403
// @router /get-id-types [get]
func (c *SystemController) GetIdTypes() {
	// v := c.Ctx.Input.GetData("user")
	// userData, err := v.(*responses.UsersOri)

	// fmt.Printf("Type of v: %T\n", v)
	// fmt.Printf("Value of v: %+v\n", v)

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

	var isSuccess bool = false
	var message string = "An error occurred"

	idTypes := functions.GetIdTypes(&c.Controller, query, fields, sortby, order, offset, limit)

	// json.Unmarshal(getBranchResp.Branches, &v)

	if idTypes.StatusCode == 200 {
		logs.Info("Response is 200")

		isSuccess = true
		message = "ID Types fetched successfully"
		var resp responses.IDTypesGatewayResponseDTO = responses.IDTypesGatewayResponseDTO{Success: isSuccess, Result: idTypes.IdTypes, StatusDesc: message}
		c.Data["json"] = resp
	} else {
		logs.Error("An error occurred fetching id types from api ", idTypes.StatusDesc)
		message = "Id type fetch error"
		var resp responses.IDTypesGatewayResponseDTO = responses.IDTypesGatewayResponseDTO{Success: isSuccess, Result: nil, StatusDesc: message}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
