package api

import (
	"api-jasa-pengiriman/api/admin"
	auth "api-jasa-pengiriman/api/middleware"
	"api-jasa-pengiriman/api/ongkir"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	OngkirController *ongkir.Controller
	AdminControlller *admin.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	// ongkir
	e.POST("/cost", controller.OngkirController.GetCost)
	e.POST("/cekresi", controller.OngkirController.GetResi)
	// public
	e.POST("/city", controller.AdminControlller.GetCityByName)
	//admin
	g := e.Group("/admin")
	g.POST("", controller.AdminControlller.CreateAdmin)
	g.POST("/token", controller.AdminControlller.GetToken)
	g.PUT("/:id", controller.AdminControlller.UpdateAdmin)
	g.GET("/:id", controller.AdminControlller.GetAdminByID)
	// admin using jwt
	g.DELETE("/:id", controller.AdminControlller.DeleteAdmin, auth.SetupAuthenticationJWT())
	g.DELETE("/city/:id", controller.AdminControlller.RemoveCity, auth.SetupAuthenticationJWT())
	g.POST("/city", controller.AdminControlller.CreateCity, auth.SetupAuthenticationJWT())
	g.PUT("/city/:id", controller.AdminControlller.UpdateCity, auth.SetupAuthenticationJWT())
	g.GET("", controller.AdminControlller.GetAdmins, auth.SetupAuthenticationJWT())
	g.GET("/city", controller.AdminControlller.GetAllCity, auth.SetupAuthenticationJWT())
}
