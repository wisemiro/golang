build:
	docker build -t auth-go .

dev:
	docker run -p 8081:8081 -ti auth-go