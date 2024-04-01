package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetSportmen(token string) ([]byte, int, error) {

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.AutomaticEnv()

	sport_host, _ := viper.Get("SPORT_HOST").(string)
	sport_path, _ := viper.Get("SPORT_PATH_CREATE").(string)

	sport_url := fmt.Sprintf("%s%s", sport_host, sport_path)

	req, err := http.NewRequest(http.MethodGet, sport_url, nil)
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

	body_sport, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error in read Body")
	}

	defer res.Body.Close()

	return body_sport, res.StatusCode, nil
}
