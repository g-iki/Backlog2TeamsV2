package backlog

import (
	"time"
)

type Connections struct {
	TeamsWebhookURL string `dynamo:"TeamsWebhookURL,omitempty"`
	Notes           string `dynamo:"Notes,omitempty"`
}

type Backlog struct {
	ID          string        `dynamo:"ID"`
	Connections []Connections `dynamo:"Connections"`
	Owner       string        `dynamo:"Owner,omitempty"`
	CreateDate  time.Time     `dynamo:"CreateDate,omitempty"`
	UpdateDate  time.Time     `dynamo:"UpdateDate,omitempty"`
}

func (c *Connections) SetTeamsWebhookURL(s string) {
	c.TeamsWebhookURL = s
}

func (c *Connections) GetTeamsWebhookURL() string {
	return c.TeamsWebhookURL
}

func (c *Connections) SetNotes(s string) {
	c.Notes = s
}

func (c *Connections) GetNotes() string {
	return c.Notes
}

func (b *Backlog) AddConnections(c Connections) {
	b.Connections = append(b.Connections, c)
}
func (b *Backlog) GetConnections() []Connections {
	return b.Connections
}
func (b *Backlog) SetConnections(con []Connections) {
	b.Connections = con
}
func (b *Backlog) SetID(s string) {
	b.ID = s
}

func (b *Backlog) GetID() string {
	return b.ID
}

func (b *Backlog) SetOwner(s string) {
	b.Owner = s
}

func (b *Backlog) GetOwner() string {
	return b.Owner
}

func (b *Backlog) SetCreateDate(t time.Time) {
	b.CreateDate = t
}

func (b *Backlog) GetCreateDate() time.Time {
	return b.CreateDate
}

func (b *Backlog) SetUpdateDate(t time.Time) {
	b.UpdateDate = t
}

func (b *Backlog) GetUpdateDate() time.Time {
	return b.UpdateDate
}
