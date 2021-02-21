# Create new project and setup folders

Create your directory structure under $GOPATH/src.

```
cd $GOPATH/src
mkdir romanservice
mkdir -p romanservice/romanNumerals
mkdir -p romanservice/romanserver
cd romanservice
touch README.md
```
# Init Module

only after the *.go files are created

```
go mod init
go mod tidy
```

```
go fmt ./romanserver
go fmt ./romanNumerals
go install ./romanserver
```


# Run

```
bash -c '$(go env GOPATH)/bin/romanserver'

curl http://localhost:8000/roman_number/9
curl -X GET "http://localhost:8000/roman_number/5"
curl -X GET "http://localhost:8000/roman_number/12"
curl -X GET "http://localhost:8000/random_resource/3" # Invalid resource
```

# Git

```
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:dliulabs/romanservice.git
git push -u origin main
```

# Supervisor

```
brew reinstall supervisor
brew services start supervisor -v
ls /usr/local/homebrew/etc/supervisord.conf
ls /usr/local/homebrew/etc/supervisor.d/*.ini
sudo cp /Users/David/go/src/romanservice/goproject.conf /usr/local/homebrew/etc/supervisor.d/
```

# Basic HTTP Handler

```
cd basichandler
go run basichandler.go

curl -X GET http://localhost:8000/hello
```

# CustomServeMux

```
cd customservemux
go run customMux.go

curl -X GET http://localhost:8000
curl -X GET http://localhost:8000/randomFloat
curl -X GET http://localhost:8000/randomInt  
```

# Using httprouter

## Install

```
go get github.com/julienschmidt/httprouter
```
