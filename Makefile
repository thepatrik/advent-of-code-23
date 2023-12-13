.DEFAULT_GOAL:= all

all: 1 

.PHONY: 1
1:
	cd one && go test -v