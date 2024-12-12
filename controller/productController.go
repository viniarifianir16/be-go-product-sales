package controller

import (
	"be-go-product-sales/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type productInput struct {
	NamaBarang       string    `json:"nama_barang" binding:"required"`
	Stok             int       `json:"stok" binding:"required"`
	JumlahTerjual    int       `json:"jumlah_terjual" binding:"required"`
	TanggalTransaksi time.Time `json:"tanggal_transaksi" binding:"required"`
	JenisBarang      string    `json:"jenis_barang" binding:"required"`
}

// GetAllProduct godoc
// @Summary Get All Product.
// @Description Get a list of Product.
// @Tags Product
// @Produce json
// @Success 200 {object} []models.Product
// @Router /product [get]
func GetAllProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product []models.Product
	if err := db.Find(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// AddProduct godoc
// @Summary Add New Product.
// @Description Add a new Product.
// @Tags Product
// @Param Body body productInput true "The body to create a new Product"
// @Produce json
// @Success 200 {object} models.Product
// @Router /product [post]
func AddProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// var product models.Product

	// validasi input
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create product
	product := models.Product{
		NamaBarang:       input.NamaBarang,
		Stok:             input.Stok,
		JumlahTerjual:    input.JumlahTerjual,
		TanggalTransaksi: input.TanggalTransaksi,
		JenisBarang:      input.JenisBarang,
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// UpdateProduct godoc
// @Summary Update Product.
// @Description Update product by id.
// @Tags Product
// @Param id path string true "Product ID"
// @Param Body body productInput true "The body to update an Product"
// @Produce json
// @Success 200 {object} models.Product
// @Router /product/{id} [patch]
func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var product models.Product
	var input productInput

	// validasi input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cari id product dari input
	if err := db.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// update product
	product.NamaBarang = input.NamaBarang
	product.Stok = input.Stok
	product.JumlahTerjual = input.JumlahTerjual
	product.TanggalTransaksi = input.TanggalTransaksi
	product.JenisBarang = input.JenisBarang

	if err := db.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete one product.
// @Description Delete a product by id.
// @Tags Product
// @Param id path string true "Product ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted succesfully"})
}
