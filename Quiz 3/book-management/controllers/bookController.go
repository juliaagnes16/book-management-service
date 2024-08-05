package controllers

import (
	"database/sql"
	"net/http"
	"time"

	"main.go/database"
	"main.go/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	rows, err := database.DB.Query("SELECT * FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	if book.TotalPage > 100 {
		book.Thickness = "tebal"
	} else {
		book.Thickness = "tipis"
	}

	book.CreatedAt = time.Now()
	book.ModifiedAt = time.Now()

	_, err := database.DB.Exec("INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)", book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedAt, book.CreatedBy, book.ModifiedAt, book.ModifiedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	query := `SELECT title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id = $1`

	var book models.Book
	row := database.DB.QueryRow(query, id)

	if err := row.Scan(
		&book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price,
		&book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt,
		&book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy,
	); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve book"})
		}
		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingBook models.Book
	err := database.DB.QueryRow("SELECT id FROM books WHERE id = $1", id).Scan(&existingBook.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	book.ModifiedAt = time.Now()
	book.ModifiedBy = "postgre"

	_, err = database.DB.Exec(`
        UPDATE books SET
            title = $1,
            description = $2,
            image_url = $3,
            release_year = $4,
            price = $5,
            total_page = $6,
            thickness = $7,
            category_id = $8,
            modified_at = $9,
            modified_by = $10
        WHERE id = $11
    `, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.ModifiedAt, book.ModifiedBy, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var existingBook models.Book
	err := database.DB.QueryRow("SELECT id FROM books WHERE id = $1", id).Scan(&existingBook.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = database.DB.Exec("DELETE FROM books WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
