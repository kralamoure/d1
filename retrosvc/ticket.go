package retrosvc

import (
	"context"
	"time"

	"github.com/kralamoure/retro"
)

func (svc Service) CreateTicket(ctx context.Context, t retro.Ticket) (string, error) {
	return svc.repo.CreateTicket(ctx, t)
}

func (svc Service) DeleteTickets(ctx context.Context, before time.Time) (int, error) {
	return svc.repo.DeleteTickets(ctx, before)
}

func (svc Service) UseTicket(ctx context.Context, id string) (retro.Ticket, error) {
	return svc.repo.UseTicket(ctx, id)
}

func (svc Service) Tickets(ctx context.Context) (map[string]retro.Ticket, error) {
	return svc.repo.Tickets(ctx)
}

func (svc Service) Ticket(ctx context.Context, id string) (retro.Ticket, error) {
	return svc.repo.Ticket(ctx, id)
}
