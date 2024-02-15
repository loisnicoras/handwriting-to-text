
# Introduction
The primary purpose of the Handwriting Text Extractor is to streamline
and modernize the process of transcribing handwritten content. Whether
you have lecture notes, meeting minutes, or personal jottings, this
application enables users to convert their physical notes into editable
and searchable text with just a few clicks.

The Handwriting Text Extractor is an innovative application designed to
effortlessly convert handwritten text from images into digital, editable
content. Leveraging a powerful and user-friendly API developed by
experts.

## Index
- [Application Architecture](#application-architecture)
- [Text Extraction API](#text-extraction-api)
  - [OCR Functionality](#ocr-functionality)
  - [API Access](#api-access)
  - [Supported Image Formats](#supported-image-formats)  
  - [Multi-Language Support](#multi-language-support)
  - [Output Formats](#output-formats)
  - [Configuration and API Key](#configuration-and-api-key)
  - [Usage and Pricing](#usage-and-pricing)
  - [Documentation and Support](#documentation-and-support)
- [API endpoints](#api-endpoints)
- [Examples of API Requests and Responses](#examples-of-api-requests-and-responses)
  - [Request](#request)
  - [Response Body](#response-body)
- [API key for Google Cloud Platform (GCP)](#api-key-for-google-cloud-platform-gcp)
- [Costs](#costs)

# Application Architecture
![Application architecture](doc/media/architecture.png)

# Text Extraction API
- Google Cloud Handwriting Detection with Optical Character Recognition (OCR) is a service provided by Google Cloud Platform (GCP) that allows users to extract text from handwritten documents, such as scanned images or photos. Here's an overview of how it works:

### OCR Functionality:
-   The OCR functionality is designed to deliver high accuracy in text extraction, even from challenging handwritten documents. It employs state-of-the-art machine learning algorithms and large-scale training datasets to achieve robust performance. 

### API Access:
-   Users can access the Handwriting Detection OCR service programmatically through REST API endpoints provided by Google Cloud Platform. This allows for seamless integration into applications and workflows.

### Supported Image Formats:
-   Google Cloud Handwriting Detection OCR supports various image formats, including JPEG, PNG, and PDF. Users can upload scanned images or photos containing handwritten text for processing.

### Multi-Language Support:
-   The OCR functionality supports multiple languages and scripts, allowing it to recognize handwritten text in various languages worldwide. This includes languages with complex scripts like Chinese, Japanese, and Arabic.

### Output Formats: 
- The extracted text is returned in structured formats such as plain text strings or JSON responses. This facilitates easy integration with downstream applications and systems for further processing and analysis.

### Configuration and API Key:
-  Users need to sign up for an OCR.space account and obtain an API key to access the OCR service. The API key is usually used for authentication when making requests to the OCR API.

### Usage and Pricing:
-   Google Cloud Platform imposes usage limits on the number of OCR requests per month. The exact limits depend on the pricing plan chosen by the user.

### Documentation and Support:
-   The official documentation for Google Cloud Vision API, including Handwriting Detection OCR, is available on the Google Cloud website. It provides comprehensive guides, tutorials, reference documentation, and examples to help users understand and utilize the OCR functionality effectively.

# API key for Google Cloud Platform (GCP)
1. Go to the Google Cloud Console: Visit the Google Cloud Console website at https://console.cloud.google.com/.

2. Create a new project (if necessary): If you don't already have a project, create one by clicking on the project drop-down menu at the top of the page and selecting "New Project". Follow the prompts to create a new project.

3. Select your project: After creating or selecting your project, make sure it's selected in the project drop-down menu at the top of the page.

4. Navigate to the API & Services Credentials page: Click on the hamburger menu (☰) in the upper left corner, then navigate to "APIs & Services" > "Credentials".

5. Create credentials: On the "Credentials" page, click on the "Create credentials" dropdown and select "API key".

6. Copy your API key: Once the API key is created, it will be displayed on the screen. Copy the API key and use it in your application or script.

# Examples of API Requests and Responses

#### Request

```bash
# Set your API key
API_KEY="YOUR_API_KEY"

# Set the API endpoint
API_ENDPOINT="https://vision.googleapis.com/v1/images:annotate?key=${API_KEY}"

# Base64 encode the image file
BASE64_IMAGE=$(base64 -w 0 img.jpeg)

# Construct the JSON request payload
JSON_PAYLOAD=$(cat <<EOF
{
  "requests": [
    {
      "image": {
        "content": "${BASE64_IMAGE}"
      },
      "features": [
        {
          "type": "DOCUMENT_TEXT_DETECTION"
        }
      ]
    }
  ]
}
EOF
)

# Send the request using curl
curl -s -X POST -H "Content-Type: application/json" --data-binary "${JSON_PAYLOAD}" "${API_ENDPOINT}" 
```

#### Response Body
<details><summary markdown="span">Show body response</summary>

```json
{
  "responses": [
    {
      "textAnnotations": [
        {
          "locale": "en",
          "description": "happy",
          "boundingPoly": {
            "vertices": [
              {
                "x": 11,
                "y": 12
              },
              {
                "x": 51,
                "y": 12
              },
              {
                "x": 51,
                "y": 27
              },
              {
                "x": 11,
                "y": 27
              }
            ]
          }
        },
        {
          "description": "happy",
          "boundingPoly": {
            "vertices": [
              {
                "x": 11,
                "y": 12
              },
              {
                "x": 51,
                "y": 12
              },
              {
                "x": 51,
                "y": 27
              },
              {
                "x": 11,
                "y": 27
              }
            ]
          }
        }
      ],
      "fullTextAnnotation": {
        "pages": [
          {
            "property": {
              "detectedLanguages": [
                {
                  "languageCode": "en",
                  "confidence": 1
                }
              ]
            },
            "width": 68,
            "height": 35,
            "blocks": [
              {
                "boundingBox": {
                  "vertices": [
                    {
                      "x": 11,
                      "y": 12
                    },
                    {
                      "x": 51,
                      "y": 12
                    },
                    {
                      "x": 51,
                      "y": 27
                    },
                    {
                      "x": 11,
                      "y": 27
                    }
                  ]
                },
                "paragraphs": [
                  {
                    "boundingBox": {
                      "vertices": [
                        {
                          "x": 11,
                          "y": 12
                        },
                        {
                          "x": 51,
                          "y": 12
                        },
                        {
                          "x": 51,
                          "y": 27
                        },
                        {
                          "x": 11,
                          "y": 27
                        }
                      ]
                    },
                    "words": [
                      {
                        "property": {
                          "detectedLanguages": [
                            {
                              "languageCode": "en",
                              "confidence": 1
                            }
                          ]
                        },
                        "boundingBox": {
                          "vertices": [
                            {
                              "x": 11,
                              "y": 12
                            },
                            {
                              "x": 51,
                              "y": 12
                            },
                            {
                              "x": 51,
                              "y": 27
                            },
                            {
                              "x": 11,
                              "y": 27
                            }
                          ]
                        },
                        "symbols": [
                          {
                            "boundingBox": {
                              "vertices": [
                                {
                                  "x": 11,
                                  "y": 12
                                },
                                {
                                  "x": 19,
                                  "y": 12
                                },
                                {
                                  "x": 19,
                                  "y": 27
                                },
                                {
                                  "x": 11,
                                  "y": 27
                                }
                              ]
                            },
                            "text": "h",
                            "confidence": 0.98002976
                          },
                          {
                            "boundingBox": {
                              "vertices": [
                                {
                                  "x": 20,
                                  "y": 12
                                },
                                {
                                  "x": 28,
                                  "y": 12
                                },
                                {
                                  "x": 28,
                                  "y": 27
                                },
                                {
                                  "x": 20,
                                  "y": 27
                                }
                              ]
                            },
                            "text": "a",
                            "confidence": 0.98148566
                          },
                          {
                            "boundingBox": {
                              "vertices": [
                                {
                                  "x": 27,
                                  "y": 12
                                },
                                {
                                  "x": 35,
                                  "y": 12
                                },
                                {
                                  "x": 35,
                                  "y": 27
                                },
                                {
                                  "x": 27,
                                  "y": 27
                                }
                              ]
                            },
                            "text": "p",
                            "confidence": 0.97381717
                          },
                          {
                            "boundingBox": {
                              "vertices": [
                                {
                                  "x": 35,
                                  "y": 12
                                },
                                {
                                  "x": 43,
                                  "y": 12
                                },
                                {
                                  "x": 43,
                                  "y": 27
                                },
                                {
                                  "x": 35,
                                  "y": 27
                                }
                              ]
                            },
                            "text": "p",
                            "confidence": 0.9486064
                          },
                          {
                            "property": {
                              "detectedBreak": {
                                "type": "LINE_BREAK"
                              }
                            },
                            "boundingBox": {
                              "vertices": [
                                {
                                  "x": 41,
                                  "y": 12
                                },
                                {
                                  "x": 51,
                                  "y": 12
                                },
                                {
                                  "x": 51,
                                  "y": 27
                                },
                                {
                                  "x": 41,
                                  "y": 27
                                }
                              ]
                            },
                            "text": "y",
                            "confidence": 0.92424244
                          }
                        ],
                        "confidence": 0.9616363
                      }
                    ],
                    "confidence": 0.9616363
                  }
                ],
                "blockType": "TEXT",
                "confidence": 0.9616363
              }
            ],
            "confidence": 0.9616363
          }
        ],
        "text": "happy"
      }
    }
  ]
}
```

</details>
<br/>

# Costs
-   **Amazon EC2 Instances:**
   
    -   Our application will run on an EC2 instance of type t2.micro.
    -   The price for t2.micro instances is approximately \$0.0116/hour.
    -   If the application runs continuously for 24 hours a day for a month, the cost would be:
        -   \$0.0116/hour \* 24 hour/day \* 30 days = \$8.35
-   **Estimated Total Cost for One Month:**
    
    -   AWS Cost + API Cost = \$8.35 + \[API Specific Cost\]