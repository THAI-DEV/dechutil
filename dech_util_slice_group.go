package dechutil

import (
	"strconv"
	"strings"
)

// Flow arrang group for general slice
// Step 1 : ConvertAnySliceToIndexString
// Step 2 : Arrang group as ...
// Step 3 : TransformGroupDimension
// Step 4 : Loop slice to use index string

func ArrangGroupAsSeparate(inputSlice []string, groupNum int, isAsymmetryMember bool) []string {
	/*
		input := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09"}

		output := ArrangGroupAsSeparate(input, 2, true)
		Result : group1 = 01 , 02 , 03 , 04 , 05
				 group2 = 06 , 07 , 08 , 09

		output := ArrangGroupAsSeparate(input, 2, false)
		Result : group1 = 01 , 02 , 03 , 04 , 09
				 group2 = 05 , 06 , 07 , 08

	*/
	if groupNum > len(inputSlice) {
		groupNum = len(inputSlice)
	}

	totalMember := len(inputSlice)
	memberInGroup := totalMember / groupNum
	beyondMember := totalMember % groupNum

	resultList := []string{}

	if memberInGroup > 0 && len(inputSlice) >= groupNum {
		str := ""
		ind := 0
		for i := 0; i < groupNum; i++ {
			str = ""

			for j := 0; j < memberInGroup; j++ {
				if j < memberInGroup-1 {
					str = str + inputSlice[ind] + " , "
				} else {
					str = str + inputSlice[ind]
				}

				ind++
			}

			resultList = append(resultList, str)

		}

		if beyondMember > 0 {
			for i := 0; i < beyondMember; i++ {
				for j := 0; j < groupNum; j++ {
					if ind == len(inputSlice) {
						break
					}
					str = resultList[j] + " , " + inputSlice[ind]
					ind++
					resultList[j] = str
				}

			}
		}

		if beyondMember > 0 && isAsymmetryMember {
			newResultList := []string{}
			ind := 0

			for i := 0; i < len(resultList); i++ {
				str := ""
				list := convertStringParamToSlice(
					resultList[i], ",")
				for j := 0; j < len(list); j++ {
					if j < len(list)-1 {
						str = str + inputSlice[ind] + " , "
					} else {
						str = str + inputSlice[ind]
					}

					ind++
				}
				newResultList = append(newResultList, str)

			}

			resultList = newResultList
		}

	}

	return resultList
}

func ArrangGroupAsDistribute(inputSlice []string, groupNum int) []string {
	/*
		input := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09"}

		output := ArrangGroupAsDistribute(input, 2)
		Result : group1 = 01 , 03 , 05 , 07 , 09
				 group2 = 02 , 04 , 06 , 08
	*/
	resultList := []string{}

	if groupNum > len(inputSlice) {
		groupNum = len(inputSlice)
	}

	ind := 0
	if groupNum > 0 && len(inputSlice) >= groupNum {
		for i := 0; i < groupNum; i++ {
			resultList = append(resultList, inputSlice[ind])
			ind++
		}

		str := ""
		for i := 0; i < len(inputSlice); i++ {
			for j := 0; j < groupNum; j++ {
				if ind == len(inputSlice) {
					break
				}
				str = resultList[j] + " , " + inputSlice[ind]
				ind++

				resultList[j] = str
			}

		}

	}

	return resultList
}

// one dimension slice to two dimension slice
func TransformGroupDimension(inputSlice []string) [][]string {
	result := [][]string{}

	for _, v := range inputSlice {
		rowData := strings.Split(v, ",")
		result = append(result, rowData)
	}

	return result
}

// convert data slice to index string slice
func ConvertAnySliceToIndexString[T sliceInf](inputSlice []T) []string {
	result := []string{}

	for i := range inputSlice {
		result = append(result, strconv.Itoa(i))
	}

	return result
}

func convertStringParamToSlice(strParam string,
	delimiter string) []string {
	/*
	 strMsg = "DM-01-01-01-01, MPS001"
	 delimiter = ","
	 resukt = [DM-01-01-01-01] [MPS001]

	*/

	resultList := []string{}

	strArr := []string{}
	if len(strParam) > 0 && len(delimiter) > 0 {
		strParam = strings.TrimSpace(strParam)
		delimiter = strings.TrimSpace(delimiter)
		strArr = strings.Split(strParam, delimiter)
	}

	if len(strArr) > 0 {
		for _, s := range strArr {
			if len(s) > 0 && len(strings.TrimSpace(delimiter)) > 0 {
				resultList = append(resultList, strings.TrimSpace(s))
			}
		}

	}

	return resultList
}

// func computeIndexGroupSlice(numLen int, numStep int) [][]int {
// 	/*
// 		input := []int{1,2,3,4,5,6,7,8,9,10}

// 		info := ComputeIndexGroupSlice(len(input), 6)
// 		for i := 0; i < len(info); i++ {
// 			beginInd := info[i][0]
// 			endInd := info[i][1]
// 			data := input[beginInd:endInd]
// 			fmt.Println(data)
// 		}

// 		output
// 			[1 2 3 4 5 6]
// 			[7 8 9 10]
// 	*/
// 	f := float64(numLen) / float64(numStep)
// 	f = math.Ceil(f)
// 	numRow := int(f)

// 	result := [][]int{}
// 	var a = [2]int{0, 0}
// 	beginInd := 0
// 	endInd := 0
// 	for i := 0; i < numRow; i++ {
// 		factor := i * numStep
// 		a[0] = beginInd + factor
// 		a[1] = endInd + numStep + factor
// 		if a[1] > numLen {
// 			a[1] = numLen
// 		}
// 		result = append(result, []int{a[0], a[1]})
// 	}

// 	return result
// }
