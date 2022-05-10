package gslice

// SliceStrTo convert string slice to interface{} slice
func SliceStrTo(sts []string) []interface{} {
	uns := make([]interface{}, len(sts))
	for _, v := range sts {
		uns = append(uns, v)
	}
	return uns
}

// SliceInt64To convert int64 slice to interface{} slice
func SliceInt64To(sts []int64) []interface{} {
	uns := make([]interface{}, len(sts))
	for _, v := range sts {
		uns = append(uns, v)
	}
	return uns

}

// StrSliceContains ....
func StrSliceContains(sl []string, str string) bool {
	for _, s := range sl {
		if s == str {
			return true
		}
	}
	return false
}

// Int64SliceContains ...
func Int64SliceContains(sl []int64, i int64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}
