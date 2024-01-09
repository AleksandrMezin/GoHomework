package aggregator

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig()
	if len(feeds) == 0 {
		t.Error("Expected feeds to be populated, got empty feeds")
	}

	if refresh == 0 {
		t.Error("Expected refresh interval to be set, got 0")
	}
}

func TestFetchNewsPeriodically(t *testing.T) {
	go fetchNewsPeriodically()
}
