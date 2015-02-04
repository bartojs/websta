build:
	GOPATH=`pwd` go install webstatic
	GOPATH=`pwd` GOOS=windows GOARCH=386 go install webstatic
	GOPATH=`pwd` GOOS=linux GOARCH=386 go install webstatic

run:
	bin/webstatic -port 8000 -dir test
