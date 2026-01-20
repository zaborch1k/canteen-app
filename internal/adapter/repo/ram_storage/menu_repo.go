package ram_storage

import (
	"time"

	domMenu "canteen-app/internal/domain/menu"
	"canteen-app/internal/usecase"
)

type MenuRepo struct {
	Menu map[time.Time]domMenu.DayMenu
}

var _ usecase.MenuRepository = (*MenuRepo)(nil)

func NewMenuRepo() *MenuRepo {
	return &MenuRepo{
		Menu: make(map[time.Time]domMenu.DayMenu),
	}
}

func (mr *MenuRepo) GetMenuByDate(date time.Time) (*domMenu.DayMenu, error) {
	if menu, ok := mr.Menu[date]; ok {
		return &menu, nil
	}
	return &domMenu.DayMenu{}, usecase.ErrMenuNotFound
}
