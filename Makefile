run: 
	go run main.go serve

docker-compose:
	docker-compose up -d postgres

migrate:
	go run main.go migrate

import:
	go run main.go import
