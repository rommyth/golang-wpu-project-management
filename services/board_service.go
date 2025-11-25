package services

import (
	"project-management/models"
	"project-management/repositories"
)

type BoardService interface {
}

type boardService struct {
	boardRepo repositories.BoardRepository
	userRepo  repositories.UserRepository
}

// Constructor
func NewBoardService(
	boardRepo repositories.BoardRepository,
	userRepo repositories.UserRepository,
) BoardService {
	return &boardService{
		boardRepo: boardRepo,
		userRepo:  userRepo,
	}
}

func (s *boardService) Create(board *models.Board) error {
	return s.boardRepo.Create(board)
}
