package usecase

import (
	"context"
	"server/ent"

	"github.com/samber/lo"
)

type SearchResult struct {
	Labels     []*Label
	Statuses   []*Status
	Priorities []*Priority
}

func (u *usecase) Search(ctx context.Context) (*SearchResult, error) {
	search, err := u.TodoRepositoryInterface.Search(ctx)
	if err != nil {
		return nil, err
	}

	return &SearchResult{
		Labels: lo.Map(search.Labels, func(item *ent.Label, index int) *Label {
			return &Label{
				ID:    item.ID,
				Value: item.Value,
			}
		}),
		Statuses: lo.Map(search.Statuses, func(item *ent.Status, index int) *Status {
			return &Status{
				ID:    item.ID,
				Value: item.Value,
			}
		}),
		Priorities: lo.Map(search.Priorities, func(item *ent.Priority, index int) *Priority {
			return &Priority{
				ID:   item.ID,
				Name: item.Name,
			}
		}),
	}, nil
}
