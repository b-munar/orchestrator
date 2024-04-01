package handlers

import (
	"encoding/json"
	"fiber-orchestrator/structs"
	"fiber-orchestrator/utils/concurrent"
	"fiber-orchestrator/utils/request"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	data, statusCode, err := request.LoginUser(c)
	response_user := structs.Response{Data: data, StatusCode: statusCode, Err: err}

	var UserJson structs.UserResponse
	err_1 := json.Unmarshal(response_user.Data, &UserJson)
	if err_1 != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	response_sport, response_sub := concurrent.GoLogin(c, UserJson.Auth.Token)

	if response_sport.Err != nil || response_user.Err != nil || response_sub.Err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	if response_sport.StatusCode != 200 || response_user.StatusCode != 202 || response_sub.StatusCode != 200 {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}
	var SportmenJson structs.SportmenResponse
	err_2 := json.Unmarshal(response_sport.Data, &SportmenJson)
	if err_2 != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	var SubJson structs.Subscription
	err_3 := json.Unmarshal(response_sub.Data, &SubJson)
	if err_3 != nil {
		return c.Status(400).JSON(fiber.Map{"status": "failure"})
	}

	return c.Status(202).JSON(fiber.Map{"sportmen": SportmenJson.Sportmen, "auth": UserJson.Auth, "subscription": SubJson.Sub})
}
