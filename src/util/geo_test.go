/**
 * Created by elvizlai on 2015/8/28 10:03
 * Copyright © PubCloud
 */
package util
import (
	"testing"
	"fmt"
)


func TestGeo(t *testing.T) {
	result:=InfoGeoByIP("222.175.234.10")
	fmt.Println(result)

	fmt.Println(result.Get("regionName").MustString()+","+result.Get("city").MustString())
	fmt.Println()
	fmt.Println(result.Get("lon").MustFloat64())
	fmt.Println(result.Get("lat").MustFloat64())

}