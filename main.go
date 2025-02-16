package main // Déclare le package principal, point d'entrée de l'application.

import (
	"ConvertisseurFichiersGo/api" // Importe le package "api" pour démarrer l'API.
	"ConvertisseurFichiersGo/cli" // Importe le package "cli" pour gérer le mode ligne de commande (CLI).
	"os"                          // Importe le package "os" pour accéder aux arguments de la ligne de commande.
)

// Fonction principale du programme.
func main() {
	// Vérifie si un argument est fourni et s'il est égal à "cli".
	if len(os.Args) > 1 && os.Args[1] == "cli" {
		cli.StartCLI(os.Args[2:]) // Démarre l'interface en ligne de commande.
	} else {
		api.StartAPI() // Sinon, démarre l'API web.
	}
}
