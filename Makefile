SHELL = /bin/bash

#Foreground colors
_nc=$(shell tput sgr0)
_red=$(shell tput setaf 1)

.PHONY: snapshot
snapshot:
	@echo "Running snapshot release just for sure"
	@goreleaser release --rm-dist --snapshot
	@echo "Running SNAPSHOT release successfull"

.PHONY: check
check: snapshot
	@echo "Running checks: linter, git untracked/uncommited/unpushed changes, "
	@golangci-lint run
	@git diff-index --exit-code HEAD || (echo "$(_red)Uncommited changes$(_nc)" && exit 1)
	@[ $$(git ls-files -o --exclude-standard | wc -l) -eq 0 ] || (echo "$(_red)Untracked files$(_nc)" && git ls-files -o --exclude-standard && exit 1)
	@[ $$(git log --branches --not --remotes | wc -l) -eq 0 ] || (echo "$(_red)Unpushed commits$(_nc)" && git log --branches --not --remotes && exit 1)
	
PHONY: tag
tag: check
ifndef VERSION
	@echo "$(_red)VERSION must be specified$(_nc)" && echo "use make <target> VERSION=..." && exit 1
endif
	@echo "Create git tag: $(VERSION) if non-existent"
	@[ $$(git tag --list "$(VERSION)" | wc -l) -eq 1 ] || (git tag -a $(VERSION) -m "Release $(VERSION)" && git push origin $(VERSION))

.PHONY: release
release: tag
	@echo "makingrelease into github repo"
	@goreleaser release --rm-dist
