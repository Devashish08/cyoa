# Choose Your Own Adventure (CYOA)

A web application that lets you create and serve "Choose Your Own Adventure" style stories. This project renders a story from a JSON file and presents it as an interactive web experience where users can make choices that affect the story's progression.

## Features

- Serves interactive story adventures via a web interface
- Customizable HTML templates for story presentation
- Configurable URL path handling
- Supports complex story arcs with multiple paths and endings
- Clean, responsive design with CSS styling

## Installation

### Prerequisites

- Go 1.16 or higher

### Steps

1. Clone the repository:
   ```
   git clone https://github.com/Devashish08/cyoa.git
   cd cyoa
   ```

2. Build the application:
   ```
   go build ./cmd/cyoaweb
   ```

## Usage

### Running the Web Server

Run the web application with default settings:

```
./cyoaweb
```

This will start the server on port 3001 and use the default `story.json` file.

### Command Line Options

The application accepts the following command-line flags:

- `-port`: The port to run the web server on (default: 3001)
- `-file`: The JSON file containing the story (default: "story.json")

Example with custom options:

```
./cyoaweb -port=8080 -file=custom_story.json
```

### Creating Your Own Stories

Stories are defined in JSON format. The structure should follow this pattern:

```json
{
  "start-arc": {
    "title": "Your Story Title",
    "story": [
      "First paragraph of the story.",
      "Second paragraph of the story."
    ],
    "options": [
      {
        "text": "Option text displayed to the user",
        "arc": "next-arc-key"
      },
      {
        "text": "Another option",
        "arc": "different-arc-key"
      }
    ]
  },
  "next-arc-key": {
    // ... next part of the story
  }
}
```

## Library Usage

You can also use this as a package in your Go applications:

```go
package main

import (
  "os"
  "net/http"
  
  "github.com/Devashish08/cyoa"
)

func main() {
  // Open the JSON story file
  f, err := os.Open("story.json")
  if err != nil {
    panic(err)
  }
  
  // Parse the story
  story, err := cyoa.JsonStory(f)
  if err != nil {
    panic(err)
  }
  
  // Create a handler and serve
  handler := cyoa.NewHandler(story)
  http.ListenAndServe(":8080", handler)
}
```

### Customization Options

The handler can be customized using functional options:

```go
handler := cyoa.NewHandler(
  story,
  cyoa.WithTemplate(customTemplate),
  cyoa.WithPathFunc(customPathFunction),
)
```

## Example

The included example story "The Little Blue Gopher" demonstrates various story paths and options. To try it:

1. Run the web server
2. Navigate to http://localhost:3001/story/intro in your browser
3. Follow the story and make choices to explore different paths

## License

[MIT License](LICENSE)
