tests:
	go test `go list ./... | grep -v "/vendor/"`

coverage:
	overalls -project=github.com/thedodd/barge -covermode=atomic -debug -ignore=.git,vendor > /dev/null 2>&1
	go tool cover -html=overalls.coverprofile -o coverage.html
	# Open in your browser with `open coverage.html`.

cover_profiles = `find . -name '*.coverprofile'`
clean:
	if [ ! -z $(cover_profiles) ]; then rm $(cover_profiles); fi
	if [ -f coverage.html ]; then rm coverage.html; fi
