# TwitchHub Golang Backend

#### setup development environment:

1.  clone the repository:
    ```shell script
    git clone https://github.com/Bardiesel/twitchhub-go-backend.git
    ```
2.  install packages:
    ```shell script
    go install
    ```
3.  config `.env` file:
    ```shell script
    cp .env.example .env
    ```
4.  run the project:
    ```shell script
    go run main.go
    ```
    _you may add `-migrate` at the end of the command to migrate tables to database._

#### base structure for new services:

```
- root
    * controllers
        service_controller_name.go

    * models
        service_model_name.go

    * routes
        service_route_name.go

    * service
        service_name.go
```

> important: remember to add service model to `AutoMigrate` inside `main.go`.
