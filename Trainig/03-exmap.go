package main

import "fmt"

func main() {
	fname := make(map[string]string)
	fname["fateme"] = "Fateme"
	fname["parto"] = "Parto"
	fname["shima"] = "Shima"
	fmt.Println(fname)
	salary := map[string]int{
		"Alishah":  40000,
		"Mehrabi":  30000,
		"Rezazade": 25000,
	}
	for key, value := range salary {
		fmt.Printf("salary[%s] = %d\n", key, value)
	}
	fmt.Println(salary)
	delete(salary, "Rezazade")
	fmt.Println(salary)
	fmt.Println("=====================================")
	score := map[string]int{
		"math":      76,
		"chemistry": 40,
		"physics":   90,
	}

	fmt.Println(score)
	lname := map[string]string{
		"danesh":    "Danesh",
		"farokhzad": "Farokhzad",
		"graei":     "Graei",
	}
	fmt.Println(lname)
	value, ok := lname["danesh"]
	if ok == true {
		fmt.Println("the value of", value, "is available")
		return
	}
	fmt.Println("not found")

}
