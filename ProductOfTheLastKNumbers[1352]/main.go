package main

func main() {}

// Ovo resenje radi sa malim brojem brojeva
//type Value struct{
//  num int64
//  sum int64
//}
//type ProductOfNumbers struct{
//  val []Value
//  lastZero int
//  noZeroSum int64
//}
//
//func Constructor() ProductOfNumbers{
//  return ProductOfNumbers{[]Value{},-1,1}
//}
//
//func (this *ProductOfNumbers) Add(num int) {
//  if len(this.val)==0{
//    if num == 0{
//      this.noZeroSum = 1
//      this.lastZero = 0
//      this.val = append(this.val, Value{int64(num), 1})
//    }else{
//        this.noZeroSum = int64(num)
//        this.lastZero = -1
//        this.val = append(this.val, Value{int64(num), int64(num)})
//    }
//  }else{
//    if num ==0 {
//      this.lastZero = len(this.val)
//      this.val = append(this.val, Value{int64(num), this.noZeroSum})
//    }else{
//      this.val = append(this.val, Value{int64(num), int64(num)*this.noZeroSum})
//      this.noZeroSum *= int64(num)
//    }
//  }
//
//}
//
//func (this *ProductOfNumbers) GetProduct (k int) int{
//    if this.lastZero > len(this.val) - k{
//        return 0
//    }
//    if this.val[len(this.val)-k].num == 0{
//        return 0
//    }
//    if len(this.val)-k-1 < 0 && this.lastZero ==-1{
//        return int(this.noZeroSum)
//    }
//    if this.val[len(this.val)-k-1].sum == 0{
//        return 0
//    }
//    return int(this.noZeroSum/this.val[len(this.val)-k-1].sum)
//
//}

//Ovo resenje radi sa svim primerima

type ProductOfNumbers struct {
	val []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{[]int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.val = []int{1}
	} else {
		this.val = append(this.val, this.val[len(this.val)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(this.val) {
		return 0
	}
	return this.val[len(this.val)-1] / this.val[len(this.val)-k-1]
}
