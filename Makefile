export CLOUDFLARE_ZONE = 663e42bb29abec71fd4fa45f82dfadd7

PORT ?= 8080

ifeq ($(strip $(shell git status --porcelain)),)
export VERSION = `git show -s --format=%cd --date=format:'%Y%m%dt%H%M%S' HEAD`-`git rev-parse --short HEAD`
else
export VERSION = `date --utc +%Y%m%dt%H%M%S`-dev
endif

build:
	embedfiles -out=shared/looks/yaml_files.go -pkg=looks looks.yml tags.yml
	$(MAKE) -C services/default/website clean build

run:
	dev_appserver.py \
		--port=$(PORT) \
		--default_gcs_bucket_name looks-wtf.appspot.com \
		services/dispatch.yaml \
		services/default/app.yaml \
		services/slackoauth/app.yaml \
		services/slackcommandlook/app.yaml \
		services/slackcommandlooks/app.yaml \
		services/slackactions/app.yaml

push: build
	gcloud app deploy \
		--project looks-wtf \
		--version $(VERSION) \
		--promote \
		services/dispatch.yaml \
		services/default/app.yaml \
		services/slackoauth/app.yaml \
		services/slackcommandlook/app.yaml \
		services/slackcommandlooks/app.yaml \
		services/slackactions/app.yaml

cdn:
	curl -X DELETE "https://api.cloudflare.com/client/v4/zones/$(CLOUDFLARE_ZONE)/purge_cache" \
		-H "X-Auth-Email: $(CLOUDFLARE_EMAIL)" \
		-H "X-Auth-Key: $(CLOUDFLARE_CLIENT_API_KEY)" \
		-H "Content-Type: application/json" \
		--data '{"purge_everything":true}'

deploy: build push cdn

setup:
	go get 4d63.com/embedfiles
