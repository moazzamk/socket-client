package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"socket-client/socketclient"
	"log"
)


func ShowForm(rs http.ResponseWriter, req *http.Request, param httprouter.Params) {
	var errMessage string
	var ret string
	var err error

	if req.FormValue(`host`) != `` {
		ret, err = socketclient.GetResponse(
			req.FormValue(`host`),
			`<ops>` + req.FormValue(`input`) + `</ops>`)

	}

	if err != nil {
		errMessage = err.Error()
	}

	rs.Header().Add(`Content-Type`, `text/html`)
	rs.Write([]byte(`
		<head>
			<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" />
			<style type="text/css">
				form {
					margin: 20 20px;
				}
			</style>
		</head>
		
		<form action="/" method="post">
			<div class="form-group">
				<label for="host"><h4>Host:</h4></label> 
				<input type="text" name="host" id="host" class="form-control" value="` + req.FormValue(`host`) + `"/> <br/>
			</div>
			<div class="form-group">
				<label for="input"><h4>Input:</h4></label> 
				<textarea name="input" rows="10" class="form-control" id="input">` + req.FormValue(`input`) + `</textarea><br/>
			</div>
			<input type="submit" value="GO" />
		</form>

		<pre> ` + ret + errMessage + `</pre>

		<script
  			src="https://code.jquery.com/jquery-3.3.1.min.js"
  			integrity="sha256-FgpCb/KJQlLNfOu91ta32o/NMZxltwRo8QtmkMRdAu8="
  			crossorigin="anonymous"></script>
		<script type="text/javascript" src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js"></script>
		<script type="text/javascript" src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.bundle.min.js"></script>

	`))
}

func main() {
	router := httprouter.New()
	router.GET(`/`, ShowForm)
	router.POST(`/`, ShowForm)

	log.Fatal(http.ListenAndServe(`:8080`, router))
}
