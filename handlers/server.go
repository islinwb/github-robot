package handlers

import (
	"github.com/golang/glog"
	"net/http"

	"github.com/google/go-github/github"
	"io"
	"io/ioutil"
)

// Server implements http.Handler. It validates incoming GitHub webhooks and
// then dispatches them to the handlers accordingly.
type Server struct {
	WebHookSecret []byte
	GithubClient  *github.Client
}

// ServeHTTP validates an incoming webhook and invoke its handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//payload, err := github.ValidatePayload(r, s.WebHookSecret)
	payload, err := github.ValidatePayload(r, s.WebHookSecret)
	if err != nil {
		glog.Errorf("Invalid payload: %v", err)
		return
	}
	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		glog.Errorf("Failed to parse webhook")
		return
	}
	//fmt.Fprint(w, "Received a webhook event")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("fail to read: %v", err)
	}
	glog.Infof("body: %v", body)

	switch event.(type) {
	case *github.IssueEvent:
		go s.handleIssueEvent(s.GithubClient)
	case *github.IssueCommentEvent:
		go s.handleIssueCommentEvent(body)
	case *github.PullRequest:
		go s.handlePullRequestEvent(r)
	case *github.PullRequestComment:
		go s.handlePullRequestCommentEvent(r)

	}
}
