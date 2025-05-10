
.PHONY: update-subtree

update-subtree:
	git subtree pull --prefix=fe/dashboard/src/proto origin gen-js --squash
