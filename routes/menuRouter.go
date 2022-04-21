package routes

import(
	controller "restaurent-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/Menus", controller.GetMenus())
	incomingRoutes.GET("/Menus/:Menus_id", controller.GetMenu())
	incomingRoutes.POST("/Menus", controller.CreateMenu())
	incomingRoutes.PATCH("/Menus/:Menu_id", controller.UpdateMenu()) //update to food item
}