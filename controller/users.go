package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"minitask2/models"
	"minitask2/utils"
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
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Router /users [get]
func GetAlluser(ctx *gin.Context) {
	result := utils.RedisClient.Exists(context.Background(), ctx.Request.RequestURI)
	if result.Val() == 0 {
		search := ctx.DefaultQuery("search", "")
		users := models.FindAllUser(search)
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "List User",
			Result:  users,
		})
		encoded, err := json.Marshal(users)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "Gagal Get user",
				Errors:  err.Error(),
			})
			return
		}
		utils.RedisClient.Set(context.Background(), ctx.Request.RequestURI, string(encoded), 0)

	} else {
		data := utils.RedisClient.Get(context.Background(), ctx.Request.RequestURI)
		str := data.Val()
		users := []models.Profile{}
		json.Unmarshal([]byte(str), &users)
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "List User",
			Result:  users,
		})
	}

}

// @Description list One users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} string "string"
// @Param id path int true "User ID"
// @Router /users/{id} [get]
func GetOne(ctx *gin.Context) {
	result := utils.RedisClient.Exists(context.Background(), ctx.Request.RequestURI)
	if result.Val() == 0 {
		id := ctx.Param("id")
		users := models.FindUserById(id)
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "List User",
			Result:  users,
		})
		encoded, err := json.Marshal(users)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, models.Response{
				Success: false,
				Message: "Gagal Get user",
				Errors:  err.Error(),
			})
			return
		}
		utils.RedisClient.Set(context.Background(), ctx.Request.RequestURI, string(encoded), 0)
	} else {
		data := utils.RedisClient.Get(context.Background(), ctx.Request.RequestURI)
		str := data.Val()
		users := []models.Profile{}
		json.Unmarshal([]byte(str), &users)
		ctx.JSON(http.StatusOK, models.Response{
			Success: true,
			Message: "List User",
			Result:  users,
		})
	}
}

// @Description Create a user from JSON input
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.Profile true "User Data"
// @Success 201 {object} models.Response
// @Router /users [post]
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
	_, err = models.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}
	utils.RedisClient.Del(context.Background(), "/users")

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "User created",
	})

}

// @Description update  user from JSON input
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 201 {object} models.Response
// @Router /users/{id} [delete]
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
	utils.RedisClient.Del(context.Background(), "/users")

	ctx.JSON(200, models.Response{
		Success: true,
		Message: "Berhasil update user",
	})
}

// @Description Delete  user
// @Tags Users
// @Accept json
// @Produce json
// @Success 201 {object} models.Response
// @Router /users/{id} [patch]
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
	utils.RedisClient.Del(context.Background(), "/users")

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Berhasil hapus user",
	})

}
