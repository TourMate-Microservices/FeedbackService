package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"tourmate/feedback-service/constant/noti"
	"tourmate/feedback-service/interface/repo"
	"tourmate/feedback-service/model/dto/request"
	"tourmate/feedback-service/model/entity"
)

type feedbackRepo struct {
	db     *sql.DB
	logger *log.Logger
}

func InitializeFeedbackRepo(db *sql.DB, logger *log.Logger) repo.IFeedbackRepo {
	return &feedbackRepo{
		db:     db,
		logger: logger,
	}
}

// CreateFeedback implements repo.IFeedbackRepo.
func (f *feedbackRepo) CreateFeedback(feedback entity.Feedback, ctx context.Context) error {
	var query string = "INSERT INTO " + feedback.GetFeedbackTable() +
		"(customerId, tourGuideId, createdDate, " +
		"content, rating, isDeleted, " +
		"updatedAt, invoiceId) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8)"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, feedback.GetFeedbackTable()) + "CreateFeedback - "

	if _, err := f.db.Exec(query, feedback.CustomerId, feedback.TourGuideId, feedback.CreatedDate,
		feedback.Content, feedback.Rating, feedback.IsDeleted,
		feedback.UpdatedAt, feedback.InvoiceId); err != nil {

		f.logger.Println(errLogMsg + err.Error())
		return errors.New(noti.INTERNALL_ERR_MSG)
	}

	return nil
}

// GetFeedbackById implements repo.IFeedbackRepo.
func (f *feedbackRepo) GetFeedbackById(id int, ctx context.Context) (*entity.Feedback, error) {
	var res entity.Feedback
	var query string = "SELECT * FROM " + res.GetFeedbackTable() + " WHERE feedbackId = $1"
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, res.GetFeedbackTable()) + "GetFeedbackById - "

	if err := f.db.QueryRow(query, id).Scan(
		&res.FeedbackId, &res.CustomerId, &res.TourGuideId, &res.CreatedDate,
		&res.Content, &res.Rating, &res.IsDeleted, &res.UpdatedAt, &res.InvoiceId); err != nil {

		if err == sql.ErrNoRows {
			return nil, nil
		}

		f.logger.Println(errLogMsg + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return &res, nil
}

// GetFeedbacks implements repo.IFeedbackRepo.
func (f *feedbackRepo) GetFeedbacks(req request.GetFeedbacksRequest, ctx context.Context) (*[]entity.Feedback, int, int, error) {
	var res []entity.Feedback
	var table string = res[0].GetFeedbackTable()
	var limitRecords int = res[0].GetFeedbackLimitRecords()

	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, table) + "GetFeedbacks - "
	var queryCondition string = fmt.Sprintf(
		" WHERE customerId = '%d' AND tourGuideId = '%d' AND invoiceId = '%d' LOWER(content) LIKE LOWER('%%%s%%') AND isDeleted = '%b'",
		req.CustomerId,
		req.TourGuideId,
		req.InvoiceId,
		req.Request.Keyword,
		req.IsDeleted,
	)
	var orderCondition string = generateOrderCondition(req.Request.FilterProp, req.Request.Order)
	var query string = generateRetrieveQuery(table, queryCondition+orderCondition, limitRecords, req.Request.Page, false)

	rows, err := f.db.Query(query)
	if err != nil {
		f.logger.Println(errLogMsg + err.Error())
		return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
	}

	for rows.Next() {
		var x entity.Feedback
		if err := rows.Scan(
			&x.FeedbackId, &x.CustomerId, &x.TourGuideId, &x.CreatedDate,
			&x.Content, &x.Rating, &x.IsDeleted, &x.UpdatedAt, &x.InvoiceId,
		); err != nil {

			f.logger.Println(errLogMsg + err.Error())
			return nil, 0, 0, errors.New(noti.INTERNALL_ERR_MSG)
		}

		res = append(res, x)
	}

	// Track total records in table
	var totalRecords int
	f.db.QueryRow(generateRetrieveQuery(table, queryCondition, limitRecords, req.Request.Page, true)).Scan(&totalRecords)

	return &res, caculateTotalPages(totalRecords, limitRecords), totalRecords, nil
}

// UpdateFeedback implements repo.IFeedbackRepo.
func (f *feedbackRepo) UpdateFeedback(feedback entity.Feedback, ctx context.Context) error {
	var errLogMsg string = fmt.Sprintf(noti.REPO_ERR_MSG, feedback.GetFeedbackTable()) + "Updatefeedback - "
	var query string = "UPDATE " + feedback.GetFeedbackTable() + " SET content = $1, rating = $2, isDeleted = $3, updatedAt = $4 WHERE feedbackId = $5"

	res, err := f.db.Exec(query, feedback.Content, feedback.Rating, feedback.IsDeleted, feedback.UpdatedAt, feedback.FeedbackId)

	var INTERNALL_ERR_MSGMsg error = errors.New(noti.INTERNALL_ERR_MSG)

	if err != nil {
		f.logger.Println(errLogMsg + err.Error())
		return INTERNALL_ERR_MSGMsg
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		f.logger.Println(errLogMsg + err.Error())
		return INTERNALL_ERR_MSGMsg
	}

	if rowsAffected == 0 {
		return errors.New(fmt.Sprintf(noti.UNDEFINED_OBJECT_WARN_MSG, feedback.GetFeedbackTable()))
	}

	return nil
}
