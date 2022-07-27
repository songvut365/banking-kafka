package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepository repositories.AccountRepository
}

func NewAccountEventHandler(accountRepository repositories.AccountRepository) EventHandler {
	return accountEventHandler{accountRepository: accountRepository}
}
func (eventHandler accountEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	// Open Account Event
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := &events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}

		err = eventHandler.accountRepository.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(event)

	// DepositFundEvent
	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := &events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount, err := eventHandler.accountRepository.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount.Balance += event.Amount

		err = eventHandler.accountRepository.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(event)

	// Withdraw Fund Event
	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():
		event := &events.WithdrawFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount, err := eventHandler.accountRepository.FindByID(event.ID)
		if err != nil {
			log.Println(err)
			return
		}

		bankAccount.Balance -= event.Amount

		err = eventHandler.accountRepository.Save(bankAccount)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(event)

	// Close Account Event
	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := &events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Println(err)
			return
		}

		err = eventHandler.accountRepository.Delete(event.ID)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(event)

	default:
		log.Println("no event handler")
	}
}
