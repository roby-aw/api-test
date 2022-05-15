package ongkir_test

import (
	"api-jasa-pengiriman/business/ongkir"
	"os"
	"testing"
)

var service ongkir.Service
var resi string
var hasilongkir ongkir.CostResponse
var hasilresi ongkir.CekResiBinderByte
var iniresult ongkir.Results
var insertdataongkir ongkir.Ongkir
var dataresi ongkir.Resi

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestGetCost(t *testing.T) {
	t.Run("Expect get cost", func(t *testing.T) {
		hasil, _ := service.GetCost(&insertdataongkir)
		if hasil.Rajaongkir.Origin_Details.City_name != insertdataongkir.Asal {
			t.Error("Error data")
		}
		if hasil.Rajaongkir.Origin_Details.City_name != insertdataongkir.Tujuan {
			t.Error("Error data")
		}
	})
}

func TestGetResi(t *testing.T) {
	t.Run("Expect get cost", func(t *testing.T) {
		hasil, _ := service.GetResi(&dataresi)
		if hasil.Message != "Success" {
			t.Error("Failed get resi")
		}
	})
}

func setup() {
	hasilongkir = ongkir.CostResponse{
		struct {
			Origin_Details ongkir.Origin_Details
			Results        []ongkir.Results "json:\"results\""
		}{
			Origin_Details: ongkir.Origin_Details{
				City_id:     "1",
				Province_id: "1",
				Province:    "testoriginprovince",
				Tipe:        "testtipe",
				City_name:   "testorigincity",
				Postal_code: "35235",
			},
			Results: []ongkir.Results{
				{
					Code: "TestCode",
					Name: "TestName",
					Costs: []ongkir.Cost{
						{
							Service:     "testservice",
							Description: "testdescription",
							Cost: []struct {
								Value int    "json:\"value\""
								ETD   string "json:\"etd\""
								Note  string "json:\"note\""
							}{
								{
									Value: 1,
									ETD:   "string",
									Note:  "note",
								},
							},
						},
					},
				},
			},
		},
	}
	resi = "3123"
	hasilresi.Message = "Success"
	hasilresi.Status = 200
	hasilresi.Data = struct {
		Summary struct {
			Awb     string "json:\"awb\""
			Courier string "json:\"courier\""
			Service string "json:\"service\""
			Status  string "json:\"status\""
			Weight  string "json:\"weight\""
		} "json:\"summary\""
		History []ongkir.History "json:\"history\""
	}{
		Summary: struct {
			Awb     string "json:\"awb\""
			Courier string "json:\"courier\""
			Service string "json:\"service\""
			Status  string "json:\"status\""
			Weight  string "json:\"weight\""
		}{
			Awb: "test",
		},
	}
	// Cost := ongkir.Cost.Cost{1, "test", "note"}
	// costs == ongkir.Cost{"testservice", "testdescription", []cost{}}
	// iniresi := ongkir.Results{"inicode", "test", _}
	// resi = "iniresi"
	// hasilongkir = hasilongkir.Rajaongkir.Results
	// result.Code = "3123"
	repo := newInMemoryRepository()
	service = ongkir.NewService(&repo)

	insertdataongkir.Asal = "testorigincity"
	insertdataongkir.Tipe_Asal = "test"
	insertdataongkir.Tujuan = "testorigincity"
	insertdataongkir.Tipe_Tujuan = "tipetujuan"
	insertdataongkir.Berat = "testberat"

	dataresi.Resi = "3123"
	dataresi.Kurir = "test"
}

type inMemoryRepository struct {
	DataOngkir map[string]ongkir.CostResponse

	DataResi map[string]ongkir.CekResiBinderByte
}

func newInMemoryRepository() inMemoryRepository {
	var repo inMemoryRepository
	repo.DataOngkir = make(map[string]ongkir.CostResponse)
	repo.DataOngkir["testorigincity"] = hasilongkir

	repo.DataResi = make(map[string]ongkir.CekResiBinderByte)
	repo.DataResi[resi] = hasilresi
	return repo
}

func (repo *inMemoryRepository) Cost(Dataongkir *ongkir.Ongkir) (*ongkir.CostResponse, error) {
	hasil := repo.DataOngkir[Dataongkir.Asal]
	return &hasil, nil
}

func (repo *inMemoryRepository) Resi(Dataresi *ongkir.Resi) (*ongkir.CekResiBinderByte, error) {
	hasil := repo.DataResi[dataresi.Resi]
	return &hasil, nil
}
