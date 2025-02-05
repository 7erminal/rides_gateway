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

func AddItem(c *beego.Controller, req requests.AddItemRequestDTO, productTypeId int64, countryCode string, branchId int64, addedBy int) (resp responses.ItemOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	// logs.Info("Sending first name ", req.BranchId)

	request := api.NewRequest(
		host,
		"/v1/items/",
		api.POST)
	request.InterfaceParams["ItemName"] = req.ProductName
	request.InterfaceParams["Description"] = ""
	request.InterfaceParams["Weight"] = ""
	request.InterfaceParams["Category"] = productTypeId
	request.InterfaceParams["AvailableSizes"] = ""
	request.InterfaceParams["AvailableColors"] = ""
	request.InterfaceParams["Quantity"] = req.Quantity
	request.InterfaceParams["ItemPrice"] = req.SellingPrice
	request.InterfaceParams["AltItemPrice"] = req.CostPrice
	request.InterfaceParams["Country"] = countryCode
	request.InterfaceParams["Branch"] = branchId
	request.InterfaceParams["CreatedBy"] = addedBy
	request.InterfaceParams["QuantityAlert"] = req.ReorderLevel

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
	var data responses.ItemOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateItem(c *beego.Controller, req requests.UpdateItemRequestDTO, countryCode string, branchId int64, addedBy int, itemId string) (resp responses.ItemOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	// logs.Info("Sending first name ", req.BranchId)

	request := api.NewRequest(
		host,
		"/v1/items/",
		api.PUT)
	request.InterfaceParams["ItemName"] = req.ProductName
	request.InterfaceParams["Description"] = ""
	request.InterfaceParams["Weight"] = ""
	request.InterfaceParams["Category"] = req.ProductTypeId
	request.InterfaceParams["AvailableSizes"] = ""
	request.InterfaceParams["AvailableColors"] = ""
	request.InterfaceParams["Quantity"] = req.Quantity
	request.InterfaceParams["ItemPrice"] = req.SellingPrice
	request.InterfaceParams["AltItemPrice"] = req.CostPrice
	request.InterfaceParams["Country"] = countryCode
	request.InterfaceParams["Branch"] = branchId
	request.InterfaceParams["CreatedBy"] = addedBy

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
	var data responses.ItemOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateItemImage(c *beego.Controller, itemId int64, imagePath string) (resp responses.ItemOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	// logs.Info("Sending first name ", req.BranchId)

	item_id := strconv.FormatInt(itemId, 10)

	request := api.NewRequest(
		host,
		"/v1/items/update-item-image/"+item_id,
		api.PUT)
	request.InterfaceParams["ImagePath"] = imagePath

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
	var data responses.ItemOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetItems(c *beego.Controller, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.ItemsOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/items/",
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.ItemsOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetItem(c *beego.Controller, itemId string) (resp responses.ItemOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/items/"+itemId,
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.ItemOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetItemsByBranch(c *beego.Controller, branchId string, query string, fields string, sortby string, order string,
	offset string, limit string) (resp responses.ItemsOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/items/branch/"+branchId,
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.ItemsOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetItemStats(c *beego.Controller, branchId string) (resp responses.ItemsStatsOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/items/get-item-stats/"+branchId,
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.ItemsStatsOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func AddCategory(c *beego.Controller, categoryImage string, categoryName string) (resp responses.CategoryOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	logs.Info("Sending file ", categoryImage)

	request := api.NewRequest(
		host,
		"/v1/categories/",
		api.POST)

	request.FileField["Image"] = categoryImage
	request.Params["CategoryName"] = categoryName
	request.Params["Icon"] = ""
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
	var data responses.CategoryOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UploadItemImage(c *beego.Controller, itemImage string) (resp responses.ItemImageOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	logs.Info("Sending file ", itemImage)

	request := api.NewRequest(
		host,
		"/v1/item-images/upload-pictures",
		api.POST)

	request.FileField["Image"] = itemImage
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
	var data responses.ItemImageOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCategories(c *beego.Controller) (resp responses.CategoriesOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/categories/",
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.CategoriesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCategory(c *beego.Controller, categoryId string) (resp responses.CategoryOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/categories/"+categoryId,
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.CategoryOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCategoryByName(c *beego.Controller, category string) (resp responses.CategoryOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/categories/name/"+category,
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.CategoryOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetItemImages(c *beego.Controller) (resp responses.ItemImagesOriResponseDTO) {
	host, _ := beego.AppConfig.String("itemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/item-images/",
		api.GET)

	// request.FileField["UserImage"] = userImage
	// request.Params["UserId"] = strconv.FormatInt(userId, 10)
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
	var data responses.ItemImagesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}
