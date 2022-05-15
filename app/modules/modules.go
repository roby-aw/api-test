package modules

import (
	"api-jasa-pengiriman/api"
	adminApi "api-jasa-pengiriman/api/admin"
	ongkirApi "api-jasa-pengiriman/api/ongkir"
	adminBusiness "api-jasa-pengiriman/business/admin"
	ongkirBusiness "api-jasa-pengiriman/business/ongkir"
	"api-jasa-pengiriman/config"
	adminRepo "api-jasa-pengiriman/repository/admin"
	ongkirRepo "api-jasa-pengiriman/repository/ongkir"
	"api-jasa-pengiriman/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {
	ongkirPermitRepository := ongkirRepo.RepositoryFactory(dbCon)
	ongkirPermitService := ongkirBusiness.NewService(ongkirPermitRepository)
	ongkirPermitController := ongkirApi.NewController(ongkirPermitService)

	adminPermitRepository := adminRepo.RepositoryFactory(dbCon)
	adminPermitService := adminBusiness.NewService(adminPermitRepository)
	adminPermitController := adminApi.NewController(adminPermitService)

	controller := api.Controller{
		OngkirController: ongkirPermitController,
		AdminControlller: adminPermitController,
	}
	return controller
}
