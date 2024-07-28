package repository

import (
	"benneighbour.com/services/b_service/entity"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GoodbyeRepository defines the interface for the Goodbye repository
type GoodbyeRepository interface {
	// Create adds a new Goodbye entity to the database
	Create(ctx context.Context, goodbye *entity.Goodbye) (*entity.Goodbye, error)

	// Update modifies an existing Goodbye entity in the database
	Update(ctx context.Context, goodbye *entity.Goodbye) (*entity.Goodbye, error)

	// FindByID retrieves a Goodbye entity by its ID from the database
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Goodbye, error)
}

// GoodbyeRepositoryImpl implements GoodbyeRepository
type GoodbyeRepositoryImpl struct {
	DB *gorm.DB
}

// NewGoodbyeRepository creates a new GoodbyeRepositoryImpl
func NewGoodbyeRepository(db *gorm.DB) GoodbyeRepository {
	return &GoodbyeRepositoryImpl{DB: db}
}

// Create adds a new Goodbye entity to the database
func (r *GoodbyeRepositoryImpl) Create(ctx context.Context, goodbye *entity.Goodbye) (*entity.Goodbye, error) {
	return goodbye, r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Create(goodbye).Error
	})
}

// Update modifies an existing Goodbye entity in the database
func (r *GoodbyeRepositoryImpl) Update(ctx context.Context, goodbye *entity.Goodbye) (*entity.Goodbye, error) {
	return goodbye, r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Save(goodbye).Error
	})
}

// FindByID retrieves a Goodbye entity by its ID from the database
func (r *GoodbyeRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entity.Goodbye, error) {
	goodbye := new(entity.Goodbye)
	err := r.DB.WithContext(ctx).Where("id = ?", id).Take(goodbye).Error
	return goodbye, err
}
