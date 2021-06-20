TAG=asia.gcr.io/itms-api/itms-api

.PHONY: noop
noop:

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test -race ./...

.PHONY: build
build:
	docker build -t $(TAG) --platform linux/x86_64 .

.PHONY: push
push:
	docker push $(TAG)

.PHONY: deploy
deploy:
	gcloud run deploy --platform managed --project=itms-api --region=asia-northeast1 --image $(TAG) itms-api
