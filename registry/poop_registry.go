package registry

import (
	"github.com/mkeort/bath-hexagonal/interface/controller"
	ip "github.com/mkeort/bath-hexagonal/interface/presenter"
	ir "github.com/mkeort/bath-hexagonal/interface/repository"
	"github.com/mkeort/bath-hexagonal/usecase/interactor"
	up "github.com/mkeort/bath-hexagonal/usecase/presenter"
	ur "github.com/mkeort/bath-hexagonal/usecase/repository"
)

func (r *registry) NewPoopPresenter() up.PoopPresenter {
	return ip.NewPoopPresenter()
}

func (r *registry) NewPoopRepository() ur.PoopRepository {
	return ir.NewPoopRepository(r.db)
}

func (r *registry) NewPoopInteractor() interactor.PoopInteractor {
	return interactor.NewPoopInteractor(r.NewPoopRepository(), r.NewPoopPresenter(), ir.NewDBRepository(r.db))
}

func (r *registry) NewPoopController() controller.PoopController {
	return controller.NewPoopController(r.NewPoopInteractor())
}
