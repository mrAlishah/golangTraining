package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sentences *bufio.Reader
func init() {
	sentences = bufio.NewReader(os.Stdin)
}

func input() string {
	str, _ := sentences.ReadString('\n')
	return strings.TrimRight(str, "\n")
}

func main() {
	var (
	 num int
	 name string
	 married bool
	)

//! Scanln ------------------------------------------------------------------------------------- 
	fmt.Println("Scanln: num name married")
	//  stops scanning when it encounters a newline. Enter
	cnt, err := fmt.Scanln(&num, &name , &married) 

	//Closure function : declaration after variables as globale variable
	Print := func (){
		fmt.Println("---------------------------------------------------------------------------")
		fmt.Printf("count:  | %T | %v | %d \n",cnt,cnt,cnt)
		fmt.Printf("error:  | %T | %v  \n",err,err)
		fmt.Printf("num:  | %T | %v | %d \n",num,num,num)
		fmt.Printf("name: | %T | %v | %s \n",name,name,name)
		fmt.Printf("married: | %T | %v | %t \n",married,married,married) //t,1 = true f,0,others chars = false
		fmt.Println("==========================================================================")
		}
	Print()


//! Scanf ------------------------------------------------------------------------------------- 
    fmt.Println("Scanf: 1:num 2:name 3:married")
	// formate rule is necessary = 1:2312 2:june 3:1
	cnt, err = fmt.Scanf("1:%d 2:%s 3:%t",&num, &name , &married) 
	//you can just use fmt.Scan(&num, &name)
	Print()	
	
//! Scan ------------------------------------------------------------------------------------- 
	fmt.Println("Scan: num name married")
	// you can type number Space string or
	// number Enter string
	cnt, err = fmt.Scan(&num, &name , &married) 
	//you can just use fmt.Scan(&num, &name)

    //just for remove scan newline inorder to do correct for input() if remove this line Enter used to finish sentences.ReadString('\n')
	fmt.Scanln() 
	Print()

//! bufio.Reader.ReadString ------------------------------------------------------------------------------------- 
	//when you need more words by space scan is not work because of every word is assigned to scanf
	// This is a sentence.
	// scan output => This 
	// bufio.Reader.ReadString => This is a sentence.
	prf := input()
	fmt.Printf("prf: | %T | %v | %s \n",prf,prf,prf)

}