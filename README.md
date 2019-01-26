# Time Tracker
Track your projects

## Features

* Todo

## Prerequisites
* [Go](https://golang.org/)
* A Firebase project with a [Realtime Database](https://firebase.google.com/docs/database/)
* A [service account](https://developers.google.com/identity/protocols/OAuth2ServiceAccount) json file

## Installation
### Using `install.sh` script
Todo

### Using `go get`
**Note**: First you need to add the environment variables for the URL to the Firebase Realtime Database (var: `TRACKER_DB_URL`) and the service account json file (var: `TRACKER_SERVICE_ACCOUNT`). You can do this manually or by running the following:
```bash
# Todo using preinstall.sh
```
Now:
```bash
$ go get -u github.com/Samuel-L/time-tracker
# If your $GOPATH/bin is in your PATH, you can simply:
$ time-tracker --version
time tracker version 0.5.0
```

## Usage
```bash
$ time-tracker --help
NAME:
   time tracker - A new cli application

USAGE:
   time-tracker [global options] command [command options] [arguments...]

VERSION:
   0.5.0

DESCRIPTION:
   Track your time working on projects!

COMMANDS:
     start    start tracking a project
     stop     stop tracking a project
     day      view your day
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
