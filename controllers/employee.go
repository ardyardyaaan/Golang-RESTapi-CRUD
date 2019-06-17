package controllers

import (
	"net/http"

	"../structs"
	"github.com/gin-gonic/gin"
)

//menampilkan semua data Employee(GET)
func (idb *InDB) GetEmployees(c *gin.Context) {
	var (
		employees []structs.Employee
		result    gin.H
	)

	idb.DB.Find(&employees)
	if len(employees) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": employees,
			"count":  len(employees),
		}
	}
	c.JSON(http.StatusOK, result)
}

//menampilkan satu data Employee berdasarkan ID(GET)
func (idb *InDB) GetEmployee(c *gin.Context) {
	var (
		employee structs.Employee
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&employee).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": employee,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

//memasukan satu data Employee ke database(POST)
func (idb *InDB) CreateEmployee(c *gin.Context) {
	var (
		employee structs.Employee
		result   gin.H
	)
	nomor_karyawan := c.PostForm("nomor_karyawan")
	nama_karyawan := c.PostForm("nama_karyawan")
	departemen := c.PostForm("departemen")
	jabatan := c.PostForm("jabatan")
	alamat := c.PostForm("alamat")
	employee.Nomor_Karyawan = nomor_karyawan
	employee.Nama_Karyawan = nama_karyawan
	employee.Departemen = departemen
	employee.Jabatan = jabatan
	employee.Alamat = alamat
	idb.DB.Create(&employee)
	result = gin.H{
		"result": employee,
	}
	c.JSON(http.StatusOK, result)
}

//edit satu data Employee menggunakan ID(PUT)
func (idb *InDB) UpdateEmployee(c *gin.Context) {
	// id := c.Param("id")
	id := c.Query("id")
	nomor_karyawan := c.PostForm("nomor_karyawan")
	nama_karyawan := c.PostForm("nama_karyawan")
	departemen := c.PostForm("departemen")
	jabatan := c.PostForm("jabatan")
	alamat := c.PostForm("alamat")
	var (
		employee    structs.Employee
		newEmployee structs.Employee
		result      gin.H
	)

	err := idb.DB.First(&employee, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	newEmployee.Nomor_Karyawan = nomor_karyawan
	newEmployee.Nama_Karyawan = nama_karyawan
	newEmployee.Departemen = departemen
	newEmployee.Jabatan = jabatan
	newEmployee.Alamat = alamat
	err = idb.DB.Model(&employee).Updates(newEmployee).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

//menghapus satu data Employee dari database(DELETE)
func (idb *InDB) DeleteEmployee(c *gin.Context) {
	var (
		employee structs.Employee
		result   gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&employee, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&employee).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
