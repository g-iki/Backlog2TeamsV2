package controller

import (
	"Backlog2Teams/src/model/backlog"
	response "Backlog2Teams/src/model/response"
	"log"

	"net/http"
	"strings"
	"time"

	utils "Backlog2Teams/src/utils"

	"github.com/gin-gonic/gin"
)

type BacklogController struct{}

func (bc BacklogController) Index(c *gin.Context) {
	log.Printf("backlog/Index called")
	res := new(response.Response)

	b := new(backlog.Backlog)
	list := b.ReadAll()
	// con := make([]backlog.Connections, 0)
	// for i := 0; i < len(list); i++ {
	// 	list[i].SetConnections(con)
	// }
	res.SetData(list)
	res.SetMessage("Success")

	c.String(http.StatusOK, utils.JsonToString(res))
	return
}

func (bc BacklogController) Create(c *gin.Context) {
	log.Printf("backlog/Create called")
	res := new(response.Response)

	var req backlog.Backlog
	c.BindJSON(&req)

	b := new(backlog.Backlog)

	if req.GetID() == "" {
		res.SetMessage("ID cannot be null")
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	b.SetID(req.GetID())
	b.SetOwner(req.GetOwner())
	b.SetCreateDate(time.Now())
	b.SetUpdateDate(time.Now())

	for _, v := range req.Connections {
		con := new(backlog.Connections)
		con.SetTeamsWebhookURL(v.GetTeamsWebhookURL())
		con.SetNotes(v.GetNotes())
		b.AddConnections(*con)
	}

	err := b.Create()

	if err != nil {
		if strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			res.SetMessage("Duplicate Error")
			c.String(http.StatusBadRequest, utils.JsonToString(res))
		}
		return
	}

	res.SetData(b)
	res.SetMessage("Success, new Item inserted")
	c.String(http.StatusOK, utils.JsonToString(res))
	return
}

func (bc BacklogController) Show(c *gin.Context) {
	log.Printf("backlog/Show called")
	res := new(response.Response)

	id := c.Params.ByName("id")

	if id == "" {
		res.SetMessage("ID cannot be null")
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	var req backlog.Backlog
	c.BindJSON(&req)

	b := new(backlog.Backlog)
	backlog, err := b.GetItem(id)

	if err != nil {
		res.SetMessage(err.Error())
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	res.SetData(backlog)
	res.SetMessage("Success")
	c.String(http.StatusOK, utils.JsonToString(res))

	return
}

func (bc BacklogController) Update(c *gin.Context) {
	log.Printf("backlog/Update called")
	res := new(response.Response)

	var req backlog.Backlog
	c.BindJSON(&req)
	id := c.Params.ByName("id")

	if id == "" {
		res.SetMessage("ID cannot be null")
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	b := new(backlog.Backlog)
	reged, err := b.GetItem(id)

	if err != nil {
		res.SetMessage(err.Error())
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	for _, v := range req.Connections {
		isNew := true
		for i, j := range reged.Connections {
			if v.GetTeamsWebhookURL() == j.GetTeamsWebhookURL() {
				reged.Connections[i].SetNotes(v.GetNotes())
				isNew = false
				break
			}
		}
		if isNew {
			con := new(backlog.Connections)
			con.SetTeamsWebhookURL(v.GetTeamsWebhookURL())
			con.SetNotes(v.GetNotes())
			reged.AddConnections(*con)
		}
	}
	if req.GetOwner() != "" {
		reged.SetOwner(req.GetOwner())
	}
	reged.SetUpdateDate(time.Now())
	updated, err := reged.Update(id)
	if err != nil {
		res.SetMessage(err.Error())
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}
	res.SetData(updated)
	res.SetMessage("Success, record Update(Connection Update)")
	c.String(http.StatusOK, utils.JsonToString(res))

	return
}

func (bc BacklogController) Remove(c *gin.Context) {
	log.Printf("backlog/Remove called")
	res := new(response.Response)

	var req backlog.Backlog
	c.BindJSON(&req)
	id := c.Params.ByName("id")

	if id == "" {
		res.SetMessage("ID cannot be null")
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	b := new(backlog.Backlog)
	upd := new(backlog.Backlog)
	reged, err := b.GetItem(id)

	upd.SetID(b.GetID())
	upd.SetOwner(b.GetOwner())
	upd.SetCreateDate(b.GetCreateDate())
	upd.SetUpdateDate(time.Now())

	if err != nil {
		res.SetMessage(err.Error())
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	for _, v := range reged.Connections {
		isDel := false
		for _, j := range req.Connections {
			if v.TeamsWebhookURL == j.GetTeamsWebhookURL() {
				isDel = true
				break
			}
		}
		if !isDel {
			con := new(backlog.Connections)
			con.SetTeamsWebhookURL(v.GetTeamsWebhookURL())
			con.SetNotes(v.GetNotes())
			upd.AddConnections(*con)
		}
	}

	updated, err := upd.Update(id)
	if err != nil {
		res.SetMessage(err.Error())
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}

	res.SetData(updated)
	res.SetMessage("Success, record Update(Connection Remove)")
	c.String(http.StatusOK, utils.JsonToString(res))

	return
}

func (bc BacklogController) Delete(c *gin.Context) {
	log.Printf("backlog/Delete called")
	res := new(response.Response)

	id := c.Params.ByName("id")
	b := new(backlog.Backlog)

	backlog, err := b.Delete(id)

	if err != nil {
		res.SetMessage(err.Error())
		c.String(http.StatusBadRequest, utils.JsonToString(res))
		return
	}
	res.SetData(backlog)
	res.SetMessage("Success, record Delete")
	c.String(http.StatusOK, utils.JsonToString(res))

	return
}
