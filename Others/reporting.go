package Others

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func createReportFolder() {

	if _, err := os.Stat("./Others/reports"); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir("./Others/reports", 1)
			if err != nil && !os.IsExist(err) {
				log.Fatal(err)
			}
		} else {
			// other error
		}
	}
}

func ReportData(data []byte, rep string) {
	createReportFolder()

	file, err := os.Create("./Others/reports/report-" + rep + ".txt")
	handleError(err)

	length, err := io.WriteString(file, "REPORT - "+time.Now().Format("01-02-2006")+" ################# \n"+string(data))

	handleError(err)
	fmt.Printf("Data reporting done! - %d characters.\n", length)
	defer file.Close()
}

func ReadReport(date string) {

	data, err := ioutil.ReadFile("./Others/report-" + date + ".txt")
	if err != nil {
		handleError(errors.New("dosya bulunamadı"))
	}

	fmt.Printf("Read that:  %s", string(data))
}

func ShowDisplayableReports() {

	files, err := ioutil.ReadDir("./Others/reports")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("~~ Show reports ./reports ~~")

	for _, file := range files {
		f := file.Name()
		sReport := f[0:6]
		sTxt := f[len(f)-4:]

		if sReport == "report" && sTxt == ".txt" {
			fmt.Println(file.Name())
		}

	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
