# Duolingo Domination Discord bot

## Table of contents

- [Architecture](#architecture)
- [Build and deploy the app](#build-and-deploy-the-app)

## Architecture

The bot is written in Go. It uses mostly standard libraries but also the [Canvas
library from Taco de Wolff](https://github.com/tdewolff/canvas).

[main.go](main.go) is the entry point of the application. It registers the HTTP
handlers then start the HTTP server on port 8080.

The HTTP handlers are defined in [server/handlers.go](server/handlers.go).
They do not care about most of the logic, they just parse the request, call various
functions from other packages, then respond with the result.

The [discord](discord) package contains the logic to interact with the Discord API.
[discord/security.go](discord/security.go) is here to make request's signature
verification easier.

The [canvas](canvas) package is used to draw the leaderboard once the data
have been fetched from the Duolingo API. After finishing the drawing, the
image is saved to a file named "leaderboard.png".

## Build and deploy the app

Run `go build -o executable`.

ℹ️ If you're are build for an other OS or arch,
make sure to define GOOS and GOARCH. For example, to build for linux on amd64,
run `CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o executable`.

For the application to run properly, the asset folder needs to be alongside
the executable, so don't forget it if you're moving the executable to a container
or other folder.




