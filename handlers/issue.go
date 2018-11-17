package handlers

import (
	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

type GithubIssue github.Issue

func (s *Server) handleIssueEvent(client *github.Client) {
	glog.Infof("Received an Issue Event")

}

func (s *Server) handleIssueCommentEvent(client *github.Client) {
	glog.Infof("Received an IssueComment Event")
}

func (issue *GithubIssue) TagLabel(client *github.Client, labels []string) {
}
