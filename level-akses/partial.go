/*
Disini akan menjelaskan bagaimana cara mengakses properti yang masih dalam 1 package 
*/

package main 

import f "fmt"


func SayHello(s string) {
	f.Println("Nama saya adalah",s)
}