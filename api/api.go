package api

import (
	"ConvertisseurFichiersGo/handlers"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// handleFileUpload gère l'upload de fichiers et les conversions
func handleFileUpload(w http.ResponseWriter, r *http.Request, action func(string, string) error, successMessage string) {
	// Lire le fichier uploadé
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "❌ Impossible de lire le fichier", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Lire le format de sortie
	format := r.FormValue("format")
	if format == "" {
		http.Error(w, "❌ Format de sortie manquant", http.StatusBadRequest)
		return
	}

	// Créer un fichier temporaire pour stocker le fichier uploadé
	tempFile, err := os.CreateTemp("", "upload-*"+filepath.Ext(header.Filename))
	if err != nil {
		http.Error(w, "❌ Erreur lors de la création du fichier temporaire", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tempFile.Name())

	// Copier le contenu du fichier uploadé dans le fichier temporaire
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "❌ Erreur lors de la copie du fichier", http.StatusInternalServerError)
		return
	}

	// Exécuter l'action (conversion, compression, etc.)
	outputFile := tempFile.Name() + "." + format
	if err := action(tempFile.Name(), outputFile); err != nil {
		http.Error(w, fmt.Sprintf("❌ Erreur : %v", err), http.StatusInternalServerError)
		return
	}

	// Ajouter un message de succès dans l'en-tête
	w.Header().Set("X-Success-Message", successMessage)

	// Renvoyer le fichier converti
	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(outputFile))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, outputFile)
}

// handleCompression gère la compression de fichiers
func handleCompression(w http.ResponseWriter, r *http.Request) {
	source := r.FormValue("source")
	target := r.FormValue("target")
	if source == "" || target == "" {
		http.Error(w, "❌ Paramètres manquants : source ou target", http.StatusBadRequest)
		return
	}

	if err := handlers.CompressFiles(source, target); err != nil {
		http.Error(w, fmt.Sprintf("❌ Erreur lors de la compression : %v", err), http.StatusInternalServerError)
		return
	}

	// Ajouter un message de succès dans l'en-tête
	w.Header().Set("X-Success-Message", "✅ Compression réussie ! 🎉")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Le fichier a été compressé avec succès."))
}

// handleDecompression gère la décompression de fichiers
func handleDecompression(w http.ResponseWriter, r *http.Request) {
	source := r.FormValue("source")
	target := r.FormValue("target")
	if source == "" || target == "" {
		http.Error(w, "❌ Paramètres manquants : source ou target", http.StatusBadRequest)
		return
	}

	if err := handlers.DecompressFiles(source, target); err != nil {
		http.Error(w, fmt.Sprintf("❌ Erreur lors de la décompression : %v", err), http.StatusInternalServerError)
		return
	}

	// Ajouter un message de succès dans l'en-tête
	w.Header().Set("X-Success-Message", "✅ Décompression réussie ! 🎉")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Le fichier a été décompressé avec succès."))
}

// StartAPI démarre le serveur HTTP
func StartAPI() {
	// Endpoint pour la conversion de fichiers texte
	http.HandleFunc("/convert/text", func(w http.ResponseWriter, r *http.Request) {
		handleFileUpload(w, r, func(source, target string) error {
			extSource := filepath.Ext(source)
			extTarget := filepath.Ext(target)

			switch {
			case extSource == ".csv" && extTarget == ".json":
				return handlers.ConvertCSVtoJSON(source, target)
			case extSource == ".json" && extTarget == ".xml":
				return handlers.ConvertJSONtoXML(source, target)
			default:
				return fmt.Errorf("unsupported conversion: %s → %s", extSource, extTarget)
			}
		}, "✅ Conversion réussie ! 🎉")
	})

	// Endpoint pour la conversion d'images
	http.HandleFunc("/convert/image", func(w http.ResponseWriter, r *http.Request) {
		handleFileUpload(w, r, handlers.ConvertImageFormat, "✅ Conversion d'image réussie ! 🎉")
	})

	// Endpoint pour la compression de fichiers
	http.HandleFunc("/compress", handleCompression)

	// Endpoint pour la décompression de fichiers
	http.HandleFunc("/decompress", handleDecompression)

	// Endpoint racine
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenue sur l'API de conversion ! 🚀\n")) // Message de bienvenue pour la racine
	})

	port := ":8080"                           // Définit le port du serveur
	fmt.Println("Starting server on", port)   // Affiche un message de démarrage du serveur
	log.Fatal(http.ListenAndServe(port, nil)) // Démarre le serveur et logge une erreur fatale en cas d'échec
}
