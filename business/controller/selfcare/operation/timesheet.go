package operation

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/jffp113/selfcare/business/config"
	"golang.org/x/net/html"
)

type registerOp struct {
	config    config.Config
	timesheet config.Timesheet
	login     Doer
	client    *http.Client
	host      string
}

func NewRegisterOp(cfg config.Config, timesheet config.Timesheet, host string) Doer {
	client := getClient()

	return &registerOp{
		config:    cfg,
		timesheet: timesheet,
		host:      host,
		client:    client,
		login:     NewLoginOperation(cfg, client, host),
	}
}

func (r *registerOp) Do() error {
	ctx, ok := r.config.Contexts[r.config.Context]
	if !ok {
		return fmt.Errorf("context '%s' not present", r.config.Context)
	}

	if err := r.login.Do(); err != nil {
		return err
	}

	for _, e := range r.timesheet.Entries {
		v, ok := r.config.Projects[e.Project]
		if !ok {
			return fmt.Errorf("not found project: %s", e.Project)
		}

		entry := entry{
			client:        v.Client,
			cpackage:      v.Package,
			projetos:      v.Project,
			task:          e.Tasks,
			serviceLevels: v.ServiceLevel,
			options:       v.Option,
			spent:         e.Time,
			timezone:      ctx.Timezone,
			startDate:     e.Day.Time,
			endDate:       e.Day.Time,
		}

		if err := r.registerHours(entry); err != nil {
			return err
		}
	}

	return nil
}

/*
func (r *registerOp) doLogin(ctx config.Context) error {
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
}*/

type entry struct {
	client        string
	cpackage      string
	projetos      string
	serviceLevels string
	options       string
	spent         float64
	task          []string
	timezone      string
	endDate       time.Time
	startDate     time.Time
}

func (r *registerOp) registerHours(e entry) error {
	resp, err := r.client.Get(timesheetURL(r.host))
	if err != nil {
		return err
	}

	node, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	n := htmlquery.FindOne(node, "//form/input")
	csrf := n.Attr[2].Val

	time := timeToString(e.startDate, e.timezone)

	resp, err = r.client.PostForm(saveTimesheetUrl(r.host), url.Values{
		"_csrf":                []string{csrf},
		"clientid":             []string{},
		"hourid":               []string{},
		"select_client":        []string{e.client},
		"select_cpackage":      []string{e.cpackage},
		"select_projetos":      []string{e.projetos},
		"select_ServiceLevels": []string{e.serviceLevels},
		"select_options":       []string{e.options},
		"hours_spent":          []string{fmt.Sprintf("%f", e.spent)},
		"task":                 []string{joinTasks(e.task)},
		"endDate":              []string{time},
		"startDate":            []string{time},
	})

	if err != nil {
		return err
	}

	bs, _ := io.ReadAll(resp.Body)

	if !strings.Contains(string(bs), time) {
		return fmt.Errorf("error registering hour")
	}

	return nil
}

func joinTasks(s []string) string {
	return strings.Join(s, "\n")
}

func timeToString(t time.Time, timezone string) string {
	return fmt.Sprintf("%s 00:00:00 %s", t.Format("Mon Jan 02 2006"), timezone)
}
