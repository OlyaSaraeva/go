package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type indexPage struct {
	Title         string
	Subtitle      string
	FeaturedPosts []featuredPostData
	MostPosts     []mostPostData
}

type featuredPostData struct {
	Title          string `db:"title"`
	Subtitle       string `db:"subtitle"`
	Emblema        string `db:"Emblema"`
	EmblemaTitle   string `db:"EmblemaTitle"`
	Outt           string `db:"Outt"`
	BlockDirection string `db:"BlockDirection"`
	Author         string `db:"Author"`
	AuthorImg      string `db:"AuthorImg"`
	PublishDate    string `db:"PublishDate"`
}

type mostPostData struct {
	TitleMost    string `db:"title"`
	Text         string `db:"subtitle"`
	Background   string `db:"Background"`
	AuthorMost   string `db:"Author"`
	AuthorImg    string `db:"AuthorImg"`
	PublishDate  string `db:"PublishDate"`
	SizeSmall    string `db:"SizeSmall"`
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := featuredPosts(db)
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
			FeaturedPosts: posts,
			MostPosts: postsmost,
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
}

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