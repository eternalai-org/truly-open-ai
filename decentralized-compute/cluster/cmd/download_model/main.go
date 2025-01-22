package main

import (
	"flag"
	"log"
	"os"
	"solo/pkg/zip_hf_model_to_light_house"
)

var bashExec = flag.String("bash_exec", "/usr/bin/bash", "Path to the Bash executable")
var hfDir = flag.String("hf_dir", "/root/.cache/huggingface/hub", "Directory for Hugging Face models")
var hash = flag.String("hash", "", "Hash in Light House (contains model info after upload to Light House). Required if action is 'download'.")

func main() {
	flag.Parse()
	zip_hf_model_to_light_house.BASH_EXEC = *bashExec
	if len(*hfDir) == 0 {
		log.Fatalf("Must specify hf_dir:%v", *hfDir)
	}
	if _, err := os.Stat(*hfDir); os.IsNotExist(err) {
		log.Fatalf("Directory does not exist:%v", *hfDir)
	}
	err := zip_hf_model_to_light_house.DownloadHFModelFromLightHouse(*hash, *hfDir)
	if err != nil {
		log.Fatal(err)
	}

}
