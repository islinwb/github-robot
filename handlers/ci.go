package handlers

import (
	"fmt"
	"github.com/golang/glog"
	"io"
	"net/http"
	"os"
)

const (
	CircleCIurl  = "https://circleci.com/hooks/github"
	ContentType  = "application/x-www-form-urlencoded"
	CircleAPIURL = "https://circleci.com/api/v1.1/project/github/islinwb/test/build?circle-token="
)

func SendToCI(body []byte) {
	// currently only support circle ci
	CIRCLE_API_USER_TOKEN := os.Getenv("CIRCLE_API_USER_TOKEN")
	glog.Info("going to send test request to circle ci")
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf(CircleAPIURL+CIRCLE_API_USER_TOKEN), body)
	if err != nil {
		glog.Errorf("%v", err)
	}
	req.Header.Set("Content-Type", ContentType)
	resp, err := client.Do(req)
	glog.Infof("%v", resp)
}
