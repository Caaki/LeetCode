package main

func minOperations(s string) int{
	if len(s)==1 {
		return 0
	}
	count1 := 0
	count2 := 0
	v1 := byte('0')
	v2 := byte('1')
	
	for i :=0; i < len(s); i++ {
		if v1 != s[i]{
			count1++
		}

		if v2!=s[i]	{
			count2++
		}
		v1,v2 = v2,v1
	}

	if count1 > count2{
		return count2
	}
	return count1
}
