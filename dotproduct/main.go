package main
import (
	. "encoding/json"
	. "fmt"
	. "os"
)
func main(){var U []byte;var V []byte;var u []int;var v []int;Scanln(&U);Scanln(&V);e:=Unmarshal(U,&u);if e!=nil{Println("error: ",e)};e=Unmarshal(V,&v);if e!=nil{Println("error: ",e)};l:=len(u);if l!=len(v){Println("Vectors are different sizes");Exit(1)};s:=0;for i:=0;i<l;i++{s=s+u[i]*v[i]};Println(s)}
