package retrosvc

import (
	"context"
	"time"

	"github.com/kralamoure/retro"
)

func (svc Service) CreateTicket(ctx context.Context, t retro.Ticket) (string, error) {
	return svc.storer.CreateTicket(ctx, t)
}

func (svc Service) DeleteTickets(ctx context.Context, before time.Time) (int, error) {
	return svc.storer.DeleteTickets(ctx, before)
}

func (svc Service) UseTicket(ctx context.Context, id string) (retro.Ticket, error) {
	return svc.storer.UseTicket(ctx, id)
}

func (svc Service) Tickets(ctx context.Context) (map[string]retro.Ticket, error) {
	return svc.storer.Tickets(ctx)
}

func (svc Service) Ticket(ctx context.Context, id string) (retro.Ticket, error) {
	return svc.storer.Ticket(ctx, id)
}
