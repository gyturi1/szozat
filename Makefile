SHELL = /bin/bash

#Foreground colors
_nc=$(shell tput sgr0)
_red=$(shell tput setaf 1)

.PHONY: snapshot
snapshot:
	@goreleaser release --rm-dist --snapshot

.PHONY: check
check:
	@golangci-lint run
	@git diff-index --exit-code HEAD || (echo "$(_red)Uncommited changes$(_nc)" && exit 1)
	@[ $$(git ls-files -o --exclude-standard | wc -l) -eq 0 ] || (echo "$(_red)Untracked files$(_nc)" && git ls-files -o --exclude-standard && exit 1)
	@$(MAKE) snapshot
	
PHONY: release
release: check
	@git tag -a $(VERSION) -m "Release $(VERSION)"
	@git push origin $(VERSION)
	@goreleaser release --rm-dist
