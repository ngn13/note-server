SRCS = $(wildcard *.go) $(wildcard */*.go)
PREFIX = usr

note-server: $(SRCS)
	go build

install:
	install -Dm755 note-server $(DESTDIR)/$(PREFIX)/bin/note-server
	install -dm655 static $(DESTDIR)/$(PREFIX)/lib/note-server/static
	install -dm655 views $(DESTDIR)/$(PREFIX)/lib/note-server/views

uninstall:
	rm $(DESTDIR)/$(PREFIX)/bin/note-server
	rm -r $(DESTDIR)/$(PREFIX)/lib/note-server

format:
	gofmt -s -w .

.PHONY: install uninstall format
