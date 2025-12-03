package utils

import (
	"project-management/models"

	"github.com/google/uuid"
)

func SortListPosition(list []models.List, order []uuid.UUID) []models.List {
	ordered := make([]models.List, 0, len(order))

	listMap := make(map[uuid.UUID]models.List)
	for _, l := range list {
		listMap[l.PublicID] = l
	}

	// urutkan sesuai order
	for _, id := range order {
		if list, ok := listMap[id]; ok {
			ordered = append(ordered, list)
		}
	}

	return ordered
}
