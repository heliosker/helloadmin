package login_record

import (
	"context"
	"helloadmin/api"
	"helloadmin/internal/service"
	"time"
)

type LoginRecordService interface {
	Create(ctx context.Context, record *LoginRecordRequest) error
	Search(ctx context.Context, request *LoginRecordFindRequest) (*LoginRecordResponse, error)
}

func NewService(service *service.Service, loginRecordRepository LoginRecordRepository) LoginRecordService {
	return &loginRecordService{
		Service:               service,
		loginRecordRepository: loginRecordRepository,
	}
}

type loginRecordService struct {
	*service.Service
	loginRecordRepository LoginRecordRepository
}

func (lrs *loginRecordService) Create(ctx context.Context, req *LoginRecordRequest) error {
	model := Model{
		Ip:           req.Ip,
		Os:           req.Os,
		Email:        req.Email,
		Browser:      req.Browser,
		Platform:     req.Platform,
		ErrorMessage: req.ErrorMessage,
		UpdatedAt:    time.Now(),
		CreatedAt:    time.Now(),
	}
	if err := lrs.loginRecordRepository.Create(ctx, &model); err != nil {
		return err
	}
	return nil
}

func (lrs *loginRecordService) Search(ctx context.Context, request *LoginRecordFindRequest) (*LoginRecordResponse, error) {
	var result LoginRecordResponse
	count, records, err := lrs.loginRecordRepository.Search(ctx, request)
	if err != nil {
		return nil, err
	}
	result.Items = make([]LoginRecordItem, 0)
	if count > 0 {
		for _, record := range *records {
			result.Items = append(result.Items, LoginRecordItem{
				Ip:           record.Ip,
				Os:           record.Os,
				Email:        record.Email,
				Browser:      record.Browser,
				Platform:     record.Platform,
				ErrorMessage: record.ErrorMessage,
				UpdatedAt:    record.UpdatedAt.Format(time.DateTime),
				CreatedAt:    record.CreatedAt.Format(time.DateTime),
			})
		}
	}
	result.Pagination = api.Pagination{
		Page:  request.Page,
		Size:  request.Size,
		Count: int(count),
	}
	return &result, nil
}
