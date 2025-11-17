package api

import (
	"fmt"
	"net/http"
	"os"
	"post-management/pkg/dto"

	"github.com/goccy/go-json"
)

func UserCountById(id int64) (int64, error) {
	url := fmt.Sprintf("%s:%s/api/users/services/count/%d",
		os.Getenv("API_USER_HOST"),
		os.Getenv("API_USER_PORT"),
		id,
	)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	body := new(dto.ApiWebResponse[dto.ApiUserCountResponse])
	err = json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("code: %d, msg: %s", resp.StatusCode, *body.Error)
	}
	return body.Data.Total, nil
}
