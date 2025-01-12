package main

import (
	"context"
	"fmt"
	"log"
	"os"

	speech "cloud.google.com/go/speech/apiv1"
	speechpb "cloud.google.com/go/speech/apiv1/speechpb"
	"google.golang.org/api/option"
)

func SpeechToText() {
	ctx := context.Background()

	// Configure the client with your Google Cloud Platform credentials.
	client, err := speech.NewClient(ctx, option.WithCredentialsFile("/path/to/your/credentials.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Read the file.
	audioData, err := os.ReadFile("./assets/cual-es-la-fecha-cumple.mp3")
	if err != nil {
		log.Fatalf("Failed to read audio file: %v", err)
	}

	// Configure the request
	req := &speechpb.RecognizeRequest{
		Config: &speechpb.RecognitionConfig{
			Encoding:                   speechpb.RecognitionConfig_MP3,
			SampleRateHertz:            16000,
			LanguageCode:               "es-ES",
			AudioChannelCount:          1,
			EnableAutomaticPunctuation: true,
		},
		Audio: &speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{
				Content: audioData,
			},
		},
	}

	// Send the request.
	resp, err := client.Recognize(ctx, req)
	if err != nil {
		log.Fatalf("Failed to recognize: %v", err)
	}

	// Process the response.
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Transcript: %s\n", alt.Transcript)
			fmt.Printf("Confidence: %f\n", alt.Confidence)

			// We can add a translation feature here.
			translatedText := translateText(alt.Transcript, "es")
			fmt.Printf("Translated Text: %s\n", translatedText)
		}
	}
}

// invented function to translate text
func translateText(text, targetLang string) string {

	// Implement the translation logic here, for example, using Google Translate API.
	// This is just a placeholder for a marker position.
	return "Texto traducido al español."
}
