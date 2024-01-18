package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",                               // 이 주소는 Google 콘솔에 등록된 리디렉션 URL과 일치해야 합니다.
		ClientID:     "", 																		  // 여기에 Google 클라이언트 ID를 넣습니다.
		ClientSecret: "",                                      									  // 여기에 Google 클라이언트 비밀번호를 넣습니다.
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

// Google 로그인 페이지로 리다이렉트하는 핸들러
func GoogleLoginHandler(c *gin.Context) {
	// Google OAuth2 URL을 생성합니다.
	url := googleOauthConfig.AuthCodeURL("state")
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Google 로그인 콜백을 처리하는 핸들러
func GoogleAuthCallback(c *gin.Context) {
	// URL 쿼리에서 코드를 추출합니다.
	code := c.Query("code")

	// 코드를 사용하여 토큰을 교환합니다.
	token, err := googleOauthConfig.Exchange(c.Request.Context(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Code exchange failed",
		})
		return
	}

	// 토큰을 사용하여 HTTP 클라이언트를 생성합니다.
	client := googleOauthConfig.Client(c.Request.Context(), token)

	// Google 사용자 정보 API에 요청을 보냅니다.
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user info",
		})
		return
	}
	defer resp.Body.Close()

	// 응답 데이터를 파싱합니다.
	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode user info",
		})
		return
	}

	// 사용자 정보를 응답으로 반환합니다.
	c.JSON(http.StatusOK, gin.H{
		"email": userInfo.Email,
		"name":  userInfo.Name,
	})
}
