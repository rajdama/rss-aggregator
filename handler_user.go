package main

import "net/http"

/*
We make this function/handler part of apiconfig because we have to also pass reference to aur datbase in order
to interact with it and the function signature of handler i.e number of arguments and return type cannot
be changed hence we pass our refernce to database by adding this function to apiconfig struct
*/

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
