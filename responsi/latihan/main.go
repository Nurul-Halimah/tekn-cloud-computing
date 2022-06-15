package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err =
		gorm.Open(sqlite.Open("./Responsi.db"), &gorm.Config{})
	if err != nil {
		panic("Gagal Conect Ke Database")
	}
}

type (
	mahasiswa struct {
		ID    int     `json:"ID"`
		NAMA  string  `json:"NAMA"`
		NIM int  `json:"NIM"`
		ALAMAT  string `json:"ALAMAT"`
		JK string `json:"JK"`
	}
)

func fetchAllMahasiswa(c *gin.Context) {
	var model []mahasiswa

	db.Find(&model)

	if len(model) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
	}

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": model})
}

func main() {

	router := gin.Default()
	v1 := router.Group("/api/mahasiswa")
	{
		v1.GET("", fetchAllMahasiswa)
	}
	router.Run(":20001")