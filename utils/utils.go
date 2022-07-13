package utils

func CheckSliceEquality [K comparable] (arr1, arr2 []K) bool {
	if arr1 == nil || arr2 == nil || len(arr1) != len(arr2){
		return false 
	}
	
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}