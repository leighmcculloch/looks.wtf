export CLOUDFLARE_ZONE = 663e42bb29abec71fd4fa45f82dfadd7

deploy:
	$(MAKE) build
	$(MAKE) -C service deploy
	$(MAKE) cdn

build:
	cp *.yml website/data/
	cp *.yml service/data/
	$(MAKE) -C website clean build

cdn:
	curl -X DELETE "https://api.cloudflare.com/client/v4/zones/$(CLOUDFLARE_ZONE)/purge_cache" \
		-H "X-Auth-Email: $(CLOUDFLARE_EMAIL)" \
		-H "X-Auth-Key: $(CLOUDFLARE_CLIENT_API_KEY)" \
		-H "Content-Type: application/json" \
		--data '{"purge_everything":true}'
