package mainApp


import (
	"os"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"./../database/"

	"strings"
	"github.com/havard/Oblig2/cmd/database"
)

func HandlerRequest(w http.ResponseWriter, r * http.Request) {
	//mongodb://<dbuser>:<dbpassword>@ds141274.mlab.com:41274/cloudimt2681

	url := strings.Split(r.URL.Path, "/")
	if len(url) != 3 {
		http.Error(w, "Splitscreen sadness", 400)
	}else if url[2] == "" {
		http.Error(w, "this is Sparta!", 400)
	}else{
		db := database.MgoDB {
			"mongodb://admin:imt2681@dds141274.mlab.com:41274/cloudimt2681",
			"cloudimt2681",
			"currency",
			"webhooks",
		}
		switch r.Method {
		case "GET":
			res, err := db.GetWebHook(url[2])
			if err != nil {
				fmt.Printf("something went wrong with return value of GetWebHook func %v", err)
			}
			json.NewEncoder(w).Encode(res)

		default:
			http.Error(w, "you have joined the darkside, darthVader is your father now", 400)
		}
	}
}
func RegistrationOfNewWebHook(w http.ResponseWriter, r * http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Something went wrong with reading webhook input %v", err)
	}
	defer r.Body.Close()
	payUrDues := database.Webhookers{}
	err = json.Unmarshal(body, &payUrDues)
	if err != nil {
		fmt.Printf("somethin went wrong in externalInput/Unmarshal body to Payload %v", err)
	}

	mydb := database.MgoDB {
		"mongodb://admin:imt2681@dds141274.mlab.com:41274/cloudimt2681",
		"cloudimt2681",
		"currency",
		"webhooks",
	}
	Hid, err := mydb.AddWebHook(payUrDues)
	if err != nil {
		fmt.Printf("Returned error from AddWebhook func %v", err)
	}
	fmt.Printf("WebHookId was added %v", Hid)

}
func LatestCurrency(w http.ResponseWriter, r * http.Request) {
	/*content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("something went wrong %v", err)
	}
	defer r.Body.Close()
	tmp := database.Webhookers{}
*/
fmt.Fprintln("Not Implementet yet")

}
func AverageCurrency(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln("Not Implementet yet")

}

func Addemdum(w http.ResponseWriter, r * http.Request) {


}


func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/root", HandlerRequest)
	http.HandleFunc("/root", RegistrationOfNewWebHook)
	http.HandleFunc("/root/latest", LatestCurrency)
	http.HandleFunc("/root/average", AverageCurrency)
	http.HandleFunc("/root/evaluationtrigger", Addemdum)


	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}