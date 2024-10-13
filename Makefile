run:
	@GOOS=darwin GOARCH=arm64 go build -o ./builds/annotate-arm64 ./annotation-tool
	@ls images | xargs ./builds/annotate-arm64
	
build:
	CGO_ENABLED=1 CC=clang GOOS=darwin  GOARCH=arm64 go build -ldflags "-linkmode external -s -w '-extldflags=-mmacosx-version-min=13.0.0'" -o ./builds/annotate-arm64 ./annotation-tool
	CGO_ENABLED=1 CC=clang GOOS=darwin  GOARCH=amd64 go build -ldflags "-linkmode external -s -w '-extldflags=-mmacosx-version-min=13.0.0'" -o ./builds/annotate-686 ./annotation-tool
	CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -o ./builds/annotate-win64.exe ./annotation-tool
	CGO_ENABLED=1 CC="i686-w64-mingw32-gcc" GOOS=windows GOARCH=386 go build -o ./builds/annotate-win32.exe ./annotation-tool
