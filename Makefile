gorun:
	go run ./cmd/main.go

dynorestart:
    heroku dyno:restart --app authent-service
    heroku dyno:restart --app users-sservice
    heroku dyno:restart --app comment-sservice
    heroku dyno:restart --app book-sservice