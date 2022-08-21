#!/bin/bash -x

docker compose exec backend goose.sh up
docker compose exec backend go run main.go job loaddata -f /backend/jobs/loaddata/data/images.csv -t csv -m images
docker compose exec backend go run main.go job loaddata -f /backend/jobs/loaddata/data/about_topics.csv -t csv -m about_topics
docker compose exec backend go run main.go job loaddata -f /backend/jobs/loaddata/data/languages.csv -t csv -m languages
docker compose exec backend go run main.go job loaddata -f /backend/jobs/loaddata/data/works.csv -t csv -m works
