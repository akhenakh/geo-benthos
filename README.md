# Geo-Benthos

[Benthos](benthos.dev/) plugins to transform geographic coordinates from a stream.

This repo contains multiple Benthos plugins as Go modules, that you can build on demand (see `cmd/geo-benthos`).

Benthos is the swiss army of stream processing: Benthos solves common data engineering tasks such as transformations, integrations, and multiplexing with declarative and unit testable configuration. 



Note that the h3 plugin is using a [CGO free version](https://github.com/akhenakh/goh3).

## Transform latitude and longitude into an Uber h3 cell

Use `h3` with the following parameters: `latitude`, `longitude`, `resolution`.

An example `position.json`:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

A `h3.yaml` pipeline.

```yaml
input:
  file:
    paths: ["testdata/position.json"]
    codec: all-bytes

pipeline:
  threads: 1
  processors:
  - mapping: |
      root = this
      root.h3 = h3(this.lat, this.lng, 5)

output:
  label: "out"
  stdout:
    codec: lines
```

Enrich the input with the h3 cell:

```sh
go build -o geo-benthos ./cmd/geo-benthos
./geo-benthos -c testdata/h3.yaml
{"h3":"851fb467fffffff","id":42,"lat":48.86,"lng":2.34}
```

## Transform latitude and longitude into a Google s2 cell

Use `s2` with the following parameters: `latitude`, `longitude`, `resolution`.

An example `position.json`:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

A `s2.yaml` pipeline.

```yaml
input:
  file:
    paths: ["testdata/position.json"]
    codec: all-bytes

pipeline:
  threads: 1
  processors:
  - mapping: |
      root = this
      root.s2 = s2_object(this.lat, this.lng, 15)

output:
  label: "out"
  stdout:
    codec: lines
```

Enrich the input with the s2 cell:

```sh
go build -o geo-benthos ./cmd/geo-benthos
./geo-benthos -c testdata/s2.yaml
{"id":42,"lat":48.86,"lng":2.34,"s2":"2/033303031301002"}
```

## Get the timezone for a given latitude and longitude

Use `tz` with the following parameters: `latitude`, `longitude`.

An example `position.json`:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

A `tz.yaml` pipeline.

```yaml
input:
  file:
    paths: ["testdata/position.json"]
    codec: all-bytes

pipeline:
  threads: 1
  processors:
  - mapping: |
      root = this
      root.tz = tz(this.lat, this.lng)

output:
  label: "out"
  stdout:
    codec: lines
```

Enrich the input with the h3 cell:

```sh
go build -o geo-benthos ./cmd/geo-benthos
./geo-benthos -c testdata/tz.yaml
{"tz":"Europe/Paris","id":42,"lat":48.86,"lng":2.34}
```

## Live Testing

Run this command and point your browser to http://localhost:4195/

```sh
./geo-benthos blobl server --no-open --host 0.0.0.0 --input-file ./testdata/position.json -m testdata/s2_mapping.txt   
```


## TODO

- [ ] s2 shape index to perform PIP
- [ ] spatialite lookup to perform PIP
- [ ] random points in a rect
- [X] lat lng to h3
- [X] lat lng to s2
- [X] lat lng to tz