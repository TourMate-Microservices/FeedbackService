package repo

import (
	"context"
	"tourmate/feedback-service/model/dto/request"
	"tourmate/feedback-service/model/entity"
)

type IFeedbackRepo interface {
	GetFeedbacks(req request.GetFeedbacksRequest, ctx context.Context) (*[]entity.Feedback, int, error)
	GetFeedbackById(id string, ctx context.Context) (*entity.Feedback, error)
	CreateFeedback(Feedback entity.Feedback, ctx context.Context) error
	UpdateFeedback(Feedback entity.Feedback, ctx context.Context) error
}
