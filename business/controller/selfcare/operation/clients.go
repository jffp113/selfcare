package operation

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/jffp113/selfcare/business/config"
	"golang.org/x/net/html"
)

type Client struct {
	Name string
	Id   string
}

type clientsOp struct {
	config    config.Config
	timesheet config.Timesheet
	login     Doer
	client    *http.Client
	host      string
}

func NewClientOp(cfg config.Config, host string) *clientsOp {
	client := getClient()

	return &clientsOp{
		config: cfg,
		host:   host,
		client: client,
		login:  NewLoginOperation(cfg, client, host),
	}
}

func (r *clientsOp) Do() ([]Client, error) {
	err := r.login.Do()
	if err != nil {
		return nil, err
	}

	var clients []Client

	userId, err := r.getUserId()
	if err != nil {
		return nil, err
	}

	resp, err := r.client.Get(getClientsUrl(r.host, userId))
	if err != nil {
		return nil, err
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, &clients)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *clientsOp) getUserId() (string, error) {
	resp, err := r.client.Get(baseUrl(r.host))
	if err != nil {
		return "", err
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	n := htmlquery.Find(node, "//a")

	return getUserIdAnchor(n), nil
}

func getUserIdAnchor(nodes []*html.Node) string {
	var a *html.Node
	for _, node := range nodes {
		if len(node.Attr) < 3 {
			continue
		}

		if strings.Contains(node.Attr[2].Val, "/user") {
			a = node
			break
		}
	}

	url := a.Attr[2].Val
	id := strings.Split(url, "/")[2]
	return id
}
