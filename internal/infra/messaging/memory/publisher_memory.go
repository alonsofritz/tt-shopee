package memory

import (
	"fmt"

	"github.com/alonsofritz/tt-shopee/internal/domain/model"
	"github.com/alonsofritz/tt-shopee/internal/infra/messaging"
)

type TicketPublisherMemory struct {
	queue chan model.Ticket
}

func NewTicketPublisherMemory(bufferSize int) messaging.TicketPublisher {
	return &TicketPublisherMemory{
		queue: make(chan model.Ticket, bufferSize),
	}
}

func (p *TicketPublisherMemory) Publish(ticket model.Ticket) error {
	select {
	case p.queue <- ticket:
		fmt.Printf("Ticket enviado para fila em memória: %+v\n", ticket)
	default:
		fmt.Println("Fila cheia! Ticket descartado ou reter para retry")
	}
	return nil
}

func (p *TicketPublisherMemory) Close() {
	close(p.queue)
}
