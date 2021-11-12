
.PHONY: clean

default: upper.plugin

clean:
	rm upper.plugin

upper.plugin: plugImpl.go
	$(GO) build -o upper.plugin -buildmode=plugin plugImpl.go

install: upper.plugin
	mkdir -p ../plugins
	cp upper.plugin ../plugins
