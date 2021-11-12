
.PHONY: clean

default: lower.plugin

clean:
	rm lower.plugin

lower.plugin: plugImpl.go
	$(GO) build -o ../plugins/lower.plugin -buildmode=plugin plugImpl.go

install: lower.plugin
	cp lower.plugin ../plugins
