package tab

import "github.com/cqrs-domain-ddd/v1/src/domain/tab/model"

type IHandler struct {
	Handle (model.OpenTab)
}

type Handler struct{}

func (h Handler) handle(tab model.OpenTab) model.TabOpened {
	item := model.TabOpened{tab.Guid,
		tab.TableNumber,
		tab.Waiter}
	return item
}
