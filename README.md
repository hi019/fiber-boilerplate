## fiber-boilerplate
fiber-boilerplate is a boilerplate for Golang web services (rest APIs, templating...) built using [Fiber](https://github.com/gofiber/fiber). Its domain-driven design was inspired by [Gorsk](https://github.com/ribice/gorsk). Here are some of the boilerplate's out of the box features:

* **Fast  prototyping** - Uses Facebook's ORM software, [Ent](https://github.com/facebook/ent)
* **Fast** - Fiber [benchmarks](https://github.com/gofiber/fiber#-benchmarks) show that it's the fastest Golang HTTP framework, handling around 35k req/s. That's much more than net/http.
* **Extendable** - The boilerplate's domain driven file structure makes organization of large projects easy

## Note
This boilerplate uses Fiber v15, which is currently in beta and will be released to stable on September 15th. Its possible there will be breaking changes between now and that date. 
In that case, I will continue to update this project. 

### TODO
- [ ] More unit tests
- [ ] Incorporate Google's [Wire](https://github.com/google/wire) for Dependency Injection
- [ ] Add user authentication 