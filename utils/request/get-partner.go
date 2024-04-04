package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetPartner(token string) ([]byte, int, error) {

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.AutomaticEnv()

	partner_host, _ := viper.Get("PARTNER_HOST").(string)
	partner_path, _ := viper.Get("PARTNER_PATH_GET").(string)

	partner_url := fmt.Sprintf("%s%s", partner_host, partner_path)

	req, err := http.NewRequest(http.MethodGet, partner_url, nil)
	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not create request")
	}

	bearerToken := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", bearerToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error making http request")
	}

	body_partner, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error in read Body")
	}

	defer res.Body.Close()

	return body_partner, res.StatusCode, nil
}
