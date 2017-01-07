package controllers

import (
	"kyunooh/gonte/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type NoteController struct {
	beego.Controller
}


// @Title Create
// @Description create object
// @Param body body models.Note true "The note content"
// @Success 200 {string} models.Note.Id
// @Failure 403 body is empty
// @router / [post]
func (n *NoteController) Post() {
  var note models.Note
	json.Unmarshal(n.Ctx.Input.RequestBody, &note)
	nid := models.AddNote(note)
	n.Data["json"] = map[string]string{"nid": nid}
	n.ServeJSON()
}

// @Title Get
// @Description find note by nid
// @Param nid path string true "The key for static block"
// @Success 200 {string} models.Note
// @Failure 403 :nid is empty
// @router /:nid [get]
func (n *NoteController) Get() {
  nid := n.GetString(":nid")
	if nid != "" {
		note, err := models.GetNote(nid)
		if err != nil {
			n.Data["json"] = err.Error()
		} else {
			n.Data["json"] = note
		}
	}
	n.ServeJSON()
}

// @Title Update
// @Description update the note
// @Param nid path string true " The nid you want to update"
// @Param body body models.Note true "body for noteu content"
// @Success 200 {object} models.Note
// @failure 403 :nid is not int
// @router /:nid [put]
func (n *NoteController) Put() {
	nid := n.GetString(":nid")
	if nid != "" {
		var note models.Note
		json.Unmarshal(n.Ctx.Input.RequestBody, &note)
		nn, err := models.UpdateNote(nid, &note)
		if err != nil {
			n.Data["json"] = err.Error()
		} else {
			n.Data["json"] = nn
		}
	}
	n.ServeJSON()
}

// @Title Delete
// @Description delete the note
// @Param nid path string true "the nid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 nid is empty
// @router /:nid [delete]
func (n * NoteController) Delete() {
	nid := n.GetString(":nid")
	models.DeleteNote(nid)
	n.Data["json"] = "delete success!"
	n.ServeJSON()
}


