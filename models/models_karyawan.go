package models

type Karyawan struct {
	ID        int    `json:"id" gorm:"primary_key"`
	KdNama    string `json:"kd_nama"`
	Nama      string `json:"nama"`
	AwalMasuk string `json:"awal_masuk"`
}
