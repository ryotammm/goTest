package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"goTest/app/config"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/lib/pq"
)

var Db *sql.DB

var err error

/*
const (
	tableNameUser            = "users"
	tableNameSession         = "sessions"
	tableNamePracticeContent = "practicecontents"
	tableNameTag             = "tags"
)
*/

func init() {

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	/*	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
			if err != nil {
				log.Fatalln(err)
			}
			//users table 作成
			cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		        id INTEGER PRIMARY KEY AUTOINCREMENT,
		        uuid STRING NOT NULL UNIQUE,
		        name STRING,
		        email STRING,
		        password STRING,
		        created_at DATETIME)`, tableNameUser)

			Db.Exec(cmdU)

			//sessions table 作成
			cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		        id INTEGER PRIMARY KEY AUTOINCREMENT,
		        uuid STRING NOT NULL UNIQUE,
				name STRING NOT NULL,
		        email STRING,
		        user_id INTEGER,
		        created_at DATETIME)`, tableNameSession)

			Db.Exec(cmdS)

			cmdP := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				user_id INTEGER,
				prefecture STRING NOT NULL,
				place STRING NOT NULL,
				strat_time STRING  NOT NULL,
				end_time STRING  NOT NULL,
				scale INTEGER ,
				tags STRING,
				describe TEXT,
				uuid STRING NOT NULL UNIQUE,
				created_at DATETIME)`, tableNamePracticeContent)

			Db.Exec(cmdP)

			//tags
			cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
					id INTEGER ,
					tag STRING
					)`, tableNameTag)
			Db.Exec(cmdT)



	*/
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext

}
