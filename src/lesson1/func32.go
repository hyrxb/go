package lesson1

import "fmt"

func main(){
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "冬季"
	countryCapitalMap["India"] = "新德里"

	for country := range countryCapitalMap{
		fmt.Println(country,"首都是",countryCapitalMap[country])
	}

	captical,ok := countryCapitalMap["美国"]

	if(ok){
		fmt.Println("美国的首都是",captical)
	}else{
		fmt.Println("美国的首都都不存在")
	}
}
