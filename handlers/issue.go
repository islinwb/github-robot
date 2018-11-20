package handlers

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/google/go-github/github"
	"io"
	"io/ioutil"
	"strings"
)

type GithubIssue github.Issue

func (s *Server) handleIssueEvent(client *github.Client) {
	glog.Infof("Received an Issue Event")

}

func (s *Server) handleIssueCommentEvent(body []byte) {
	glog.Infof("Received an IssueComment Event")

	var prc github.IssueComment
	err := json.Unmarshal(body, &prc)
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
		SendToCI(body)
	} else if testReg.MatchString(comment) {
		// TODO: trigger particular job(s)
	}

}

func (issue *GithubIssue) TagLabel(client *github.Client, labels []string) {
}
