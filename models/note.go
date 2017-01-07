package models

import (
	"errors"
	"strconv"
	"time"
)

var (
	NoteList map[string]*Note
)

type Note struct {
  Id       string
	Author   User
	Title    string
	Contents string
}

func init() {
  NoteList = make(map[string]*Note)
}


func AddNote(n Note) string {
	n.Id = "note_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	NoteList[n.Id] = &n
	return n.Id
}

func GetNote(nid string) (n *Note, err error) {
  if n, ok := NoteList[nid]; ok {
    return n, nil
	}
	return nil, errors.New("Note not exists")
}

func getAllNotes() map[string]*Note {
  return NoteList
}

func UpdateNote(nid string, nn *Note) (a *Note, err error) {
	if n, ok := NoteList[nid]; ok {
		if nn.Title != "" {
			n.Title = nn.Title
		}
		if n.Contents != "" {
			n.Contents = nn.Contents
		}
		return n, nil
	}
	return nil, errors.New("Note Not Exist")
}

func DeleteNote(nid string) {
  delete(NoteList, nid)
}
