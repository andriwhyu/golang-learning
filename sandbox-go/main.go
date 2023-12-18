package main

import "fmt"

//func computeHash(b []byte, seed uint32) uint32 {
//	hash := fnv.New32a()
//
//	fmt.Println(hash)
//	// the implementation fnv.Write() does not return an error, see hash/fnv/fnv.go
//	_, _ = hash.Write(i32tob(seed))
//	_, _ = hash.Write(b)
//	return hash.Sum32()
//	//return 0
//}
//
//// i32tob converts a seed to a byte array to be used as part of fnv.Write()
//func i32tob(val uint32) []byte {
//	r := make([]byte, 4)
//	for i := uint32(0); i < 4; i++ {
//		r[i] = byte((val >> (8 * i)) & 0xff)
//	}
//	return r
//}

//const (
//	gigaBytesConverter = 1073741824
//	terraByteConverter = 1099511627776
//)
//
//func byteToGigabytes(volume float32) (string, float32) {
//	convertedVolume := volume / gigaBytesConverter
//	return "Gigabyte", convertedVolume
//}
//
//func byteToTerabyte(volume float32) (string, float32) {
//	convertedVolume := volume / terraByteConverter
//
//	return "Terabyte", convertedVolume
//}
//
//func teraBytetoByte(volume float32) float32 {
//	convertedVolume := volume * terraByteConverter
//	return convertedVolume
//}
//
//func gigaBytetoByte(volume float32) float32 {
//	convertedVolume := volume * gigaBytesConverter
//	return convertedVolume
//}

func main() {
	//var b []byte = []byte{97}
	//fmt.Println(computeHash(b, 32))
	//var abc uint32
	//
	//abc = uint32(50 * 163.84)
	//
	//fmt.Println(abc)
	//
	//s := os.Args[1]
	//repeatText := strings.Repeat("1", 3)
	//fmt.Println(s, repeatText)
	//
	//for i := 0; i < 100; i++ {
	//	ib := byte(i)
	//	traceID := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, ib, ib, ib, ib, ib, ib, ib, ib}
	//
	//	fmt.Println(traceID)
	//}

	fmt.Println([]byte("accesslog.justice.var.log.containers.justice-platform-service-79b6887bfb-24nbh_justice_service-42aa1aad82cef9e9aa4c4d36e8610e9d507a7f467da4f7b6bea89fc95768d29d.log"))
	fmt.Println([]byte("accesslog.justice.var.log.containers.justice-iam-service-bdbdb6657-79d52_justice_service-81c67c88ae98a774fbacc9ffa9d41323b0e5163d0c2a8f1fb1620b4931be70bf.log"))
	//var (
	//	currentVolumePerHourinByte float32
	//	targetVolumePerHourinByte  float32
	//	reducePercentage           float32
	//)
	//
	//volumePermonthTerrabyte := float32(160.0)
	//reducePercentage = float32(30)
	//
	//currentVolumePerHourinByte = teraBytetoByte(volumePermonthTerrabyte) / 730
	//
	//fmt.Println("---Current Volume per hour---")
	//dataType, volume := byteToGigabytes(currentVolumePerHourinByte)
	//fmt.Println(volume, dataType)
	//
	//fmt.Println()
	//fmt.Println("---Target Volume per hour---")
	//fmt.Println("Decreasing percentage:", reducePercentage)
	//targetVolumePerHourinByte = currentVolumePerHourinByte - (currentVolumePerHourinByte * reducePercentage / 100)
	//dataType, volume = byteToGigabytes(targetVolumePerHourinByte)
	//fmt.Println("Total Target Volume:", volume, dataType)
	//fmt.Println(157.0 / 4.0 / 113.0)
	//fmt.Println(byteToTerabyte(gigaBytetoByte(157 * 730)))

	//text := "promtail-2xlsv_monitoring_promtail-caf67b3d797f53eaecff37cc06e0bdb7986328a9623f42e849e945c2b107aecc"
	//regex := regexp.MustCompile("(_(.+)(<=|promtail|postgresql12|fluentd|kafka|mongodb-([a-zA-Z0-9]))(.+))")
	//fmt.Println(regex.ReplaceAllString(text, ""))

}
