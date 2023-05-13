package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"encoding/base64"
	"io"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title         string
	Subtitle      string
	Posts    	  []*postListData
	Featured 	  []*postfeacheListData
}

type postData struct {
	Title        string `db:"title"`
	Subtitle     string `db:"subtitle"`
	Background   string `db:"Background"`
	AuthorMost   string `db:"Author"`
	Content      string `db:"content"`
}

type postListData struct {
	PostID         string `db:"post_id"`
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	BlockDirection string `db:"BlockDirection"`
	Emblema        string `db:"Emblema"`
	EmblemaTitle   string `db:"EmblemaTitle"`
	Outt           string `db:"Outt"`
	Author         string `db:"Author"`
	AuthorImg      string `db:"AuthorImg"`
	PublishDate    string `db:"PublishDate"`
	Background     string `db:"Background"`
	SizeSmall      string `db:"SizeSmall"`
	PostURL        string // URL ордера, на который мы будем переходить для конкретного поста
}

type postfeacheListData struct {
	PostID         string `db:"post_id"`
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	BlockDirection string `db:"BlockDirection"`
	Emblema        string `db:"Emblema"`
	EmblemaTitle   string `db:"EmblemaTitle"`
	Outt           string `db:"Outt"`
	Author         string `db:"Author"`
	AuthorImg      string `db:"AuthorImg"`
	PublishDate    string `db:"PublishDate"`
	Background     string `db:"Background"`
	SizeSmall      string `db:"SizeSmall"`
	PostURL        string // URL ордера, на который мы будем переходить для конкретного поста
}

type adminPage struct {
	Title         string
}

type createPostRequest struct {
	Title   string `json:"title"`  
	Content string `json:"content"`
 }
 
func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postsData, err := posts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		postsfeacheData, err := postsfeache(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
			return                                      // Не забываем завершить выполнение ф-ии
		}

		data := indexPage{
			Title:         "Let's do it together.",
			Subtitle: "We travel the world in search of stories. Come along for the ride.",
			Posts: postsData,
			Featured: postsfeacheData,
		}

		err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}

		log.Println("Request completed successfully")
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"] // Получаем postID в виде строки из параметров урла

		postID, err := strconv.Atoi(postIDStr) // Конвертируем строку postID в число
		if err != nil {
			http.Error(w, "Invalid post id", 403)
			log.Println(err)
			return
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "Post not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		err = ts.Execute(w, post)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

// Возвращаем не просто []postListData, а []*postListData - так у нас получится подставить PostURL в структуре
func posts(db *sqlx.DB) ([]*postListData, error) {
	const query = `
		SELECT
			post_id,
			title,
			subtitle,
			PublishDate,
			Author,
			AuthorImg,
			Background,
			Outt,
			EmblemaTitle,
			Emblema,
			BlockDirection,
			SizeSmall
		FROM `  + "`post`" + 
		`WHERE featured = 0`
		
	// Такое объединение строк делается только для таблицы order, т.к. это зарезерированное слово в SQL, наряду с SELECT, поэтому его нужно заключить в ``

	var posts []*postListData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                  // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range posts {
		post.PostURL = "/post/" + post.PostID // Формируем исходя из ID post'a в базе
	}

	fmt.Println(posts)

	return posts, nil
}

func postsfeache(db *sqlx.DB) ([]*postfeacheListData, error) {
	const query = `
		SELECT
			post_id,
			title,
			subtitle,
			PublishDate,
			Author,
			AuthorImg,
			Background,
			Outt,
			EmblemaTitle,
			Emblema,
			BlockDirection,
			SizeSmall
		FROM `  + "`post`" + 
		`WHERE featured = 1`
		
	// Такое объединение строк делается только для таблицы order, т.к. это зарезерированное слово в SQL, наряду с SELECT, поэтому его нужно заключить в ``

	var posts []*postfeacheListData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                  // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range posts {
		post.PostURL = "/post/" + post.PostID // Формируем исходя из ID post'a в базе
	}

	fmt.Println(posts)

	return posts, nil
}

// Получает информацию о конкретном ордере из базы данных
func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			content,
			subtitle,
			Background
		FROM
			` + "`post`" +
		`WHERE
		post_id = ?
	`
	// В SQL-запросе добавились параметры, как в шаблоне. ? означает параметр, который мы передаем в запрос ниже

	var post postData

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&post, query, postID)
	if err != nil {
		return postData{}, err
	}

	return post, nil
}

func login(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/login.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := adminPage{
		Title:         "Admin",
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("pages/admin.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	data := adminPage{
		Title:         "Admin",
	}

	err = ts.Execute(w, data) // Заставляем шаблонизатор вывести шаблон в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
	reqData, err := io.ReadAll(r.Body) // Прочитали тело запроса с reqData в виде массива байт
       if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошбики в консоль
		return 
       }

	   imgAuthor, err := base64.StdEncoding.DecodeString(reqData.authorImg)
	   img, err := base64.StdEncoding.DecodeString(reqData.PostImage)
	   file, err := os.Create("static/img/" + reqData.PostImageName)

	   _, err = file.Write(img) // Записываем контент картинки в файл
	   _, err = file.Write(imgAuthor)

       var req createPostRequest  // Заранее объявили переменную  createOrderRequest

       err =  json.Unmarshal(reqData, &req) // Отдали reqData и req на парсинг библиотеке json
       if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
       }
	}
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	const query = `
		INSERT INTO
			post
		(
			title,
			content
		)
		VALUES
		(
			?,
			?
		)
	`
 
	_, err := db.Exec(query, req.Title, req.Content) // Сами данные передаются через аргументы к ф-ии Exec
	return err
 }
 

