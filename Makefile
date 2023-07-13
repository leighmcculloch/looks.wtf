HUGO?=go run -tags extended github.com/gohugoio/hugo@v0.115.2

dev: clean
	kill $$(lsof -t -i :9000) $$(lsof -t -i :8788) || true
	$(HUGO) server --port 9000 &
	wrangler pages dev --proxy 9000 \
		-e SLACK_CLIENT_ID=a \
		-e SLACK_CLIENT_SECRET=b \
		-e SLACK_VERIFICATION_TOKEN=c

build: clean
	$(HUGO)

clean:
	rm -fr public
