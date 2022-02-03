package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/manishdangi98/banking/errs"
	"github.com/manishdangi98/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountId, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last instered id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")

	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepository {
	return AccountRepositoryDb{dbClient}
}
