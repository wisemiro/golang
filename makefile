build:
	docker build -t auth-go .

dev:
	docker run -p 8081:8081 -ti auth-go

proto:
	protoc --go_out=. --go_opt=paths=source_relative user.proto 


# protoc --go_out=. *.proto
# Set-ExecutionPolicy Unrestricted