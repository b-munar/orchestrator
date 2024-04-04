package concurrent

import (
	"fiber-orchestrator/structs"
	"fiber-orchestrator/utils/request"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/google/uuid"
)

func GoRegisterSportman(c *fiber.Ctx, UserUuid uuid.UUID, role int32) (structs.Response, structs.Response, structs.Response) {
	response_sport := make(chan structs.Response)
	response_user := make(chan structs.Response)
	response_sub := make(chan structs.Response)
	go func() {
		data, statusCode, err := request.CreateSportmen(c, UserUuid)
		response_sport <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	go func() {
		data, statusCode, err := request.CreateUser(c, UserUuid, role)
		response_user <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	go func() {
		data, statusCode, err := request.GetSub()
		response_sub <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	return <-response_sport, <-response_user, <-response_sub
}

func GoDelete(c *fiber.Ctx, UserUuid uuid.UUID) {
	request.DeleteSportmen(c, UserUuid)
	request.DeleteUser(c, UserUuid)
}
