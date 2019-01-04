deploy:
	$(MAKE) build
	$(MAKE) -C service deploy

clean:
	$(MAKE) -C website clean
	$(MAKE) -C service clean

build: clean
	cp *.yml website/data/
	cp *.yml service/data/
	$(MAKE) -C website build
	cp -r website/build service/static
