package libs

import "log"



	const (
		Red    = "\033[31m"
		Green  = "\033[32m"
		Yellow = "\033[33m"
		Reset  = "\033[0m"
	)

	type LoggerP struct {
		Filename  string
	}


	func (lp *LoggerP) InfoLogger(message string) {
		log.Println("LOCATION: "+ lp.Filename + "  " + Green + "INFO: " + message + Reset)
		log.SetFlags(log.LstdFlags)
	}


	func (lp *LoggerP) WarningLogger(message string) {
		log.Println("LOCATION: "+ lp.Filename + "  " + Yellow + "WARNING: " + message + Reset)
		log.SetFlags(log.LstdFlags)
	}



	func (lp *LoggerP) ErrorLogger(message string, error error ) {
		log.Println("LOCATION: "+ lp.Filename + "  " + Red + "ERROR: " + message + Reset + error.Error())
		log.SetFlags(log.LstdFlags)
	}
