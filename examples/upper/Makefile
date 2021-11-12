
.PHONY: clean

default: upper.plugin

clean:
	rm upper.plugin

upper.plugin: plugImpl.go
	$(GO) build -o ../plugins/upper.plugin -buildmode=plugin plugImpl.go

install: upper.plugin
	cp upper.plugin ../plugins
