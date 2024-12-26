package handlers

import (
	"net/http"

	"github.com/Gierdiaz/Echo-golang/database"
	"github.com/Gierdiaz/Echo-golang/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	var users []models.User
	err := database.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao buscar usuários"})
	}
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Dados inválidos"})
	}

	query := `INSERT INTO users (name, email, password) VALUES (:name, :email, :password) RETURNING id`
	err := database.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Erro ao criar usuário"})
	}

	return c.JSON(http.StatusCreated, user)
}
