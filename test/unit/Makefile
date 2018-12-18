TEST=go test

.SILENT: test

test:
	$(TEST) -v ./...

	for pkg in ../../src/*/ ; do	\
		$(TEST) -coverprofile=coverage.out -coverpkg $$(basename $$pkg) ; \
		go tool cover -func=coverage.out ; \
		go tool cover -html=coverage.out ; \
	done							\

	rm coverage.out
