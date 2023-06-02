
package main
	
import (
	"fmt"
	"encoding/json"
)
	
// declaring a struct
type Human struct{
		
	Name string
	Age int
	Address string
}
	

func main() {

	human1 := Human{"Ankit", 23, "New Delhi"}
		
	human_enc, err := json.Marshal(human1)
		
	if err != nil {
			
	
		fmt.Println(err)
	}
		

	fmt.Println(string(human_enc))
		

	human2 := []Human{
		{Name: "Rahul", Age: 23, Address: "New Delhi"},
		{Name: "Priyanshi", Age: 20, Address: "Pune"},
		{Name: "Shivam", Age: 24, Address: "Bangalore"},
	}
		
	// encoding into JSON format
	human2_enc, err := json.Marshal(human2)
		
		if err != nil {
	
			fmt.Println(err)
		}
		
	fmt.Println()
		fmt.Println(string(human2_enc))
}
