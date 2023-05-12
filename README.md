# geo-benthos
A benthos geo plugin to transform coordinates

## Transform latitude and longitude into a h3 cell

Use `h3_object` with the following params latitude, longitude, resolutiom.

### Example

`position.json` contains position:

```js
{"id":42, "lat": 48.86, "lng": 2.34}
```

An `h3.yaml` pipeline.
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
