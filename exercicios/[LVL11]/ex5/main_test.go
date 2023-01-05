package main

import "testing"  

func TestHello(test *testing.T) {
    checkMessageIsCorrect := func(test *testing.T, result,expected string){
     test.Helper()
     if result != expected {
        test.Errorf("\nResultado: '%s', \nResultado Esperado: '%s'", result, expected)
    }
     

    }
    test.Run("Fala Hello para Gabriel", func(test *testing.T){
        result := Hello("Gabriel","")
        expected := "Hello Gabriel"
        checkMessageIsCorrect(test,result,expected)
    })
    test.Run("'Test' como padrão para String vazia passada por parametro",func(test *testing.T){
        result := Hello("","")
        expected := "Hello Test"
        checkMessageIsCorrect(test,result,expected)
    })
    test.Run("Fala Hello em Português", func(test *testing.T){
        result := Hello("Gabriel",portugueseBrazil)
        expected := "Olá Gabriel"
        checkMessageIsCorrect(test,result,expected)
    })
    test.Run("Fala Hello em Espanhol", func(test *testing.T){
        result := Hello("Gabriel",spanish)
        expected := "Hola Gabriel"
        checkMessageIsCorrect(test,result,expected)
    })
     
	
}
