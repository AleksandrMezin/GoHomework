package server

import (
	"News/pkg/database"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// News представляет структуру новости с PubDateStr
type News struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PubDate     time.Time `json:"pubDate"`
	SourceURL   string    `json:"sourceURL"`
	PubDateStr  string    `json:"pubDateStr"`
}

// getNewsHandler обрабатывает запрос на получение новостей
func getNewsHandler(w http.ResponseWriter, r *http.Request) {
	countParam := r.URL.Query().Get("count")
	count, err := strconv.Atoi(countParam)
	if err != nil {
		http.Error(w, "Invalid count parameter", http.StatusBadRequest)
		return
	}

	newsList, err := database.GetNews(count)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting news: %v", err), http.StatusInternalServerError)
		return
	}

	//Преобразование времени в строковый формат
	//News/pkg/database
	var formattedNewsList []News
	for i := range newsList {
		formattedNews := News{
			Title:       newsList[i].Title,
			Description: newsList[i].Description,
			PubDate:     newsList[i].PubDate,
			SourceURL:   newsList[i].SourceURL,
			PubDateStr:  newsList[i].PubDate.Format("2006-01-02 15:04:05"),
		}
		formattedNewsList = append(formattedNewsList, formattedNews)
	}

	//Преобразование в JSON и отправка клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formattedNewsList)
}

// StartServer запускает веб-сервер
func StartServer() {
	http.HandleFunc("/api/news", getNewsHandler)
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8080", nil)
}
