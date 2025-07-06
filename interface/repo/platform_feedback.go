package repo

import (
	"context"
	"tourmate/feedback-service/model/dto/request"
	"tourmate/feedback-service/model/entity"
)

type IPlatformFeedbackRepo interface {
	GetPlatformFeedbacks(req request.GetPlatformFeedbacksRequest, ctx context.Context) (*[]entity.PlatformFeedback, int, error)
	GetPlatformFeedbackById(id string, ctx context.Context) (*entity.PlatformFeedback, error)
	CreatePlatformFeedback(PlatformFeedback entity.PlatformFeedback, ctx context.Context) error
	UpdatePlatformFeedback(PlatformFeedback entity.PlatformFeedback, ctx context.Context) error
}
