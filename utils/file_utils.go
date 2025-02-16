package utils // Déclare le package "utils", ce qui signifie que ce fichier appartient au package utilitaire du projet.

import (
	"os"            // Importe le package "os" pour interagir avec le système de fichiers.
	"path/filepath" // Importe le package "path/filepath" pour manipuler les chemins de fichiers de manière portable.
)

// EnsureDirectoryExists vérifie si un répertoire existe et le crée s'il n'existe pas.
func EnsureDirectoryExists(path string) error {
	return os.MkdirAll(path, os.ModePerm) // Crée le répertoire et tous ses répertoires parents avec les permissions maximales.
}

// GetFileExtension retourne l'extension d'un fichier à partir de son chemin.
func GetFileExtension(path string) string {
	return filepath.Ext(path) // Utilise filepath.Ext pour extraire et retourner l'extension du fichier, y compris le point (ex: ".txt").
}
