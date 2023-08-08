package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetCampaignByID(campaignID int) ([]Transaction, error)
	GetByUserID(userID int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetCampaignByID(CampaignID int) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", CampaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

func (r *repository) GetByUserID(userID int) ([]Transaction, error){
	var transactions []Transaction

	err := r.db.Preload("Campaign.CampaignImages","campaign_images.is_primary = 1").Where("user_id = ?", userID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}