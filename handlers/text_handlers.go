package handlers

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"os"
)

// ConvertCSVtoJSON convertit un fichier CSV en fichier JSON
func ConvertCSVtoJSON(csvFilePath, jsonFilePath string) error {
	file, err := os.Open(csvFilePath) // Ouvre le fichier CSV
	if err != nil {
		return err // Retourne une erreur si l'ouverture du fichier échoue
	}
	defer file.Close() // Ferme le fichier CSV à la fin de la fonction

	reader := csv.NewReader(file)    // Crée un lecteur CSV
	records, err := reader.ReadAll() // Lit toutes les lignes du fichier CSV
	if err != nil {
		return err // Retourne une erreur si la lecture échoue
	}

	jsonFile, err := os.Create(jsonFilePath) // Crée le fichier JSON
	if err != nil {
		return err // Retourne une erreur si la création du fichier échoue
	}
	defer jsonFile.Close() // Ferme le fichier JSON à la fin de la fonction

	jsonEncoder := json.NewEncoder(jsonFile) // Crée un encodeur JSON
	jsonEncoder.SetIndent("", "  ")          // Définit l'indentation pour une meilleure lisibilité
	return jsonEncoder.Encode(records)       // Encode les données CSV en JSON
}

// ConvertCSVtoXML convertit un fichier CSV en fichier XML
func ConvertCSVtoXML(csvFilePath, xmlFilePath string) error {
	file, err := os.Open(csvFilePath) // Ouvre le fichier CSV
	if err != nil {
		return err // Retourne une erreur si l'ouverture du fichier échoue
	}
	defer file.Close() // Ferme le fichier CSV à la fin de la fonction

	reader := csv.NewReader(file)    // Crée un lecteur CSV
	records, err := reader.ReadAll() // Lit toutes les lignes du fichier CSV
	if err != nil {
		return err // Retourne une erreur si la lecture échoue
	}

	xmlFile, err := os.Create(xmlFilePath) // Crée le fichier XML
	if err != nil {
		return err // Retourne une erreur si la création du fichier échoue
	}
	defer xmlFile.Close() // Ferme le fichier XML à la fin de la fonction

	xmlEncoder := xml.NewEncoder(xmlFile) // Crée un encodeur XML
	xmlEncoder.Indent("", "  ")           // Définit l'indentation pour une meilleure lisibilité

	// Encapsuler les données CSV dans une structure XML
	type Row struct {
		Columns []string `xml:"Column"`
	}
	type Data struct {
		XMLName xml.Name `xml:"Records"`
		Rows    []Row    `xml:"Row"`
	}

	var data Data
	for _, record := range records {
		data.Rows = append(data.Rows, Row{Columns: record}) // Ajoute chaque ligne CSV à la structure XML
	}

	return xmlEncoder.Encode(data) // Encode les données en XML
}

// ConvertJSONtoXML convertit un fichier JSON en fichier XML
func ConvertJSONtoXML(jsonFilePath, xmlFilePath string) error {
	file, err := os.Open(jsonFilePath) // Ouvre le fichier JSON
	if err != nil {
		return err // Retourne une erreur si l'ouverture du fichier échoue
	}
	defer file.Close() // Ferme le fichier JSON à la fin de la fonction

	var data interface{}
	jsonDecoder := json.NewDecoder(file) // Crée un décodeur JSON
	if err := jsonDecoder.Decode(&data); err != nil {
		return err // Retourne une erreur si le décodage échoue
	}

	xmlFile, err := os.Create(xmlFilePath) // Crée le fichier XML
	if err != nil {
		return err // Retourne une erreur si la création du fichier échoue
	}
	defer xmlFile.Close() // Ferme le fichier XML à la fin de la fonction

	xmlEncoder := xml.NewEncoder(xmlFile) // Crée un encodeur XML
	xmlEncoder.Indent("", "  ")           // Définit l'indentation pour une meilleure lisibilité
	return xmlEncoder.Encode(data)        // Encode les données JSON en XML
}
