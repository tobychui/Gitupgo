package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-git/go-git/v5"
)

/*
	Kōshin
	author: tobychui

	General purpose automated Git based update system
	Spinoff project for the ArozOS system
*/

type ConfigInfo struct {
	Gitrepo    string `json:"gitrepo"`
	Folder     string `json:"folder,omitempty"`
	PreScript  string `json:"pre-script,omitempty"`
	PostScript string `json:"post-script,omitempty"`
	Interval   int    `json:"interval,omitempty"`
}

var (
	UsingConfig ConfigInfo
)

func main() {
	//Setup flags
	config := flag.String("config", "./config.json", "The system configuration file")
	updateOnBoot := flag.Bool("b", true, "Enable update on application startup")
	flag.Parse()

	//Start the main system
	if !fileExists(*config) {
		panic("*ERROR* Configuration not exists!")
	}

	//Load the config and start doing stuffs
	content, err := ioutil.ReadFile(*config)
	if err != nil {
		panic(err)
	}

	//Parse the config content
	UsingConfig = ConfigInfo{
		Gitrepo:    "https://github.com/tobychui/arozos",
		PreScript:  "",
		PostScript: "",
		Interval:   86400, //1 day by default
	}
	err = json.Unmarshal(content, &UsingConfig)
	if err != nil {
		panic(err)
	}

	log.Println("Starting Koshin Utilties...")
	//Start doing things
	if *updateOnBoot {
		PerformUpdate()
	}

	//Register a timer to start doing update
	log.Println("Ticker update started. Updating every " + strconv.Itoa(UsingConfig.Interval) + " seconds")
	ticker := time.NewTicker(time.Duration(UsingConfig.Interval) * time.Second)
	done := make(chan bool)

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			log.Println("Running Update Process")
			PerformUpdate()
		}
	}
}

func PerformUpdate() {
	//Execute the start script if exists
	if UsingConfig.PreScript != "" {
		if fileExists(UsingConfig.PreScript) {
			cmd := exec.Command(UsingConfig.PreScript)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Println("*PRESCRIPT* Failed to run ", string(out))
				log.Println("Continuing the update process")
			} else {
				fmt.Println("*PRESCRIPT* Execution completed ", string(out))
			}

		} else {
			log.Println("*PRESCRIPT* Script not exists. Skipping")
		}
	}

	if fileExists(UsingConfig.Folder) && fileExists(filepath.Join(UsingConfig.Folder, ".git")) {
		//Pull the repo
		log.Println("Updating the system with Git Pull Operation")
		log.Println("Targeting git directory: ", filepath.Join(UsingConfig.Folder))
		r, err := git.PlainOpen(filepath.Join(UsingConfig.Folder))
		if err != nil {
			//This happens when network failed or git pull already up to date. Just continue next cycle
			log.Println(err.Error())
			log.Println("Waiting next update cycle...")
			return
		}

		// Get the working directory for the repository
		w, err := r.Worktree()
		if err != nil {
			log.Println("*ERROR* Failed to analysis work tree. Is the git repo corrupted?")
			log.Println(err.Error())
			return
		}

		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			log.Println(err.Error())
			log.Println("Waiting next update cycle...")
			return
		}

		//Print the commit head
		ref, _ := r.Head()
		commit, err := r.CommitObject(ref.Hash())
		if err == nil {
			fmt.Println(commit)
		}

	} else {
		//Git clone it anyway
		log.Println("Updating the system with Git Clone Operation")
		_, err := git.PlainClone(UsingConfig.Folder, false, &git.CloneOptions{
			URL:               UsingConfig.Gitrepo,
			Progress:          os.Stdout,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})

		if err != nil {
			log.Println("*ERROR* Unable to clone git repo. Are you sure the repo is accessiable from this client?")
			log.Println(err)
		} else {
			log.Println("Clone completed. Waiting for another update cycle...")
		}

	}

	//Execute the postscript
	if UsingConfig.PostScript != "" {
		if fileExists(UsingConfig.PostScript) {
			cmd := exec.Command(UsingConfig.PostScript)
			out, err := cmd.CombinedOutput()
			if err != nil {
				log.Println("*POSTSCRIPT* Failed to run ", string(out))
				log.Println("Continuing the update process")
			} else {
				fmt.Println("*POSTSCRIPT* Execution completed ", string(out))
			}

		} else {
			log.Println("*POSTSCRIPT* Script not exists. Skipping")
		}
	}

	//Cycle completed
}

//Check if a file exists given its filename
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
