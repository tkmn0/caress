all:
	$(MAKE) build
	$(MAKE) setup
build:
	go build -ldflags="-s -w"  -buildmode=c-shared -o ./_dist/mac/libcaress.dylib

setup:
	install_name_tool -change "/usr/local/opt/opus/lib/libopus.0.dylib" "@rpath/libopus.0.dylib" ./_dist/mac/libcaress.dylib
	install_name_tool -change "/usr/local/lib/librnnoise-nu.1.dylib" "@rpath/librnnoise-nu.1.dylib" ./_dist/mac/libcaress.dylib
	install_name_tool -add_rpath @loader_path/. ./_dist/mac/libcaress.dylib
	otool -L ./_dist/mac/libcaress.dylib

clean:
	rm -rf ./_dist/mac/*