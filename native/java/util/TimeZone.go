package util

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_tz(getSystemTimeZoneID, "getSystemTimeZoneID", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;")
}

func _tz(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/util/TimeZone", name, desc, method)
}

// private static native String getSystemTimeZoneID(String javaHome, String country);
// (Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;
func getSystemTimeZoneID(frame *rtda.Frame) {
	//javaHomeObj := frame.GetRefVar(0)
	countryObj := frame.GetRefVar(1)

	//for osx read system timezone
	file, err := os.Open("/usr/share/zoneinfo/zone.tab")

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var timezone string
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '#' {
			continue
		}
		zone := strings.Split(line, "\t")
		if zone[0] == heap.JSToGoStr(countryObj) {
			timezone = zone[2]
			break
		}
	}

	location, _ := time.LoadLocation(timezone)
	zoneID := frame.GetRuntime().JSFromGoStr(location.String())
	frame.PushRef(zoneID)
}
