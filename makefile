dev:
		go run main.go

test-article:
		richgo test ./repository

console:
		docker-compose up -d