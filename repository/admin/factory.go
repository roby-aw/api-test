package admin

import (
	"api-jasa-pengiriman/business/admin"
	"api-jasa-pengiriman/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) admin.Repository {
	adminRepo := NewMysqlRepository(dbCon.Mysql)
	return adminRepo
}
