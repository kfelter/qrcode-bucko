# qrcode-bucko

A simple, fast QR code generator web application built with Go.

## Live Demo

**Production URL**: [https://qrcode.kfelter.com/](https://qrcode.kfelter.com/)

## Features

- Generate QR codes from any URL
- Clean, modern UI with responsive design
- Direct links to generated QR code images
- Simple to deploy and use

## How It Works

1. Enter any URL you want to convert to a QR code
2. The application generates a QR code image
3. View and download the QR code image
4. Share the QR code link with others

## Technology Stack

- Backend: Go
- QR Code Generation: [github.com/skip2/go-qrcode](https://github.com/skip2/go-qrcode)
- Frontend: HTML, CSS (with Poppins font from Google Fonts)

## Development Setup

### Prerequisites

- Go 1.22.0 or higher

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/kfelter/qrcode-bucko.git
   cd qrcode-bucko
```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
    ```bash
    go run cmd/server/main.go
    ```
4. Open your browser and navigate to `http://localhost:8888`


### Environment Variables
PORT: Server port (default: 8888)
BASE_URL: Base URL for the application (default: http://localhost:8888)
Deployment
You can deploy this application to any platform that supports Go applications. Set the appropriate environment variables for your production environment.

### Author
Kyle Felter