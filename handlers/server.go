package handlers

import (
	"fmt"
	"github.com/golang/glog"
	"net/http"

	"github.com/google/go-github/github"
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
	payload, err := github.ValidatePayload(r, []byte("1ec31a1faae0ee37f1d8472941783092c841c130"))
	if err != nil {
		glog.Errorf("Invalid payload: %v", err)
		return
	}
	glog.Infof("received a payload: %v", payload)
	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		glog.Errorf("Failed to parse webhook")
		return
	}
	fmt.Fprint(w, "Received a webhook event")

	switch event.(type) {
	case *github.IssueEvent:
		go s.handleIssueEvent(s.GithubClient)
	case *github.IssueCommentEvent:
		go s.handleIssueCommentEvent(s.GithubClient)
	case *github.PullRequest:
		go s.handlePullRequestEvent(r)
	case *github.PullRequestComment:
		go s.handlePullRequestCommentEvent(r)

	}
}
