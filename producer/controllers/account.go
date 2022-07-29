package controllers

import (
	"log"
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return accountController{accountService}
}

func (controller accountController) OpenAccount(c *fiber.Ctx) error {
	command := commands.OpenAccountCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	id, err := controller.accountService.OpenAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "open account success",
		"id":      id,
	})
}

func (controller accountController) DepositFund(c *fiber.Ctx) error {
	command := commands.DepositFundCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = controller.accountService.DepositFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deposit fund success",
	})
}

func (controller accountController) WithdrawFund(c *fiber.Ctx) error {
	command := commands.WithdrawFundCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = controller.accountService.WithdrawFund(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "withdraw fund success",
	})
}

func (controller accountController) CloseAccount(c *fiber.Ctx) error {
	command := commands.CloseAccountCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return err
	}

	err = controller.accountService.CloseAccount(command)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "close account success",
	})
}
