docs:
	swag init -g server/http.go --parseDependency
run-rest:
	go run main.go rest
fmt:
	go fmt ./... 
coverage:
	go test ./... -coverprofile cover.out
test:
	go test ./...
postgres:
	docker run --name postgis  -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgis/postgis
createdb:
	docker exec -it postgis createdb --username=postgres --owner=postgres gopher_net
startdb:
	docker start postgis