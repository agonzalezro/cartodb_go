CartoDB client for Golang
=========================

[![GoDoc](https://godoc.org/github.com/agonzalezro/cartodb_go?status.png)](https://godoc.org/github.com/agonzalezro/cartodb_go)

**THIS IS A WORK IN PROGRESS**

The cartodb_go project is a Python client for the [CartoDB SQL
API](http://developers.cartodb.com/documentation/sql-api.html) that supports
[authentication using API
key](http://developers.cartodb.com/documentation/sql-api.html#authentication).
CartoDB also support OAuth authentication, but it's not implemented here yet
(and it's not going to be soon).

Installation
------------

    go get github.com/agonzalezro/cartodb_go

Examples
--------

You can find examples of usage on the `examples/` folder of this repo.

In some of them you will need to set some environment variables. For example:

    CARTODB_API_KEY=123 go run examples/sql/api_key/example.go

How it works?
-------------

You can check the documentation of the project
[here](https://godoc.org/github.com/agonzalezro/cartodb_go).

Basically the library is going to take care of the client creation for you, but
after doing a request (Sql() || Req()) you will need to unmarshall the
response.Body or do whatever you want with it.

TODO
----

- the Req() method for the API Key client was not tested yet.
- in some point in the future it could be nice to have OAuth authentication too.
- tests & CI.
- throw different errors when the http status code is not 200?
