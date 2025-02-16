package handlers

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// CompressFiles compresse les fichiers d'un répertoire source dans un fichier ZIP cible
func CompressFiles(source, target string) error {
	zipFile, err := os.Create(target) // Crée le fichier ZIP cible
	if err != nil {
		return err // Retourne l'erreur si la création du fichier échoue
	}
	defer zipFile.Close() // Ferme le fichier ZIP à la fin de la fonction

	archive := zip.NewWriter(zipFile) // Crée un nouvel écrivain ZIP
	defer archive.Close()             // Ferme l'écrivain ZIP à la fin de la fonction

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Retourne l'erreur si la lecture du fichier échoue
		}

		header, err := zip.FileInfoHeader(info) // Crée un en-tête de fichier ZIP
		if err != nil {
			return err // Retourne l'erreur si la création de l'en-tête échoue
		}

		header.Name = path // Définit le nom du fichier dans l'archive
		if info.IsDir() {
			header.Name += "/" // Ajoute un slash si c'est un répertoire
		} else {
			header.Method = zip.Deflate // Utilise la méthode de compression Deflate pour les fichiers
		}

		writer, err := archive.CreateHeader(header) // Crée un écrivain pour le fichier dans l'archive
		if err != nil {
			return err // Retourne l'erreur si la création de l'écrivain échoue
		}

		if !info.IsDir() { // Si ce n'est pas un répertoire
			file, err := os.Open(path) // Ouvre le fichier source
			if err != nil {
				return err // Retourne l'erreur si l'ouverture du fichier échoue
			}
			defer file.Close()             // Ferme le fichier source à la fin de la fonction
			_, err = io.Copy(writer, file) // Copie le contenu du fichier source dans l'archive
			return err                     // Retourne l'erreur si la copie échoue
		}
		return nil // Retourne nil si tout s'est bien passé
	})
}

// DecompressFiles décompresse les fichiers d'un fichier ZIP source dans un répertoire cible
func DecompressFiles(source, target string) error {
	reader, err := zip.OpenReader(source) // Ouvre le fichier ZIP source
	if err != nil {
		return err // Retourne l'erreur si l'ouverture du fichier échoue
	}
	defer reader.Close() // Ferme le lecteur ZIP à la fin de la fonction

	for _, file := range reader.File { // Parcourt chaque fichier dans l'archive
		path := filepath.Join(target, file.Name) // Définit le chemin du fichier dans le répertoire cible
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm) // Crée le répertoire si c'est un répertoire
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			return err // Retourne l'erreur si la création du répertoire échoue
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode()) // Crée le fichier cible
		if err != nil {
			return err // Retourne l'erreur si la création du fichier échoue
		}
		defer outFile.Close() // Ferme le fichier cible à la fin de la fonction

		rc, err := file.Open() // Ouvre le fichier dans l'archive
		if err != nil {
			return err // Retourne l'erreur si l'ouverture du fichier échoue
		}
		defer rc.Close() // Ferme le fichier dans l'archive à la fin de la fonction

		_, err = io.Copy(outFile, rc) // Copie le contenu du fichier dans l'archive vers le fichier cible
		if err != nil {
			return err // Retourne l'erreur si la copie échoue
		}
	}
	return nil // Retourne nil si tout s'est bien passé
}
