run:
	@go build -o ./builds/annotate-arm ./annotation-tool 

	@ls images | xargs ./builds/annotate-arm
	
build:
	@go build -o ./builds/annotate-arm ./annotation-tool 
	GOOS=windows GOARCH=386 go build -o ./builds/annotate-win ./annotation-tool
