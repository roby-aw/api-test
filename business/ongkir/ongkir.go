package ongkir

type Kota struct {
	ID                 uint      `gorm:"primaryKey"`
	Rajaongkir_city_id int       `json:"rajaongkir_city_id"`
	Kota_Nama          string    `json:"kota_nama"`
	Postal_code        string    `json:"postal_code"`
	Tipe               string    `json:"tipe"`
	Province_ID        int       `json:"province_id"`
	Provinsi           Provinces `json:"provinces" gorm:"foreignKey:Province_ID;references:ID"`
}

type Provinces struct {
	ID       uint   `gorm:"primaryKey"`
	Province string `json:"provinsi_nama"`
}

type Ongkir struct {
	Asal        string `json:"asal" validate:"required"`
	Tipe_Asal   string `json:"tipe_asal" validate:"required"`
	Tujuan      string `json:"tujuan" validate:"required"`
	Tipe_Tujuan string `json:"tipe_tujuan" validate:"required"`
	Berat       string `json:"berat" validate:"required"`
}

type CostResponse struct {
	Rajaongkir struct {
		Origin_Details Origin_Details
		Results        []Results `json:"results"`
	} `json:"rajaongkir"`
}
type Results struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Costs []Cost `json:"costs"`
}

type Cost struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []struct {
		Value int    `json:"value"`
		ETD   string `json:"etd"`
		Note  string `json:"note"`
	} `json:"cost"`
}

type Resi struct {
	Resi  string `json:"resi" validate:"required"`
	Kurir string `json:"kurir" validate:"required"`
}

type CekResiBinderByte struct {
	Status  int    `json:"status"`
	Message string `json:"Message"`
	Data    struct {
		Summary struct {
			Awb     string `json:"awb"`
			Courier string `json:"courier"`
			Service string `json:"service"`
			Status  string `json:"status"`
			Weight  string `json:"weight"`
		} `json:"summary"`
		History []History `json:"history"`
	} `json:"data"`
}
type History struct {
	Date string `json:"date"`
	Desc string `json:"desc"`
}

type Origin_Details struct {
	City_id     string
	Province_id string
	Province    string
	Tipe        string
	City_name   string
	Postal_code string
}
