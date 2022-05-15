package admin

import (
	"api-jasa-pengiriman/business/admin"
	adminBusiness "api-jasa-pengiriman/business/admin"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service adminBusiness.Service
}

func NewController(service adminBusiness.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Get All admin
// @description Get all admin with data
// @tags Admin using Token JWT
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} []admin.Admin
// @Router /admin [get]
func (Controller *Controller) GetAdmins(c echo.Context) error {
	users, err := Controller.service.GetAdmins()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// Create godoc
// @Summary Get Admin By ID
// @description Get Admin By ID
// @tags admin
// @Accept json
// @Produce json
// @Param id path int true "anything id"
// @Success 200 {object} admin.Admin
// @Router /admin/{id} [get]
func (Controller *Controller) GetAdminByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	admin, err := Controller.service.GetAdminByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, admin)
}

// Create godoc
// @Summary Create admin
// @description create admin with data
// @tags admin
// @Accept json
// @Produce json
// @Param admin body admin.AdminSwagger true "admin"
// @Success 201 {object} admin.Admin
// @Failure 400 {object} map[string]interface{}
// @Router /admin [post]
func (Controller *Controller) CreateAdmin(c echo.Context) error {
	admin := admin.Admin{}
	c.Bind(&admin)
	admins, err := Controller.service.CreateAdmin(&admin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed",
			"Error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, admins)
}

// Create godoc
// @Summary Get All City
// @description Get all city for admin using jwt
// @tags Admin using Token JWT
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} []ongkir.Kota
// @Failure 400
// @Router /admin/city [get]
func (Controller *Controller) GetAllCity(c echo.Context) error {
	users, err := Controller.service.GetAllCity()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

// Create godoc
// @Summary Get Token
// @description Get token for admin
// @tags admin
// @Accept json
// @Produce json
// @Param admin body admin.InputAdminToken true "admin"
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/token [post]
func (Controller *Controller) GetToken(c echo.Context) error {
	var request adminBusiness.Admin

	c.Bind(&request)
	token, err := Controller.service.GetToken(&request)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"token":   token,
	})
}

// Create godoc
// @Summary Delete Admin
// @description delete data admin
// @tags admin
// @Accept json
// @Produce json
// @Param id path int true "id admin"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/{id} [delete]
func (Controller *Controller) DeleteAdmin(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := Controller.service.DeleteAdmin(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"messages": "success delete admin",
	})
}

// Create godoc
// @Summary Update Admin
// @description update data admin
// @tags Admin using Token JWT
// @Accept json
// @Produce json
// @Param id path int true "id admin"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param admin body admin.AdminSwagger true "admin"
// @Success 200 {object} map[string]interface{}
// @Failure 400
// @Router /admin/{id} [PUT]
func (Controller *Controller) UpdateAdmin(c echo.Context) error {
	var admin *admin.Admin
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&admin)
	admin, err := Controller.service.UpdateAdmin(id, admin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"data":     admin,
	})
}

// Create godoc
// @Summary Create Data City
// @description Create data city for admin
// @tags Admin using Token JWT
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param admin body adminBusiness.Kota true "Kota"
// @Success 200 {object} adminBusiness.Kota
// @Failure 400
// @Router /admin/city [post]
func (Controller *Controller) CreateCity(c echo.Context) error {
	var kota *adminBusiness.Kota
	c.Bind(&kota)
	kota, err := Controller.service.CreateCity(kota)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"data":     kota,
	})
}

// Create godoc
// @Summary Get data city by name dan tipe
// @description Get data city by name dan tipe for public
// @tags city
// @Accept json
// @Produce json
// @Param data body admin.GetCityById true "Kota"
// @Success 200 {object} adminBusiness.Kota
// @Failure 400
// @Router /city [post]
func (Controller *Controller) GetCityByName(c echo.Context) error {
	data := admin.GetCityById{}
	c.Bind(&data)
	kota, err := Controller.service.GetCityByName(&data)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, kota)
}

// Create godoc
// @Summary Update Data City
// @description Update data city for admin
// @tags Admin using Token JWT
// @Accept json
// @Produce json
// @Param id path int true "id city"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param admin body adminBusiness.Kota true "Kota"
// @Success 200 {object} adminBusiness.Kota
// @Failure 400
// @Router /admin/city/{id} [put]
func (Controller *Controller) UpdateCity(c echo.Context) error {
	var datakota *admin.Kota
	id, _ := strconv.Atoi(c.Param("id"))
	c.Bind(&datakota)
	datakota, err := Controller.service.UpdateCity(id, datakota)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"data":     datakota,
	})
}

// Create godoc
// @Summary Delete Data City
// @description Delete data city for admin
// @tags Admin using Token JWT
// @Accept json
// @Produce json
// @Param id path int true "id city"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} adminBusiness.Kota
// @Failure 400
// @Router /admin/city/{id} [delete]
func (Controller *Controller) RemoveCity(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	kota, err := Controller.service.DeleteCity(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success deleted kota",
		"data":     kota,
	})
}
