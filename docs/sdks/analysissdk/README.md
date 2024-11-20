# AnalysisSDK
(*Analysis*)

## Overview

### Available Operations

* [AnalyzeAiVisionGeneral](#analyzeaivisiongeneral) - Analyze - AI Vision General
* [AnalyzeAiVisionModeration](#analyzeaivisionmoderation) - Analyze - AI Vision Moderation
* [AnalyzeAiVisionTagging](#analyzeaivisiontagging) - Analyze - AI Vision Tagging
* [AnalyzeCaptioning](#analyzecaptioning) - Analyze - Captioning
* [AnalyzeCldFashion](#analyzecldfashion) - Analyze - Cld-Fashion
* [AnalyzeCldText](#analyzecldtext) - Analyze - Cld-Text
* [AnalyzeCoco](#analyzecoco) - Analyze - Coco
* [AnalyzeGoogleTagging](#analyzegoogletagging) - Analyze - Google Tagging
* [AnalyzeHumanAnatomy](#analyzehumananatomy) - Analyze - Human Anatomy
* [AnalyzeLvis](#analyzelvis) - Analyze - Lvis
* [AnalyzeShopClassifier](#analyzeshopclassifier) - Analyze - Shop Classifier
* [AnalyzeUnidet](#analyzeunidet) - Analyze - Unidet
* [~~AnalyzeURI~~](#analyzeuri) - Analyze an asset :warning: **Deprecated**

## AnalyzeAiVisionGeneral

The General mode serves a wide array of applications by providing detailed answers to diverse questions about an image. Users can inquire about any aspect of an image, such as identifying objects, understanding scenes, or interpreting text within the image.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `cloudName`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | The name of your Cloudinary cloud                                                                    | your-cloud-name                                                                                      |
| `analyzeAIVisionGeneralRequest`                                                                      | [components.AnalyzeAIVisionGeneralRequest](../../models/components/analyzeaivisiongeneralrequest.md) | :heavy_check_mark:                                                                                   | A JSON object containing request parameters                                                          |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.AnalyzeAiVisionGeneralResponse](../../models/operations/analyzeaivisiongeneralresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeAiVisionModeration

The Moderation mode accepts multiple questions about an image, to which the response provides concise answers of "yes," "no," or "unknown." This functionality allows for a nuanced evaluation of whether the image adheres to specific content policies, creative specs, or aesthetic criteria.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeAiVisionModeration(ctx, "your-cloud-name", components.AnalyzeAIVisionModerationRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
        RejectionQuestions: []string{
            "Does this image contain any violent activity?",
            "Is there any nudity in the image?",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeAIVisionModerationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |                                                                                                            |
| `cloudName`                                                                                                | *string*                                                                                                   | :heavy_check_mark:                                                                                         | The name of your Cloudinary cloud                                                                          | your-cloud-name                                                                                            |
| `analyzeAIVisionModerationRequest`                                                                         | [components.AnalyzeAIVisionModerationRequest](../../models/components/analyzeaivisionmoderationrequest.md) | :heavy_check_mark:                                                                                         | A JSON object containing request parameters                                                                |                                                                                                            |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |                                                                                                            |

### Response

**[*operations.AnalyzeAiVisionModerationResponse](../../models/operations/analyzeaivisionmoderationresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeAiVisionTagging

The Tagging mode accepts a list of tag names along with their corresponding descriptions. If the image matches the description, which may encompass various elements, it will be appropriately tagged. This approach enables customers to align with their own brand taxonomy, offering a dynamic, flexible, and open method for image classification.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeAiVisionTagging(ctx, "your-cloud-name", components.AnalyzeAIVisionTaggingRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
        TagDefinitions: []components.TagDefinitions{
            components.TagDefinitions{
                Name: "cigarettes",
                Description: "Does the image contain a pack of cigarettes?",
            },
            components.TagDefinitions{
                Name: "furniture",
                Description: "Does the image contain any tables, chairs, couches or sofas?",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeAIVisionTaggingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |                                                                                                      |
| `cloudName`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | The name of your Cloudinary cloud                                                                    | your-cloud-name                                                                                      |
| `analyzeAIVisionTaggingRequest`                                                                      | [components.AnalyzeAIVisionTaggingRequest](../../models/components/analyzeaivisiontaggingrequest.md) | :heavy_check_mark:                                                                                   | A JSON object containing request parameters                                                          |                                                                                                      |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |                                                                                                      |

### Response

**[*operations.AnalyzeAiVisionTaggingResponse](../../models/operations/analyzeaivisiontaggingresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeCaptioning

Provides a caption for an image.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeCaptioning(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCaptioningResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeCaptioningResponse](../../models/operations/analyzecaptioningresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeCldFashion

Analyze an image using the [Cld-Fashion](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's fashion model is specifically dedicated to items of clothing. The response includes attributes of the clothing identified, for example whether the garment contains pockets, its material and the fastenings used.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeCldFashion(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCldFashionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeCldFashionResponse](../../models/operations/analyzecldfashionresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeCldText

Analyze an image using the [Cld-Text](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's text model tells you if your image includes text, and where it's located. Used with image tagging, you can then search for images that contain blocks of text. Used with object-aware cropping, you can choose to keep only the text part, or specify a crop that avoids the text.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeCldText(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCldTextResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeCldTextResponse](../../models/operations/analyzecldtextresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeCoco

Analyze an image using the [Coco](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. The [Common Objects in Context](https://cocodataset.org/) model contains just 80 common objects.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeCoco(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeCocoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeCocoResponse](../../models/operations/analyzecocoresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeGoogleTagging

Provides tags for an image using Google's tagging service.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeGoogleTagging(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeGoogleTaggingResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeGoogleTaggingResponse](../../models/operations/analyzegoogletaggingresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeHumanAnatomy

Analyze an image using the [Human Anatomy](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's human anatomy model identifies parts of the human body in an image. It works best when the majority of a human body is detected in the image.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeHumanAnatomy(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeHumanAnatomyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeHumanAnatomyResponse](../../models/operations/analyzehumananatomyresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeLvis

Analyze an image using the [Lvis](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. The [Large Vocabulary Instance Segmentation](https://www.lvisdataset.org/) model contains thousands of general objects.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeLvis(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeLvisResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeLvisResponse](../../models/operations/analyzelvisresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeShopClassifier

Analyze an image using the [Shop Classifier](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. Cloudinary's shop classifier model detects if the image is a product image taken in a studio, or if it's a natural image.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeShopClassifier(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeShopClassifierResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeShopClassifierResponse](../../models/operations/analyzeshopclassifierresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## AnalyzeUnidet

Analyze an image using the [Unidet](https://cloudinary.com/documentation/cloudinary_ai_content_analysis_addon#ai_based_image_captioning) content-aware detection model. The [UniDet](https://github.com/xingyizhou/UniDet) model is a unified model, combining a number of object models, including [Objects365](https://www.objects365.org/overview.html), which focuses on diverse objects in the wild.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeUnidet(ctx, "your-cloud-name", components.BaseAnalyzeRequest{
        Source: components.CreateSourceURI(
            components.URI{
                URI: "https://res.cloudinary.com/demo/image/upload/sample.jpg",
            },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeUnidetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `cloudName`                                                                    | *string*                                                                       | :heavy_check_mark:                                                             | The name of your Cloudinary cloud                                              | your-cloud-name                                                                |
| `baseAnalyzeRequest`                                                           | [components.BaseAnalyzeRequest](../../models/components/baseanalyzerequest.md) | :heavy_check_mark:                                                             | A JSON object containing request parameters                                    |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.AnalyzeUnidetResponse](../../models/operations/analyzeunidetresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |

## ~~AnalyzeURI~~

Analyzes an asset with the requested analysis type.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

```go
package main

import(
	"os"
	"github.com/cloudinary/analysis-go/models/components"
	analysisgo "github.com/cloudinary/analysis-go"
	"context"
	"log"
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
    res, err := s.Analysis.AnalyzeURI(ctx, "your-cloud-name", components.AnalyzeURIRequest{
        URI: analysisgo.String("https://res.cloudinary.com/demo/image/upload/sample.jpg"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AnalyzeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `cloudName`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | The name of your Cloudinary cloud                                            | your-cloud-name                                                              |
| `analyzeURIRequest`                                                          | [components.AnalyzeURIRequest](../../models/components/analyzeurirequest.md) | :heavy_check_mark:                                                           | A JSON object containing request parameters                                  |                                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.AnalyzeURIResponse](../../models/operations/analyzeuriresponse.md), error**

### Errors

| Error Type                                                                                              | Status Code                                                                                             | Content Type                                                                                            |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| sdkerrors.ErrorResponse                                                                                 | 400, 401, 403, 404, 407, 413, 414, 415, 422, 429, 431, 500, 501, 502, 503, 505, 506, 507, 508, 510, 511 | application/json                                                                                        |
| sdkerrors.Timeout                                                                                       | 408, 504                                                                                                | application/json                                                                                        |
| sdkerrors.SDKError                                                                                      | 4XX, 5XX                                                                                                | \*/\*                                                                                                   |