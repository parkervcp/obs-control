package main

import (
	"log"

	obsws "github.com/christopher-dG/go-obs-websocket"
	flag "github.com/spf13/pflag"
)

var (
	c = obsws.Client{Host: "localhost", Port: 4444}

	job   string
	scene string
)

func main() {
	flag.StringVarP(&job, "job", "j", "get_scenes", "job to run against the obs-websocket")
	flag.StringVarP(&scene, "scene", "s", "", "scene to transition to")
	flag.Parse()

	log.Printf("scene '%s' specified", scene)

	if err := c.Connect(); err != nil {
		log.Fatal(err)
	}
	defer c.Disconnect()

	switch job {
	case "get_scenes":
		getScenes()
	case "change_preview":
		setPreview(scene)
	case "get_transitions":
		getTransitions()
	case "transition":
		transition()
	}
}

func getScenes() {
	// make a request for scened
	gsreq := obsws.NewGetSceneListRequest()
	if err := gsreq.Send(c); err != nil {
		log.Fatal(err)
	}

	// receive response
	resp, err := gsreq.Receive()
	if err != nil {
		log.Fatal(err)
	}

	// parse and print response
	for _, each := range resp.Scenes {
		log.Println("response:", each["name"])
	}
}

func getTransitions() {
	// make call to get all available transitions
	gtreq := obsws.NewGetTransitionListRequest()
	if err := gtreq.Send(c); err != nil {
		log.Fatal(err)
	}

	// get transitions response
	resp, err := gtreq.Receive()
	if err != nil {
		log.Fatal(err)
	}

	// list avalable
	log.Println("transitions:", resp.Transitions)
}

func setPreview(scene string) {
	if scene == "" {
		log.Fatal("no scene set to transition to")
	} else {
		log.Print("checking if scene exists")

	}

	sreq := obsws.NewSetPreviewSceneRequest(scene)
	if err := sreq.Send(c); err != nil {
		log.Fatal(err)
	}
}

func transition() {
	// a transition request
	treq := obsws.NewTransitionToProgramRequest(nil, "Fade", 300)
	if err := treq.Send(c); err != nil {
		log.Fatal(err)
	}

	// transition request response
	resp, err := treq.Receive()
	if err != nil {
		log.Fatal(err)
	}

	// log status
	log.Println("transitions:", resp.Status_)
}
