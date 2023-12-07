package dechutil

type UnitByteType int

const (
	TB UnitByteType = iota
	GB
	MB
	KB
)

// 1TB = 1*1024*1024*1024*1024
// 1GB = 1*1024*1024*1024
// 1MB = 1*1024*1024
// 1KB = 1*1024

func ByteToUnitByte(numValue uint64, unit UnitByteType) float64 {
	var divisor float64

	switch unit {
	case 0:
		divisor = 1024.0 * 1024.0 * 1024.0 * 1024.0
	case 1:
		divisor = 1024.0 * 1024.0 * 1024.0
	case 2:
		divisor = 1024.0 * 1024.0
	case 3:
		divisor = 1024.0
	}

	result := float64(numValue) / divisor
	// s := dechutil.FormatCommaFloat(f, numDecimal)

	return result
}

func UnitByteToByte(numValue float64, unit UnitByteType) uint64 {
	var divisor float64

	switch unit {
	case 0:
		divisor = 1024.0 * 1024.0 * 1024.0 * 1024.0
	case 1:
		divisor = 1024.0 * 1024.0 * 1024.0
	case 2:
		divisor = 1024.0 * 1024.0
	case 3:
		divisor = 1024.0

	}

	result := numValue * divisor

	return uint64(result)
}
