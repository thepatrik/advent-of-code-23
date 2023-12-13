.DEFAULT_GOAL:= all

all: 1 2 3 4

.PHONY: 1
1:
	cd one && go test -v

.PHONY: 2
2:
	cd two && go test -v

.PHONY: 3
3:
	cd three && go test -v

.PHONY: 4
4:
	cd four && go test -v