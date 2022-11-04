package v1

import (
	util "gin_consult/pkg/utils"
	"gin_consult/service"
	"github.com/gin-gonic/gin"
)

func ShowMoney(c *gin.Context) {
	util.LogrusObj.Infoln("ces")
	showMoneyService := service.ShowMoneyService{}
	claim, _ := util.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showMoneyService); err == nil {
		res := showMoneyService.Show(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		util.LogrusObj.Infoln(err)
	}
}
