package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github-robot/handlers"

	"encoding/json"
	"github.com/golang/glog"
	"github.com/google/go-github/github"
	"github.com/spf13/pflag"
	"golang.org/x/oauth2"
)

type WebHookServer struct {
	Address    string
	Port       int64
	ConfigFile string
}

type Config struct {
	GitHubToken   string `json:"git_hub_token"`
	WebhookSecret string `json:"webhook_secret"`
	CircleCIToken string `json:"circle_ci_token"`
}

func NewWebHookServer() *WebHookServer {
	s := WebHookServer{
		Address:    "0.0.0.0",
		Port:       3000,
		ConfigFile: "/etc/github-robot/config.json",
	}
	return &s
}

func (s *WebHookServer) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Address, "address", s.Address, "IP address to serve, 0.0.0.0 by default")
	fs.Int64Var(&s.Port, "port", s.Port, "Port to listen on, 3000 by default")
	fs.StringVar(&s.ConfigFile, "config-file", s.ConfigFile, "Config file.")
}

func (s *WebHookServer) Run() {
	configContent, err := ioutil.ReadFile(s.ConfigFile)
	if err != nil {
		glog.Fatal("Could not read config file: %v", err)
	}
	var config Config
	err = json.Unmarshal(configContent, &config)
	if err != nil {
		glog.Fatal("fail to unmarshal: %v", err)
	}
	oauthSecret := config.GitHubToken
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(oauthSecret)},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	// Return 200 on / for health checks.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	webhookSecret := []byte(config.WebhookSecret)
	webHookHandler := handlers.Server{
		WebHookSecret: webhookSecret,
		GithubClient:  client,
	}
	//setting handler
	http.HandleFunc("/hook", webHookHandler.ServeHTTP)

	address := s.Address + ":" + strconv.FormatInt(s.Port, 10)
	//starting server
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Println(err)
	}
}

func main() {
	s := NewWebHookServer()
	s.AddFlags(pflag.CommandLine)

	s.Run()
}
