package repositories

import (
	"project-management/config"
	"project-management/models"

	"github.com/google/uuid"
)

type ListRepository interface {
	Create(list *models.List) error
	Update(list *models.List) error
	Delete(publicID string) error
	UpdatePosition(boardPublidID string, position []string) error
	GetPosition(listPublicID string) ([]uuid.UUID, error)
	FindByBoardID(boardID string) ([]models.List, error)
	FindByPublicID(publicID string) (*models.List, error)
	FindByID(id uint) (*models.List, error)
}

type listRepository struct{}

func NewListRepository() ListRepository {
	return &listRepository{}
}

func (r *listRepository) Create(list *models.List) error {
	return config.DB.Create(list).Error
}

func (r *listRepository) Update(list *models.List) error {
	return config.DB.Model(&models.List{}).
		Where("public_id = ?", list.PublicID).
		Updates(map[string]interface{}{
			"title": list.Title,
		}).Error
}

func (r *listRepository) Delete(publicID string) error {
	return config.DB.Where("public_id = ?", publicID).Delete(&models.List{}).Error
}

func (r *listRepository) UpdatePosition(boardPublidID string, position []string) error {
	return config.DB.Model(&models.ListPosition{}).
		Where("board_internal_id = (SELECT internal_id FROM boards WHERE public_id = ?)", boardPublidID).
		Updates(map[string]interface{}{
			"list_order": position,
		}).Error
}

func (r *listRepository) GetPosition(listPublicID string) ([]uuid.UUID, error) {
	var position models.CardPosition
	err := config.DB.Joins("JOIN lists ON list.internal_id = card.positions.list_internal_id").
		Where("list.public_id = ?", listPublicID).Scan(&position).Error
	return position.CardOrder, err
}

func (r *listRepository) FindByBoardID(boardID string) ([]models.List, error) {
	var list []models.List
	err := config.DB.Where("board_public_id = ?", boardID).Order("internal_id ASC").Find(&list).Error
	return list, err
}

func (r *listRepository) FindByPublicID(publicID string) (*models.List, error) {
	var list models.List
	err := config.DB.Where("public_id = ?", publicID).First(&list).Error
	return &list, err
}

func (r *listRepository) FindByID(id uint) (*models.List, error) {
	var list models.List
	err := config.DB.First(&list, id).Error
	return &list, err
}
