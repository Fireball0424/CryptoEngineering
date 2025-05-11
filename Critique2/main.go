package main

import (
	"Critique2/gcm"
	"Critique2/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type PageData struct {
	CipherText string
	Tag        string
	PlainText  string
}

var result PageData

func renderPage(w http.ResponseWriter) {
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, _ := template.ParseFiles(tmplPath)
	tmpl.Execute(w, result)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderPage(w)
	})

	http.HandleFunc("/encrypt", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		IV := r.FormValue("IV")
		plainText := r.FormValue("PlainText")
		AAD := r.FormValue("AAD")

		IVbits := utils.StringToBits(IV, false)
		plainTextBits := utils.StringToBits(plainText, false)
		AADbits := utils.StringToBits(AAD, false)

		CipherTextBits, TagBits := gcm.GCM_AE(IVbits, plainTextBits, AADbits)

		result.CipherText = fmt.Sprint("CipherText: ", utils.BitsToString(CipherTextBits, true), "\n")
		result.Tag = fmt.Sprint("Tag: ", utils.BitsToString(TagBits, true), "\n")

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/decrypt", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		IV := r.FormValue("IV")
		CipherText := r.FormValue("CipherText")
		AAD := r.FormValue("AAD")
		Tag := r.FormValue("Tag")

		IVbits := utils.StringToBits(IV, false)
		CipherTextBits := utils.StringToBits(CipherText, true)
		AADbits := utils.StringToBits(AAD, false)
		Tagbits := utils.StringToBits(Tag, true)

		log.Println(CipherTextBits)

		PlainTextBits := gcm.GCM_AD(IVbits, CipherTextBits, AADbits, Tagbits)

		if PlainTextBits == nil {
			result.PlainText = "Authenticated Fail\n"
		} else {
			result.PlainText = fmt.Sprint("PlainText: ", utils.BitsToString(PlainTextBits, false), "\n")
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.ListenAndServe(":8080", nil)
}
