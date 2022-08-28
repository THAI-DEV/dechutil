package dechutil

import "strings"

func SeparateGroupSlice(inputSlice []string, groupNum int, isAsymmetryMember bool) []string {
	/*
		input := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09"}

		output := SeparateGroupSlice(input, 2, true)
		Result : group1 = 01 , 02 , 03 , 04 , 05
				 group2 = 06 , 07 , 08 , 09

		output := SeparateGroupSlice(input, 2, false)
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
				list := convertStringParam2Slice(
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

func DistributeGroupSlice(inputSlice []string, groupNum int) []string {
	/*
		input := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09"}

		output := DistributeGroupSlice(input, 2)
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

func convertStringParam2Slice(strParam string,
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
