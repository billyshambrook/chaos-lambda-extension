EXTENSIONS := $(shell find extensions -name Dockerfile -maxdepth 2 -exec dirname {} \;)

.PHONY: build-extensions
build-extensions:
	@for extension in $(EXTENSIONS) ; do \
		(cd $$extension ; docker build -t extension-$$(basename $$extension) . ; cd -) ; \
	done

.PHONY: build-functions
build-functions:
	sam build

.PHONY: build
build: build-extensions build-functions
