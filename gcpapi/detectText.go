package gcpapi

import (
	"bytes"
	"context"
	"log"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

func TextInPicture(inputImge []byte) string {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	ioImage := bytes.NewReader(inputImge)
	image, _ := vision.NewImageFromReader(ioImage)

	response, err := client.DetectDocumentText(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	newText := strings.ReplaceAll(response.Text, "\n", " ")
	return newText
}
