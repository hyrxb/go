package main

import "fmt"

func main(){
	fmt.Println(" 4 > 3:结果：",4 > 3) //4 > 3:结果： true
	fmt.Println("４!=３：结果:",4!=3)

	fmt.Println("!( 4 > 3 ) 结果：",!(4 >3) )

	fmt.Println("!(4 != 3) 结果:",!(4!=3))

	//&&

	fmt.Println("true && ture 结果：",true && true)
	fmt.Println("true && false 结果：",true && false)

    // ||

    fmt.Println("true || true 结果",true || true)
	fmt.Println("true || false 结果",true || false)



	/**
	 4 > 3:结果： true
	４!=３：结果: true
	!( 4 > 3 ) 结果： false
	!(4 != 3) 结果: false
	true && ture 结果： true
	true && false 结果： false
	true || true 结果 true
	true || false 结果 true
	 */
}