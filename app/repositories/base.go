package repositories

import (
	"gorm.io/gorm"
	"main.go/app/models"
)

type IKBaseRepository struct {
	db *gorm.DB
}

func NewIKBaseRepository(db *gorm.DB) *IKBaseRepository {
	return &IKBaseRepository{db: db}
}

// Add implements ModelIKLocationURLRepository.
func (i *IKBaseRepository) Create(input *models.Base) error {
	err := i.db.Model(&models.Base{}).Create(input).Error
	if err != nil {
		return err
	}
	return nil
}

// List implements ModelIKLocationURLRepository.
func (i *IKBaseRepository) List(limit, skip int) ([]*models.Base, error) {
	var base []*models.Base
	err := i.db.Model(&models.Base{}).Where("deleted_at IS NULL").Offset(skip).Limit(limit).Find(&base).Error
	if err != nil {
		return nil, err
	}
	return base, nil
}

func (i *IKBaseRepository) Delete(id uint) error {
	result, err := i.GetById(id)
	if err != nil {
		return err
	}
	i.db.Model(&models.Base{}).Delete(result)
	return nil
}

func (i *IKBaseRepository) GetById(id uint) (*models.Base, error) {
	var base models.Base
	err := i.db.Model(&models.Base{}).Where("deleted_at IS NULL").First(&base, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return &base, nil
}

func (i *IKBaseRepository) Update(id uint, input *models.Base) error {
	result, err := i.GetById(id)
	if err != nil {
		return err
	}
	result.Name = input.Name
	result.Age = input.Age
	err = i.db.Save(&result).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *IKBaseRepository) Count() (int64, error) {
	var count int64
	err := i.db.Model(&models.Base{}).Count(&count).Error
	if err != nil {
		return 0, nil
	}
	return count, nil
}
