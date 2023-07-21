dev: clean
	kill $$(lsof -t -i :3000) $$(lsof -t -i :8788) || true
	cd website && deno task serve &
	wrangler pages dev --proxy 3000 \
		-e SLACK_CLIENT_ID=a \
		-e SLACK_CLIENT_SECRET=b \
		-e SLACK_VERIFICATION_TOKEN=c

build: clean
	cd website && deno task build

clean:
	rm -fr website/_site
