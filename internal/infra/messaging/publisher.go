package rabbitmq

import "github.com/alonsofritz/tt-shopee/internal/domain/model"

type TicketPublisher interface {
	Publish(ticket model.Ticket) error
}
