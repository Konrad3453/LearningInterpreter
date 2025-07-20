package main

import "interpreter/token"

func main() {
	// Create a token to test the package
	t := token.Token{
		Type:    token.IDENT,
		Literal: "test",
	}
	
	// Print some token information
	println("Token Type:", string(t.Type))
	println("Token Literal:", t.Literal)
	println("Available token types:")
	println("ILLEGAL:", token.ILLEGAL)
	println("EOF:", token.EOF)
	println("FUNCTION:", token.FUNCTION)
}
