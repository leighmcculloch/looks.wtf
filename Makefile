export CLOUDFLARE_ZONE = 663e42bb29abec71fd4fa45f82dfadd7

deploy:
	$(MAKE) build
	$(MAKE) -C service deploy
	$(MAKE) cdn

clean:
	$(MAKE) -C website clean
	$(MAKE) -C service clean

build: clean
	cp *.yml website/data/
	cp *.yml service/data/
	$(MAKE) -C website build
	cp -r website/build service/static

cdn:
	curl -X DELETE "https://api.cloudflare.com/client/v4/zones/$(CLOUDFLARE_ZONE)/purge_cache" \
		-H "X-Auth-Email: $(CLOUDFLARE_EMAIL)" \
		-H "X-Auth-Key: $(CLOUDFLARE_CLIENT_API_KEY)" \
		-H "Content-Type: application/json" \
		--data '{"purge_everything":true}'
