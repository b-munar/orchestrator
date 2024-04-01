package request

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetSub() ([]byte, int, error) {

	viper.SetConfigFile(".env")

	viper.ReadInConfig()

	viper.AutomaticEnv()

	sub_host, _ := viper.Get("SUB_HOST").(string)

	sub_path, _ := viper.Get("SUB_PATH_LIST").(string)

	sub_url := fmt.Sprintf("%s%s", sub_host, sub_path)

	req_sub, err := http.NewRequest(http.MethodGet, sub_url, nil)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("could not create request")
	}

	req_sub.Header.Set("Content-Type", "application/json")

	res_sub, err := http.DefaultClient.Do(req_sub)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error making http request")
	}

	body_sub, err := ioutil.ReadAll(res_sub.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, errors.New("error in read Body")
	}

	defer res_sub.Body.Close()

	return body_sub, res_sub.StatusCode, nil
}
