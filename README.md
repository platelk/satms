# SATMS
Simple And Tiny Message Server

## Disclaimer

Warning : This project is a pet project and shouldn't be used for a real product/project

## Description

SATMS is a messaging server that enable communication of client through WebSocket mainly but can be accessed by HTTP call.

The message struct is quite simple :

- topic : Topic of the message
- to : Id of the client you want to send a message to
- from : Id of the sender
- body : content of the message

Actual the JSON format is used, but others could supported

## How to install ?

SATMS use go language so standard golang tools can be use:

    go get -d github.com/platelk/satms
    cd $GOPATH/src/github.com/platelk/satms
    go build main.go
    ./main

## How to run the example ?

The client example is wrote in [Dart](https://www.dartlang.org/), so first you have to download and install Dart, then :

    cd $GOPATH/src/github.com/platelk/satms/example/client
    pub get
    pub build
    <you_favorite_browser> build/web/main.html

Note: You have to launch the SATMS service first
Note2: A build script is available in the client directory

## Author

- Kévin PLATEL : platel.kevin@gmail.com