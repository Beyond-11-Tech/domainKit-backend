# DomainKit, by Beyond11 technologies

this is the working backend for the domainKit application.

for api documentation see docs/api.md

## development guide
the backend is developed in Go with Gin as the web framework. all calls to check domain information is currently handled by os terminal calls to the Dig command

#### required deps
to develop, or build the application you will need Go installed with on at least version 1.24.2 or greater.
once installed running `go mod tidy` will install all required dependencies

#### running locally
once the deps are installed, the app can then be developed and once ready, ran locally. this is done by using the command `go run .`. the app does **not** have hot reloading, so once a change is made, you will need to stop the app and restart it using the same command.

## build steps
to build the app for distribution, you will want to run the command `go build .`

this will build the app using the system default environments of the OS so for now make sure the os matches the build destination (this will be fixed in a future update)

the application needs the flags `webKey` & `appKey` to be passed accepted passwords for the web app and mobile app to use to authenticate with the API server, they can be passed to the app in the same way other flags from other apps are passed.