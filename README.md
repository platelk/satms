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

## Author

- Kévin PLATEL : platel.kevin@gmail.com