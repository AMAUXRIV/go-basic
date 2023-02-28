

# Property modifier public dan private 👓
Kapan sebuah struct, fungsi, atau method bisa diakses dari package lain dan kapan tidak.
Dalam sebuah package pasti kita menuliskan banyak hal seperti var, struct, looping .etc

jika dalam 1 program terdapat lebh 1 package atau ada package lain selain main.go 
maka komponen dalam package lain tersebut tidak bisa diakses secara bebas dari file yang ber packge main karena tiap komponen memiliki hak akses

di golang ada 2 akses yakni :
exported / public > bebas digunakan dalam package lain
unexported / private > komponen hanya bisa diakses dalam package yang sama , bisa 1 file ataupun 1 folder

penentuan ini sangat penting

di golang character case ialah penanda jika diawali dengan capital (Ujang) maka dia public
jika low (ujang) maka private 

