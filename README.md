# github.com/cloudinary/analysis-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/cloudinary/analysis-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/cloudinary/analysis-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasy.com/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasy.com/docs/advanced-setup/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README

<!-- Start Summary [summary] -->
## Summary

Analyze API (Beta): Use the Analyze API to analyze any external asset and return details based on the type of analysis requested.

Currently supports the following analysis options:
  * [AI Vision - Tagging](https://cloudinary.com/documentation/cloudinary_ai_vision_addon#tagging_mode)
  * [AI Vision - Moderation](https://cloudinary.com/documentation/cloudinary_ai_vision_addon#moderation_mode)
  * [AI Vision - General](https://cloudinary.com/documentation/cloudinary_ai_vision_addon#general_mode)
  * [Captioning](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning)
  * [Cld Fashion](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Cld Text](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Coco](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Google tagging](https://cloudinary.com/documentation/google_auto_tagging_addon)
  * [Human Anatomy](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Lvis](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Shop Classifier](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)
  * [Unidet](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#supported_content_aware_detection_models)

  **Notes**:
  * The Analyze API is currently in development and is available as a Public Beta, which means we value your feedback, so please feel free to [share any thoughts with us](https://support.cloudinary.com/hc/en-us/requests/new).
  * The analysis options require an active subscription to the relevant add-on. Learn more about [registering for add-ons](https://cloudinary.com/documentation/cloudinary_add_ons#registering_for_add_ons).

  The API supports both Basic Authentication using your Cloudinary API Key and API Secret (which can be found on the Dashboard page of your [Cloudinary Console](https://console.cloudinary.com/pm)) or OAuth2 ([Contact support](https://support.cloudinary.com/hc/en-us/requests/new) for more information regarding OAuth).
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/cloudinary/analysis-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Analysis](docs/sdks/analysissdk/README.md)

* [AnalyzeAiVisionGeneral](docs/sdks/analysissdk/README.md#analyzeaivisiongeneral) - Analyze - AI Vision General
* [AnalyzeAiVisionModeration](docs/sdks/analysissdk/README.md#analyzeaivisionmoderation) - Analyze - AI Vision Moderation
* [AnalyzeAiVisionTagging](docs/sdks/analysissdk/README.md#analyzeaivisiontagging) - Analyze - AI Vision Tagging
* [AnalyzeCaptioning](docs/sdks/analysissdk/README.md#analyzecaptioning) - Analyze - Captioning
* [AnalyzeCldFashion](docs/sdks/analysissdk/README.md#analyzecldfashion) - Analyze - Cld-Fashion
* [AnalyzeCldText](docs/sdks/analysissdk/README.md#analyzecldtext) - Analyze - Cld-Text
* [AnalyzeCoco](docs/sdks/analysissdk/README.md#analyzecoco) - Analyze - Coco
* [AnalyzeGoogleTagging](docs/sdks/analysissdk/README.md#analyzegoogletagging) - Analyze - Google Tagging
* [AnalyzeHumanAnatomy](docs/sdks/analysissdk/README.md#analyzehumananatomy) - Analyze - Human Anatomy
* [AnalyzeLvis](docs/sdks/analysissdk/README.md#analyzelvis) - Analyze - Lvis
* [AnalyzeShopClassifier](docs/sdks/analysissdk/README.md#analyzeshopclassifier) - Analyze - Shop Classifier
* [AnalyzeUnidet](docs/sdks/analysissdk/README.md#analyzeunidet) - Analyze - Unidet
* [~~AnalyzeURI~~](docs/sdks/analysissdk/README.md#analyzeuri) - Analyze an asset :warning: **Deprecated**


</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	analysisgo "github.com/cloudinary/analysis-go"
	"github.com/cloudinary/analysis-go/models/components"
	"github.com/cloudinary/analysis-go/retry"
	"log"
	"models/operations"
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
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.AnalyzeAIVisionGeneralResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	analysisgo "github.com/cloudinary/analysis-go"
	"github.com/cloudinary/analysis-go/models/components"
	"github.com/cloudinary/analysis-go/retry"
	"log"
	"os"
)

func main() {
	s := analysisgo.New(
		analysisgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `AnalyzeAiVisionGeneral` function may return the following errors:

| Error Type              | Status Code                                                                                             | Content Type     |
| ----------------------- | ------------------------------------------------------------------------------------------------------- | ---------------- |
| sdkerrors.ErrorResponse | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json |
| sdkerrors.Timeout       | 408, 504                                                                                                | application/json |
| sdkerrors.SDKError      | 4XX, 5XX                                                                                                | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	analysisgo "github.com/cloudinary/analysis-go"
	"github.com/cloudinary/analysis-go/models/components"
	"github.com/cloudinary/analysis-go/models/sdkerrors"
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

		var e *sdkerrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.Timeout
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
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
		analysisgo.WithServerURL("https://api.cloudinary.com/v2"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name        | Type   | Scheme       | Environment Variable  |
| ----------- | ------ | ------------ | --------------------- |
| `BasicAuth` | http   | HTTP Basic   | `ANALYSIS_BASIC_AUTH` |
| `OAuth2`    | oauth2 | OAuth2 token | `ANALYSIS_O_AUTH2`    |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/cloudinary/analysis-go&utm_campaign=go)
