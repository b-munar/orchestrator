package handlers

import (
	"fiber-orchestrator/structs"
	"fiber-orchestrator/utils/concurrent"

	"github.com/gofiber/fiber/v2"

	uuid "github.com/google/uuid"

	"encoding/json"
)

func RegisterPartner(c *fiber.Ctx) error {
	UserUuid := uuid.New()

	var role int32

	role = 2

	response_partner, response_user := concurrent.GoRegisterPartner(c, UserUuid, role)

	if response_user.Err != nil || response_partner.Err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	if response_user.StatusCode != 201 || response_partner.StatusCode != 201 {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	var UserJson structs.UserResponse
	err_1 := json.Unmarshal(response_user.Data, &UserJson)
	if err_1 != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	var PartnerJson structs.PartnerResponse
	err_2 := json.Unmarshal(response_partner.Data, &PartnerJson)
	if err_2 != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	return c.Status(201).JSON(fiber.Map{"auth": UserJson.Auth, "partner": PartnerJson.Partner})

}
