package main
import "fmt"
func main() {const T=float64(1712614916);const M=29.53059*24*60*60;var g float64;fmt.Scan(&g);d:=g-T;for d<0{d=+M;};res:=d/M*360;fmt.Printf("%.2f",res);}
