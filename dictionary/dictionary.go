package dictionary

import (
	"bufio"
	"fmt"
	"os"
)

type Word []struct {
	word   string
	length int
	active bool
}

func LoadDictionary() {
	// Open our jsonFile
	dicFile, err := os.Open("assets/test.dic")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened dictionary")
	// defer the closing of our jsonFile so that we can parse it later on
	defer dicFile.Close()
	// read our opened xmlFile as a byte array.

	scanner := bufio.NewScanner(dicFile)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	//json.Unmarshal(byteValue, &users)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	//for i := 0; i < len(users.Users); i++ {
	//	fmt.Println("user Type: " + users.Users[i].Type)
	//	fmt.Println("user Age: " + strconv.Itoa(users.Users[i].Age))
	//	fmt.Println("user Name: " + users.Users[i].Name)
	//	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	//}

}
