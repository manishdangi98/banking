package service

import (
	"time"

	"github.com/manishdangi98/banking-lib/errs"
	"github.com/manishdangi98/banking/domain"
	"github.com/manishdangi98/banking/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	if err := req.Validate(); err != nil {
		return nil, err
	}
	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)
	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	//incoming request validate
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	//server side validation for checking the avialable balance in the account
	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insuffiecent balance in account")
		}
	}
	//if all is well then transaction take place
	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
