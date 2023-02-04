package presenter

import (
	"github.com/mkeort/bath-hexagonal/usecase/presenter"
)

type poopPresenter struct{}

func NewPoopPresenter() presenter.PoopPresenter {
	return &poopPresenter{}
}
