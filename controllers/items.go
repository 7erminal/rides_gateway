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

// ItemsController operations for Items
type ItemsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ItemsController) URLMapping() {
	c.Mapping("AddSalesItem", c.AddSalesItem)
	c.Mapping("UpdateItemImage", c.UpdateItemImage)
	c.Mapping("AddCategory", c.AddCategory)
	c.Mapping("GetCategories", c.GetCategories)
	c.Mapping("GetItems", c.GetItems)
	c.Mapping("GetProduct", c.GetProduct)
	c.Mapping("UpdateItem", c.UpdateItem)
	c.Mapping("AddRentalsItem", c.AddRentalsItem)
}

// AddSalesItem ...
// @Title Add Item
// @Description Add an item for sale
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.AddSalesItemRequestDTO	true		"body for Authentication content"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /add-sales-product [post]
func (c *ItemsController) AddSalesItem() {
	var v requests.AddSalesItemRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			logs.Info("Received \nProduct name: ", v.ProductName, "Branch ID:: ", v.BranchId, "Cost price:: ", v.CostPrice, "Image path:: ", v.ImagePath, "Quantity:: ", v.Quantity, "Selling price:: ", v.SellingPrice)
			proceed := true
			errorMessage := "An error occurred"
			logs.Info("Token verified!")
			logs.Info("Country is !", verifyToken.User.UserDetails.Branch)
			getBranchResp := functions.GetBranch(&c.Controller, v.BranchId)

			// if getBranchResp.StatusCode == 200 {
			// userId := verifyToken.User.UserId
			if getBranchResp.StatusCode != 200 {
				errorMessage = "Branch provided does not exist"
				proceed = false
			}

			sales_product_type_name, _ := beego.AppConfig.String("salesProductType")

			getProductTypes := functions.GetCategoryByName(&c.Controller, sales_product_type_name)

			if getProductTypes.StatusCode != 200 {
				errorMessage = "Product type provided does not exist"
				proceed = false
			}

			if proceed {
				req := requests.AddItemRequestDTO{ProductName: v.ProductName, Quantity: v.Quantity, ReorderLevel: 0, CostPrice: v.CostPrice, SellingPrice: v.SellingPrice, BranchId: v.BranchId, ImagePath: v.ImagePath}
				addItemResp := functions.AddItem(&c.Controller, req, getProductTypes.Category.CategoryId, getBranchResp.Branch.Country.CountryCode, v.BranchId, int(verifyToken.User.UserId))

				itemResp := responses.Item{}
				if addItemResp.StatusCode == 200 {
					if addItemResp.Item != nil {
						logs.Info("About to update item image")
						itemImageUpdateResp := functions.UpdateItemImage(&c.Controller, addItemResp.Item.ItemId, v.ImagePath)
						if itemImageUpdateResp.StatusCode == 200 {
							logs.Info("Successfully updated item image")
						} else {
							logs.Error("Failed update")
						}

						itemResp = responses.Item{
							ProductId:        addItemResp.Item.ItemId,
							ProductName:      addItemResp.Item.ItemName,
							Description:      addItemResp.Item.Description,
							ProductType:      addItemResp.Item.Category.CategoryName,
							ProductPrice:     float64(addItemResp.Item.ItemPrice.ItemPrice),
							ProductCostPrice: float64(addItemResp.Item.ItemPrice.AltItemPrice),
							ImagePath:        itemImageUpdateResp.Item.ImagePath,
							Quantity:         addItemResp.Item.Quantity,
							Branch:           addItemResp.Item.Branch,
						}

						isSuccess = true
					} else {
						isSuccess = false
					}

				}

				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &itemResp, StatusDesc: addItemResp.StatusDesc}

				c.Data["json"] = resp
				// } else {
				// 	var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

				// 	c.Data["json"] = resp
				// }
			} else {
				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: errorMessage}

				c.Data["json"] = resp
			}

		} else {
			var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to access this resource"}

			c.Data["json"] = resp
		}
	} else {
		logs.Info("Token not verified")
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Invalid authorization token"}

		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// AddRentalsItem ...
// @Title Add Rentals Item
// @Description Add an item for rent
// @Param	Authorization		header 	string true		"header for User"
// @Param	body		body 	requests.AddRentalItemRequestDTO	true		"body for Authentication content"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /add-rental-product [post]
func (c *ItemsController) AddRentalsItem() {
	var v requests.AddRentalItemRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			logs.Info("Received \nProduct name: ", v.ProductName, "Branch ID:: ", v.BranchId, "Image path:: ", v.ImagePath, "Quantity:: ", v.Quantity, "Rental price:: ", v.RentalPrice)
			proceed := true
			errorMessage := "An error occurred"
			logs.Info("Token verified!")
			logs.Info("Country is !", verifyToken.User.UserDetails.Branch)
			getBranchResp := functions.GetBranch(&c.Controller, v.BranchId)

			// if getBranchResp.StatusCode == 200 {
			// userId := verifyToken.User.UserId
			if getBranchResp.StatusCode != 200 {
				errorMessage = "Branch provided does not exist"
				proceed = false
			}

			sales_product_type_name, _ := beego.AppConfig.String("rentalsProductType")

			getProductTypes := functions.GetCategoryByName(&c.Controller, sales_product_type_name)

			if getProductTypes.StatusCode != 200 {
				errorMessage = "Product type provided does not exist"
				proceed = false
			}

			if proceed {
				req := requests.AddItemRequestDTO{ProductName: v.ProductName, Quantity: v.Quantity, ReorderLevel: v.ReorderLevel, CostPrice: 0, SellingPrice: v.RentalPrice, BranchId: v.BranchId, ImagePath: v.ImagePath}
				addItemResp := functions.AddItem(&c.Controller, req, getProductTypes.Category.CategoryId, getBranchResp.Branch.Country.CountryCode, v.BranchId, int(verifyToken.User.UserId))

				itemResp := responses.Item{}
				if addItemResp.StatusCode == 200 {
					if addItemResp.Item != nil {
						logs.Info("About to update item image")
						itemImageUpdateResp := functions.UpdateItemImage(&c.Controller, addItemResp.Item.ItemId, v.ImagePath)
						if itemImageUpdateResp.StatusCode == 200 {
							logs.Info("Successfully updated item image")
						} else {
							logs.Error("Failed update")
						}

						itemResp = responses.Item{
							ProductId:        addItemResp.Item.ItemId,
							ProductName:      addItemResp.Item.ItemName,
							Description:      addItemResp.Item.Description,
							ProductType:      addItemResp.Item.Category.CategoryName,
							ProductPrice:     float64(addItemResp.Item.ItemPrice.ItemPrice),
							ProductCostPrice: float64(addItemResp.Item.ItemPrice.AltItemPrice),
							ImagePath:        itemImageUpdateResp.Item.ImagePath,
							Quantity:         addItemResp.Item.Quantity,
							Branch:           addItemResp.Item.Branch,
						}

						isSuccess = true
					} else {
						isSuccess = false
					}

				}

				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &itemResp, StatusDesc: addItemResp.StatusDesc}

				c.Data["json"] = resp
				// } else {
				// 	var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

				// 	c.Data["json"] = resp
				// }
			} else {
				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: errorMessage}

				c.Data["json"] = resp
			}

		} else {
			var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "You are not authorized to access this resource"}

			c.Data["json"] = resp
		}
	} else {
		logs.Info("Token not verified")
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. Invalid authorization token"}

		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateItem ...
// @Title Update Item
// @Description Add an item
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.UpdateItemRequestDTO	true		"body for Authentication content"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /update-product/:id [put]
func (c *ItemsController) UpdateItem() {
	var v requests.UpdateItemRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	idStr := c.Ctx.Input.Param(":id")

	authorization := c.Ctx.Input.Header("Authorization")
	token := strings.Split(authorization, " ")

	logs.Info("Received \nProduct name: ", v.ProductName, "Product Type ID:: ", v.ProductTypeId, "Branch ID:: ", v.BranchId, "Cost price:: ", v.CostPrice, "Image path:: ", v.ImagePath, "Quantity:: ", v.Quantity, "Selling price:: ", v.SellingPrice)

	var isSuccess bool = false

	if token[0] == "Bearer" {
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		if verifyToken.StatusCode == 200 {
			proceed := true
			errorMessage := "An error occurred"
			logs.Info("Token verified!")
			logs.Info("Country is !", verifyToken.User.UserDetails.Branch)
			getBranchResp := functions.GetBranch(&c.Controller, v.BranchId)

			// if getBranchResp.StatusCode == 200 {
			// userId := verifyToken.User.UserId
			if getBranchResp.StatusCode != 200 {
				errorMessage = "Branch provided does not exist"
				proceed = false
			}

			categoryId := strconv.FormatInt(v.ProductTypeId, 10)

			getProductTypes := functions.GetCategory(&c.Controller, categoryId)

			if getProductTypes.StatusCode != 200 {
				errorMessage = "Product type provided does not exist"
				proceed = false
			}

			getItemResp := functions.GetItem(&c.Controller, idStr)

			if getItemResp.StatusCode != 200 {
				errorMessage = "Item provided does not exist"
				proceed = false
			}

			if proceed {
				addItemResp := functions.UpdateItem(&c.Controller, v, getBranchResp.Branch.Country.CountryCode, v.BranchId, int(verifyToken.User.UserId), idStr)

				itemResp := responses.Item{}
				if addItemResp.StatusCode == 200 {
					if addItemResp.Item != nil {
						logs.Info("About to update item image")
						itemImageUpdateResp := functions.UpdateItemImage(&c.Controller, addItemResp.Item.ItemId, v.ImagePath)
						if itemImageUpdateResp.StatusCode == 200 {
							logs.Info("Successfully updated item image")
						} else {
							logs.Error("Failed update")
						}

						itemResp = responses.Item{
							ProductId:        addItemResp.Item.ItemId,
							ProductName:      addItemResp.Item.ItemName,
							Description:      addItemResp.Item.Description,
							ProductType:      addItemResp.Item.Category.CategoryName,
							ProductPrice:     float64(addItemResp.Item.ItemPrice.ItemPrice),
							ProductCostPrice: float64(addItemResp.Item.ItemPrice.AltItemPrice),
							ImagePath:        itemImageUpdateResp.Item.ImagePath,
							Quantity:         addItemResp.Item.Quantity,
							Branch:           addItemResp.Item.Branch,
						}

						isSuccess = true
					} else {
						isSuccess = false
					}

				}

				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &itemResp, StatusDesc: addItemResp.StatusDesc}

				c.Data["json"] = resp
				// } else {
				// 	var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

				// 	c.Data["json"] = resp
				// }
			} else {
				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: errorMessage}

				c.Data["json"] = resp
			}

		} else {
			var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

			c.Data["json"] = resp
		}
	} else {
		logs.Info("Token not verified")
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An error occurred"}

		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UpdateImage ...
// @Title UpdateImage
// @Description Update User's Image
// @Param	Authorization		header 	string true		"header for User"
// @Param	Image		formData 	file	true		"Item Image"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 body is empty
// @router /upload-product-image [post]
func (c *ItemsController) UpdateItemImage() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			image, header, err := c.GetFile("Image")

			if err != nil {
				var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
				c.Data["json"] = resp
			} else {
				logs.Info("Success response received")
				isSuccess = false
				respCode, filePath := functions.SaveImage(&c.Controller, "Image", image, *header)

				if respCode == 200 {
					itemImage := functions.UploadItemImage(&c.Controller, filePath)

					if itemImage.StatusCode == 200 {
						logs.Info("Item image returned: ", itemImage.Value)

						isSuccess = true

						var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: itemImage.Value, StatusDesc: itemImage.StatusDesc}
						c.Data["json"] = resp
					} else {
						var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
						c.Data["json"] = resp
					}
				} else {
					var resp responses.ItemImageResponseDTO = responses.ItemImageResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "Failed to upload file. Tmp"}
					c.Data["json"] = resp
				}

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

