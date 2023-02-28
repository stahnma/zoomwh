#!/bin/bash

curl -X POST -H "Content-Type: application/json" -d @examples/participant_left.json http://localhost:9999/zoom
