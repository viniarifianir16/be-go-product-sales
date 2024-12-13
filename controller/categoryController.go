package controller

import (
	"be-go-product-sales/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryInput struct {
	JenisBarang string `json:"jenis_barang" binding:"required"`
}

// GetAllCategory godoc
// @Summary Get All Category.
// @Description Get a list of Category.
// @Tags Category
// @Produce json
// @Success 200 {object} []models.Category
// @Router /category [get]
func GetAllCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category []models.Category
	if err := db.Find(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// AddCategory godoc
// @Summary Add New Category.
// @Description Add a new Category.
// @Tags Category
// @Param Body body categoryInput true "The body to create a new Category"
// @Produce json
// @Success 200 {object} models.Category
// @Router /category [post]
func AddCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// var category models.Category

	// validasi input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create category
	category := models.Category{
		JenisBarang: input.JenisBarang,
	}

	if err := db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
// @Summary Update Category.
// @Description Update category by id.
// @Tags Category
// @Param id path string true "Category ID"
// @Param Body body categoryInput true "The body to update an Category"
// @Produce json
// @Success 200 {object} models.Category
// @Router /category/{id} [patch]
func UpdateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var category models.Category
	var input categoryInput

	// validasi input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cari id category dari input
	if err := db.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// update category
	category.JenisBarang = input.JenisBarang

	if err := db.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory godoc
// @Summary Delete one category.
// @Description Delete a category by id.
// @Tags Category
// @Param id path string true "Category ID"
// @Produce json
// @Success 200 {object} map[string]boolean
// @Router /category/{id} [delete]
func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	if err := db.Delete(&models.Category{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted succesfully"})
}
