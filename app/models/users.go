package models

import (
	"log"
	"time"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

type Session struct {
	ID       int
	UUID     string
	Name     string
	Email    string
	UserID   int
	CreateAt time.Time
}

func (u *User) CreateUser() (err error) {
	// user = User{}
	cmd := `insert into users (
        uuid,
        name,
        email,
        password,
        created_at) values ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}

	/*スキャンすrる*/

	return err
}

func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `select id, uuid,name,email,password,created_at from users where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)
	return user, err
}

func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = $1, email = $2 where id = $3`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (u *User) DeleteUser() (err error) {

	cmd := `delete  from users where id = $1`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetUserByEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select id ,uuid,name,email,password,created_at from users where email = $1`
	err = Db.QueryRow(cmd, email).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)

	return user, err
}

//session作成
func (u *User) CreateSession() (session Session, err error) {
	session = Session{}
	cmd1 := `insert into sessions(
		uuid,
		name,
		email,
		user_id,
		created_at) values($1,$2,$3,$4,$5)`
	_, err = Db.Exec(cmd1, createUUID(), u.Name, u.Email, u.ID, time.Now())
	if err != nil {
		log.Println(err)
	}

	cmd2 := `select id, uuid, name,email, user_id, created_at from sessions where  uuid = $1 `

	err = Db.QueryRow(cmd2, u.UUID).Scan(

		&session.ID,
		&session.UUID,
		&session.Name,
		&session.Email,
		&session.UserID,
		&session.CreateAt)

	return session, err

}

func (sess *Session) CheckSession() (valid bool, err error) {

	cmd := `select id, uuid,name, email,user_id, created_at from sessions where uuid = $1 `
	err = Db.QueryRow(cmd, sess.UUID).Scan(
		&sess.ID,
		&sess.UUID,
		&sess.Name,
		&sess.Email,
		&sess.UserID,
		&sess.CreateAt)

	if err != nil {
		valid = false
		return
	}
	if sess.ID != 0 {
		valid = true
	}
	return valid, err
}

func (sess *Session) DeleteSessionByUUID() (err error) {
	cmd := `delete from sessions where uuid = $1`
	_, err = Db.Exec(cmd, sess.UUID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

func (sess *Session) GetUserBySession() (user User, err error) {
	user = User{}
	cmd := `select id, uuid,name,email,created_at FROM users where id = $1`
	err = Db.QueryRow(cmd, sess.UserID).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.CreatedAt)

	return user, err

}

func GetUserByEmailAndPassWord(email, password string) (user User, err error) {
	user = User{}
	cmd := `select id , uuid , name , email, password, created_at from users where email = $1  and password = $2`

	err = Db.QueryRow(cmd, email, password).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt)
	return user, err
}

func NameEdit(id int, name string) error {

	cmd := `update users set name = $1 where id = $2 `
	_, err = Db.Exec(cmd, name, id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

// func (sess *Session) NameEditSession(id int, name string) (err error) {

// 	cmd := `update sessions set name = $1  where id = $2`
// 	_, err = Db.Exec(cmd, name, id)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return err
// }

func NameEditSession(id int, name string) (err error) {

	cmd := `update sessions set name = $1  where user_id = $2`
	_, err = Db.Exec(cmd, name, id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func EmailEdit(id int, email string) error {

	cmd := `update users set email = $1 where id = $2 `
	_, err = Db.Exec(cmd, email, id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (sess *Session) EmailEditSession(id int, email string) (err error) {

	cmd := `update sessions set email = $1  where id = $2`
	_, err = Db.Exec(cmd, email, id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func CheckPassWord(id int, pass string) (user User, err error) {
	user = User{}
	cmd := `select id, password  from users where id = $1 and password = $2 `

	err = Db.QueryRow(cmd, id, pass).Scan(
		&user.ID,
		&user.PassWord,
	)

	return user, err

}

func PasswordEdit(id int, pass string) (err error) {
	cmd := `update users set password = $1 where id = $2 `
	_, err = Db.Exec(cmd, pass, id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

//Email存在チェック
func ExistEmail(email string) (user User, err error) {
	user = User{}
	cmd := `select  email  from users where email = $1 `

	err = Db.QueryRow(cmd, email).Scan(

		&user.Email,
	)

	return user, err
}
