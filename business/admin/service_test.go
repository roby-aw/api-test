package admin_test

import (
	"api-jasa-pengiriman/business/admin"
	"api-jasa-pengiriman/business/ongkir"
	"errors"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var service admin.Service
var admin1, admin2, admin3, UpdateAdmin admin.Admin
var kota1, kota2, kota3 ongkir.Kota
var adminkota1, adminkota2, adminkota3, insertadminkota, updateadminkota admin.Kota
var insertkota, updatekota, failedkota ongkir.Kota

var insertSpec, updateSpec, failedSpec, errorspec admin.Admin
var errorFindID int
var getkotabyid admin.GetCityById

// var creator, updater string
var errorInsert error = errors.New("error on insert")
var errorFind error = errors.New("error on find")

func TestMain(m *testing.M) {
	setup()
	setupCity()
	os.Exit(m.Run())
}

func TestGetAdminByID(t *testing.T) {
	t.Run("Expect found the content", func(t *testing.T) {
		foundContent, _ := service.GetAdminByID(admin1.ID)
		if !reflect.DeepEqual(*foundContent, admin1) {
			t.Error("Expect content has to be equal with content1", foundContent, admin1)
		}
	})

	t.Run("Expect not found the content", func(t *testing.T) {
		content, err := service.GetAdminByID(90)

		if err != nil {
			t.Error("Expect error is nil. Error: ", err)
		} else if content != nil {
			t.Error("Expect content must be not found (nil)")
		}
	})
}

func TestGetAdminAll(t *testing.T) {
	t.Run("Expect found the admin", func(t *testing.T) {
		admins, _ := service.GetAdmins()

		if len(admins) != 3 {
			t.Error("Expect admin length must be 3")
			t.FailNow()
		}

		if reflect.DeepEqual(admins[0], admin1) {
			if !reflect.DeepEqual(admins[1], admin2) {
				t.Error("Expect 2nd admin is equal to admin2")
			}
		} else if reflect.DeepEqual(admins[0], admin2) {
			if !reflect.DeepEqual(admins[1], admin1) {
				t.Error("Expect 2nd admin is equal to admin")
			}
		} else {
			t.Error("Expect admin is admin1 and admin2")
		}
	})
}

func TestDeleteAdmin(t *testing.T) {
	t.Run("Expect delete admin2", func(t *testing.T) {
		err := service.DeleteAdmin(admin2.ID)
		if err == fmt.Errorf("") {
			t.Error("error delete")
		}
		admin, _ := service.GetAdmins()
		if len(admin) != 2 {
			t.Error("Error")
		}
	})
}
func TestInsertAdmin(t *testing.T) {
	t.Run("Expect Insert the admin", func(t *testing.T) {
		createadmins, err := service.CreateAdmin(&insertSpec)
		if err != nil {
			t.Error("Cannot insert admin")
		}
		if createadmins.ID != 3 {
			t.Error("Expect 2nd admin is equal to admin2")
		}
		NewAdmins, _ := service.GetAdminByID(createadmins.ID)
		if NewAdmins == nil {
			t.Error("Expect admins is not nil after inserted")
			t.FailNow()
		}
		GetAllAdmin, _ := service.GetAdmins()
		if reflect.DeepEqual(GetAllAdmin[2], createadmins) {
			if !reflect.DeepEqual(GetAllAdmin[0], admin1) {
				t.Error("Expect 3rd admin is equal to insertadmin")
			}
		}
		if reflect.DeepEqual(GetAllAdmin[2].ID, createadmins.ID) {
			if !reflect.DeepEqual(GetAllAdmin[0], admin1) {
				t.Error("Expect 1st admin is equal to admin1")
			}
		}
	})
}

func TestUpdateAdmin(t *testing.T) {
	t.Run("Expect Insert the admin", func(t *testing.T) {
		fmt.Println(updateSpec)
		admin, err := service.UpdateAdmin(updateSpec.ID, &updateSpec)
		fmt.Println(admin)
		if err != nil {
			t.Error("Error Update")
		}
		if admin.Name != updateSpec.Name {
			t.Error("Name tidak update")
		}
		if admin.Username != updateSpec.Username {
			t.Error("Username tidak update")
		}
		if admin.Email != updateSpec.Email {
			t.Error("Email tidak update")
		}
		if admin.Password != updateSpec.Password {
			t.Error("Password tidak update")
		}
	})
}

func TestGetCityAll(t *testing.T) {
	t.Run("Expect found the All City", func(t *testing.T) {
		city, _ := service.GetAllCity()

		if len(city) != 3 {
			t.Error("Expect city lenght must be 3")
		}
		if reflect.DeepEqual(city[0], kota1) {
			if !reflect.DeepEqual(city[1], kota2) {
				t.Error("Expect city1 is kota1")
			}
		}
	})
}

func TestRemoveCity(t *testing.T) {
	t.Run("Test Update Kota", func(t *testing.T) {
		kota, err := service.DeleteCity(adminkota2.ID)
		fmt.Println(kota)
		if err != nil {
			t.Error("Delete Error")
		}
		if kota.Kota_Nama != adminkota2.Kota_Nama {
			t.Error("Data Error")
		}
	})
}
func TestInsertCityAdmin(t *testing.T) {
	t.Run("Expect found the All City", func(t *testing.T) {
		kota, _ := service.CreateCity(&insertadminkota)
		fmt.Println(kota.Kota_Nama)
		if kota.Kota_Nama != insertadminkota.Kota_Nama {
			t.Error("Error")
		}
	})
}

func TestUpdateKota(t *testing.T) {
	t.Run("Test Update Kota", func(t *testing.T) {
		kota, _ := service.UpdateCity(updateadminkota.ID, &updateadminkota)
		if kota.Kota_Nama != updateadminkota.Kota_Nama {
			t.Error("Error")
		}
	})
}

func TestGetKotaByID(t *testing.T) {
	t.Run("Test get kota by id", func(t *testing.T) {
		hasil, err := service.GetCityByName(&getkotabyid)
		if err != nil {
			t.Error("terdapat error")
		}
		fmt.Println(hasil.Kota_Nama)
		fmt.Println(getkotabyid.Nama_kota)
		if hasil.Kota_Nama != getkotabyid.Nama_kota {
			t.Error("data tidak ditemukan")
		}
	})
}

func setup() {
	//initialize admin 1
	admin1.ID = 1
	admin1.Email = "testemail@gmail.com"
	admin1.Username = "testusername"
	admin1.Name = "testname"
	admin1.Password = "testpassword"

	//initialize admin 2
	admin2.ID = 2
	admin2.Email = "testemail2@gmail.com"
	admin2.Username = "testusername2"
	admin2.Name = "testname2"
	admin2.Password = "testpassword2"

	//initialize admin 3
	admin3.ID = 3
	admin3.Email = "testemail3@gmail.com"
	admin3.Username = "testusername3"
	admin3.Name = "testname3"
	admin3.Password = "testpassword3"

	repo := newInMemoryRepository()
	service = admin.NewService(&repo)

	insertSpec.ID = 3
	insertSpec.Name = "InsertName"
	insertSpec.Email = "insertemail@gmail.com"
	insertSpec.Username = "InsertUsername"
	insertSpec.Password = "InsertPassword"

	updateSpec.ID = 1
	updateSpec.Name = "updatename"
	updateSpec.Email = "updateemail@gmail.com"
	updateSpec.Username = "UpdateUsername"
	updateSpec.Password = "updatePassword"

	// failedSpec.Name = ""
	// failedSpec.Description = "Failed Description"
	// failedSpec.Tags = []string{}

	// errorSpec.Name = "Error Content"
	// errorSpec.Description = "Error Description"
	// errorSpec.Tags = []string{}

	// creator = "creator"
	// updater = "updater"

	errorFindID = 3235
}

func setupCity() {
	//initialize admin 1
	kota1.ID = 1
	kota1.Rajaongkir_city_id = 1
	kota1.Kota_Nama = "KotaTest1"
	kota1.Postal_code = "15534"
	kota1.Province_ID = 1

	//initialize admin 2
	kota2.ID = 2
	kota2.Rajaongkir_city_id = 2
	kota2.Kota_Nama = "KotaTest2"
	kota2.Postal_code = "15534"
	kota2.Province_ID = 2

	//initialize admin 3
	kota3.ID = 3
	kota3.Rajaongkir_city_id = 3
	kota3.Kota_Nama = "KotaTest3"
	kota3.Postal_code = "15534"
	kota3.Province_ID = 3

	//initialize adminkota1
	adminkota1.ID = 1
	adminkota1.Rajaongkir_city_id = 1
	adminkota1.Kota_Nama = "testname1"
	adminkota1.Postal_code = 15530
	adminkota1.Province_ID = 1
	adminkota1.Tipe = "testipe"

	adminkota2.ID = 2
	adminkota2.Rajaongkir_city_id = 2
	adminkota2.Kota_Nama = "testname2"
	adminkota2.Postal_code = 15530
	adminkota2.Province_ID = 2
	adminkota2.Tipe = "testipe"

	adminkota3.ID = 3
	adminkota3.Rajaongkir_city_id = 3
	adminkota3.Kota_Nama = "testname3"
	adminkota3.Postal_code = 15530
	adminkota3.Province_ID = 3
	adminkota3.Tipe = "testipe"

	repo := newInMemoryRepository()
	service = admin.NewService(&repo)

	insertadminkota.ID = 4
	insertadminkota.Rajaongkir_city_id = 4
	insertadminkota.Kota_Nama = "insertnamakota"
	insertadminkota.Postal_code = 153363
	insertadminkota.Province_ID = 4
	insertadminkota.Tipe = "inserttipe"

	updateadminkota.ID = 1
	updateadminkota.Rajaongkir_city_id = 4
	updateadminkota.Kota_Nama = "updatenamakota"
	updateadminkota.Postal_code = 153363
	updateadminkota.Province_ID = 4
	updateadminkota.Tipe = "updatetipe"

	getkotabyid.Nama_kota = "KotaTest1"
	getkotabyid.Tipe_kota = "testtipe"

	errorFindID = 3235
}

type inMemoryRepository struct {
	adminByID        map[int]admin.Admin
	AllAdmin         []admin.Admin
	AllCity          []ongkir.Kota
	AllAdminKota     []admin.Kota
	AllAdminKotaByID map[int]admin.Kota
}

func newInMemoryRepository() inMemoryRepository {
	var repo inMemoryRepository
	repo.adminByID = make(map[int]admin.Admin)
	repo.adminByID[admin1.ID] = admin1
	repo.adminByID[admin2.ID] = admin2
	repo.adminByID[admin3.ID] = admin3

	repo.AllAdmin = []admin.Admin{}
	repo.AllAdmin = append(repo.AllAdmin, admin1)
	repo.AllAdmin = append(repo.AllAdmin, admin2)
	repo.AllAdmin = append(repo.AllAdmin, admin3)

	repo.AllCity = []ongkir.Kota{}
	repo.AllCity = append(repo.AllCity, kota1)
	repo.AllCity = append(repo.AllCity, kota2)
	repo.AllCity = append(repo.AllCity, kota3)

	repo.AllAdminKota = []admin.Kota{}
	repo.AllAdminKota = append(repo.AllAdminKota, adminkota1)
	repo.AllAdminKota = append(repo.AllAdminKota, adminkota2)
	repo.AllAdminKota = append(repo.AllAdminKota, adminkota3)

	repo.AllAdminKotaByID = make(map[int]admin.Kota)
	repo.AllAdminKotaByID[adminkota1.ID] = adminkota1
	repo.AllAdminKotaByID[adminkota2.ID] = adminkota2
	repo.AllAdminKotaByID[adminkota3.ID] = adminkota3

	return repo
}

func (repo *inMemoryRepository) FindAdminByID(id int) (*admin.Admin, error) {
	if id == errorFindID {
		return nil, errorFind
	}

	content, ok := repo.adminByID[id]
	if !ok {
		return nil, nil
	}

	return &content, nil
}

func (repo *inMemoryRepository) FindAdmins() (admins []admin.Admin, err error) {
	admins = repo.AllAdmin
	return admins, err
}

func (repo *inMemoryRepository) InsertAdmin(Admins *admin.Admin) (*admin.Admin, error) {
	if Admins.Name == errorspec.Name {
		return nil, errorInsert
	}
	repo.AllAdmin = append(repo.AllAdmin, *Admins)
	repo.adminByID[Admins.ID] = *Admins

	return Admins, nil
}

func (repo *inMemoryRepository) RemoveAdmin(id int) error {
	id = id - 1
	repo.AllAdmin = append(repo.AllAdmin[:id], repo.AllAdmin[id+1:]...)
	return fmt.Errorf("")
}

func (repo *inMemoryRepository) RenewAdmin(id int, admin *admin.Admin) (*admin.Admin, error) {
	admins, ok := repo.adminByID[id]
	fmt.Println(admins)
	if !ok {
		return nil, nil
	}
	admins.Email = admin.Email
	admins.Username = admin.Username
	admins.Name = admin.Name
	admins.Password = admin.Password

	return &admins, nil
}

func (repo *inMemoryRepository) FindAllCity() (city []ongkir.Kota, err error) {
	city = repo.AllCity
	return city, err
}

func (repo *inMemoryRepository) InsertCity(kota *admin.Kota) (*admin.Kota, error) {
	repo.AllAdminKota = append(repo.AllAdminKota, insertadminkota)
	fmt.Println(kota)
	return kota, fmt.Errorf("")
}

func (repo *inMemoryRepository) CreateToken(Admins *admin.Admin) (string, error) {
	return "", fmt.Errorf("")
}

func (repo *inMemoryRepository) RenewCity(id int, datakota *admin.Kota) (*admin.Kota, error) {
	kota, ok := repo.AllAdminKotaByID[id]
	if !ok {
		return nil, nil
	}
	kota.Kota_Nama = datakota.Kota_Nama
	kota.Rajaongkir_city_id = datakota.Rajaongkir_city_id
	kota.Postal_code = datakota.Postal_code
	kota.Tipe = datakota.Tipe
	kota.Province_ID = datakota.Province_ID
	return &kota, fmt.Errorf("")
}

func (repo *inMemoryRepository) RemoveCity(id int) (*admin.Kota, error) {
	id = id - 1
	var deletecity admin.Kota
	deletecity = repo.AllAdminKota[id]
	repo.AllAdminKota = append(repo.AllAdminKota[:id], repo.AllAdminKota[id+1:]...)
	if len(repo.AllAdminKota) != 2 {
		return nil, fmt.Errorf("Len must be 2")
	}
	return &deletecity, nil
}

func (repo *inMemoryRepository) FindCityByName(data *admin.GetCityById) (*ongkir.Kota, error) {
	for _, hasil := range repo.AllCity {
		fmt.Println(hasil.Kota_Nama)
		fmt.Println(data.Nama_kota)
		if hasil.Kota_Nama == data.Nama_kota {
			return &hasil, nil
		}
	}
	return nil, fmt.Errorf("data tidak ditemukan")
}
