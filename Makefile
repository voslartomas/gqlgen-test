generate:
	go generate ./... && go run models/model_tags.go

dev:
	gin --port 8080 run server.go

run:
	go run server.go