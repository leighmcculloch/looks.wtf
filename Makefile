clean:
	$(MAKE) -C website clean
	$(MAKE) -C service clean

build: clean
	docker build -t lookswtf .

run: build
	docker run -it -p 8080:8080 lookswtf

deploy:
	fly deploy --local-only --strategy bluegreen
