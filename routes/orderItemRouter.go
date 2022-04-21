package routes

import (
	controller "restaurent-management-golang/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItem-order/:order_id", controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/OrderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/OrderItems/:OrderItem_id", controller.UpdateTable()) //update to food item
}
