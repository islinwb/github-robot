package handlers

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/google/go-github/github"
	"io/ioutil"
	"net/http"
	"strings"
)

type GithubIssue github.Issue

func (s *Server) handleIssueEvent(client *github.Client) {
	glog.Infof("Received an Issue Event")

}

func (s *Server) handleIssueCommentEvent(r *http.Request) {
	glog.Infof("Received an IssueComment Event")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("fail to read request body: %v", err)
	}
	var prc github.IssueComment
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

func (issue *GithubIssue) TagLabel(client *github.Client, labels []string) {
}
