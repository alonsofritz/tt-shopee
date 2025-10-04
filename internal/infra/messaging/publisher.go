package messaging

import "github.com/alonsofritz/tt-shopee/internal/domain/model"

type TicketPublisher interface {
	Publish(ticket model.Ticket) error
	StartConsumer(handler func(ticket model.Ticket))
	Close()
}
