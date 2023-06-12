package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func getDate(date time.Time) (filedate, titlename string) {
	filedate = date.Format("2006_01_02")
	titlename = date.Format("2006-Jan-02")
	return
}

func createDailyContent(date string) string {
	title := fmt.Sprintf("# Reunión y quehaceres %s\n", date)
	dailytitle := fmt.Sprintf("\n## Daily\n")
	questions := fmt.Sprintf("\n### Preguntas\n\n- [ ] \n- [ ] \n")
	dailymatters := fmt.Sprintf("\n### Cuestiones tratadas\n\n- \n- \n")
	projects := fmt.Sprintf("\n## Proyectos\n")
	todo := fmt.Sprintf("\n### To-do\n\n- [ ]\n- [ ]\n")
	return title + dailytitle + questions + dailymatters + projects + todo
}

func main() {

	dailyNote := flag.Bool("daily", false, "creates a daily note type with the current date")
	typeDocs := flag.Bool("docs", false, "creates a general doc file")

	flag.Parse()
	fmt.Printf("daily=%t, docs=%t", *dailyNote, *typeDocs)

	if *typeDocs && *dailyNote {
		fmt.Fprintf(os.Stderr, "\n%sCan't create both daily and docs files at the same time\n", colorRed)
		os.Exit(1)
	}

	fmt.Printf("\nAbout to create your file.\n")

	filename := ""
	var filecontent string
	if *dailyNote {
		filedate, titlename := getDate(time.Now())

		filename = fmt.Sprintf("%s.md", filedate)
		filecontent += createDailyContent(titlename)
	} else {
		filename = "README.md"
		filecontent = "# Title\n"
	}

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)

	_, err = w.WriteString(filecontent)

	if err != nil {
		panic(err)
	}
	w.Flush()
}
