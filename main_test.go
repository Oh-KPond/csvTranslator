package main

import "testing"

func Test_getLanguage(t *testing.T) {
	t.Run("should return 'German' for German File", func(t *testing.T) {
		lang := getLanguage("German Word Frequency toBe Translated - Sheet1.csv")

		if lang != "German" {
			t.Error("Did not return the correct language")
		}
	})

	t.Run("should return 'French' for French File", func(t *testing.T) {
		lang := getLanguage("French Word Frequency toBe Translated - Sheet1.csv")

		if lang != "French" {
			t.Error("Did not return the correct language")
		}
	})
}

func Test_getMergeKey(t *testing.T) {
	mergeKey := getMergeKey()

	if mergeKey["English"]["block"] != "block" {
		t.Errorf("This is the wrong key: %v", mergeKey)
	}

	if mergeKey["German"]["Block"] != "block" {
		t.Errorf("This is the wrong key: %v", mergeKey)
	}
}
