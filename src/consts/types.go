package consts

var TypeMap = map[string]int{
	"tree": 0,
	"color": 1,
	"time": 2,
}

func GetUtilTypeByString(utype string) int {
	if val, ok := TypeMap[utype]; ok {
		return val
	}
	return -1
}

func GetUtilTypeByInt(utype int) string {
	for key, val := range TypeMap {
		if val == utype {
			return key
		}
	}
	return ""
}
