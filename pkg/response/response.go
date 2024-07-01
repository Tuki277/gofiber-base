package response

import (
	"math"
	"strconv"

	dto "main.go/app/dtos"
)

func ResponseData(data interface{}, metadata *dto.ListResponse, err error, code int) dto.Response {
	if err != nil {
		return dto.Response{
			Code:    code,
			Message: err.Error(),
			Data:    data,
			Error:   err,
			Success: false,
		}
	}
	return dto.Response{
		Code:     code,
		Message:  "successfully",
		Data:     data,
		Metadata: metadata,
		Error:    err,
		Success:  true,
	}
}

func ResponseMetaData(count int, limit, skip string) (dto.ListResponse, error) {
	if limit != "" && skip != "" {
		convertLimit, err := strconv.Atoi(limit)
		if err != nil {
			return dto.ListResponse{}, err
		}
		convertSkip, err := strconv.Atoi(skip)
		if err != nil {
			return dto.ListResponse{}, err
		}
		return dto.ListResponse{
			TotalItems:   count,
			ItemsPerPage: convertLimit,
			TotalPages:   int(math.Ceil(float64(count) / float64(convertLimit))),
			CurrentPage:  int(int(math.Floor(float64(convertSkip)/float64(convertLimit)) + 1)),
		}, nil
	}
	return dto.ListResponse{
		TotalItems:   count,
		ItemsPerPage: 10,
		TotalPages:   int(math.Ceil(float64(count) / float64(10))),
		CurrentPage:  int(int(math.Floor(float64(0)/float64(10)) + 1)),
	}, nil
}
