package services

type Token struct {
	AccessToken  string `json: "access_token"`
	RefreshToken string `json: "refresh_token"`
}

type TokenServices struct{}

func (svc TokenServices) CreateToken() (token Token) {

}
