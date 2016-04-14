package quay

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/quay/billing"
	"github.com/coreos/go-quay/quay/build"
	"github.com/coreos/go-quay/quay/discovery"
	"github.com/coreos/go-quay/quay/error"
	"github.com/coreos/go-quay/quay/image"
	"github.com/coreos/go-quay/quay/logs"
	"github.com/coreos/go-quay/quay/organization"
	"github.com/coreos/go-quay/quay/permission"
	"github.com/coreos/go-quay/quay/prototype"
	"github.com/coreos/go-quay/quay/repository"
	"github.com/coreos/go-quay/quay/repositorynotification"
	"github.com/coreos/go-quay/quay/repotoken"
	"github.com/coreos/go-quay/quay/robot"
	"github.com/coreos/go-quay/quay/search"
	"github.com/coreos/go-quay/quay/secscan"
	"github.com/coreos/go-quay/quay/tag"
	"github.com/coreos/go-quay/quay/team"
	"github.com/coreos/go-quay/quay/trigger"
	"github.com/coreos/go-quay/quay/user"
)

// Default client HTTP client.
var Default = NewHTTPClient(nil)

// NewHTTPClient creates a new client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Client {
	if formats == nil {
		formats = strfmt.Default
	}
	transport := httptransport.New("quay.io", "/", []string{"https"})
	return New(transport, formats)
}

// New creates a new client client
func New(transport client.Transport, formats strfmt.Registry) *Client {
	cli := new(Client)
	cli.Transport = transport

	cli.Billing = billing.New(transport, formats)

	cli.Build = build.New(transport, formats)

	cli.Discovery = discovery.New(transport, formats)

	cli.Error = error.New(transport, formats)

	cli.Image = image.New(transport, formats)

	cli.Logs = logs.New(transport, formats)

	cli.Organization = organization.New(transport, formats)

	cli.Permission = permission.New(transport, formats)

	cli.Prototype = prototype.New(transport, formats)

	cli.Repository = repository.New(transport, formats)

	cli.Repositorynotification = repositorynotification.New(transport, formats)

	cli.Repotoken = repotoken.New(transport, formats)

	cli.Robot = robot.New(transport, formats)

	cli.Search = search.New(transport, formats)

	cli.Secscan = secscan.New(transport, formats)

	cli.Tag = tag.New(transport, formats)

	cli.Team = team.New(transport, formats)

	cli.Trigger = trigger.New(transport, formats)

	cli.User = user.New(transport, formats)

	return cli
}

// Client is a client for client
type Client struct {
	Billing *billing.Client

	Build *build.Client

	Discovery *discovery.Client

	Error *error.Client

	Image *image.Client

	Logs *logs.Client

	Organization *organization.Client

	Permission *permission.Client

	Prototype *prototype.Client

	Repository *repository.Client

	Repositorynotification *repositorynotification.Client

	Repotoken *repotoken.Client

	Robot *robot.Client

	Search *search.Client

	Secscan *secscan.Client

	Tag *tag.Client

	Team *team.Client

	Trigger *trigger.Client

	User *user.Client

	Transport client.Transport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Client) SetTransport(transport client.Transport) {
	c.Transport = transport

	c.Billing.SetTransport(transport)

	c.Build.SetTransport(transport)

	c.Discovery.SetTransport(transport)

	c.Error.SetTransport(transport)

	c.Image.SetTransport(transport)

	c.Logs.SetTransport(transport)

	c.Organization.SetTransport(transport)

	c.Permission.SetTransport(transport)

	c.Prototype.SetTransport(transport)

	c.Repository.SetTransport(transport)

	c.Repositorynotification.SetTransport(transport)

	c.Repotoken.SetTransport(transport)

	c.Robot.SetTransport(transport)

	c.Search.SetTransport(transport)

	c.Secscan.SetTransport(transport)

	c.Tag.SetTransport(transport)

	c.Team.SetTransport(transport)

	c.Trigger.SetTransport(transport)

	c.User.SetTransport(transport)

}
