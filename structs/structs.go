package structs

import "github.com/jinzhu/gorm"

type Employee struct {
	gorm.Model
	Nomor_Karyawan string
	Nama_Karyawan  string
	Departemen     string
	Jabatan        string
	Alamat         string
}
