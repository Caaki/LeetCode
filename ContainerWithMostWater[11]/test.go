package main


func maxArea2(height []int) int {
	
	l:=0
	r:=len(height)-1
	maxVal:=0
	
	for l<=r {
		if height[l] == 0{
			l++
			continue
		}
		if height[r]  == 0{
			r--
			continue
		}
		
		nVal :=calculateContent(height[l], height[r], r-l)
		if maxVal < nVal{
			maxVal = nVal
		}
        	if height[l] < height[r] {
		        l++
	        }else {
		        r--
	    }
	}

	return maxVal
}

func calculateContent(a, b, w int ) int {
	if a < b {
		return a*w
	}else {
		return b*w
	}
}
