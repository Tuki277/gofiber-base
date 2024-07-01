package services

import (
	"strconv"

	dto "main.go/app/dtos"
	"main.go/app/models"
	"main.go/app/repositories"
)

type IkBaseService struct {
	repo *repositories.IKBaseRepository
}

func NewIkBaseService(repo *repositories.IKBaseRepository) *IkBaseService {
	return &IkBaseService{repo: repo}
}

func (s *IkBaseService) Create(input *dto.BaseRequest) error {
	models := models.Base{
		Name: input.Name,
		Age:  input.Age,
	}
	err := s.repo.Create(&models)
	if err != nil {
		return err
	}
	return nil
}

func (s *IkBaseService) GetList(limit, skip string) ([]*dto.BaseResponse, int64, error) {
	convertLimit, convertSkip := 10, 0
	if limit != "" {
		if parsedLimit, err := strconv.Atoi(limit); err == nil {
			convertLimit = parsedLimit
		} else {
			return nil, 0, err
		}
	}
	if skip != "" {
		if parsedSkip, err := strconv.Atoi(skip); err == nil {
			convertSkip = parsedSkip
		} else {
			return nil, 0, err
		}
	}

	// Fetch list from the repository
	result, err := s.repo.List(convertLimit, convertSkip)
	if err != nil {
		return nil, 0, err
	}

	// Fetch count from the repository
	count, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}

	// Convert and return results
	return dto.ToBases(result), count, nil
}

func (s *IkBaseService) GetById(id string) (*dto.BaseResponse, error) {
	convertId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	result, err := s.repo.GetById(uint(convertId))
	if err != nil {
		return nil, err
	}
	return dto.ToBase(result), nil
}

func (s *IkBaseService) Delete(id string) error {
	convertId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	err = s.repo.Delete(uint(convertId))
	if err != nil {
		return err
	}
	return nil
}

func (s *IkBaseService) Update(id string, input *dto.BaseRequest) error {
	convertId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	models := models.Base{
		Name: input.Name,
		Age:  input.Age,
	}
	err = s.repo.Update(uint(convertId), &models)
	if err != nil {
		return err
	}
	return nil
}
