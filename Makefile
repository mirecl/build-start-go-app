TIME=`date +%Y-%m-%dT%H:%M:%S`
COMMIT=`git log --pretty=format:'%h' -n 1`
REPO=`git config --get remote.origin.url`
BRANCH=`git symbolic-ref --short HEAD`

build:
	@echo "Build starting..."
	@go build -ldflags "-s \
	-X main.Repo=${REPO} \
	-X main.Branch=${BRANCH} \
	-X main.Commit=${COMMIT} \
	-X main.Time=${TIME}" \
	-o server