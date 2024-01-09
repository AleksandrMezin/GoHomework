package database

import (
	"testing"
	"time"
)

func TestSaveAndGetNews(t *testing.T) {
	news := News{
		Title:       "Test News",
		Description: "This is a test news",
		PubDate:     time.Now(),
		SourceURL:   "https://example.com",
	}

	err := SaveNews(news)
	if err != nil {
		t.Errorf("Error saving news: %v", err)
	}

	newsList, err := GetNews(1)
	if err != nil {
		t.Errorf("Error getting news: %v", err)
	}

	if len(newsList) != 1 {
		t.Errorf("Expected 1 news, got %d", len(newsList))
	}

	if newsList[0].Title != news.Title {
		t.Errorf("Expected title: %s, got: %s", news.Title, newsList[0].Title)
	}
}
