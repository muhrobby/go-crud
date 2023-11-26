package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-crud/database"
	"github.com/muhrobby/go-crud/model"
	"github.com/muhrobby/go-crud/utils"
)

func ShowUser(c *fiber.Ctx) error {

	var user []model.User

	db := database.ConnectDB()

	err := db.Find(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal menampilkan semua data",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "berhasil ambil data",
		"data":    user,
	})
}

func ShowByIdUser(c *fiber.Ctx) error {
	var user []model.User

	uuid := c.Params("uuid")

	db := database.ConnectDB()

	db.Find(&user, "uuid= ?", uuid)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})

}

func CreateUser(c *fiber.Ctx) error {
	type UserCreate struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	var CreateUser UserCreate

	user := new(model.User)

	err := c.BodyParser(&CreateUser)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error mengamnil data",
			"status":  "error",
			"data":    nil,
		})
	}

	errV := validator.New().Struct(CreateUser)

	if errV != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": errV.Error(),
		})
	}

	hash, errHash := model.HashPassword(CreateUser.Password)
	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  errHash,
			"message": "failed to hash",
		})
	}

	user.Username = CreateUser.Username
	user.Email = CreateUser.Email
	user.Password = hash
	user.Uuid = utils.RandomString(25) + utils.Init()

	errCreated := database.ConnectDB().Create(&user).Error

	if errCreated != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Gagal membuat data",
			"error":   errCreated.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"data":    CreateUser,
		"message": "Berhasil, data tersimpan",
	})

}

func UpdateUser(c *fiber.Ctx) error {
	var UpdUsr model.UpdateUser

	err := c.BodyParser(&UpdUsr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal ambil data",
			"error":   err.Error(),
		})

	}

	errV := validator.New().Struct(UpdUsr)
	if errV != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": errV.Error(),
		})
	}

	var user model.User

	uuid := c.Params("uuid")

	db := database.ConnectDB()

	errDb := db.Find(&user, "uuid = ?", uuid).Error

	if errDb != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": errDb.Error(),
		})
	}

	user.Names = UpdUsr.Names

	db.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "berhasil update data",
		"data":    UpdUsr,
	})

}

func DeleteUser(c *fiber.Ctx) error {

	var DltUsr model.DeleteUser

	err := c.BodyParser(&DltUsr)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "gagal ambil data",
			"error":   err.Error(),
		})
	}

	errV := validator.New().Struct(DltUsr)
	if errV != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": errV.Error(),
		})
	}

	var user model.User

	uuid := c.Params("uuid")

	db := database.ConnectDB()

	db.Find(&user, "uuid = ?", uuid)

	isValid := model.ComparePassword(user.Password, DltUsr.Password)

	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid password",
		})
	}

	db.Delete(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "ok",
		"message": "success delete user",
		"data":    user,
	})
}
