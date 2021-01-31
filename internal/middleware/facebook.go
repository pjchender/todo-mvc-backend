package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pjchender/todo-mvc-backend/global"
	"github.com/pjchender/todo-mvc-backend/internal/database"
	"github.com/pjchender/todo-mvc-backend/internal/service"
	"github.com/pjchender/todo-mvc-backend/pkg/app"
	"github.com/pjchender/todo-mvc-backend/pkg/errMsg"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type AuthWithFacebook struct {
	DB *database.GormDatabase
}

type FacebookAuth struct {
	AppID   string `json:"app_id"`
	IsValid bool   `json:"is_valid"`
	UserID  string `json:"user_id"`
}

type FacebookAuthResp struct {
	Data FacebookAuth `json:"data"`
}

func (f *AuthWithFacebook) CheckFacebookLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		inputToken := ctx.GetHeader("fb-client-token")
		if inputToken == "" {
			app.SuccessOrAbort(ctx, http.StatusUnauthorized, errors.New(errMsg.InvalidToken))
		}

		facebookAuth, err := isUserLogin(inputToken, global.AuthSetting.Facebook.AppToken)
		if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
			log.Println("[middleware] facebook - isUserLogin failed: ", err)
			return
		}

		if !facebookAuth.IsValid {
			app.SuccessOrAbort(ctx, http.StatusUnauthorized, errors.New("facebook is not login"))
			return
		}

		svc := service.New(ctx, f.DB)
		userID, err := strconv.ParseInt(facebookAuth.UserID, 10, 64)
		if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
			log.Println("[middleware] facebook - strconv.ParseInt failed: ", err)
			return
		}

		param := service.FirstOrCreateUserRequest{
			FacebookUserID: uint(userID),
		}
		user, err := svc.FirstOrCreateUser(param)
		if success := app.SuccessOrAbort(ctx, http.StatusInternalServerError, err); !success {
			log.Println("[user] svc.FirstOrCreateUser failed: ", err)
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func isUserLogin(inputToken, appToken string) (*FacebookAuth, error) {
	baseURL := "https://graph.facebook.com/debug_token"
	endPoint := fmt.Sprintf("%s?input_token=%s&access_token=%s", baseURL, inputToken, appToken)

	log.Println("endPoint", endPoint)

	res, err := http.Get(endPoint)
	if err != nil {
		return nil, err
	}

	facebookAuthResp := FacebookAuthResp{}
	err = json.NewDecoder(res.Body).Decode(&facebookAuthResp)
	if err != nil {
		return nil, err
	}

	return &facebookAuthResp.Data, nil
}
