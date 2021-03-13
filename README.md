# apassphrase

## About

A Go web application that returns JSON object containing an Passphrase or Emojiphrase

## API Endpoints

GET:
    /passphrase - Returns a JSON object {"passphrase": String}
    /emojiphrase - Returns a JSON object {"emojiphrase": String, "emojis": String}

## Environment Variables

PORT - HTTP Port for application to listen for requests on
