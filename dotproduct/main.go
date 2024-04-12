package main
import (
	"encoding/json"
	"fmt"
	"os"
)
func main(){
	var ub []byte
	var vb []byte
	var u []int
	var v []int
	fmt.Scanln(&ub)
	fmt.Scanln(&vb)
	e:=json.Unmarshal(ub,&u)
	if e!=nil{fmt.Println("error: ",e)}
	e=json.Unmarshal(vb,&v)
	if e!=nil{fmt.Println("error: ",e)}
	vs:=len(u)
	if vs!=len(v){fmt.Println("Vectors are different sizes");os.Exit(1)}
	s:=0
	for i:=0;i<vs;i++{s=s+u[i]*v[i]}
	fmt.Println(s)
}
