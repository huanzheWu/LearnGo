package converters

import (
	"strconv"
	"strings"
)

// IntSliceToCSVString converts a slice of integers to a 'comma separated values' string
func IntSliceToCSVString(ints []uint32) (string, error) {

	b := make([]string, len(ints))
	for i, e := range ints {
		b[i] = strconv.FormatUint(uint64(e), 10)
	}

	result := strings.Join(b, ",")

	return result, nil
}

//该函数把 一个用逗号分隔的整数字符串转换成整数切片
func CSVStringToIntSlice(s string) ([]uint32, error) {
	//将,替换成“ ”
	sanitizedString := strings.Replace(s, ",", " ", -1)
	trimmedString := strings.Trim(sanitizedString, " ")
	numbersSlice := strings.Split(trimmedString, " ")
	ints := []uint32{}
	for _, number := range numbersSlice {
		i, err := strconv.ParseUint(number, 32, 0)
		if err != nil {
			return nil, err
		}
		ints = append(ints, uint32(i))
	}

	return ints, nil
}
