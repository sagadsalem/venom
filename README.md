<p align="center">
<img width="300px" src="https://github.com/sagadsalem/venom/blob/master/static/img/logo1.png"
style="border: 1px double #e6e1e1" align="center" width="200" >
</p>


# Venom

it's a simple mini Golang starter app done by using some packages with MVC structure :) highly inspired by Laravel

## Installation

Use [Docker](https://www.docker.com/) to install Venom.

```
 docker build -t starter .
 docker run -p 8080:8080 starter
```
Or using [GoModules](https://blog.golang.org/using-go-modules).

```
 export GO111MODULE=on
 go run main.go
```

## Thanks

Thanks for all these packages for all awesome work they do

[gorilla/mux](github.com/gorilla/mux)

[jinzhu/gorm](github.com/jinzhu/gorm)

[joho/godotenv](github.com/joho/godotenv)

[unrolled/render](github.com/unrolled/render)

## License
[MIT](https://choosealicense.com/licenses/mit/)
