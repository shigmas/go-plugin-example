
.PHONY: all clean
TARGETS := plugins/a/a.plugin plugins/b/b.plugin cmd/driver/driver

all: plugins/a/a.plugin plugins/b/b.plugin main

clean:
	@rm $(TARGETS)

plugins/a/a.plugin:
	$(MAKE) -C plugins/a

plugins/b/b.plugin:
	$(MAKE) -C plugins/b

main:
	$(MAKE) -C 	cmd/driver

test:
	$(warning $(FOO))
