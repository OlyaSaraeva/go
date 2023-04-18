package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title         string
	Subtitle      string
	//FeaturedPosts []featuredPostData
	Posts     []*postListData
}

/* type featuredPostData struct {
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	Emblema        string `db:"Emblema"`
	EmblemaTitle   string `db:"EmblemaTitle"`
	Outt           string `db:"Outt"`
	BlockDirection string `db:"BlockDirection"`
	Author         string `db:"Author"`
	AuthorImg      string `db:"AuthorImg"`
	PublishDate    string `db:"PublishDate"`
} */

type postData struct {
	TitleMost    string `db:"title"`
	Text         string `db:"subtitle"`
	Background   string `db:"Background"`
	AuthorMost   string `db:"Author"`
	//AuthorImg    string `db:"AuthorImg"`
	//PublishDate  string `db:"PublishDate"`
	//SizeSmall    string `db:"SizeSmall"`
	//Content      string `db:"content"`
	//Title          string `db:"title"`
	
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
	TitleMost    string `db:"title"`
	Text         string `db:"subtitle"`
	PostURL        string // URL ордера, на который мы будем переходить для конкретного поста
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postsData, err := posts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		/* posts, err := featuredPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		postsmost, err := mostPosts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		} */


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
			//MostPosts: postsmost,
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
			BlockDirection
		FROM `  + "`post`" + 
		`WHERE featured = 1`
		
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

// Получает информацию о конкретном ордере из базы данных
func postByID(db *sqlx.DB, postID int) (postData, error) {
	const query = `
		SELECT
			title,
			content
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

/* 
func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			PublishDate,
			Author,
			AuthorImg,
			Outt,
			EmblemaTitle,
			Emblema,
			BlockDirection
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []featuredPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func mostPosts(db *sqlx.DB) ([]mostPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			PublishDate,
			Author,
			AuthorImg,
			Background,
			SizeSmall
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var postsmost []mostPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&postsmost, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return postsmost, nil
} */

/* func featuredPosts(db *sqlx.DB) ([]featuredPostData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			publish_date
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []featuredPostData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil

	/* return []featuredPostData{
		{
		   Title:       "The Road Ahead",
           Subtitle:    "The road ahead might be paved - it might not be.",
		   Emblema:     "",
		   EmblemaTitle : "",
		   Outt : "block__text-block_out",
		   BlockDirection : "block_left",
           Author:      "Mat Vogels",
           AuthorImg:   "/static/img/mat_vogels.png",
           PublishDate: "September 25, 2015",
		},
		{
			Title:    "From Top Down",
			Subtitle: "Once a year, go someplace you’ve never been before.",
			Emblema: "block__emblema",
			EmblemaTitle : "Adventute",
			Outt : "",
			BlockDirection : "block_right",
			Author:      "William Wong",
            AuthorImg:   "/static/img/william_wong.png",
            PublishDate: "9/25/2015",
		},
	}
} */

/* func mostPosts() []mostPostData { */
	/* return []mostPostData{
		{
			TitleMost:       "Still Standing Tall",
           Text:    "Life begins at the end of your comfort zone.",
		   Background : "/static/img/syill_standin.jpg",
           AuthorMost:      "William Wong",
           AuthorImg:   "/static/img/william_wong.png",
           PublishDate: "9/25/2015",
		   SizeSmall: "card__text-block_size_small",
		},
		{
			TitleMost:    "Sunny Side Up",
			Text: "No place is ever as bad as they tell you it’s going to be.",
			Background : "/static/img/sunny_side.jpg",
			AuthorMost:      "Mat Vogels",
            AuthorImg:   "/static/img/mat_vogels.png",
            PublishDate: "9/25/2015",
			SizeSmall: "",
		},
		{
			TitleMost:    "Water Falls",
			Text: "We travel not to escape life, but for life not to escape us.",
			Background : "/static/img/water_falls.jpg",
			AuthorMost:      "Mat Vogelsg",
            AuthorImg:   "/static/img/mat_vogels.png",
            PublishDate: "9/25/2015",
			SizeSmall: "",
		},
		{
			TitleMost:    "Through the Mist",
			Text: "Travel makes you see what a tiny place you occupy in the world.",
			Background : "/static/img/through.jpg",
			AuthorMost:      "William Wong",
            AuthorImg:   "/static/img/william_wong.png",
            PublishDate: "9/25/2015",
			SizeSmall: "",
		},
		{
			TitleMost:    "Awaken Early",
			Text: "Not all those who wander are lost.",
			Background : "/static/img/awaken.jpg",
			AuthorMost:      "Mat Vogels",
            AuthorImg:   "/static/img/mat_vogels.png",
            PublishDate: "9/25/2015",
			SizeSmall: "card__text-block_size_small",
		},
		{
			TitleMost:    "Try it Always",
			Text: "The world is a book, and those who do not travel read only one page.",
			Background : "/static/img/try_it_always.jpg",
			AuthorMost:      "Mat Vogels",
            AuthorImg:   "/static/img/mat_vogels.png",
            PublishDate: "9/25/2015",
			SizeSmall: "",
		},
	}
 }*/