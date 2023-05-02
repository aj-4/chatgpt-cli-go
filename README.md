# ChatGPT Golang CLI

* Starter CLI implementation in Golang for Chat GPT
* Currently uses GPT-3.5-Turbo & works the same as web with plain text output

## Prerequisites
1. To use ChatGPT, you'll need an OpenAI API key. You can sign up for an API key at https://beta.openai.com/signup/.
2. You will need Golang installed https://go.dev/doc/install

## Installation
Clone the repository:
```sh
git clone https://github.com/your-username/chatgpt.git
```
Change into the directory:
```sh
cd chatgpt
```
Download the required dependencies:
```sh
go mod download
```

Build the executable:
```sh
make build API_KEY=<your-api-key>
```

## Usage
To start the ChatGPT app, run the following command:

```sh
make run
```

Avoid committing your binary with
```sh
make clean
```


To exit the app, press ctrl+c or q.
