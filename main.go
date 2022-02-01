package main

import (
	"fmt"
	"time"
)

//Affiche les résultats de l'algorithme
func PrintResult(m Mariages, started time.Time, proposants []Agent, disposants []Agent){
	elapsed := time.Since(started)
	fmt.Println()
	fmt.Println("Algorithm Result:")

	m.Debug()
	stable := m.IsStable(proposants, disposants)
	if stable{
		fmt.Println("Result is stable.")
	} else {
		fmt.Println("Result is not stable !")
	}
	fmt.Printf("Execution time: %s \n", elapsed)
}

//Affiche les agents générés
func printAgents(proposants []Agent, disposants []Agent){
	fmt.Println("Agents Generated:");
	fmt.Println()
	fmt.Println("Proposants:")
	for i := 0; i < len(proposants); i++{
		fmt.Println(proposants[i])
	}
	fmt.Println()
	fmt.Println("Disposants:")
	for i := 0; i < len(disposants); i++{
		fmt.Println(disposants[i])
	}
	fmt.Println()
}

//Demande à l'utilisateur de saisir N (taille du problème)
func askUserN() int{
	n := -1
	fmt.Println("Enter n, the number of agent in each group (0 to exit): ")
	fmt.Scan(&n)

	return n
}

//Demande à l'utilisateur quel algorithme il veut utiliser
func askUserAlgorithm() int{
	var n int
	fmt.Println("Which algorithm do you want to use?")
	fmt.Println("0 - Dynamique Libre Algorithm")
	fmt.Println("1 - Acception Immediate Algorithm")
	fmt.Println("2 - Acception Differee Algorithm")
	fmt.Println("3 - Top Trading cycles algorithm")
	fmt.Println("4 - Exit program")

	fmt.Scan(&n)

	return n
}

func main() {
	fmt.Println("Welcome to TP3 - IA04 Program.")
	fmt.Println("Loading prenoms from 'Prenoms.csv'. This may take a while...")
	prenoms := ExtractPrenoms()	//Extraction des prénoms cela peut prendre un peu de temps
	for{
		fmt.Println()
		fmt.Println("Let's start a new test!")
		n := askUserN()
		for n < 0{
			fmt.Println("n must be a positive number.")
			n = askUserN()
		}
		if(n == 0){	//on arrête le programme
			break
		}

		fmt.Println()
		fmt.Println("Generating two groups of", n, "agents. Please wait...")
		proposants, disposants := GenerateAgents(prenoms, n)	//génération des proposants et des disposants
		printAgents(proposants, disposants)

		choice := askUserAlgorithm()
		for choice < 0 || choice > 4{
			fmt.Println("You must enter a number between 0 and 4.")
			choice = askUserAlgorithm()
		}

		if(choice == 4){	//on arrête le programme
			break
		}

		switch choice{	//On execute un algorithme selon le choix de l'utilisateur
		case 0:
			fmt.Println("Executing Dynamique Libre Algorithm, please wait...")
			mariages := DynamiqueLibreAlgorithm(proposants, disposants)
			PrintResult(mariages, time.Now(), proposants, disposants)
		case 1:
			fmt.Println("Executing Acceptation Immediate Algorithm, please wait...")
			mariages := AcceptationImmediateAlgorithm(proposants, disposants)
			PrintResult(mariages, time.Now(), proposants, disposants)
		case 2:
			fmt.Println("Executing Acceptation Differee Algorithm, please wait...")
			mariages := make(Mariages)
			mariages = AcceptationDiffereeAlgorithm(TableauAgTOtableauPointeursAg(proposants), TableauAgTOtableauPointeursAg(disposants), mariages)
			PrintResult(mariages, time.Now(),proposants, disposants)
		case 3:
			fmt.Println("Executing Top Trading Cycles Algorithm, please wait...")
			mariages := make(Mariages)
			mariages = AcceptationDiffereeAlgorithm(TableauAgTOtableauPointeursAg(proposants), TableauAgTOtableauPointeursAg(disposants), mariages)
			PrintResult(mariages, time.Now(),proposants, disposants)
		}

	}
	fmt.Println("Bye.")	//Au revoir
}
