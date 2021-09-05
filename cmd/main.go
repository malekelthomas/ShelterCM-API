package main

import "github.com/malekelthomas/ShelterCM-API/pkg/router"

func main() {
	s := router.NewServer()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}

}
