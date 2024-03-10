# Duolingo Domination Discord bot

## Table of contents

- [Architecture](#architecture)
- [Build and deploy the app](#build-and-deploy-the-app)

## Architecture

The bot is written in Go. It uses mostly standard libraries but also the [Canvas
library from Taco de Wolff](https://github.com/tdewolff/canvas).

The [discord](discord) package contains the logic to interact with the Discord API.
[discord/security.go](discord/security.go) is here to make request's signature
verification easier.

The [canvas](canvas) package is used to draw the leaderboard once the data
have been fetched from the Duolingo API. After finishing the drawing, the
image is saved to a file named "leaderboard.png".

## Development

Check out the development framework documentation [here](https://github.com/scaleway/serverless-functions-go).

## Build and deploy the app

### Deploy to Scaleway cloud functions

Create a zip file using `zip -r function.zip .` and upload it to the Scaleway console.

In the "handler" field, put `Handle`.


