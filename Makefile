.DEFAULT_GOAL:= all

all: 1 2

.PHONY: 1
1:
	cd one && go test -v

.PHONY: 2
2:
	cd two && go test -v
