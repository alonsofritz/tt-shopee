package service

import (
	"errors"

	"github.com/alonsofritz/tt-shopee/internal/domain/model"
	"github.com/alonsofritz/tt-shopee/internal/domain/repository"
	"github.com/alonsofritz/tt-shopee/internal/infra/messaging"
)

type TicketService struct {
	ShowRepo  repository.ShowRepository
	UserRepo  repository.UserRepository
	Publisher messaging.TicketPublisher
}

func (ts *TicketService) ProcessTicket(ticket model.Ticket) error {
	if _, err := ts.ShowRepo.Exists(ticket.ShowID); err != nil {
		return errors.New("show not found")
	}

	if _, err := ts.UserRepo.Exists(ticket.UserID); err != nil {
		return errors.New("user not found")
	}

	if err := ts.Publisher.Publish(ticket); err != nil {
		return err
	}

	return nil
}
