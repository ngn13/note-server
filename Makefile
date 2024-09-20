SRCS = $(wildcard *.go) $(wildcard */*.go)
PREFIX = usr

note-server: $(SRCS)
	go build

install:
	install -dm755 $(DESTDIR)/$(PREFIX)/bin/note-server
	install -dm755 $(DESTDIR)/$(PREFIX)/lib/note-server/static
	install -dm755 $(DESTDIR)/$(PREFIX)/lib/note-server/views
	install -m755 note-server $(DESTDIR)/$(PREFIX)/bin/note-server
	cp -r static/* $(DESTDIR)/$(PREFIX)/lib/note-server/static
	cp -r views/*  $(DESTDIR)/$(PREFIX)/lib/note-server/views

uninstall:
	rm $(DESTDIR)/$(PREFIX)/bin/note-server
	rm -r $(DESTDIR)/$(PREFIX)/lib/note-server

format:
	gofmt -s -w .

.PHONY: install uninstall format
