CartoDB client for Golang
=========================

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

WIP
===

Things working
--------------

- sql method when user the api key client.

Things NOT working
------------------

- easy... the rest.

TODO
----

- Tests & CI
- Godoc

Doubts
------
Trying to mimic the
[cartodb-python](https://github.com/Vizzuality/cartodb-python) usage, I am not
pretty sure that I can unmarshall the json response so easily, I need to
investigate that.
