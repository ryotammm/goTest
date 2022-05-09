package models

import (
	"log"

	"github.com/google/uuid"
)

type Tags struct {
	ID  uuid.UUID
	Tag string
}

func GetTags(tagWord []string) (tagsP []Tags, err error) {

	for _, v := range tagWord {
		cmd := `select id, tag from tags where tag = $1`
		rows, err := Db.Query(cmd, v)
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			var tag Tags
			err = rows.Scan(&tag.ID,
				&tag.Tag)
			if err != nil {
				log.Fatalln(err)
			}
			tagsP = append(tagsP, tag)
		}
		rows.Close()
	}
	return tagsP, err
}

func CreateTags(uid uuid.UUID, tags []string) (err error) {

	cmd := `insert into tags(
	id,
	tag)values($1,$2)`

	for _, tag := range tags {
		_, err = Db.Exec(cmd, uid, tag)

		if err != nil {
			log.Fatalln(err)
		}

	}
	return err
}
