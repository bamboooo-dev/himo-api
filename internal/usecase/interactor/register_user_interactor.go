package interactor

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/bamboooo-dev/himo-outgame/internal/domain/model"
	"github.com/bamboooo-dev/himo-outgame/internal/registry"
	himo_repo "github.com/bamboooo-dev/himo-outgame/internal/usecase/repository/himo"
	"github.com/bamboooo-dev/himo-outgame/pkg/grpcmiddleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-gorp/gorp"
)

// RegisterUserInteractor ユーザーを登録するアプリケーションサービス
type RegisterUserInteractor struct {
	userRepo himo_repo.UserRepository
}

// NewRegistUserInteractor constructs RegisterUserInteractor
func NewRegistUserInteractor(r registry.Registry) *RegisterUserInteractor {
	return &RegisterUserInteractor{
		userRepo: r.NewUserRepository(),
	}
}

// Call は受け取ったニックネームでユーザーを登録する関数
func (r *RegisterUserInteractor) Call(ctx context.Context, db *gorp.DbMap, nickName string) (string, error) {
	user := model.User{
		Nickname: nickName,
	}
	user, err := r.userRepo.Create(ctx, db, user)
	if err != nil {
		return "", err
	}

	signKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), grpcmiddleware.AuthClaim{
		UserID: strconv.Itoa(int(user.ID)),
		StandardClaims: jwt.StandardClaims{
			// 1年後
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
	})
	ss, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}
