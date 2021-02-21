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



