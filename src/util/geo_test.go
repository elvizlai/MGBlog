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
	b := []byte{5, 6, 7}
	fmt.Println(b)
	wm(&b)
	fmt.Println(b)

}

func wm(picBytes *([]byte)) {
	n := []byte{1, 2, 3, 4}
	pic := append(*picBytes,n...)
	picBytes = &pic
	fmt.Println(picBytes)
}