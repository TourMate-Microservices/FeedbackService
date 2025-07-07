package businesslogic

import (
	"context"
	"tourmate/feedback-service/model/dto/request"
	"tourmate/feedback-service/model/dto/response"
	"tourmate/feedback-service/model/entity"
)

type IPlatformFeedbackService interface {
	GetPlatformFeedbacks(req request.GetPlatformFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error)
	GetPlatformFeedbackById(id int, ctx context.Context) (*entity.PlatformFeedback, error)
	CreatePlatformFeedback(req request.CreatePlatformFeedbackRequest, ctx context.Context) error
	UpdatePlatformFeedback(req request.UpdatePlatformFeedbackRequest, ctx context.Context) error
}
