# geo-benthos
A benthos geo plugin to transform coordinates


## Example usages

`position.json` contains position:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

## Transform latitude and longitude into an Uber h3 cell

Use `h3_object` with the following parameters: `latitude`, `longitude`, `resolution`.

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
      root.h3 = h3_object(this.lat, this.lng, 5)

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

Use `s2_object` with the following parameters: `latitude`, `longitude`, `resolution`.

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
