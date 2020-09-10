## fiber-boilerplate
fiber-boilerplate is a boilerplate for Golang web services (rest APIs, templating...) built using [Fiber](https://github.com/gofiber/fiber). Its domain-driven design was inspired by [Gorsk](https://github.com/ribice/gorsk). Here are some of the boilerplate's out of the box features:

* **Fast  prototyping** - Uses Facebook's ORM software, [Ent](https://github.com/facebook/ent)
* **Fast** - Fiber [benchmarks](https://github.com/gofiber/fiber#-benchmarks) show that it's the fastest Golang HTTP framework, handling around 35k req/s. That's much more than net/http.
* **Extendable** - The boilerplate's domain driven file structure makes organization of large projects easy

## Note
This boilerplate uses Fiber v15, which is currently in beta and will be released to stable on September 15th. Its possible there will be breaking changes between now and that date. 
In that case, I will continue to update this project. 

## Usage
After cloning this repository, run `go mod download` to download the dependencies. This project uses [Cobra](https://github.com/spf13/cobra) for its command line interface. You can do `go run main.go serve` to start the server with the default arguments (port 3000, sqlite) or `go run main.go --help` to see the available arguments. [Air](https://github.com/cosmtrek/air) also comes preconfigured, which enables easy reloading on change. To use Air, first see its [installation](https://github.com/cosmtrek/air#installation) then run `air -c air.toml`. To customize the command that is ran by Air, open `air.toml` and change `full_bin`. 

Available routes:
* `POST /signup` with body `{email: string, password: string}`
* `POST /login` with body `{email: string, password: string}`

To use in your own projects, make sure you use your editors Find & Replace feature to replace all refrences to `fiber-boilerplate` with your project's name. You should also rename the directory. 

## Directory Structure
```
.
├── air.toml # air configuration  
├── cmd # commands
│   ├── root.go   
│   └── serve.go
├── ent # ent (ORM) configuration. This is auto-generated, you only need to edit files in `schema` 
│   ├── schema
│   │   └── user.go
├── main.go
├── pkg
│   └── api # api modules
│       ├── api.go
│       └── user # user module
│           ├── service.go # user service defintion, includes dependencies, interfaces and an initialization function that returns an instance of the service
│           ├── user.go # the actual implementation of the user service
│           └── web
│               ├── web.go # user module handlers
│               └── web_test.go # handler tests
└── README.md
```

### TODO
- [ ] More unit tests
- [ ] Incorporate Google's [Wire](https://github.com/google/wire) for Dependency Injection
- [ ] Add user authentication 

