package pipmem

import (
	"encoding/json"
	"os"

	"github.com/benthosdev/benthos/v4/public/bloblang"
	"github.com/golang/geo/s2"
	"github.com/peterstace/simplefeatures/geom"
)

type Index struct {
	*s2.ShapeIndex
	*s2.ContainsPointQuery
}

func init() {
	pipObjSpec := bloblang.NewPluginSpec().
		Param(bloblang.NewStringParam("path"))

	err := bloblang.RegisterFunctionV2(
		"pipmem", pipObjSpec, func(args *bloblang.ParsedParams) (bloblang.Function, error) {
			path, err := args.GetString("path")
			if err != nil {
				return nil, err
			}

			return func() (interface{}, error) {
				f, err := os.Open(path)
				if err != nil {
					return nil, err
				}

				var fc geom.GeoJSONFeatureCollection

				if json.NewDecoder(f).Decode(&fc); err != nil {
					return nil, err
				}

				return &Index{}, nil
			}, nil
		})
	if err != nil {
		panic(err)
	}

	intoObjectSpec := bloblang.NewPluginSpec().
		Param(bloblang.NewFloat64Param("lat")).
		Param(bloblang.NewFloat64Param("lng"))

	err = bloblang.RegisterMethodV2("query", intoObjectSpec, func(args *bloblang.ParsedParams) (bloblang.Method, error) {
		lat, err := args.GetFloat64("lat")
		if err != nil {
			return nil, err
		}

		lng, err := args.GetFloat64("lng")
		if err != nil {
			return nil, err
		}

		return func(v interface{}) (interface{}, error) {
			p := s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lng))

			shapes := idx.ContainsPointQuery.ContainingShapes(p)
		}, nil
	})
	if err != nil {
		panic(err)
	}
}
