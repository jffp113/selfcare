package operation

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/jffp113/selfcare/business/config"
	"golang.org/x/net/html"
)

type loginOperation struct {
	config config.Config
	client *http.Client
	host   string
}

func NewLoginOperation(config config.Config, client *http.Client, host string) *loginOperation {
	return &loginOperation{
		config: config,
		client: client,
		host:   host,
	}
}

func (r *loginOperation) Do() error {
	ctx, ok := r.config.Contexts[r.config.Context]
	if !ok {
		return fmt.Errorf("context '%s' not present", r.config.Context)
	}

	resp, err := r.client.Get(loginUrl(r.host))
	if err != nil {
		return err
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	n := htmlquery.FindOne(node, "//form/input")
	csrf := n.Attr[2].Val

	resp, err = r.client.PostForm(signInUrl(r.host), url.Values{
		"_csrf":    []string{csrf},
		"username": []string{ctx.Username},
		"password": []string{ctx.Password},
	})

	if err != nil {
		return err
	}

	bs, _ := io.ReadAll(resp.Body)

	if strings.Contains(string(bs), "Invalid password") {
		return fmt.Errorf("invalid credentials")
	}

	return nil
}
