package providers

import (
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/controller"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/repository"
	"github.com/Caknoooo/go-gin-clean-starter/modules/user/service"
	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/samber/do"
	"gorm.io/gorm"
)

// RegisterUserModule 注册 User 模块的所有依赖
func RegisterUserModule(injector *do.Injector) {
	do.Provide(injector, func(i *do.Injector) (repository.UserRepository, error) {
		db := do.MustInvokeNamed[*gorm.DB](i, constants.DB)
		return repository.NewUserRepository(db), nil
	})

	do.Provide(injector, func(i *do.Injector) (service.UserService, error) {
		db := do.MustInvokeNamed[*gorm.DB](i, constants.DB)
		userRepo := do.MustInvoke[repository.UserRepository](i)
		return service.NewUserService(userRepo, db), nil
	})

	do.Provide(injector, func(i *do.Injector) (controller.UserController, error) {
		userService := do.MustInvoke[service.UserService](i)
		return controller.NewUserController(i, userService), nil
	})
}
