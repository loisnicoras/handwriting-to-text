
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
- [Costs](#costs)

# Application Architecture
![Application architecture](doc/media/architecture.png)

# Text Extraction API
- Google Cloud Handwriting Detection with Optical Character Recognition (OCR) is a service provided by Google Cloud Platform (GCP) that allows users to extract text from handwritten documents, such as scanned images or photos. Here's an overview of how it works:

#### OCR Functionality:
-   The OCR functionality is designed to deliver high accuracy in text extraction, even from challenging handwritten documents. It employs state-of-the-art machine learning algorithms and large-scale training datasets to achieve robust performance. 

#### API Access:
-   Users can access the Handwriting Detection OCR service programmatically through REST API endpoints provided by Google Cloud Platform. This allows for seamless integration into applications and workflows.

#### Supported Image Formats:
-   Google Cloud Handwriting Detection OCR supports various image formats, including JPEG, PNG, and PDF. Users can upload scanned images or photos containing handwritten text for processing.

#### Multi-Language Support:
-   The OCR functionality supports multiple languages and scripts, allowing it to recognize handwritten text in various languages worldwide. This includes languages with complex scripts like Chinese, Japanese, and Arabic.

#### Output Formats: 
- The extracted text is returned in structured formats such as plain text strings or JSON responses. This facilitates easy integration with downstream applications and systems for further processing and analysis.

#### Configuration and API Key:
-  Users need to sign up for an OCR.space account and obtain an API key to access the OCR service. The API key is usually used for authentication when making requests to the OCR API.

#### Usage and Pricing:
-   Google Cloud Platform imposes usage limits on the number of OCR requests per month. The exact limits depend on the pricing plan chosen by the user.

#### Documentation and Support:
-   The official documentation for Google Cloud Vision API, including Handwriting Detection OCR, is available on the Google Cloud website. It provides comprehensive guides, tutorials, reference documentation, and examples to help users understand and utilize the OCR functionality effectively.

# API endpoints 
-   Documentarea tuturor endpoint-urilor API pentru extragerea textului.
-   Parametri de intrare și formatul așteptat pentru cererile API.

# Examples of API Requests and Responses

- Request


- Response



# Costs
-   **Amazon EC2 Instances:**
   
    -   Our application will run on an EC2 instance of type t2.micro.
    -   The price for t2.micro instances is approximately \$0.0116/hour.
    -   If the application runs continuously for 24 hours a day for a month, the cost would be:
        -   \$0.0116/hour \* 24 hour/day \* 30 days = \$8.35
-   **Estimated Total Cost for One Month:**
    
    -   AWS Cost + API Cost = \$8.35 + \[API Specific Cost\]