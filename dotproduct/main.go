package main
import (
	"encoding/json"
	"fmt"
	"os"
)
func main(){var U []byte;var V []byte;var u []int;var v []int;fmt.Scanln(&U);fmt.Scanln(&V);e:=json.Unmarshal(U,&u);if e!=nil{fmt.Println("error: ",e)};e=json.Unmarshal(V,&v);if e!=nil{fmt.Println("error: ",e)};l:=len(u);if l!=len(v){fmt.Println("Vectors are different sizes");os.Exit(1)};s:=0;for i:=0;i<l;i++{s=s+u[i]*v[i]};fmt.Println(s)}