// AddProductType ...
// @Title Add Product Type
// @Description Add Product type
// @Param	Authorization		header 	string true		"header for User"
// @Param	CategoryImage		formData 	file	true		"Category Image"
// @Param	CategoryName		formData 	string	true		"Category name"
// @Success 200 {object} responses.CategoryResponseDTO
// @Failure 403 body is empty
// @router /add-product-type [post]
func (c *ItemsController) AddCategory() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			image, header, err := c.GetFile("CategoryImage")

			if err != nil {
				var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "No file uploaded"}
				c.Data["json"] = resp
			} else {
				logs.Info("Success response received")
				isSuccess = false
				respCode, filePath := functions.SaveImage(&c.Controller, "CategoryImage", image, *header)

				if respCode == 200 {

					categoryName := c.Ctx.Input.Query("CategoryName")

					categoryResp := functions.AddCategory(&c.Controller, filePath, categoryName)

					if categoryResp.StatusCode == 200 {

						isSuccess = true

						var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: categoryResp.Category, StatusDesc: categoryResp.StatusDesc}
						c.Data["json"] = resp
					} else {
						var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
						c.Data["json"] = resp
					}
				} else {
					var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. File upload failed"}
					c.Data["json"] = resp
				}
			}
		} else {
			var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.CategoryResponseDTO = responses.CategoryResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetProductTypes ...
// @Title Get Product Types
// @Description Get Product types
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.CategoriesResponseDTO
// @Failure 403 body is empty
// @router /get-product-types [get]
func (c *ItemsController) GetCategories() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			logs.Info("Success response received")
			isSuccess = true

			categoryResponse := functions.GetCategories(&c.Controller)

			if categoryResponse.StatusCode == 200 {
				logs.Info("Product types returned: ", categoryResponse.Categories)

				isSuccess = true

				var resp responses.CategoriesResponseDTO = responses.CategoriesResponseDTO{Success: isSuccess, Result: categoryResponse.Categories, StatusDesc: "Product types fetched successfully"}
				c.Data["json"] = resp
			} else {
				var resp responses.CategoriesResponseDTO = responses.CategoriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}

		} else {
			var resp responses.CategoriesResponseDTO = responses.CategoriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.CategoriesResponseDTO = responses.CategoriesResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetItems ...
// @Title Get Items
// @Description Get Items
// @Param	Authorization		header 	string true		"header for User"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} responses.ItemsResponseDTO
// @Failure 403 body is empty
// @router /get-products [get]
func (c *ItemsController) GetItems() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

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
			logs.Info("Success response received")
			isSuccess = false

			logs.Info("User data received ", verifyToken.User.UserDetails.Branch)

			if verifyToken.User.UserDetails.Branch != nil {
				branchId := strconv.FormatInt(verifyToken.User.UserDetails.Branch.BranchId, 10)

				// Depending on the role, fetch items
				var getItemsResp responses.ItemsOriResponseDTO
				if verifyToken.User.Role.Role == "SUPER_ADMIN" {
					getItemsResp = functions.GetItems(&c.Controller, query, fields, sortby, order, offset, limit)
				} else {
					getItemsResp = functions.GetItemsByBranch(&c.Controller, branchId, query, fields, sortby, order, offset, limit)
				}

				if getItemsResp.StatusCode == 200 {
					logs.Info("Items returned: ", getItemsResp.Items)

					items := []responses.Item{}
					if getItemsResp.Items != nil || len(*getItemsResp.Items) > 0 {
						for _, item := range *getItemsResp.Items {
							itemT := responses.Item{
								ProductId:        item.ItemId,
								ProductName:      item.ItemName,
								Description:      item.Description,
								ProductType:      item.Category.CategoryName,
								ProductPrice:     float64(item.ItemPrice.ItemPrice),
								ProductCostPrice: float64(item.ItemPrice.AltItemPrice),
								ImagePath:        item.ImagePath,
								Quantity:         item.Quantity,
								Branch:           item.Branch,
							}

							items = append(items, itemT)
						}

					} else {
						items = []responses.Item{}
					}

					isSuccess = true

					data := responses.ItemsData{}
					data.Data = &items
					data.Count = len(items)

					var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: &data, StatusDesc: getItemsResp.StatusDesc}
					c.Data["json"] = resp
				} else {
					var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
					c.Data["json"] = resp
				}
			} else {
				logs.Error("User is not linked to a branch")
				var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred. User is not linked to a branch"}
				c.Data["json"] = resp
			}

		} else {
			var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.ItemsResponseDTO = responses.ItemsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetProduct ...
// @Title Get Product
// @Description Get Product
// @Param	Authorization		header 	string true		"header for User"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.ItemResponseDTO
// @Failure 403 body is empty
// @router /get-product/:id [get]
func (c *ItemsController) GetProduct() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			logs.Info("Success response received")
			isSuccess = false

			idStr := c.Ctx.Input.Param(":id")

			itemResp := functions.GetItem(&c.Controller, idStr)

			item := responses.Item{}
			if itemResp.StatusCode == 200 {
				// logs.Info("Categories returned: ", categoryResponse.Categories)
				item = responses.Item{
					ProductId:        itemResp.Item.ItemId,
					ProductName:      itemResp.Item.ItemName,
					Description:      itemResp.Item.Description,
					ProductType:      itemResp.Item.Category.CategoryName,
					ProductPrice:     float64(itemResp.Item.ItemPrice.ItemPrice),
					ProductCostPrice: float64(itemResp.Item.ItemPrice.AltItemPrice),
					ImagePath:        itemResp.Item.ImagePath,
					Quantity:         itemResp.Item.Quantity,
					Branch:           itemResp.Item.Branch,
				}

				isSuccess = true

				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: &item, StatusDesc: itemResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}

		} else {
			var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.ItemResponseDTO = responses.ItemResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
