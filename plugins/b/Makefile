
.PHONY: clean

default: b.plugin

clean:
	rm b.plugin

b.plugin: pluglower.go
	$(GO) build -o b.plugin -buildmode=plugin pluglower.go
