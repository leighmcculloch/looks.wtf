run:
	dev_appserver.py --default_gcs_bucket_name looks-wtf.appspot.com dispatch.yaml slack/app.yaml slackoauth/app.yaml slackcommandlook/app.yaml

deploy:
	gcloud app deploy --project looks-wtf --promote dispatch.yaml slack/app.yaml slackoauth/app.yaml slackcommandlook/app.yaml
