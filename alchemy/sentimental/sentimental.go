package main

import (
	"fmt"
	"strings"
)

// Simple sentiment analysis function
func analyzeSentiment(text string) string {
	positiveWords := []string{"good", "great", "excellent"}
	negativeWords := []string{"bad", "terrible", "awful"}

	// Convert text to lowercase and split into words
	words := strings.Split(strings.ToLower(text), " ")

	// Check for positive and negative words
	positiveCount := 0
	negativeCount := 0
	for _, word := range words {
		for _, positiveWord := range positiveWords {
			if strings.Contains(word, positiveWord) {
				positiveCount++
			}
		}
		for _, negativeWord := range negativeWords {
			if strings.Contains(word, negativeWord) {
				negativeCount++
			}
		}
	}

	// Determine sentiment based on counts
	if positiveCount > negativeCount {
		return "Positive"
	} else if negativeCount > positiveCount {
		return "Negative"
	} else {
		return "Neutral"
	}
}

func main() {
	text := "I love this product, it's great! but the service is terrible."
	sentiment := analyzeSentiment(text)
	fmt.Println("Sentiment:", sentiment)
}
