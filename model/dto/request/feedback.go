package request

type GetFeedbacksRequest struct {
	Request     SearchPaginationRequest `json:"request"`
	CustomerId  int                     `json:"customer_id" form:"customer_id" validate:"min=1"`
	TourGuideId int                     `json:"tour_guide_id" form:"tour_guide_id" validate:"min=1"`
	InvoiceId   int                     `json:"invoice_id" form:"invoice_id" validate:"min=1"`
	Rating      int                     `json:"rating" form:"rating" validate:"min=1"`
	IsDeleted   *bool                   `json:"is_deleted" form:"is_deleted"`
}
