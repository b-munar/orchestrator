package concurrent

import (
	"fiber-orchestrator/structs"
	"fiber-orchestrator/utils/request"

	"github.com/gofiber/fiber/v2"
)

func GoLoginSportmen(c *fiber.Ctx, token string) (structs.Response, structs.Response) {
	response_sport := make(chan structs.Response)
	response_sub := make(chan structs.Response)
	go func() {
		data, statusCode, err := request.GetSportmen(token)
		response_sport <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	go func() {
		data, statusCode, err := request.GetSub(token)
		response_sub <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	return <-response_sport, <-response_sub

}
