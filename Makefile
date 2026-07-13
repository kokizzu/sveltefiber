GO ?= go
GOVULNCHECK ?= govulncheck
NPM ?= npm
WEB_DIR ?= svelte
CMD ?=
.PHONY: test frontend-build frontend-audit verify-dependency-security run vulncheck

test:
	$(GO) test ./...
	$(NPM) --prefix $(WEB_DIR) run build:prod

frontend-build:
	$(NPM) --prefix $(WEB_DIR) run build:prod

frontend-audit:
	$(NPM) --prefix $(WEB_DIR) audit

verify-dependency-security:
	bash ./scripts/verify-dependency-security.sh
	$(NPM) --prefix $(WEB_DIR) audit

vulncheck:
	$(GOVULNCHECK) ./...

run:
	@test -n "$(CMD)" || (echo "usage: make run CMD='go test ./...'" >&2; exit 2)
	$(CMD)
