all: test

test:
	gotest -v $$(go list ./... | grep -v generated)

plan :
	cdktf plan

apply :
	cdktf apply
