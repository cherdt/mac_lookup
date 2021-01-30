package main
import ( 
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"os"
)

func main() {

	type MacAddress struct {
                Mac string "mac"
                Company string "company"
		Address string "address"
		MacStart string "mac_start"
		MacEnd string "mac_end"
		MacType string "type"
		Country string "country"
        }
	var m MacAddress

	resp, err := http.Get("https://api.maclookup.app/v2/macs/" + os.Args[1] + "/company/name")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	    fmt.Println(string(body))
	}

	if json.Valid(body) != true {
		fmt.Println("Response was not valid JSON")
	} else {
		// Unmarshal (deserialize) the JSON response
		jsonerr := json.Unmarshal(body, &m)
		if jsonerr != nil {
			fmt.Println(jsonerr)
		} else {
			fmt.Println(m.Company)
		}
	}
}
