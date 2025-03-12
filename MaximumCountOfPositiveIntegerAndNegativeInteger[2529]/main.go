package main

func man (){
  

}
func maximumCount(nums []int) int {
  l,r, found := binarySearch(nums)
  if l==-1 || r==-1{
    return len(nums)
  }
  if found {
      if l > len(nums)-r-1{
    return l
  }else{
    return len(nums)-r-1
  }
  }

  if l > len(nums)-r{
    return l+1
  }else{
    return len(nums)-r
  }

}

func binarySearch(nums []int) (l int, r int,found bool ){
  l=0
  r= len(nums)-1
  found = false
  for l<=r{
    mid := l+(r-l) /2
    if nums[mid] == 0{
      l=mid
      r=mid
      found =true
      break
    }
    if nums[mid] < 0{
      l = mid+1
    }else{
      r = mid-1
    }
  }

  if !found && (nums[0]>0 || nums[len(nums)-1] < 0){
    if nums[0]<0{
      return -1,-1,found
    }else{
      return -1,-1,found
    }
  }

  finl:=false
  finr:=false
  if found {
  for !finl|| !finr{
    if l>=1 && nums[l-1] == 0{
      l-=1
    }
    if r < len(nums)-1 && nums[r+1]==0{
      r+=1
    }
    if l==0 || nums[l-1] !=0{
      finl=true
    }
    if r==len(nums)-1 || nums[r+1] !=0{
      finr=true
    }
  }
  }else{
    for !finl|| !finr {
    if l>=1 && nums[l] >0{
      l-=1
    }
    if r < len(nums)-1 && nums[r]<0{
      r+=1
    }
    if nums[l]<0 || l==0{
      finl=true
    }
    if r==len(nums)-1 || nums[r] >0{
      finr=true
    }
    }
      return l,r,found
  }
  return l,r,found

}
