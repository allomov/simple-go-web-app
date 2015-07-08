# Description

Super simple web app written in go lang. The main purpose of this app is to test CF installations.

# How to run 

Run following:
```
go install github.com/allomov/simple-go-web-app
$GOPATH/bin/simple-go-web-app
```

# Deploying to CF with null-buildpack

```
go get -d github.com/allomov/simple-go-web-app
cd $GOPATH/github.com/allomov/simple-go-web-app
./bin/build
cf push
```
