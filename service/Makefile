PORT ?= 8080

ifeq ($(strip $(shell git status --porcelain)),)
export VERSION = `git show -s --format=%cd --date=format:'%Y%m%dt%H%M%S' HEAD`-`git rev-parse --short HEAD`
else
export VERSION = `date --utc +%Y%m%dt%H%M%S`-dev
endif

clean:
	rm -fR static

datastore:
	gcloud beta emulators datastore start

run:
	$$(gcloud beta emulators datastore env-init) && \
		cd scripts && \
		go run local_setup.go && \
		cd .. && \
		go run .

deploy:
	gcloud app deploy \
		--project looks-wtf \
		--version $(VERSION) \
		--promote \
		app.yaml
