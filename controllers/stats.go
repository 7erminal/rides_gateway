package controllers

import (
	"AMC_gateway/controllers/functions"
	"AMC_gateway/structs/responses"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// StatsController operations for Stats
type StatsController struct {
	beego.Controller
}

// URLMapping ...
func (c *StatsController) URLMapping() {
	c.Mapping("GetGeneralStats", c.GetGeneralStats)
}

// GetGeneralStats ...
// @Title Get Stats Per Branch
// @Description Get Stats of Branch
// @Param	Authorization		header 	string true		"header for User"
// @Success 200 {object} responses.ItemsResponseDTO
// @Failure 403 body is empty
// @router /get-stats [get]
func (c *StatsController) GetGeneralStats() {
	authorization := c.Ctx.Input.Header("Authorization")

	token := strings.Split(authorization, " ")

	var isSuccess bool = false

	if token[0] == "Bearer" {
		logs.Info("Token is ", token[1])
		verifyToken := functions.VerifyToken(&c.Controller, token[1])

		logs.Info("Success response")

		if verifyToken.StatusCode == 200 {
			logs.Info("Success response received")
			branchidStr := strconv.FormatInt(verifyToken.User.UserDetails.Branch.BranchId, 10)

			getItemStatsResp := functions.GetItemStats(&c.Controller, branchidStr)

			if getItemStatsResp.StatusCode == 200 {
				logs.Info("Item stats returned: ")
				isSuccess = true

				itemStats := responses.StatsDTO{}
				if getItemStatsResp.Stats != nil {

					itemStats = *getItemStatsResp.Stats
				}

				isSuccess = true

				var resp responses.ItemsStatsResponseDTO = responses.ItemsStatsResponseDTO{Success: isSuccess, Result: &itemStats, StatusDesc: getItemStatsResp.StatusDesc}
				c.Data["json"] = resp
			} else {
				var resp responses.ItemsStatsResponseDTO = responses.ItemsStatsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
				c.Data["json"] = resp
			}

		} else {
			var resp responses.ItemsStatsResponseDTO = responses.ItemsStatsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
			c.Data["json"] = resp
		}
	} else {
		var resp responses.ItemsStatsResponseDTO = responses.ItemsStatsResponseDTO{Success: isSuccess, Result: nil, StatusDesc: "An Error occurred"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
