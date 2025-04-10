all: test

test:
	gotest -v $$(go list ./... | grep -v generated)

get:
	cdktf get

plan :
	cdktf plan

apply :
	cdktf apply
