package db

import (
	"database/sql"
	"errors"
	"log"
	"tourmate/feedback-service/constant/noti"
	db_server "tourmate/feedback-service/repository/db_server"
)

func ConnectDB(logger *log.Logger, server db_server.ISQLServer) (*sql.DB, error) {
	// Open database connection
	cnn, err := sql.Open(server.GetSQLServer(), server.GetCnnStr())

	if err != nil {
		logger.Println(noti.DB_CONNECTION_ERR_MSG + err.Error())
		return nil, errors.New(noti.INTERNALL_ERR_MSG)
	}

	return cnn, nil
}
