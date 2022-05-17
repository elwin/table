# table
`table` is a utility tool to print CSV or JSON documents in a table format.

## Installation
### Homebrew
```console
brew install elwin/tols/table
```

### Go
Requires a [go toolchain installation](https://golang.org/doc/install) to be present.
```console
go get github.com/elwin/table
```

## Usage
```console
$ echo 'id,name,price
  1,apple,15
  2,banana,10' | table
+----+--------+-------+
| ID |  NAME  | PRICE |
+----+--------+-------+
|  1 | apple  |    15 |
|  2 | banana |    10 |
+----+--------+-------+
```

By default, CSV will be assumed. Alternatively, JSON can be used by specifying `--format json` or `-f json`:
```console
$ echo '[
  {
    "id": "1",
    "name": "apple",
    "price": "15"
  },
  {
    "id": "2",
    "name": "banana",
    "price": "10"
  }
]' | table --format json
+----+--------+-------+
| ID |  NAME  | PRICE |
+----+--------+-------+
|  1 | apple  |    15 |
|  2 | banana |    10 |
+----+--------+-------+
```

Instead of reading from stdin we can also specify a file using `-i` or `--input-file`:
```console
$ table --input-file testfiles/sample.csv
+----+--------+-------+
| ID |  NAME  | PRICE |
+----+--------+-------+
|  1 | apple  |    15 |
|  2 | banana |    10 |
+----+--------+-------+
```

## Limitations
### Ordering in JSON results
When using a JSON document as input, the headers are sorted alphabetically. This is due to the usage of
`map[string]string` when un-marshalling, which otherwise gives non-deterministic results.

### JSON document format
The JSON documents needs to contain a list as the top-level structure and non-nested dictionaries as elements of this
list.
