package controller

import (
	"fmt"
	"minitask2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("endpointhit")
		ctx.Next()

	}
}

// @Description list all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Router /users [get]
func GetAlluser(ctx *gin.Context) {
	search := ctx.DefaultQuery("search", "")
	users := models.FindAllUser(search)
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List User",
		Result:  users,
	})

}
func GetOne(ctx *gin.Context) {
	id := ctx.Param("id")
	users := models.FindUserById(id)
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List User",
		Result:  users,
	})
}
func CreateUser(ctx *gin.Context) {
	var user models.Profile

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	newUser, err := models.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User created",
		Result:  newUser,
	})
}
func UpdateUserCTRL(ctx *gin.Context) {
	id := ctx.Param("id")
	var input models.Profile

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
			Errors:  err.Error(),
		})
		return
	}

	err := models.UpdateUser(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Gagal update user",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Berhasil update user",
	})
}
func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Gagal update user",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Berhasil hapus user",
	})

}
