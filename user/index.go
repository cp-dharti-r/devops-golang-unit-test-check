package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Db *sqlx.DB
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func New(db *sqlx.DB) *Repository {
	return &Repository{Db: db}
}

// create user
func (repository *Repository) Create(c *gin.Context) {
	input := User{}
	err := c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = repository.Db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", input.Name, input.Email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// get user
func (repository *Repository) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user User
	row := repository.Db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	err = row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)
}

// update user
func (repository *Repository) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	input := User{}
	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var user User
	row := repository.Db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err = row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	result, err := repository.Db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", input.Name, input.Email, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}

// delete user
func (repository *Repository) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var user User
	row := repository.Db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)
	err = row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	repository.Db.Exec("DELETE FROM users WHERE id = ?", id)

	c.JSON(http.StatusOK, gin.H{})
}
