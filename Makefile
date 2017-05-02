.DEFAULT_GOAL := help

GRAPHEME_BREAK_PROP := http://www.unicode.org/Public/9.0.0/ucd/auxiliary/GraphemeBreakProperty.txt
GRAPHEME_BREAK_TEST := http://www.unicode.org/Public/9.0.0/ucd/auxiliary/GraphemeBreakTest.txt

.PHONY: help
help: ## show this message
	@grep -E '^[a-zA-Z0-9_./-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-23s\033[0m %s\n", $$1, $$2}'

all: property.go property_test.go grapheme_break_test.go ## generate sources

property.go: tmp/GraphemeBreakProperty.txt property.py ## property.go
	printf '// Code generated from "$(GRAPHEME_BREAK_PROP)"; DO NOT EDIT\n\n' > $@
	./property.py main < $< >> $@
	go fmt $@
	go generate

property_test.go: tmp/GraphemeBreakProperty.txt property.py ## property.go
	printf '// Code generated from "$(GRAPHEME_BREAK_PROP)"; DO NOT EDIT\n\n' > $@
	./property.py property_test < $< >> $@
	go fmt $@
	go generate

tmp/GraphemeBreakProperty.txt: tmp
	curl -o $@ $(GRAPHEME_BREAK_PROP)

grapheme_break_test.go: tmp/GraphemeBreakTest.txt grapheme_break_test.py ## grapheme_break_test.go
	printf '// Code generated from "$(GRAPHEME_BREAK_TEST)"; DO NOT EDIT\n\n' > $@
	./grapheme_break_test.py main < $< >> $@
	go fmt $@

tmp/GraphemeBreakTest.txt: tmp
	curl -o $@ $(GRAPHEME_BREAK_TEST)

tmp:
	mkdir -p tmp

.PHONY: coverage
coverage: c.out ## coverage
	go tool cover -html=c.out

c.out: property.go break.go break_test.go grapheme_break_test.go
	go test -v -run=. -coverprofile=$@

.PHONY: bench ## benchmark
bench:
	go test -bench=.
