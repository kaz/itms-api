.PHONY: noop
noop:

.PHONY: deploy
deploy:
	gcloud functions deploy ItmsApi \
	--region asia-northeast1 \
	--project itms-api \
	--runtime go113 \
	--entry-point Handler \
	--trigger-http \
	--allow-unauthenticated \
	--env-vars-file env.yml
