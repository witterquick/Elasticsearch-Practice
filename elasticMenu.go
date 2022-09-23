import (
	"bufio"
	"fmt"
	"os"
	"github.com/elastic/go-elasticsearch/v8"
)

var es, _ = elasticsearch.NewDefaultClient()

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("0) Exit")
		fmt.Println("1) Load spacecraft")
		fmt.Println("2) Get spacecraft")
		option := ReadText(reader, "Enter option")
		if option == "0" {
			Exit()
		} else if option == "1" {
			LoadData()
		} else {
			fmt.Println("Invalid option")
		}
	}
}