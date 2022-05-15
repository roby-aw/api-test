package ongkir

import (
	"api-jasa-pengiriman/business/ongkir"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service ongkir.Service
}

func NewController(service ongkir.Service) *Controller {
	return &Controller{
		service: service,
	}
}

// Create godoc
// @Summary Get Cost Ongkir
// @description Get Cost Ongkir from raja ongkir
// @tags Get ongkir and check resi
// @Accept json
// @Produce json
// @Param Ongkir body ongkir.Ongkir true "ongkir"
// @Success 200 {object} []ongkir.Results
// @Failure 400 {object} map[string]interface{}
// @Router /cost [post]
func (Controller *Controller) GetCost(c echo.Context) error {
	ongkir := ongkir.Ongkir{}
	c.Bind(&ongkir)
	ongkirs, err := Controller.service.GetCost(&ongkir)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, ongkirs.Rajaongkir.Results)
}

// Create godoc
// @Summary Get Detail Resi
// @description Get detail resi from binderbyte
// @tags Get ongkir and check resi
// @Accept json
// @Produce json
// @Param Resi body ongkir.Resi true "Resi"
// @Success 200 {object} []ongkir.CekResiBinderByte
// @Failure 400 {object} map[string]interface{}
// @Router /cekresi [post]
func (Controller *Controller) GetResi(c echo.Context) error {
	Resi := ongkir.Resi{}
	c.Bind(&Resi)
	HasilResi, err := Controller.service.GetResi(&Resi)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, HasilResi.Data)
}
