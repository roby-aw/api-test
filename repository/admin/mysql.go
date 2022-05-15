package admin

import (
	"api-jasa-pengiriman/business/admin"
	"api-jasa-pengiriman/business/ongkir"
	"api-jasa-pengiriman/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (repo *MysqlRepository) FindAdmins() (Admins []admin.Admin, err error) {
	result := repo.db.Find(&Admins)
	if result.Error != nil {
		return nil, result.Error
	}
	return Admins, nil
}

func (repo *MysqlRepository) FindAdminByID(id int) (*admin.Admin, error) {
	var admin *admin.Admin
	err := repo.db.Where("ID = ? ", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo *MysqlRepository) RemoveAdmin(id int) error {
	var admin *admin.Admin
	err := repo.db.Where("ID = ?", id).First(&admin).Error
	if err != nil {
		return err
	}
	fmt.Println(admin)
	err = repo.db.Delete(admin, id).Error
	if err != nil {
		return err
	}
	return err
}
func (repo *MysqlRepository) InsertAdmin(Admins *admin.Admin) (*admin.Admin, error) {
	err := repo.db.Create(&Admins).Error
	if err != nil {
		return nil, fmt.Errorf("failed insert data")
	}
	return Admins, nil
}

func (repo *MysqlRepository) FindAllCity() (city []ongkir.Kota, err error) {
	results := repo.db.Preload("Provinsi").Find(&city)
	if results.Error != nil {
		return nil, err
	}
	return city, nil
}

func (repo *MysqlRepository) CreateToken(Admins *admin.Admin) (string, error) {
	err := repo.db.Where("username =? AND password = ?", Admins.Username, Admins.Password).First(&Admins).Error
	if err != nil {
		return "", err
	}
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &admin.Claims{
		Username: Admins.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	SECRET_KEY := config.GetConfig().Secrettoken.Token
	token_jwt, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token_jwt, err
}

func (repo MysqlRepository) RenewAdmin(id int, admin *admin.Admin) (*admin.Admin, error) {
	err := repo.db.Model(*admin).Where("ID = ?", id).Updates(admin).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Where("ID = ?", id).First(admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (repo MysqlRepository) InsertCity(kota *admin.Kota) (*admin.Kota, error) {
	err := repo.db.Create(&kota).Error
	if err != nil {
		return nil, err
	}
	return kota, err
}

func (repo MysqlRepository) RenewCity(id int, datakota *admin.Kota) (*admin.Kota, error) {
	err := repo.db.Model(datakota).Where("ID = ?", id).Updates(datakota).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Where("ID = ?", id).First(datakota).Error
	if err != nil {
		return nil, err
	}
	return datakota, nil
}

func (repo MysqlRepository) RemoveCity(id int) (*admin.Kota, error) {
	var kota *admin.Kota
	err := repo.db.Where("ID = ?", id).First(&kota).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(kota)
	err = repo.db.Delete(&kota, id).Error
	if err != nil {
		return nil, err
	}
	return kota, nil
}

func (repo MysqlRepository) FindCityByName(data *admin.GetCityById) (*ongkir.Kota, error) {
	var kota *ongkir.Kota
	err := repo.db.Where("Kota_Nama = ? AND Tipe = ?", data.Nama_kota, data.Tipe_kota).First(&kota).Error
	if err != nil {
		return nil, err
	}
	return kota, nil
}
