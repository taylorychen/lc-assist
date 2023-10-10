package api

import (
	"encoding/json"
	"fmt"

	"lc-assist/utils"
	"net/http"
	"net/url"
	"strconv"

	_ "github.com/lib/pq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db, err := utils.Connect()
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, "unable to connect to db", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// 	http.Error(w, "unable to ping db", http.StatusInternalServerError)
	//  return
	// }
	// fmt.Println("connected")

	// TODO: idk if this is right
	vals, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		msg := fmt.Sprintf("error parsing query parameters: %s", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	numStr, ok := vals["amount"]
	if !ok {
		http.Error(w, "missing parameter 'amount'", http.StatusBadRequest)
		return
	}
	num, err := strconv.Atoi(numStr[0])
	if err != nil {
		http.Error(w, "input is not a number", http.StatusBadRequest)
		return
	}

	query := "SELECT * FROM problems WHERE id <= $1"

	rows, err := db.Query(query, num)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, "query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var questions []utils.Question
	for rows.Next() {
		var id int
		var title string
		var titleSlug string
		var difficulty string
		var isPremium bool

		err = rows.Scan(&id, &title, &titleSlug, &difficulty, &isPremium)
		if err != nil {
			fmt.Printf("%v\n", err)
			http.Error(w, "access failed", http.StatusInternalServerError)
			return
		}
		// fmt.Println(id, title, titleSlug, difficulty, isPremium)
		questions = append(questions, utils.Question{Id: id, Title: title, TitleSlug: titleSlug, Difficulty: difficulty, IsPremium: isPremium})
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, "access failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	out, err := json.Marshal(questions)
	if err != nil {
		fmt.Printf("%v\n", err)
		http.Error(w, "unable to convert output to json", http.StatusInternalServerError)
		return
	}
	w.Write(out)

}
