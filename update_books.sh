#!/bin/bash
curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"