package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

var logFile = flag.String("log", "stdout", "Specify a log file to write log output to")
var envFile = flag.String("envfile", "", "Specify the environmental file to parse (usually project/.env.sample)")
var appEnvironment = flag.String("enviroment", "staging", "Specify the environment used in searching the keys in vault")
var keepdefaults = flag.Bool("keep-defaults", false, "Do NOT overwrite default values in the env file")

func init() {
	flag.Parse()
}

func main() {

	// handle all panics gracefully
	defer func() {
		if r := recover(); r != nil {
			log.SetOutput(os.Stderr)
			log.Fatal(r)
		}
	}()

	if *envFile == "" {
		flag.Usage()
		panic("No env file given to parse")
	}

	if *logFile == "stdout" {
		// output := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		log.SetOutput(os.Stdout)
	} else {

		f, err := os.OpenFile(string(*logFile), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("Cannot log file: %v", err))
		}

		log.SetOutput(f)
	}

	var wg sync.WaitGroup

	log.Println("Reading vars")

	var envVars map[string]string
	envVars, err := godotenv.Read(*envFile)

	if err != nil {
		panic(fmt.Sprintf("Cannot read env file! %v", err))
	}

	for env, defValue := range envVars {
		go createEnvVar(env, defValue, &wg)
		wg.Add(1)
	}

	wg.Wait()

}

func createEnvVar(key string, defValue string, wg *sync.WaitGroup) {

	var value string

	// create an environmental variable if the vault is found in vaul

	if *keepdefaults && defValue != "" {
		value = defValue
	} else {
		var vault bool = false

		if vault {
			// try to get the value from vault
			value = "FROM VAULT"
		} else {
			value = ""
		}

	}

	// do not create empty environmental variables
	if value != "" {
		newEnvVar := EnvVar{key: key, val: value}

		log.Println(newEnvVar.get())

	}

	wg.Done()
}
