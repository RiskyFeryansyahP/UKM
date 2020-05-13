package model

import (
	"github.com/confus1on/UKM/ent"
)

// InputAnnouncement hold value input announcement
type InputAnnouncement struct {
	Title 		string 		`json:"title"`
	Description	string 		`json:"description"`
	Time 		string 		`json:"time"`
}

// ResultAnnouncement result more announcement
type ResultAnnouncement struct {
	StatusCode	int					`json:"statuscode,omitempty"`
	Status		bool				`json:"status,omitempty"`
	Result		[]*ent.Announcement	`json:"result"`
}

// SingleResultAnnouncement result only single announcement
type SingleResultAnnouncement struct {
	StatusCode	int					`json:"statuscode,omitempty"`
	Status		bool				`json:"status,omitempty"`
	Result		*ent.Announcement	`json:"result"`
}
