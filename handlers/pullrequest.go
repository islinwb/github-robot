package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

type GithubPR github.PullRequest

func (s *Server) handlePullRequestEvent(r *http.Request) {
	glog.Infof("Received an PullRequest Event")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("fail to read request body: %v", err)
	}
	var pr github.PullRequest
	err = json.Unmarshal(b, &pr)
	if err != nil {
		glog.Errorf("fail to unmarshal: %v", err)
	}
	if pr.GetState() == "open" {

	}

}

func (s *Server) handlePullRequestCommentEvent(r *http.Request) {
	glog.Infof("Received an PullRequestComment Event")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("fail to read request body: %v", err)
	}
	var prc github.PullRequestComment
	err = json.Unmarshal(b, &prc)
	if err != nil {
		glog.Errorf("fail to unmarshal: %v", err)
	}
	comment := *prc.Body

	if labelReg.MatchString(comment) {
		labelSlice := strings.Split(comment, " ")
		if len(labelSlice) > 0 {
		}
	}

	if retestReg.MatchString(comment) {
		SendToCI(r.Body)
	} else if testReg.MatchString(comment) {
		// TODO: trigger particular job(s)
	}

}

func (pr *GithubPR) TagLabel(labels []string) {

}
