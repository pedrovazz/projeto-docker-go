package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Banda representa a estrutura da tabela no banco de dados
type Banda struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Nome       string `json:"nome"`
	Musicos    string `json:"musicos"`
	Generos    string `json:"generos"`
	Status     bool   `json:"status"`
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

var DB *gorm.DB
var err error

func main() {
	dsn := "root:root@tcp(172.17.0.2:3306)/projeto?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados!")
	}

	DB.AutoMigrate(&Banda{})

	r := gin.Default()

	r.POST("/bandas", createBanda)
	r.GET("/bandas", getBandas)
	r.GET("/bandas/:id", getBanda)
	r.PUT("/bandas/:id", updateBanda)
	r.DELETE("/bandas/:id", deleteBanda)

	r.Run(":8080")
}

func createBanda(c *gin.Context) {
	var banda Banda
	if err := c.ShouldBindJSON(&banda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&banda)
	c.JSON(http.StatusOK, banda)
}

func getBandas(c *gin.Context) {
	var bandas []Banda
	DB.Find(&bandas)
	c.JSON(http.StatusOK, bandas)
}

func getBanda(c *gin.Context) {
	id := c.Param("id")
	var banda Banda
	if result := DB.First(&banda, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banda não encontrada"})
		return
	}
	c.JSON(http.StatusOK, banda)
}

func updateBanda(c *gin.Context) {
	id := c.Param("id")
	var banda Banda
	if result := DB.First(&banda, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banda não encontrada"})
		return
	}
	if err := c.ShouldBindJSON(&banda); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Save(&banda)
	c.JSON(http.StatusOK, banda)
}

func deleteBanda(c *gin.Context) {
	id := c.Param("id")
	if result := DB.Delete(&Banda{}, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Banda não encontrada"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Banda deletada com sucesso"})
}
