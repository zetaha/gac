import (
        "fmt"
        "math/rand"
        "net/http"

        "github.com/gorilla/mux"
)

var phrases []string

func GAC(w http.ResponseWriter, r *http.Request) {
        var idx = 0
        if len(phrases) > 1 {
                idx = rand.Intn(len(phrases) - 1)
        }
        fmt.Fprintf(w, "<div style='font-size:40;text-align:center;margin-top:10%;'>"+phrases[idx]+" - GAC</div>", r.URL.Path[1:])
}

func main() {
        phrases = append(phrases, "Sono uscito sotto la pioggia e mi sono bagnato")
        phrases = append(phrases, "Non mi sono messa la crema solare e mi sono ustionata")
        phrases = append(phrases, "Mi sono rasato la barba e sembro più giovane")
        phrases = append(phrases, "Siamo andati a Luglio in Marocco e siamo morti di caldo")
        phrases = append(phrases, "A Roma era pieno di turisti")
        phrases = append(phrases, "Sono andato a Trastevere di Sabato sera ma non ho trovato posto per la macchina")
        phrases = append(phrases, "Siamo andati in discoteca ma la musica era troppo alta")
        router := mux.NewRouter()
        router.HandleFunc("/", GAC).Methods("GET")

        err := http.ListenAndServe(":8000", router)
        if err != nil {
                fmt.Println(err)
        }
}

