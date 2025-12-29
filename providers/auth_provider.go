package providers

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/controller"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/repository"
	"github.com/Caknoooo/go-gin-clean-starter/modules/auth/service"
	userRepo "github.com/Caknoooo/go-gin-clean-starter/modules/user/repository"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func RegisterAuthModule(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (repository.RefreshTokenRepository, error) {
		db := do.MustInvokeNamed[*gorm.DB](i, constants.DB)
		return repository.NewRefreshTokenRepository(db), nil
	})

	do.Provide(injector, func(i *do.Injector) (service.AuthService, error) {
		db := do.MustInvokeNamed[*gorm.DB](i, constants.DB)
		jwtService := do.MustInvokeNamed[service.JWTService](i, constants.JWTService)
		userRepository := do.MustInvoke[userRepo.UserRepository](i)
		refreshTokenRepo := do.MustInvoke[repository.RefreshTokenRepository](i)
		return service.NewAuthService(userRepository, refreshTokenRepo, jwtService, db), nil
	})

	do.Provide(injector, func(i *do.Injector) (controller.AuthController, error) {
		authService := do.MustInvoke[service.AuthService](i)
		return controller.NewAuthController(i, authService), nil
	})
}
