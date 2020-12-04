
.PHONY: clean

default: a.plugin

clean:
	rm a.plugin

a.plugin: plugupper.go
	$(GO) build -o a.plugin -buildmode=plugin plugupper.go
