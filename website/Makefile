export CLOUDFLARE_ZONE = 663e42bb29abec71fd4fa45f82dfadd7

run:
	bundle exec middleman

clean:
	rm -fR build

build:
	bundle exec middleman build

push:
	gsutil -m                              cp -a public-read build/index.html gs://looks.wtf
	gsutil -m                              cp -a public-read build/sitemap.xml gs://looks.wtf
	gsutil -m                              cp -a public-read -r build/javascripts gs://looks.wtf/javascripts
	gsutil -m -h "Content-Type: text/html" cp -a public-read `find build -type f -regex '[^.]*'` gs://looks.wtf
	gsutil -m                              cp -a public-read -r build/stylesheets gs://looks.wtf/stylesheets
	gsutil -m                              cp -a public-read -r build/javascripts gs://looks.wtf/javascripts
	gsutil web set -m index.html gs://looks.wtf

cdn:
	curl -X DELETE "https://api.cloudflare.com/client/v4/zones/$(CLOUDFLARE_ZONE)/purge_cache" \
		-H "X-Auth-Email: $(CLOUDFLARE_EMAIL)" \
		-H "X-Auth-Key: $(CLOUDFLARE_CLIENT_API_KEY)" \
		-H "Content-Type: application/json" \
		--data '{"purge_everything":true}'

deploy: clean build push cdn
