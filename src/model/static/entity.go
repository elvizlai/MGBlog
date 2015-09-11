/**
 * Created by elvizlai on 2015/8/28 10:29
 * Copyright © PubCloud
 */
package static
import "time"

type Static struct {
	IP        string
	PV        int
	Infered   bool
	City      string
	Geo       [2]float64
	Path      []string
	LastVisit time.Time
}