package repositories

import (
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/db"
	"github.com/emirhandogandemir/bitirmego/cloud-infra-rest1/models"
)

func CreateAwsAccess(awsAccessModel *models.AwsAccessModel) (*models.AwsAccessModel, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, err
	}

	result := db.Create(&awsAccessModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return awsAccessModel, nil
}

func GetAllAwsAccess() ([]*models.AwsAccessModel, error) {
	db, err := db.Connect()

	if err != nil {
		return nil, err
	}

	var awsAccessModel []*models.AwsAccessModel
	result := db.Find(&awsAccessModel)
	if result.Error != nil {
		return nil, result.Error
	}

	return awsAccessModel, nil
}