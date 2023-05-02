APP_NAME := chatgpt
API_KEY ?=

.PHONY: build clean

build:
	go build -ldflags="-s -w -X main.apiKey=$(API_KEY)" -o $(APP_NAME)

clean:
	rm -f $(APP_NAME)

run:
	./$(APP_NAME) -api_key $(API_KEY) $(ARGS)
