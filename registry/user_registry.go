package registry

import (
	"github.com/mkeort/bath-hexagonal/interface/controller"
	ip "github.com/mkeort/bath-hexagonal/interface/presenter"
	ir "github.com/mkeort/bath-hexagonal/interface/repository"
	"github.com/mkeort/bath-hexagonal/usecase/interactor"
	up "github.com/mkeort/bath-hexagonal/usecase/presenter"
	ur "github.com/mkeort/bath-hexagonal/usecase/repository"
)

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}
