package main

import "fmt"

func main(){
	var s []string
	var i []int
	s[0]="I"
	s[1]="'"
	s[2]="m"
	s[3]="go"
	s[4]="and"
	s[5]="Miku"
	s[6]="me"
	i[0]=1
	if i[0]==1 {
		fmt.Println(s[0]+s[1]+s[2]+s[3])
	}
	i[1]=2
	if i[1]==2 {
		fmt.Println(s[0]+s[1]+s[2]+s[6])
	}
	i[2]=3
	if i[2]==3{
		fmt.Println(s[0]+s[1]+s[2]+s[5])
	}
	i[3]=4
	if i[3]==4{
		fmt.Println(s[0]+s[1]+s[2]+s[3]+s[4]+s[5])
	}
}
