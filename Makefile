fmt:
	go fmt ./..

lint:
	golangci-lint run

update-changelog:
	sh scripts/changelog.sh
