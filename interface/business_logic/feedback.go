package businesslogic

import (
	"context"
	"tourmate/feedback-service/model/dto/request"
	"tourmate/feedback-service/model/dto/response"
	"tourmate/feedback-service/model/entity"
)

type IFeedbackService interface {
	GetFeedbacks(req request.GetFeedbacksRequest, ctx context.Context) (response.PaginationDataResponse, error)
	GetFeedbackById(id int, ctx context.Context) (*entity.Feedback, error)
	CreateFeedback(req request.CreateFeedbackRequest, ctx context.Context) error
	UpdateFeedback(req request.UpdateFeedbackRequest, ctx context.Context) error
	RemoveFeedback(req request.RemoveFeedbackRequest, ctx context.Context) error
}
