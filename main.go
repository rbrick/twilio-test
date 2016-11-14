package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var randomResponses = []string{"Chicken Strips", "Hotdog", "Cheeseburger"}

const (
	niceMemeResponse = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Play>http://niceme.me/nicememe.mp3</Play>
</Response>`

	randomMessageResponse = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="woman">Hello, I like %s.</Say>
</Response>`

	wubbaLubbaDubDub = `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Play>http://test.rcw.io/wubba</Play>
</Response>`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/wubba", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "static/wubbalubbadubdub.mp3")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		test := rand.Intn(35) % 2
		if test%2 != 0 {
			fmt.Fprint(w, wubbaLubbaDubDub)
		} else if test%2 == 0 {
			fmt.Fprint(w, niceMemeResponse)
		} else {
			fmt.Fprintf(w, randomMessageResponse, randomResponses[rand.Intn(len(randomResponses))])
		}
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
