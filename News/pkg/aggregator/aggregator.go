package aggregator

import (
	"News/pkg/database"
	"github.com/mmcdole/gofeed"
	"log"
	"time"
)

// fetchNews выполняет обход всех RSS-лент и сохраняет новости в базу данных
func fetchNews() {
	fp := gofeed.NewParser()
	for _, feedURL := range feeds {
		feed, err := fp.ParseURL(feedURL)
		if err != nil {
			log.Printf("Error fetching feed %s: %v", feedURL, err)
			continue
		}

		for _, item := range feed.Items {
			pubDate := time.Now()
			if item.PublishedParsed != nil {
				pubDate = *item.PublishedParsed
			}
			//News/pkg/database
			news := database.News{
				Title:       item.Title,
				Description: item.Description,
				PubDate:     pubDate,
				SourceURL:   item.Link,
			}

			//Сохранение новости в базе данных
			database.SaveNews(news)
		}
	}
}

// InitAggregator инициализирует агрегатор с использованием указанного пути к конфигурационному файлу
func InitAggregator(path string) {
	configPath = path
	LoadConfig()

	//Запуск обхода RSS в отдельной горутине
	go fetchNewsPeriodically()
}
func fetchNewsPeriodically() {
	for {
		fetchNews()
		time.Sleep(time.Duration(refresh) * time.Minute)
	}
}
