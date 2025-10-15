package main

func isSubsequence(s string, t string ) bool {
    if s == ""{
        return true
    }
	pos :=0
	count :=0
	for i:=0 ; i<len(t); i++{
		if t[i] == s[pos]{
			count++
			pos++
			if count==len(s){
				return true 
			}
		}
	}
	return false
}
