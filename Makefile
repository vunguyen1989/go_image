APP_NAME=go_image
DOCKER_IMAGE=$(APP_NAME):latest

# chạy local
local:
	go run main.go

# # build binary
# build:
# 	go build -o $(APP_NAME)

# # chạy binary
# run-bin:
# 	./$(APP_NAME)

# # build docker image
# docker:
# 	docker build -t $(DOCKER_IMAGE) .

# # chạy docker container
# run:
# 	docker run -p 8080:8080 $(DOCKER_IMAGE)

# dọn dẹp
clean:
	rm -f $(APP_NAME)