package services

import (
	"errors"
	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
	"time"
)

type PaylogForm struct {
	Date    	string 	`form: "date" json: "date" binding: "required"`
	Time    	string 	`form: "time" json: "time" binding: "required"`
	Content 	string  `form: "content" json: "content" binding: "required"`
	Cost    	int64 	`form: "cost" json: "cost" binding: "required"`
	Name		string 	`form: "name" json: "name" binding: "required"`
	Place		string	`form: "place" json: "place" binding: "required"`
	Material	string 	`form: "material" json: "material" binding: "required"`
	Ecoscore 	float64 `form: "ecoscore" json: "ecoscore" binding: "required"`
}


type PaylogServices struct{}

func (svc PaylogServices) CreatePaylog(userID uint, paylog PaylogForm) (err error) {
	// Get account
	var account models.Accounts
	result := initializers.DB.First(&account, "user_id=?", userID)
	if result.Error != nil {
		err := errors.New("Failed to get account")
		return err
	}

	// Create paylog
	result = initializers.DB.Create(&models.Paylogs{
		Date:    	paylog.Date,
		Time:    	paylog.Time,
		Content: 	paylog.Content,
		Cost:    	paylog.Cost,
		Name:		paylog.Name,
		Place:		paylog.Place,
		Material:	paylog.Material,
		Ecoscore: 	paylog.Ecoscore,
		Account_id: account.ID,
	})
	if result.Error != nil {
		err := errors.New("Failed to create paylog")
		return err
	}

	//Update account
	account.Total_ecoscore += paylog.Ecoscore
	account.Balance -= paylog.Cost
	result = initializers.DB.Save(&account)
	if result.Error != nil {
		err := errors.New("Failed to update account")
		return err
	}

	return nil
}

func (svc PaylogServices) UpdatePaylog(userID uint, paylogID uint, paylog PaylogForm) (err error) {
	// Get account
	var account models.Accounts
	result := initializers.DB.First(&account, "user_id=?", userID)
	if result.Error != nil {
		err := errors.New("Failed to get account")
		return err
	}

	// Get paylog
	var paylogModel models.Paylogs
	result = initializers.DB.First(&paylogModel, "id=?", paylogID)
	if result.Error != nil {
		err := errors.New("Failed to get paylog")
		return err
	}

	// Update paylog
	paylogModel.Date = paylog.Date
	paylogModel.Time = paylog.Time
	paylogModel.Content = paylog.Content
	paylogModel.Cost = paylog.Cost
	paylogModel.Name = paylog.Name
	paylogModel.Place = paylog.Place
	paylogModel.Material = paylog.Material
	paylogModel.Ecoscore = paylog.Ecoscore
	result = initializers.DB.Save(&paylogModel)
	if result.Error != nil {
		err := errors.New("Failed to update paylog")
		return err
	}

	//Update account
	account.Total_ecoscore += paylog.Ecoscore
	account.Balance -= paylog.Cost
	result = initializers.DB.Save(&account)
	if result.Error != nil {
		err := errors.New("Failed to update account")
		return err
	}

	return nil
}

func (svc PaylogServices) GetPaylogs(accountID uint, page int) (paylogs []models.Paylogs, err error) {
	// Setting page
	endDate := time.Now().AddDate(0, 0, -page*14)
	startDate := endDate.AddDate(0, 0, -14)
	endDateStr := endDate.Format("200601021504")
	startDateStr := startDate.Format("200601021504")

	// Get paylogs
	result := initializers.DB.Where("account_id=?", accountID).
	Where("deleted_at IS NULL").
	Where("CONCAT(Date, Time) BETWEEN ? AND ?", startDateStr, endDateStr).
	Order("CONCAT(Date, Time) DESC").
	Find(&paylogs)

	if result.Error != nil {
		err := errors.New("Failed to get paylogs")
		return paylogs, err
	}

	if result.RowsAffected == 0 {
		err := errors.New("No paylogs found")
		return paylogs, err
	}

	return paylogs, nil
}


