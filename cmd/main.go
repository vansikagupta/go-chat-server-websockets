package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	for{
		text,err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		
		text = strings.Replace(text, "\r\n", "", -1)
		fmt.Println(text)

		if strings.Compare("bye", text) == 0{
			fmt.Println("You said Bye!")
			break
		}
	}
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan(){
		text := scanner.Text()
		fmt.Println(text)

		if strings.Compare("bye", text) == 0{
			fmt.Println("You said Bye!")
			break
		}
	}
}