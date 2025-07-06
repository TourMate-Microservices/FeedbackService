package request

type GetPlatformFeedbacksRequest struct {
	Request   SearchPaginationRequest `json:"request"`
	AccountId int                     `json:"account_id" form:"account_id" validate:"min=1"`
	PaymnetId int                     `json:"paymnet_id" form:"paymnet_id" validate:"min=1"`
	Rating    int                     `json:"rating" form:"rating" validate:"min=1"`
}
