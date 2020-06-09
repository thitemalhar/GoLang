//package main
//
//import (
//	"encoding/json"
//	"fmt"
//)
//
//type agent struct {
//	First string `json:"First"`
//	Last  string `json:"Last"`
//	Age   int    `json:"Age"`
//}
//
//func main() {
//	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Iron","Last":"Man","Age":42}]`
//	bs := []byte(s)
//	fmt.Printf("%T\n", s)
//	fmt.Printf("%T\n", bs)
//
//	//agents := []agent {}
//	var agents []agent
//	err := json.Unmarshal(bs, &agents)
//	if err != nil {
//		fmt.Println("err:", err)
//	}
//
//	fmt.Println("\n All of the data: ", agents)
//
//	for _, v := range agents {
//		fmt.Println(v.First, v.Last, v.Age)
//	}
//
//}


package main

import (
	"encoding/json"
	"fmt"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	err := json.Unmarshal(byt, &dat)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dat)

	num := dat["num"]
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str0 := strs[0].(string)
	str1 := strs[1].(string)
	fmt.Println(str0)
	fmt.Println(str1)

	fmt.Println("****************************************")

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"Malhar", "Kavita", "Ojas"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	fmt.Println("****************************************")


	str := `{"page": 1, "fruits": ["apple", "peach"]}`

	//var res []response2

	res := response2{}

	json.Unmarshal([]byte(str), &res)

	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}

//provisioning RDS
// create the tables