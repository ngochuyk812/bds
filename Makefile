
.PHONY: update-subtree

update-subtree:
	git subtree pull --prefix=fe/dashboard/src/proto https://github.com/ngochuyk812/proto-bds.git gen-js --squash