run: env
	dev_appserver.py dispatch.yaml slack/app.yaml slackoauth/app.yaml slackcommandlook/app.yaml

deploy: env
	gcloud app deploy --project looks-wtf dispatch.yaml slack/app.yaml slackoauth/app.yaml slackcommandlook/app.yaml

env:
	envsubst < slackoauth/app.raw.yaml > slackoauth/app.yaml
	envsubst < slackcommandlook/app.raw.yaml > slackcommandlook/app.yaml
