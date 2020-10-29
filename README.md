# Go RESTful Yaml Server

## Setup and Installation

To setup the project, first install the module with

```sh
$ go install

... install stdout
```

Then you can start the server with

```sh
$ go run main.go

Server running at localhost:50000
```

You can also specify the port number manually with the `-p` flag.

```sh
$ go run main.go -p=8080

Server running at localhost:8080
```

## Testing

There are a few unit tests written in the `dataservice` and `validators` packages. Run `go test ./dataservice ./validators`.

To run the included python integration tests, first start the server with

```sh
$ go run main.go

Server running at localhost:50000
```

Then in another shell, run the following to create a metadata resource on the server

```sh
$ python integration-tests/createValid1.py

Response: 201 OK
id: 1
title: valid app 1
version: 1.2.3
maintainers:
- name: John
  email: john@doe.com
company: Boring Company
website: golang.org
source: github.com
license: Apache License 2.0
description: |-
  my multiline description
  which supports markdown
```

Try adjusting the `valid1.yaml` data file and re-running the `createValid.py` script to create more resources.

Then run the `testSearch.py` script to test for fetching the resources which have the `Title` equal to "valid app 1" (this could be multiple)

If 5 of the same were created the output would look like the following

```sh
$ python integration-tests/testSearch.py

Response: 200
items:
  -
    id: 2
    title: valid app 1
    version: 1.2.3
    maintainers:
    - name: John
      email: john@doe.com
    company: Boring Company
    website: golang.org
    source: github.com
    license: Apache License 2.0
    description: |-
      my multiline description
      which supports markdown
  -
    id: 1
    title: valid app 1
    version: 1.2.3
    maintainers:
    - name: John
      email: john@doe.com
    company: Boring Company
    website: golang.org
    source: github.com
    license: Apache License 2.0
    description: |-
      my multiline description
      which supports markdown
  -
    id: 5
    title: valid app 1
    version: 1.2.3
    maintainers:
    - name: John
      email: john@doe.com
    company: Boring Company
    website: golang.org
    source: github.com
    license: Apache License 2.0
    description: |-
      my multiline description
      which supports markdown
  -
    id: 3
    title: valid app 1
    version: 1.2.3
    maintainers:
    - name: John
      email: john@doe.com
    company: Boring Company
    website: golang.org
    source: github.com
    license: Apache License 2.0
    description: |-
      my multiline description
      which supports markdown
  -
    id: 4
    title: valid app 1
    version: 1.2.3
    maintainers:
    - name: John
      email: john@doe.com
    company: Boring Company
    website: golang.org
    source: github.com
    license: Apache License 2.0
    description: |-
      my multiline description
      which supports markdown
count: 5
```

## API Endpoints

There are 3 configured endpoints on the server

1. POST - `/api/metadata` - accepts yaml as the request body and will validate the schema based
2. GET - `/api/metadata/{id}` - fetches the metadata resource with id = {id}. Returns 404 if resource is not found.
3. GET - `/api/metadata?[queryparams]` - if query params are present, it will use them to filter the list of resources that match the query params. If omitted, the top 100 resources are fetched.

Note on the search capability: it assumes `AND` filtering for all query parameters, e.g. if `title=title1&version=2.0.0` then it will return all resources `WHERE title = 'title1' AND version = '2.0.0'`.
