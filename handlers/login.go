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

	if UserJson.Auth.Role == 1 {
		response_sport, response_sub := concurrent.GoLoginSportmen(c, UserJson.Auth.Token)
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

		type Activate struct {
			Activate bool `json:"activate"`
		}

		type SubscriptionActivate struct {
			Subscription Activate `json:"subscription"`
		}

		var SubAct SubscriptionActivate

		err_3 := json.Unmarshal(response_sub.Data, &SubAct)
		if err_3 != nil {
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		if SubAct.Subscription.Activate == false {
			var SubJson structs.SubscriptionList
			err_4 := json.Unmarshal(response_sub.Data, &SubJson)
			if err_4 != nil {
				return c.Status(400).JSON(fiber.Map{"status": "failure"})
			}
			return c.Status(202).JSON(fiber.Map{"sportmen": SportmenJson.Sportmen, "auth": UserJson.Auth, "subscription": SubJson.Subscriptions})

		}
	}
	if UserJson.Auth.Role == 2 {

		data, statusCode, err := request.GetPartner(UserJson.Auth.Token)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		if statusCode != 200 {
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		var PartnerJson structs.PartnerResponse
		err_2 := json.Unmarshal(data, &PartnerJson)
		if err_2 != nil {
			return c.Status(400).JSON(fiber.Map{"status": "failure"})
		}

		return c.Status(202).JSON(fiber.Map{"auth": UserJson.Auth, "partner": PartnerJson.Partner})

	}
	return c.Status(400).JSON(fiber.Map{"status": "failure"})
}
