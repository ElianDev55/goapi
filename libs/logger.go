package libs

import "log"



	const (
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Reset  = "\033[0m"
	)



	func Info(message string) {

		log.SetPrefix("INFO: ")
		log.Println(Green + "INFO: " + message + Reset)
		log.SetFlags(log.LstdFlags | log.Lshortfile)

	}


	func Warning(message string) {

		log.SetPrefix("WARNING: ")
		log.Println(Yellow + "WARNING: " + message + Reset)
		log.SetFlags(log.LstdFlags | log.Lshortfile)

	}



	func Error(message string) {

		log.SetPrefix("ERROR: ")
		log.Println(Red + "ERROR: " + message + Reset)
		log.SetFlags(log.LstdFlags | log.Lshortfile)

	}
