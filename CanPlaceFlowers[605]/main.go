package main

func canPlaceFlowers(flowerbed []int, n int ) bool{
    if n<=0 {
        return true
    }
    if len(flowerbed)==1 && flowerbed[0]==0 && n==1{
        return true
    }
	for i:=0 ; i<len(flowerbed)-1; i++{
        if i ==0 && flowerbed[i]==0 && flowerbed[i+1]==0{
            n--
			if n<=0{
				return true
			}
            continue
        } 

        if i==len(flowerbed)-2 && flowerbed[i]==0 && flowerbed[i+1]==0{
            n--
			if n<=0{
				return true
			}
            continue
        } 

		if flowerbed[i]==0 && flowerbed[i+1]==0 && flowerbed[i+2]==0{
			n--
			if n<=0{
				return true
			}
			i++
		}
	}
	
	return false
}
