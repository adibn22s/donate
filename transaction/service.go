package transaction

import (
	"errors"
	"golang/campaign"
	"golang/payment"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService payment.Service
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.User.ID != input.User.ID{
		return []Transaction{}, errors.New("Not an Owner of the campaign")
	} 
	
	transactions, err := s.repository.GetCampaignByID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error){
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil {

		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error){
	transaction := Transaction{}
	transaction.Amount = input.Amount
	transaction.CampaignID = input.CampaignID
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {

		return newTransaction, err
	}

	NewpaymentTransaction := payment.Transaction{
		ID: newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(NewpaymentTransaction, input.User)
	if err != nil {

		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}