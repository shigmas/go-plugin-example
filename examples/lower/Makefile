
.PHONY: clean

default: lower.plugin

clean:
	rm lower.plugin

lower.plugin: plugImpl.go
	$(GO) build -o lower.plugin -buildmode=plugin plugImpl.go

install: lower.plugin
	mkdir -p ../plugins
	cp lower.plugin ../plugins
