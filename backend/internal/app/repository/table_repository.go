package repositories

import (
	"gorm.io/gorm"
	"github.com/BookIT/backend/internal/app/models"
)

type tableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return &tableRepository{db: db}
}

func (r *tableRepository) CreateTable(table *models.Table) error {
	return r.db.Create(table).Error
}

func (r *tableRepository) DeleteTableByID(id uint) error {
	result := r.db.Delete(&models.Table{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *tableRepository) GetAllTables() ([]models.Table, error) {
	var tables []models.Table
	if err := r.db.Find(&tables).Error; err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *tableRepository) GetTableByID(id uint) (*models.Table, error) {
	var table models.Table
	if err := r.db.First(&table, id).Error; err != nil {
		return nil, err
	}
	return &table, nil
}

func (r *tableRepository) CreateTables(tables []models.Table) ([]uint, error) {
	err := r.db.Create(&tables).Error
	if err != nil {
		return nil, err
	}

	ids := make([]uint, 0, len(tables))
	for _, t := range tables {
		ids = append(ids, t.ID)
	}

	return ids, nil
}