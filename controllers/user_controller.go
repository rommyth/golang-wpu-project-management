package controllers

import (
	"math"
	"project-management/models"
	"project-management/services"
	"project-management/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserServices
}

func NewUserController(service services.UserServices) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Register(ctx *fiber.Ctx) error {
	// user := new(models.User)
	var user models.User

	if err := ctx.BodyParser(&user); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if err := uc.service.Register(&user); err != nil {
		return utils.BadRequest(ctx, "Gagal Mendaftar", err.Error())
	}

	var userResponse models.UserResponse
	userResponse = models.UserResponse{
		PublicID:  user.PublicID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return utils.Success(ctx, "Berhasil Mendaftar", &userResponse)
}

func (uc *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Invalid Request", err.Error())
	}

	user, err := uc.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.Unauthorize(ctx, "Login gagal", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResponse models.UserResponse
	userResponse = models.UserResponse{
		PublicID:  user.PublicID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return utils.Success(ctx, "Login Berhasil", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          &userResponse,
	})
}

func (uc *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	user, err := uc.service.GetByPublicID(id)

	if err != nil {
		return utils.NotFound(ctx, "User Tidak Ditemukan", err.Error())
	}

	var userResponse models.UserResponse
	userResponse = models.UserResponse{
		PublicID:  user.PublicID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return utils.Success(ctx, "User Ditemukan", &userResponse)
}

func (uc *UserController) GetUserPagination(ctx *fiber.Ctx) error {
	// users/page?page=1&limit=10&filter=jhon&sort=-id

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")
	offset := (page - 1) * limit

	users, total, err := uc.service.GetAllPagination(filter, sort, limit, offset)
	if err != nil {
		return utils.BadRequest(ctx, "Gagal mengambil data", err.Error())
	}

	var userResponse []models.UserResponse
	_ = copier.Copy(&userResponse, &users)

	meta := utils.PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      int(total),
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		Filter:     filter,
		Sort:       sort,
	}

	if total == 0 {
		return utils.NotFoundPagination(ctx, "User Tidak Ditemukan", &userResponse, meta)
	}

	return utils.SuccessPagination(ctx, "User Ditemukan", &userResponse, meta)
}

func (uc *UserController) UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	publicID, err := uuid.Parse(id)
	if err != nil {
		return utils.BadRequest(ctx, "Invalid ID Format", err.Error())
	}

	var user models.User
	if err := ctx.BodyParser(&user); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	user.PublicID = publicID

	if err := uc.service.Update(&user); err != nil {
		return utils.BadRequest(ctx, "Gagal mengupdate data", err.Error())
	}

	userUpdated, err := uc.service.GetByPublicID(id)
	if err != nil {
		return utils.InternalServerError(ctx, "Gagal mengambil data", err.Error())
	}

	var userResponse models.UserResponse
	err = copier.Copy(&userResponse, &userUpdated)
	if err != nil {
		return utils.InternalServerError(ctx, "Error parsing data", err.Error())
	}

	return utils.Success(ctx, "User berhasil diperbarui", &userResponse)
}

func (uc *UserController) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return utils.BadRequest(ctx, "Invalid ID Format", err.Error())
	}

	if err := uc.service.Delete(uint(id)); err != nil {
		return utils.InternalServerError(ctx, "Gagal menghapus data", err.Error())
	}

	return utils.Success(ctx, "User berhasil dihapus", nil)
}
