package handlers

import (
	"fiber-orchestrator/structs"
	"fiber-orchestrator/utils/concurrent"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	uuid "github.com/google/uuid"

	"encoding/json"
)

func Register(c *fiber.Ctx) error {
	UserUuid := uuid.New()

	RoleId := new(structs.RoleId)

	if err := c.BodyParser(RoleId); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(RoleId)

	if err1 != nil {

		return c.Status(400).JSON(fiber.Map{"status": "failure"})

	}

	if RoleId.Role == 1 {
		response_sport, response_user, response_sub := concurrent.GoRegister(c, UserUuid)

		if response_sport.Err != nil || response_user.Err != nil || response_sub.Err != nil {
			concurrent.GoDelete(c, UserUuid)
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		if response_sport.StatusCode != 201 || response_user.StatusCode != 201 || response_sub.StatusCode != 200 {
			concurrent.GoDelete(c, UserUuid)
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		var SportmenJson structs.SportmenResponse
		err := json.Unmarshal(response_sport.Data, &SportmenJson)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		var UserJson structs.UserResponse
		err_1 := json.Unmarshal(response_user.Data, &UserJson)
		if err_1 != nil {
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		var SubJson structs.Subscription
		err_2 := json.Unmarshal(response_sub.Data, &SubJson)
		if err_2 != nil {
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		return c.Status(201).JSON(fiber.Map{"sportmen": SportmenJson.Sportmen, "auth": UserJson.Auth, "subscription": SubJson.Sub})

	}

	return c.Status(400).JSON(fiber.Map{"status": "failure"})
}
