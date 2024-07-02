oms_app:
	echo "Starting OMS Server"
	cd db/ &&  sqlc generate && cd ../app
	go env -w GOPRIVATE=github.com/sharmajsr
	go mod tidy
	go build -o ./bin/oms ./app/ 