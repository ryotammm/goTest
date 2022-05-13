package models

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type Practicecontent struct {
	ID         int
	UserID     int
	Prefecture string
	Place      string
	Strat_time string
	End_time   string
	Scale      string
	Tags       string
	UUID       string
	Describe   string
	CreatedAt  time.Time
}

func (p *Practicecontent) CreatePracticecontent(userID int) (Practicecontent_uuid uuid.UUID, err error) {
	Practicecontent_uuid = createUUID()
	cmd := `insert into practicecontents(
	user_id,
	prefecture,
	place,
	strat_time,
	end_time,
	scale,
	tags,
	describe,
	uuid,
	created_at
	) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`

	_, err = Db.Exec(cmd,
		userID,
		p.Prefecture,
		p.Place,
		p.Strat_time,
		p.End_time,
		p.Scale,
		p.Tags,
		p.Describe,
		Practicecontent_uuid,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	return Practicecontent_uuid, err
}

//

func GetPracticecontent() (pracs []Practicecontent, err error) {

	cmd := `select id, user_id, prefecture, place,strat_time,end_time,scale,tags,describe,uuid,created_at FROM practicecontents`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(
			&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Tags,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)

	}
	rows.Close()
	return pracs, err
}

func GetPracticecontentByUUID(uid uuid.UUID) (prac Practicecontent) {
	prac = Practicecontent{}

	cmd := `select id, user_id, prefecture, place,strat_time,end_time,scale,tags,describe,uuid,created_at FROM practicecontents where  uuid = $1 `
	err = Db.QueryRow(cmd, uid).Scan(

		&prac.ID,
		&prac.UserID,
		&prac.Prefecture,
		&prac.Place,
		&prac.Strat_time,
		&prac.End_time,
		&prac.Scale,
		&prac.Tags,
		&prac.UUID,
		&prac.Describe,
		&prac.CreatedAt)
	return prac
}

//idで一つの記事を取得する
func GetPracticecontentByID(id int) (prac Practicecontent) {
	prac = Practicecontent{}

	cmd := `select id, user_id, prefecture, place,strat_time,end_time,scale,tags,describe,uuid,created_at FROM practicecontents where  id = $1 `
	err = Db.QueryRow(cmd, id).Scan(

		&prac.ID,
		&prac.UserID,
		&prac.Prefecture,
		&prac.Place,
		&prac.Strat_time,
		&prac.End_time,
		&prac.Scale,
		&prac.Tags,

		&prac.Describe,
		&prac.UUID,
		&prac.CreatedAt)
	return prac
}

func SearchPrefectures(pref string) (pracs []Practicecontent, err error) {
	// _
	cmd := `select id, user_id,prefecture, place, strat_time,end_time,scale, tags,describe,uuid, created_at FROM practicecontents where prefecture = $1 `
	rows, err := Db.Query(cmd, pref)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Tags,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)
	}
	return pracs, err
}
func SearchPrefecturesAndTags(pref string, tags []string) (pracs []Practicecontent, err error) {

	cmd := `select id, user_id,prefecture, place, strat_time,end_time,scale,describe, uuid, created_at FROM practicecontents where place = $1 `
	rows, err := Db.Query(cmd, pref)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)
	}
	return pracs, err
}

//都道府県のみ
func SearchPrefecturesX(pref string) (pracs []Practicecontent, err error) {

	cmd := `select id, user_id,prefecture, place, strat_time,end_time,scale, tags,describe, uuid, created_at FROM practicecontents where prefecture = $1 `
	rows, err := Db.Query(cmd, pref)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Tags,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)
	}
	return pracs, err
}

//検索フォームのみ
func SearchTagsX(tags string) (pracs []Practicecontent, err error) {

	cmd := `select id, user_id,prefecture, place, strat_time,end_time,scale,
	 tags,describe, uuid, created_at FROM practicecontents where tags like $1 or prefecture  like $2 or   place  like $3 ; `

	like := "%"

	tagsS1 := fmt.Sprintf("%s%s%s", like, tags, like)
	tagsS2 := fmt.Sprintf("%s%s%s", like, tags, like)
	tagsS3 := fmt.Sprintf("%s%s%s", like, tags, like)

	rows, err := Db.Query(cmd, tagsS1, tagsS2, tagsS3)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Tags,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)
	}
	return pracs, err

}

//検索フォーム && 都道府県
func SearchPrefecturesAndTagsX(pref string, tags string) (pracs []Practicecontent, err error) {

	cmd := `select id, user_id,prefecture, place, strat_time,end_time,scale,
	tags,describe, uuid,created_at FROM practicecontents where prefecture = $1 and  (tags like $2  or  place  like $3 ); `

	like := "%"

	tagsS1 := fmt.Sprintf("%s%s%s", like, tags, like)
	tagsS2 := fmt.Sprintf("%s%s%s", like, tags, like)

	rows, err := Db.Query(cmd, pref, tagsS1, tagsS2)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Tags,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)
	}
	return pracs, err

}

//Start 画面
func StartPrefectures() (pracs []Practicecontent, err error) {
	cmd := `select id, user_id,prefecture, place, strat_time,end_time,scale,
	tags,describe,uuid,created_at FROM practicecontents order by  created_at DESC LIMIT 50; `

	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Tags,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)
	}
	return pracs, err
}

//ユーザの投稿取得 複数
func GetPracticecontentByUserID(id int) (pracs []Practicecontent, err error) {
	cmd := `select id, user_id, prefecture, place,strat_time,end_time,scale,tags,describe,uuid,created_at FROM practicecontents where user_id = $1 order by created_at DESC `
	rows, err := Db.Query(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var prac Practicecontent
		err = rows.Scan(
			&prac.ID,
			&prac.UserID,
			&prac.Prefecture,
			&prac.Place,
			&prac.Strat_time,
			&prac.End_time,
			&prac.Scale,
			&prac.Tags,
			&prac.Describe,
			&prac.UUID,
			&prac.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		pracs = append(pracs, prac)

	}
	rows.Close()
	return pracs, err

}

func Deleterecruitment(id int) error {
	cmd := `delete from practicecontents where id = $1`

	_, err = Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func DeleterecruitmentByUUID(uuid string) error {
	cmd := `delete from practicecontents where uuid = $1`

	_, err = Db.Exec(cmd, uuid)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
