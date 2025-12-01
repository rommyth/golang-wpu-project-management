package controllers

import (
	"math"
	"project-management/models"
	"project-management/services"
	"project-management/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type BoardController struct {
	service services.BoardService
}

func NewBoardController(s services.BoardService) *BoardController {
	return &BoardController{service: s}
}

func (b *BoardController) CreateBoard(ctx *fiber.Ctx) error {
	var userID uuid.UUID
	var err error
	var board models.Board

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID, err = uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return utils.Unauthorize(ctx, "Unauthorized", err.Error())
	}

	if err := ctx.BodyParser(&board); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	board.OwnerPublicID = userID

	if err := b.service.Create(&board); err != nil {
		return utils.InternalServerError(ctx, "Gagal membuat board", err.Error())
	}

	return utils.Success(ctx, "Berhasil membuat board", &board)
}

func (b *BoardController) UpdateBoard(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")

	var board models.Board

	if err := ctx.BodyParser(&board); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "Invalid ID", err.Error())
	}

	existBoard, err := b.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "Board not found", err.Error())
	}

	board.InternalID = existBoard.InternalID
	board.PublicID = existBoard.PublicID
	board.OwnerPublicID = existBoard.OwnerPublicID
	board.OwnerID = existBoard.OwnerID
	board.CreatedAt = existBoard.CreatedAt

	if err := b.service.Update(&board); err != nil {
		return utils.BadRequest(ctx, "Gagal memperbarui board", err.Error())
	}

	return utils.Success(ctx, "Berhasil memperbarui board", &board)
}

func (b *BoardController) AddBoardMembers(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")

	var userIDs []string
	if err := ctx.BodyParser(&userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if err := b.service.AddMembers(publicID, userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal menambahkan anggota", err.Error())
	}

	return utils.Success(ctx, "Berhasil menambahkan anggota", nil)
}

func (b *BoardController) RemoveBoardMembers(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")

	var userIDs []string
	if err := ctx.BodyParser(&userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing data", err.Error())
	}

	if err := b.service.RemoveMembers(publicID, userIDs); err != nil {
		return utils.BadRequest(ctx, "Gagal menghapus anggota", err.Error())
	}

	return utils.Success(ctx, "Berhasil menghapus anggota", nil)
}

func (b *BoardController) GetMyBoardPaginate(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID, err := uuid.Parse(claims["pub_id"].(string))
	if err != nil {
		return utils.Unauthorize(ctx, "Unauthorized", err.Error())
	}

	page := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)
	filter := ctx.Query("filter", "")
	sort := ctx.Query("sort", "")
	offset := (page - 1) * limit

	board, total, err := b.service.GetAllByUserPaginate(userID.String(), filter, sort, limit, offset)
	if err != nil {
		return utils.InternalServerError(ctx, "Gagal mengambil data", err.Error())
	}

	meta := utils.PaginationMeta{
		Page:       page,
		Limit:      limit,
		Total:      int(total),
		TotalPages: int(math.Ceil(float64(total) / float64(limit))),
		Filter:     filter,
		Sort:       sort,
	}

	return utils.SuccessPagination(ctx, "Berhasil mengambil data", board, meta)

}
