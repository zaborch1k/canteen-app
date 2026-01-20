package usecase

import (
	domMenu "canteen-app/internal/domain/menu"
	"time"
)

type menuUseCase struct {
	menu MenuRepository
}

func NewMenuUseCase(menu MenuRepository) *menuUseCase {
	return &menuUseCase{menu: menu}
}

func (uc *menuUseCase) GetMenuByDate(date time.Time) (*domMenu.DayMenu, error) {
	menu, err := uc.menu.GetMenuByDate(date)
	if err != nil {
		return &domMenu.DayMenu{}, err
	}
	return menu, nil
}
