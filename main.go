package main

import (
	"encoding/csv"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"math/rand"
	"os"
	"strings"
	"syscall"
	"time"
)

var rootCmd = &cobra.Command{Use: "password-generator"}

var generatePassword = &cobra.Command{
	Use:   "generate",
	Short: "generate a password",
	Long:  "generate a long safe password to use everywhere",
	Run:   genPassword,
}

var viewPassword = &cobra.Command{
	Use:   "View",
	Short: "View all the password",
	Long:  "View All the password saved in the csv file",
	Run:   viewFuncPassword,
}

var copyPassword = &cobra.Command{
	Use:   "copy",
	Short: "copy password",
	Long:  "Copy the specified Website password value",
	Run:   copyFuncPassword,
}

func copyFuncPassword(cmd *cobra.Command, args []string) {
	if !Authenticate() {
		fmt.Println("Authentication failed. Exiting.")
		return
	}
	fmt.Print("Enter website name: ")
	var websiteName string
	fmt.Scanln(&websiteName)

	file, err := os.Open("passwords.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	found := false
	var password string
	for _, record := range records {
		if record[0] == websiteName {
			password = record[1]
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Password for website '%s' not found.\n", websiteName)
		return
	}

	err = clipboard.WriteAll(password)
	if err != nil {
		fmt.Println("Error copying password to clipboard:", err)
		return
	}

	fmt.Printf("Password for website '%s' copied to clipboard.\n", websiteName)

}

func viewFuncPassword(cmd *cobra.Command, args []string) {
	if !Authenticate() {
		fmt.Println("Authentication failed. Exiting.")
		return
	}
	file, err := os.Open("passwords.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Display the records
	if len(records) == 0 {
		fmt.Println("No passwords stored.")
		return
	}

	//fmt.Printf("%-20s%-20s%s\n", "Website", "Password", "Timestamp")
	for _, record := range records {
		fmt.Printf("%-20s%-20s%s\n", record[0], record[1], record[2])
	}
}

func init() {
	rootCmd.AddCommand(generatePassword)
	rootCmd.AddCommand(viewPassword)
	rootCmd.AddCommand(copyPassword)
	// Add the "length" flag to the "generate" command
	generatePassword.Flags().IntP("length", "l", 12, "Length of the generated password")
}

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+{}[]|;:,.<>?/~`"
	allChars         = lowercaseLetters + uppercaseLetters + digits + specialChars
)

func genPassword(cmd *cobra.Command, args []string) {

	if !Authenticate() {
		fmt.Println("Authentication failed. Exiting.")
		return
	}
	length, _ := cmd.Flags().GetInt("length")

	fmt.Print("Enter website name: ")
	var websiteName string
	fmt.Scanln(&websiteName)

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = allChars[rand.Intn(len(allChars))]
	}

	file, err := os.OpenFile("passwords.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fileInfo, _ := file.Stat()
	if fileInfo.Size() == 0 {
		header := []string{"Website", "Password", "Timestamp"}
		if err := writer.Write(header); err != nil {
			fmt.Println("Error writing headers to CSV:", err)
			return
		}
	}

	record := []string{websiteName, string(password), time.Now().Format("2006-01-02 15:04:05")}
	if err := writer.Write(record); err != nil {
		fmt.Println("Error writing to CSV:", err)
		return
	}

	fmt.Println("Password generated:", string(password))
}

var mainPassword string

func main() {

	mainPassword = "IamAFuckingGenius"
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Authenticate() bool {
	fmt.Print("Enter master password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error reading password:", err)
		return false
	}
	fmt.Println()

	enteredPassword := strings.TrimSpace(string(bytePassword))
	return enteredPassword == mainPassword
}
