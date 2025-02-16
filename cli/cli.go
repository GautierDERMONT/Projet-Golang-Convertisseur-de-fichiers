package cli

import (
	"ConvertisseurFichiersGo/handlers"
	"ConvertisseurFichiersGo/utils"
	"fmt"
)

// StartCLI lance l'interface en ligne de commande
func StartCLI(args []string) {
	if len(args) < 3 { // Vérifie si le nombre d'arguments est suffisant
		fmt.Println("Usage: cli <command> <source> <target>")                      // Affiche l'usage correct de la commande
		fmt.Println("Commands: convert-text, convert-image, compress, decompress") // Liste des commandes disponibles
		return
	}

	command := args[0] // Récupère la commande
	source := args[1]  // Récupère le fichier source
	target := args[2]  // Récupère le fichier cible

	switch command { // Traite la commande
	case "convert-text": // Conversion de texte
		if utils.GetFileExtension(source) == ".csv" && utils.GetFileExtension(target) == ".json" {
			if err := handlers.ConvertCSVtoJSON(source, target); err != nil { // Conversion CSV vers JSON
				fmt.Println("Error converting CSV to JSON:", err)
				return
			}
			fmt.Println("✅ Conversion réussie : CSV → JSON !")
		} else if utils.GetFileExtension(source) == ".json" && utils.GetFileExtension(target) == ".xml" {
			if err := handlers.ConvertJSONtoXML(source, target); err != nil { // Conversion JSON vers XML
				fmt.Println("Error converting JSON to XML:", err)
				return
			}
			fmt.Println("✅ Conversion réussie : JSON → XML !")
		} else if utils.GetFileExtension(source) == ".csv" && utils.GetFileExtension(target) == ".xml" {
			if err := handlers.ConvertCSVtoXML(source, target); err != nil { // Conversion CSV vers XML
				fmt.Println("Error converting CSV to XML:", err)
				return
			}
			fmt.Println("✅ Conversion réussie : CSV → XML !")
		} else {
			fmt.Println("❌ Conversion non supportée :", utils.GetFileExtension(source), "→", utils.GetFileExtension(target))
			return
		}

	case "convert-image": // Conversion d'image
		err := handleImageConversion(source, target) // Appelle la fonction de conversion d'image
		if err != nil {
			fmt.Println("❌ Erreur lors de la conversion d'image :", err)
			return
		}
		fmt.Println("✅ Conversion d'image réussie :", source, "→", target, "!")

	case "compress": // Compression de fichiers
		err := handlers.CompressFiles(source, target) // Appelle la fonction de compression
		if err != nil {
			fmt.Println("❌ Erreur lors de la compression :", err)
			return
		}
		fmt.Println("✅ Compression réussie :", source, "→", target, "!")

	case "decompress": // Décompression de fichiers
		err := handlers.DecompressFiles(source, target) // Appelle la fonction de décompression
		if err != nil {
			fmt.Println("❌ Erreur lors de la décompression :", err)
			return
		}
		fmt.Println("✅ Décompression réussie :", source, "→", target, "!")

	default: // Commande inconnue
		fmt.Println("❌ Commande inconnue :", command)
	}
}

// handleTextConversion gère la conversion de fichiers texte
func handleTextConversion(source, target string) error {
	srcExt := utils.GetFileExtension(source) // Récupère l'extension du fichier source
	tgtExt := utils.GetFileExtension(target) // Récupère l'extension du fichier cible

	switch {
	case srcExt == ".csv" && tgtExt == ".json": // Conversion CSV vers JSON
		return handlers.ConvertCSVtoJSON(source, target)
	case srcExt == ".json" && tgtExt == ".xml": // Conversion JSON vers XML
		return handlers.ConvertJSONtoXML(source, target)
	default: // Conversion non supportée
		return fmt.Errorf("unsupported text conversion: %s → %s", srcExt, tgtExt)
	}
}

// handleImageConversion gère la conversion de formats d'image
func handleImageConversion(source, target string) error {
	fmt.Println("Converting image:", source, "→", target) // Logge la conversion d'image

	err := handlers.ConvertImageFormat(source, target) // Appelle la fonction de conversion d'image
	if err != nil {
		fmt.Println("Debug: Error during image conversion:", err) // Logge l'erreur en cas d'échec
	}
	return err
}
