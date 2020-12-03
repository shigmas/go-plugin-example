
GO=/home/wyan/.gvm/gos/go1.13/bin/go
all: plugins/a/a.so plugins/b/b.so main

plugins/a/a.so:
	$(MAKE) -C plugins/a

plugins/b/b.so:
	$(MAKE) -C plugins/b

main:
	$(MAKE) -C 	cmd/driver

