package main

import {
	"net/http"
	"html/template"
	"log"
}

type indexPageData struct {
	Title string
	Subtitle string
}


func index(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error()) // Используем стандартный логгер для вывода ошибки в консоль
		return // Не забываем завершить выполнение ф-ии
	}
 
   // Подготовим данные для шаблона
   data := indexPageData {
	Title: "Blog for traveling",
	Subtitle: "My best blog"
   }
 
	err = ts.Execute(w, nil) // Запускаем шаблонизатор для вывода шаблона в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	} 
 }
 