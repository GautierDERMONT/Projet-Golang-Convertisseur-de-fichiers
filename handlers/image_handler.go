package handlers

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	_ "image/gif"  // Supporte le format GIF
	_ "image/jpeg" // Supporte le format JPEG
	_ "image/png"  // Supporte le format PNG
)

// ConvertImageFormat convertit une image d'un format à un autre
func ConvertImageFormat(inputPath, outputPath string) error {
	// Ouvre le fichier source
	file, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("unable to open source file: %v", err) // Retourne une erreur si l'ouverture du fichier échoue
	}
	defer file.Close() // Ferme le fichier source à la fin de la fonction

	// Détecte et décode l'image
	img, format, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("error decoding image: %v", err) // Retourne une erreur si le décodage de l'image échoue
	}
	fmt.Println("Image format detected:", format) // Affiche le format de l'image détecté

	// Crée le fichier de sortie
	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create target file: %v", err) // Retourne une erreur si la création du fichier cible échoue
	}
	defer outputFile.Close() // Ferme le fichier de sortie à la fin de la fonction

	// Détermine l'extension du fichier cible
	outputExt := strings.ToLower(filepath.Ext(outputPath))

	// Encode et sauvegarde l'image dans le bon format
	switch outputExt {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(outputFile, img, nil) // Encode l'image en JPEG
	case ".png":
		err = png.Encode(outputFile, img) // Encode l'image en PNG
	case ".gif":
		err = gif.Encode(outputFile, img, nil) // Encode l'image en GIF
	default:
		return fmt.Errorf("unsupported target format: %s", outputExt) // Retourne une erreur si le format cible n'est pas supporté
	}

	if err != nil {
		return fmt.Errorf("error encoding image: %v", err) // Retourne une erreur si l'encodage de l'image échoue
	}

	return nil // Retourne nil si tout s'est bien passé
}
