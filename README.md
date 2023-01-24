# ldb
Lightweight database wrapper around database/sql for Go developers

**Note:** Only PostgreSQL and MySQL are supported at the moment

## Why ldb?

ldb provides a very minimal interface over the existing `database/sql` package to reduce the amount of overhead needed 
when setting up database connections and making queries. It aims to be as liberal as possible so there are very little 
frills.

## Usage

Install the package:

```text
go get -u github.com/woojiahao/ldb
```