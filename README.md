# Building a Chatbot with Go Fiber and OpenAI

This project demonstrates how to build an interactive chatbot using Go Fiber (Go web framework) and OpenAI's API. The chatbot can engage in conversations and provide responses using OpenAI's powerful language models.

## Features

- Interactive web-based chat interface
- Integration with OpenAI's API
- Real-time response generation
- Go Fiber backend server
- Simple and clean user interface

## Prerequisites

- Go 1.23.6 or greater
- OpenAI API key

## Installation

1. Clone the repository:
```bash
git clone https://github.com/CodeSignal-Learn/course_building-a-chatbot-with-go-and-fiber
cd course_building-a-chatbot-with-go-and-fiber
```

2. Install the required dependencies:
```bash
go mod init main
go mod tidy
```

## Configuration

Before running the application, you need to set up your OpenAI API key. You can do this by exporting it as an environment variable:

For macOS/Linux:
```bash
export OPENAI_API_KEY='your-api-key-here'
```

For Windows (Command Prompt):
```cmd
set OPENAI_API_KEY=your-api-key-here
```

For Windows (PowerShell):
```powershell
$env:OPENAI_API_KEY='your-api-key-here'
```

## Running the Application

1. Start the Go Fiber development server:
```bash
go run ./app
```

2. Open your web browser and navigate to:
```
http://localhost:3000
```

## Usage

1. Once the application is running, you'll see a chat interface in your browser
2. Type your message in the input field
3. Press Enter or click the Send button to submit your message
4. The chatbot will process your input and provide a response

## Project Structure

```
.
├── README.md
└── app/
    ├── static/            # Static files (CSS)
    ├── views/             # HTML templates
    ├── data/              # Data storage
    ├── chat.go            # Chat data management
    ├── chat_controller.go # Route controller
    ├── chat_service.go    # Business logic service
    └── main.go            # Main Go application
```