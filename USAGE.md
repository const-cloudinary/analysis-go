<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	analysisgo "github.com/cloudinary/analysis-go"
	"github.com/cloudinary/analysis-go/models/components"
	"log"
	"os"
)

func main() {
	s := analysisgo.New(
		analysisgo.WithSecurity(components.Security{
			BasicAuth: &components.SchemeBasicAuth{
				Username: os.Getenv(""),
				Password: os.Getenv(""),
			},
		}),
	)

	ctx := context.Background()
	res, err := s.Analysis.AnalyzeAiVisionGeneral(ctx, "your-cloud-name", components.AnalyzeAIVisionGeneralRequest{
		Source: components.CreateSourceURI(
			components.URI{
				URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
			},
		),
		Prompts: []string{
			"Describe this image in detail",
			"Does this image contain an insect?",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AnalyzeAIVisionGeneralResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->