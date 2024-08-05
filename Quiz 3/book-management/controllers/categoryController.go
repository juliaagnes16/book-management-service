package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"main.go/database"
	"main.go/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category
	rows, err := database.DB.Query("SELECT * FROM categories")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		categories = append(categories, category)
	}
	c.JSON(http.StatusOK, categories)
}

func AddCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.CreatedAt = time.Now()
	category.ModifiedAt = time.Now()

	_, err := database.DB.Exec("INSERT INTO categories (name, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5)", category.Name, category.CreatedAt, category.CreatedBy, category.ModifiedAt, category.ModifiedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	query := `SELECT id, name, created_at, created_by, modified_at, modified_by FROM categories WHERE id = $1`

	var category models.Category
	row := database.DB.QueryRow(query, id)

	if err := row.Scan(
		&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy,
		&category.ModifiedAt, &category.ModifiedBy,
	); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve category"})
		}
		return
	}

	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCategory models.Category
	err := database.DB.QueryRow("SELECT id FROM categories WHERE id = $1", id).Scan(&existingCategory.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	category.ModifiedAt = time.Now()
	category.ModifiedBy = "postgre"

	_, err = database.DB.Exec(`
        UPDATE categories SET
            name = $1,
            modified_at = $2,
            modified_by = $3
        WHERE id = $4
    `, category.Name, category.ModifiedAt, category.ModifiedBy, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	var existingCategory models.Category
	err := database.DB.QueryRow("SELECT id FROM categories WHERE id = $1", id).Scan(&existingCategory.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = database.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
