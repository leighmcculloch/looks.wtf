export CLOUDFLARE_ZONE = 663e42bb29abec71fd4fa45f82dfadd7

run:
	bundle exec middleman

deploy: clean build push-github

clean:
	rm -fR build

build:
	bundle exec middleman build

push-github:
	git branch -D gh-pages 2>/dev/null | true
	git branch -D gh-pages-draft 2>/dev/null | true
	git checkout -b gh-pages-draft && \
		git add -f build && \
		git commit -m "Deploy to gh-pages" && \
		git subtree split --prefix build -b gh-pages && \
		git push --force origin gh-pages:gh-pages && \
		git checkout -

cdn:
	curl -X DELETE "https://api.cloudflare.com/client/v4/zones/$(CLOUDFLARE_ZONE)/purge_cache" \
		-H "X-Auth-Email: $(CLOUDFLARE_EMAIL)" \
		-H "X-Auth-Key: $(CLOUDFLARE_CLIENT_API_KEY)" \
		-H "Content-Type: application/json" \
		--data '{"purge_everything":true}'
