package services

import (
	"errors"
	"github.com/Eco-Led/EcoLed-Back_test/initializers"
	"github.com/Eco-Led/EcoLed-Back_test/models"
)

type RankingServices struct{}

// Ranking Service's return value
type RankingForm struct {
	Nickname		string 	`form: "nickname" json: "nickname" binding: "required"`
	Age				int64 	`form: "age" json: "age" binding: "required"`	
	TotalEcoscore 	float64 `form: "total_ecoscore" json: "total_ecoscore" binding: "required"`
	Rank 			int64 	`form: "rank" json: "rank" binding: "required"`
}

func (svc RankingServices) GetRanking() (ranking []RankingForm, err error) {
	// Sort accounts by total_ecoscore
	var accounts []models.Accounts
	result := initializers.DB.Order("total_ecoscore desc").Find(&accounts)
	if result.Error != nil {
		err := errors.New("Failed to get accounts")
		return ranking, err
	}

	// Get rankingForm date from sorted accounts
	var currentRank int 
	var lastEcoscore float64 =-1 
	for i, account := range accounts {
		var profile models.Profiles
		result := initializers.DB.Where("user_id = ?", account.User_id).First(&profile)
	
		if result.Error != nil {
			// err := errors.New("Failed to get profile")
			continue
		}

		// To distinguish same ecoscore
		if account.Total_ecoscore != lastEcoscore {
			currentRank = i + 1
			lastEcoscore = account.Total_ecoscore
		}

		// Append rankingForm
		ranking = append(ranking, RankingForm{
			Nickname:	profile.Nickname,
			Age:		int64(profile.Age),
			TotalEcoscore:	account.Total_ecoscore,
			Rank:		int64(currentRank),
		})

	}
		
	return ranking, nil
	
}