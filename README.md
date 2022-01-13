# Test with Revel (Golang)

This revel project is used for basic framework review with performance analysis that can be used as reference to decide on the framework to use in microservices. To know more information, visit [Performance Result](https://github.com/samueltan3972/framework-review)


## RUN
Install prerequisite before running the project:
- [Golang](https://go.dev/)
- [Git](http://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- Revel 
```
go get github.com/revel/revel
go get github.com/revel/cmd/revel
```

Visit [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html) if you have any problem with the installation.

To run: 

```
revel run -a test-with-revel
```

### Docker
```
docker build -t seanhao1233/test-with-revel .
docker run -p 8080:8080 seanhao1233/test-with-revel
```

## View Result

View the result at http://localhost:8080
it comes with 3 API end point
- /hello : return a simple hello world message
- /database : perform all database CRUD operation
- /fibonacci : compute and return first 5,000 fibonacci numbers



## Help

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

