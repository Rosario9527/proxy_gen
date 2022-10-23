BUILDDIR=$(shell pwd)/build
IMPORT_PATH= \
	proxy_gen/freeport \
	proxy_gen/util \
	proxy_gen/v2ray \
	proxy_gen/xtun2socks

all: ios android

ios: clean
	mkdir -p $(BUILDDIR)
	gomobile bind -o $(BUILDDIR)/golibs.framework -a -ldflags '-w -s' -target=ios $(IMPORT_PATH)


android: clean
	mkdir -p $(BUILDDIR)
	env GO111MODULE="on" gomobile bind -o $(BUILDDIR)/golibs.aar -a -v -x -androidapi 23 -tags "full" -trimpath -ldflags '-w -s' -target=android $(IMPORT_PATH)

clean:
	gomobile clean
	rm -rf $(BUILDDIR)

cleanmodcache:
	go clean -modcache
