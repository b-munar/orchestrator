package concurrent

import (
	"fiber-orchestrator/structs"
	"fiber-orchestrator/utils/request"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/google/uuid"
)

func GoRegisterPartner(c *fiber.Ctx, UserUuid uuid.UUID, role int32) (structs.Response, structs.Response) {

	response_partner := make(chan structs.Response)
	response_user := make(chan structs.Response)
	go func() {
		data, statusCode, err := request.CreatePartner(c, UserUuid)
		response_partner <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()
	go func() {
		data, statusCode, err := request.CreateUser(c, UserUuid, role)
		response_user <- structs.Response{Data: data, StatusCode: statusCode, Err: err}
	}()

	return <-response_partner, <-response_user
}
