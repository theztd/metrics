
build-linux:	
	env GOOS=linux GOARCH=amd64 go build -ldflags "-extldflags '-static'" -o build/metrics


build-rpi:
	env GOOS=linux GOARCH=arm GOARM=7 go build -ldflags "-extldflags '-static'" -o build/metrics-arm7
