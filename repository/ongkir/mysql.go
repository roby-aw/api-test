package ongkir

import (
	"api-jasa-pengiriman/business/ongkir"
	"api-jasa-pengiriman/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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

func (repo *MysqlRepository) Cost(Dataongkir *ongkir.Ongkir) (*ongkir.CostResponse, error) {
	var kotaasal ongkir.Kota
	var kotatujuan ongkir.Kota
	err := repo.db.Where("kota_nama = ? AND tipe = ?", Dataongkir.Asal, Dataongkir.Tipe_Asal).Find(&kotaasal).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Where("kota_nama = ? AND tipe = ?", Dataongkir.Tujuan, Dataongkir.Tipe_Tujuan).Find(&kotatujuan).Error
	if err != nil {
		return nil, err
	}
	url := "https://api.rajaongkir.com/starter/cost"
	var queryString string = fmt.Sprintf("origin=%d&destination=%d&weight=%s&courier=jne", kotaasal.Rajaongkir_city_id, kotatujuan.Rajaongkir_city_id, Dataongkir.Berat)
	payload := strings.NewReader(queryString)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", config.GetConfig().RajaOngkir.Api_key)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	hasil := string(body)
	var cost ongkir.CostResponse
	err = json.Unmarshal([]byte(hasil), &cost)
	if err != nil {
		return nil, err
	}
	return &cost, nil
}

func (repo *MysqlRepository) Resi(Dataresi *ongkir.Resi) (*ongkir.CekResiBinderByte, error) {
	url := fmt.Sprintf("http://api.binderbyte.com/v1/track?api_key=%s&courier=%s&awb=%s", config.GetConfig().Binderbyte.Api_key, Dataresi.Kurir, Dataresi.Resi)
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	hasil := string(body)
	var Resi ongkir.CekResiBinderByte
	err = json.Unmarshal([]byte(hasil), &Resi)
	if err != nil {
		return nil, err
	}
	fmt.Println(Resi)
	return &Resi, nil
}
