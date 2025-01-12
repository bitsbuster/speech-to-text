# Speech-to-Text

This repository contains a Go application that records audio and converts it to text using Google Cloud Speech-to-Text API.

## Prerequisites

- Go 1.23.4 or later
- Google Cloud Platform account with Speech-to-Text API enabled
- Google Cloud credentials JSON file

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/bitsbuster/speech-to-text.git
    cd speech-to-text
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

## Usage

1. Set up your Google Cloud credentials:

    Replace the path to your credentials JSON file in speechtotext.go:

    ```go
    client, err := speech.NewClient(ctx, option.WithCredentialsFile("/path/to/your/credentials.json"))
    ```

2. Record audio:

    Run the application and press 'r' to start recording and 'space' to stop:

    ```sh
    go run main.go
    ```

    The recorded audio will be saved as `output.wav`.

3. Convert speech to text:

    Uncomment the SpeechToText function call in main.go and provide the path to your audio file:

    ```go
    func main() {
        // ReadKeyboard()
        // RecordAudio()
        SpeechToText()
    }
    ```

    Run the application again to convert the audio to text:

    ```sh
    go run main.go
    ```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](https://github.com/bitsbuster/speech-to-text/blob/main/LICENSE) file for details.