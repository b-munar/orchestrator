package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"fiber-orchestrator/structs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	uuid "github.com/google/uuid"
)

func CreatePartner(c *fiber.Ctx, userUuid uuid.UUID) ([]byte, int, error) {

	//validation

	NewPartner := new(structs.Partner)

	if err := c.BodyParser(NewPartner); err != nil {
		return nil, http.StatusInternalServerError, errors.New("error parser sportman")
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(NewPartner)

	if err1 != nil {

		return nil, http.StatusInternalServerError, errors.New("error validate sportman")

	}

	NewPartner.UserId = userUuid

	JsonPartner, _ := json.Marshal(NewPartner)

	// request

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.AutomaticEnv()

	partner_host, _ := viper.Get("PARTNER_HOST").(string)
	partner_path, _ := viper.Get("PARTNER_PATH_CREATE").(string)

	partner_url := fmt.Sprintf("%s%s", partner_host, partner_path)

	req_partner, err := http.NewRequest(http.MethodPost, partner_url, bytes.NewBuffer(JsonPartner))

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not create request")
	}

	req_partner.Header.Set("Content-Type", "application/json")

	res_sport, err := http.DefaultClient.Do(req_partner)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error making http request")
	}

	body_sport, err := ioutil.ReadAll(res_sport.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error in read Body")
	}

	defer res_sport.Body.Close()

	return body_sport, res_sport.StatusCode, nil
}
