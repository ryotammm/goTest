package controllers

import (
	"crypto/md5"
	"fmt"
	"goTest/app/models"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nfnt/resize"
)

type Profile struct {
	Name  string
	Email string
	//画像
}

const (
	imageExtension = ".png"
	imagePath      = "app/views/images/"
	width          = 512
	height         = 0
)

func index(w http.ResponseWriter, r *http.Request) {

	sess, err := session(w, r)

	pracs, err2 := models.StartPrefectures()
	if err2 != nil {
		log.Fatalln(err2)
	}
	pageData := pageData(sess.Name, pracs)
	if err != nil {

		GenerateHTML(w, pageData, "layout", "public_navbar", "index", "public_navbarMobile")
		return
	} else {

		GenerateHTML(w, pageData, "layout", "private_navbar", "index", "private_navbarMobile")
	}

}
func recruitment(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	switch r.Method {
	case http.MethodGet:

		pageData := pageData(sess.Name, nil)
		if err != nil {
			GenerateHTML(w, nil, "layout", "login", "public_navbar", "public_navbarMobile")
		} else {

			if err != nil {
				log.Fatalln(err)
			}

			GenerateHTML(w, pageData, "layout", "recruitment", "private_navbar", "private_navbarMobile")

		}
	case http.MethodPost:

		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		err = r.ParseMultipartForm(32 << 20)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		practicecontent := models.Practicecontent{

			Prefecture: r.PostFormValue("prefecture"),
			Place:      r.PostFormValue("place"),
			Strat_time: r.PostFormValue("start-time"),
			End_time:   r.PostFormValue("end-time"),
			Scale:      r.PostFormValue("scale"),
			Tags:       r.PostFormValue("tags"),
			Describe:   r.PostFormValue("describe"),
		}

		//記事登録
		uid, err := practicecontent.CreatePracticecontent(sess.UserID)
		if err != nil {
			log.Println(err)
		}

		err = r.ParseMultipartForm(32 << 20)
		if err != nil {

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, _, err := r.FormFile("file")

		if err != nil {

			file_def := imagePath + "e96755e8-cedb-11ec-8604-acde48001122" + imageExtension
			reader, ferr := os.Open(file_def)
			if ferr != nil {
				fmt.Println("ファイルの読み込みエラーです", ferr)
				os.Exit(1)
			}
			img, data, derr := image.Decode(reader)
			if derr != nil {
				fmt.Println("画像の変換エラーです", derr)
				os.Exit(1)
			} else {
				fmt.Println(data, "形式のデータを得ました")
			}
			defer reader.Close()
			resizedImg := resize.Resize(width, height, img, resize.NearestNeighbor)

			// 書き出すファイル名を指定します
			path := imagePath + uid.String() + imageExtension

			dst, err := os.Create(path)
			if err != nil {

				log.Printf("err %v", err)
			}
			defer dst.Close()

			// 画像のエンコード(書き込み)
			switch data {
			case "png":
				if err := png.Encode(dst, resizedImg); err != nil {
					log.Fatal(err)
				}
			case "jpeg", "jpg":
				opts := &jpeg.Options{Quality: 100}
				if err := jpeg.Encode(dst, resizedImg, opts); err != nil {
					log.Fatal(err)
				}
			default:
				if err := png.Encode(dst, resizedImg); err != nil {
					log.Fatal(err)
				}
			}

		} else {

			// 画像を読み込み
			img, data, err := image.Decode(file)
			if err != nil {
				log.Fatalln(err)
			}

			resizedImg := resize.Resize(width, height, img, resize.NearestNeighbor)

			// 書き出すファイル名を指定します
			path := imagePath + uid.String() + imageExtension

			dst, err := os.Create(path)
			if err != nil {

				log.Printf("err %v", err)
			}
			defer dst.Close()

			// 画像のエンコード(書き込み)
			switch data {
			case "png":
				if err := png.Encode(dst, resizedImg); err != nil {
					log.Fatal(err)
				}
			case "jpeg", "jpg":
				opts := &jpeg.Options{Quality: 100}
				if err := jpeg.Encode(dst, resizedImg, opts); err != nil {
					log.Fatal(err)
				}
			default:
				if err := png.Encode(dst, resizedImg); err != nil {
					log.Fatal(err)
				}
			}

		}

		http.Redirect(w, r, "/", http.StatusFound)

	}
}

func profile(w http.ResponseWriter, r *http.Request) {

	sess, _ := session(w, r)

	profile := Profile{
		Name:  sess.Name,
		Email: sess.Email,
	}

	pracs, err := models.GetPracticecontentByUserID(sess.UserID)

	profileData := map[string]interface{}{
		"profile": profile,
		"pracs":   pracs,
	}

	if err != nil {
		log.Fatalln(err)
	}

	GenerateHTML(w, profileData, "layout", "profile", "private_navbar", "private_navbarMobile")

}

func search(w http.ResponseWriter, r *http.Request) {

	tags := r.FormValue("tags")

	prefectures := r.FormValue("prefectures")

	var pracs []models.Practicecontent
	var err error

	if len(tags) > 0 && len(prefectures) > 0 {
		//両方入力
		pracs, err = models.SearchPrefecturesAndTagsX(prefectures, tags)
		if err != nil {
			fmt.Println("err")
		}

	} else if len(tags) > 0 && len(prefectures) == 0 {
		//検索フォームのみ

		pracs, err = models.SearchTagsX(tags)

		if err != nil {
			fmt.Println("err")
		}

	} else if len(tags) == 0 && len(prefectures) > 0 {
		//都道府県のみ
		pracs, err = models.SearchPrefecturesX(prefectures)

		if err != nil {
			fmt.Println("err")
		}

	} else {
		fmt.Println("入力なし")
	}

	sess, err := session(w, r)

	pageData := pageData(sess.Name, pracs)
	if pracs == nil {
		pageData["result"] = "検索結果なし"

	}

	if err != nil {
		GenerateHTML(w, pageData, "layout", "index", "public_navbar", "public_navbarMobile")
	} else {
		GenerateHTML(w, pageData, "layout", "index", "private_navbar", "private_navbarMobile")
	}

}

func pageData(name string, pracs []models.Practicecontent) (pageData map[string]interface{}) {

	pageData = map[string]interface{}{
		"name":  name,
		"pracs": pracs,
	}
	return pageData
}

//練習募集削除
func recruitmentDelete(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	//記事id
	id := r.FormValue("id")
	idI, _ := strconv.Atoi(id)
	err = models.Deleterecruitment(idI)
	if err != nil {
		log.Fatalln(err)
	}

	http.Redirect(w, r, "/profile", http.StatusFound)

}

func PassHash(rand string) string {
	// string型を[]byte型の変更して使う
	md5 := md5.Sum([]byte(rand))

	return fmt.Sprintf("%x", md5)
}
