CartoDB client for Golang
=========================

[![GoDoc](https://godoc.org/github.com/agonzalezro/cartodb_go?status.png)](https://godoc.org/github.com/agonzalezro/cartodb_go)

**THIS IS A WORK IN PROGRESS**

The cartodb_go project is a Python client for the [CartoDB SQL
API](http://developers.cartodb.com/documentation/sql-api.html) that supports
[authentication using OAuth or API
key](http://developers.cartodb.com/documentation/sql-api.html#authentication).

Installation
------------

    go get github.com/agonzalezro/cartodb_go

Examples
--------

You can find examples of usage on the `examples/` folder of this repo.

To run the example you need to set `API_KEY` & `USERNAME` env variable with the
values for your CartoDB account. For example:

    CARTODB_API_KEY=123 CARTODB_USERNAME=alex go run examples/sql.go

How it works?
-------------

You can check the documentation of the project
[here](https://godoc.org/github.com/agonzalezro/cartodb_go).

Basically the library is going to take care of the client creation for you, but
after doing a request (Sql() || Req()) you will need to unmarshall the
response.Body or do whatever you want with it.

WIP
===

Things working
--------------

- sql method when user the api key client.

Things NOT working
------------------

- the Req() method for the API Key client was not tested.
- everything else :(

TODO
----

- tests & CI.
- throw different errors when the http status code is not 200?
