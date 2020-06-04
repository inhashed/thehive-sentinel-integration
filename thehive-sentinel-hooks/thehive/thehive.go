package thehive

import (
	"encoding/json"

	"github.com/nviso-be/thehive-sentinel-integration/thehive-sentinel-hooks/config"
)

type Details struct {
	Status string `json:"status"`
	CaseID int    `json:"caseId"`
}

type SentinelIncidentNumber struct {
	IncidentNumber int `json:"number"`
}

func (n *SentinelIncidentNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"incidentNumber": n.IncidentNumber,
	})
}

type AlertIDs struct {
	Alerts string `json:"string"`
}

func (a *AlertIDs) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"alertIDs": a.Alerts,
	})
}

type IncidentURL struct {
	IncidentURL string `json:"string"`
}

func (u *IncidentURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"incidentURL": u.IncidentURL,
	})
}

type CustomFields struct {
	SentinelIncidentNumber *SentinelIncidentNumber `json:"sentinelIncidentNumber"`
	AlertIDs               *AlertIDs               `json:"alertIds"`
	IncidentURL            *IncidentURL            `json:"IncidentURL"`
}

func (f *CustomFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"numberStruct": f.SentinelIncidentNumber,
		"alertStruct":  f.AlertIDs,
		"URLStruct":    f.IncidentURL,
	})
}

type Object struct {
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	Owner             string        `json:"owner"`
	ResolutionStatus  string        `json:"resolutionStatus"`
	ResolutionSummary string        `json:"summary"`
	Severity          int           `json:"severity"`
	CustomFields      *CustomFields `json:"customFields"`
	TLP               int           `json:"tlp"`
	Source            string        `json:"source"`
}

func (o *Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"title":             o.Title,
		"description":       o.Description,
		"owner":             o.Owner,
		"resolutionStatus":  o.ResolutionStatus,
		"resolutionSummary": o.ResolutionSummary,
		"severity":          o.Severity,
		"customFields":      o.CustomFields,
		"tlp":               o.TLP,
		"source":            o.Source,
	})
}

type Capsule struct {
	Operation    string  `json:"operation"`
	ObjectType   string  `json:"objectType"`
	ObjectID     string  `json:"objectId"`
	Details      Details `json:"details"`
	Object       Object  `json:"object"`
	Organization string  `json:"organization"`
}

func NewCapsule(c *config.Conf) *Capsule {
	object := Capsule{
		Organization: c.Organization,
	}
	return &object
}
