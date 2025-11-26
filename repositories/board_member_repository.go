package repositories

import (
	"project-management/config"
	"project-management/models"
)

type BoardMemberRepository interface {
}

type boardMemberRepository struct{}

func NewBoardMemberRepository() BoardMemberRepository {
	return &boardMemberRepository{}
}

func (r *boardMemberRepository) GetMembers(boardPublicID string) ([]models.BoardMember, error) {
	var user []models.User
	err := config.DB.Joins()
}
