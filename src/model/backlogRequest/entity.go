package backlogRequest

import (
	"encoding/json"
	"time"
)

type BacklogRequest struct {
	Created time.Time `json:"created"`
	Project struct {
		Archived          bool   `json:"archived"`
		Projectkey        string `json:"projectKey"`
		Name              string `json:"name"`
		Chartenabled      bool   `json:"chartEnabled"`
		ID                int    `json:"id"`
		Subtaskingenabled bool   `json:"subtaskingEnabled"`
	} `json:"project"`
	ID      int `json:"id"`
	Type    int `json:"type"`
	Content struct {
		Summary      string        `json:"summary"`
		KeyID        int           `json:"key_id"`
		Customfields []interface{} `json:"customFields"`
		Duedate      string        `json:"dueDate"`
		Description  string        `json:"description"`
		Priority     struct {
			Name string      `json:"name"`
			ID   interface{} `json:"id"`
		} `json:"priority"`
		Resolution struct {
			Name string      `json:"name"`
			ID   interface{} `json:"id"`
		} `json:"resolution"`
		Actualhours interface{} `json:"actualHours"`
		Issuetype   struct {
			Color        string      `json:"color"`
			Name         string      `json:"name"`
			Displayorder interface{} `json:"displayOrder"`
			ID           int         `json:"id"`
			Projectid    interface{} `json:"projectId"`
		} `json:"issueType"`
		Milestone []struct {
			Archived       bool        `json:"archived"`
			Releaseduedate string      `json:"releaseDueDate"`
			Name           string      `json:"name"`
			Displayorder   interface{} `json:"displayOrder"`
			Description    string      `json:"description"`
			ID             interface{} `json:"id"`
			Projectid      interface{} `json:"projectId"`
			Startdate      string      `json:"startDate"`
		} `json:"milestone"`
		Versions []struct {
			Archived       bool        `json:"archived"`
			Releaseduedate string      `json:"releaseDueDate"`
			Name           string      `json:"name"`
			Displayorder   interface{} `json:"displayOrder"`
			Description    string      `json:"description"`
			ID             interface{} `json:"id"`
			Projectid      interface{} `json:"projectId"`
			Startdate      string      `json:"startDate"`
		} `json:"versions"`
		Parentissueid  interface{} `json:"parentIssueId"`
		Estimatedhours interface{} `json:"estimatedHours"`
		ID             int         `json:"id"`
		Assignee       interface{} `json:"assignee"`
		Category       []struct {
			Name         string      `json:"name"`
			Displayorder interface{} `json:"displayOrder"`
			ID           interface{} `json:"id"`
		} `json:"category"`
		Startdate string `json:"startDate"`
		Status    struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"status"`
		Changes []struct {
			Field    string `json:"field"`
			OldValue string `json:"old_value,omitempty"`
			Type     string `json:"type,omitempty"`
			NewValue string `json:"new_value"`
		} `json:"changes"`
		Comment struct {
			ID      int    `json:"id"`
			Content string `json:"content"`
		} `json:"comment"`
		Link []struct {
			KeyID string `json:"key_id"`
			ID    string `json:"id"`
			Title string `json:"title"`
		} `json:"link"`
		TxID         string      `json:"tx_id"`
		Name         string      `json:"name"`
		Content      string      `json:"content"`
		Diff         interface{} `json:"diff"`
		Version      int         `json:"version"`
		Size         int         `json:"size"`
		Dir          string      `json:"dir"`
		Rev          int         `json:"rev"`
		Ref          string      `json:"ref"`
		RevisionType string      `json:"revision_type"`
		Revisions    []struct {
			Rev  string `json:"rev"`
			Link struct {
				Text string `json:"text"`
				URL  string `json:"url"`
			} `json:"link"`
			Comment string `json:"comment"`
		} `json:"revisions"`
		ChangeType    string `json:"change_type"`
		RevisionCount int    `json:"revision_count"`
		Repository    struct {
			Name        string `json:"name"`
			ID          int    `json:"id"`
			Description string `json:"description"`
		} `json:"repository"`
		Users []struct {
			Nulabaccount struct {
				Nulabid  string `json:"nulabId"`
				Name     string `json:"name"`
				Uniqueid string `json:"uniqueId"`
			} `json:"nulabAccount"`
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"users"`
		Number        int         `json:"number"`
		Issue         interface{} `json:"issue"`
		Branch        string      `json:"branch"`
		Base          string      `json:"base"`
		ReferenceDate string      `json:"reference_date"`
		StartDate     string      `json:"start_date"`
	} `json:"content"`
	Notifications []interface{} `json:"notifications"`
	Createduser   struct {
		Nulabaccount struct {
			Nulabid  string `json:"nulabId"`
			Name     string `json:"name"`
			Uniqueid string `json:"uniqueId"`
		} `json:"nulabAccount"`
		Name        string      `json:"name"`
		Mailaddress interface{} `json:"mailAddress"`
		ID          json.Number `json:"id"`
		Roletype    int         `json:"roleType"`
		Lang        string      `json:"lang"`
		Userid      interface{} `json:"userId"`
	} `json:"createdUser"`
}
