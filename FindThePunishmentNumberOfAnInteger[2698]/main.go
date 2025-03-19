package main


func main (){}

func punishmentNumber(n int) int{
  values:=[]int{1,9,10,36,45,55,82,91,99,100,235,297,369,370,379,414,657,675,703,756,792,909,918,945,964,990,991,999,1000}
  sum :=0

  for _,num := range values{
    if n < num{
      return sum
    }
    sum+= num*num
  }
  return 0
}
