clean:
	rm -f healthcheck

buildlinux: clean
	GOOS=linux GOARCH=amd64 go build -o build/healthcheck .
	docker build -t adorigi/healthcheck .

builddarwin: clean
	GOOS=darwin GOARCH=amd64 go build -o build/healthcheck .
	docker build -t adorigi/healthcheck .

kindpush: build
	kind load docker-image adorigi/healthcheck --name healthcheck

dockerpush: build
	docker push adorigi/healthcheck
