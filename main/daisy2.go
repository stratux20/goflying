---
Lecture des trames NMEA via le port série
---


package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"log"
	"strings"
)



var aisIncomingMsgChan chan string = make(chan string, 100)
var aisExitChan chan bool = make(chan bool, 1)
var aisNmeaParser = aisnmea.NMEACodecNew(ais.CodecNew(false, false))

func DaisyListen() {
	//go predTest()
	for {
		if !globalSettings.Daisy_Enabled || DaisyDev == nil {
			// wait until AIS is enabled
			time.Sleep(1 * time.Second)
			continue
		}
		// log.Printf("ais connecting...")
		aisAddr := "127.0.0.1:10110"
		conn, err := net.Dial("tcp", aisAddr)
		if err != nil { // Local connection failed.
			time.Sleep(3 * time.Second)
			continue
		}
		log.Printf("ais successfully connected")
		aisReadWriter := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
		globalStatus.AIS_connected = true

		// Make sure the exit channel is empty, so we don't exit immediately
		for len(aisExitChan) > 0 {
			<-aisExitChan
		}

		go func() {
			scanner := bufio.NewScanner(aisReadWriter.Reader)
			for scanner.Scan() {
				aisIncomingMsgChan <- scanner.Text()
			}
			if scanner.Err() != nil {
				log.Printf("ais-rx-eu connection lost: " + scanner.Err().Error())
			}
			aisExitChan <- true
		}()

	loop:
		for globalSettings.AIS_Enabled {
			select {
			case data := <-aisIncomingMsgChan:
				TraceLog.Record(CONTEXT_AIS, []byte(data))
				parseAisMessage(data)
			case <-aisExitChan:
				break loop

			}
		}
		globalStatus.AIS_connected = false
		conn.Close()
		time.Sleep(3 * time.Second)
	}
}




func startDaisy2Receiver() {
	// Configuration du port série
	c := &serial.Config{Name: "/dev/ttyACM0", Baud: 38400} // Assurez-vous que le port est correct
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Printf("Erreur lors de l'ouverture du port DAISY2+ :%v", err)
		return
	}


	// Lecture continue des trames NMEA
	reader := bufio.NewReader(s)
	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("Erreur lors de la lecture du DAISY2+ : %v", err)
			break
		}

		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "!AIVDM") { // Vérifie si c'est une trame AIS
			processAISTrame(line)
		}
	}
}


// Fonction pour traiter les trames AIS
func processAISTrame(nmea string) {

	// Exemple : Affichage de la trame brute (vous pouvez ajouter une logique pour traiter la trame ici)
	fmt.Printf("Trame AIS reçue : %s\n", nmea)

	// Convertir la trame NMEA en un message AIS et l'importer dans le flux Stratux
	convertAISToStratux(nmea)
}

func convertAISToStratux(nmea string) {
	// Utilise les fonctions existantes dans ais.go pour traiter les trames AIS
	// Cette fonction se charge de l'intégration des trames dans Stratux
	aisIncomingMsgChan <- nmea
}
