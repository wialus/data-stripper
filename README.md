# data-stripper

A tool to remove fields from CSV or ndjson formatted files.

## Building manually

    docker build -f .docker/app/Dockerfile .

## CLI arguments

* `--in` - path to an input file (must have either `.json` or `.csv` extension).
* `--out` - path to an output file. Must not exist.
* `--field` - one or more field filters (syntax is specific to the input data format).

## Running

    docker run --rm -u $(id -u):$(id -g) -v $(pwd):/data wialus/data-stripper:1.0.0 --in=/data/in.json --out=/data/out.json --field='$.field_name'

## --field syntax for NDJSON

A subset of [JSONPath](https://goessner.net/articles/JsonPath/) spec is supported:

* `$` to annotate the root object.
* `.` operator to access child properties.
* `*` to annotate all elements in an array.

Examples:

* `$.foo` would remove `foo` from every record in a file

        {"foo": "bar", "baz": 42}

    would become

        {"baz": 42}

* `$.foo.*.bar` would remove the `bar` property from evey object inside the `$.foo` array

        {"foo": [{"one": "two"}, {"bar": "three"}]}

    would become

        {"foo": [{"one": "two"}, {}]}

## --field syntax for CSV

The column name to remove.
