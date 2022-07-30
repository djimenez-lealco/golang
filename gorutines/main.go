package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Usuario struct {
	document string
	nombres  string
	celular  string
	sleep    int
	estado   string
}
type channels struct {
	done  chan bool
	start chan bool
}

func main() {
	wg := new(sync.WaitGroup)
	// canales de comunicacion para indicar si se inicio y finalizo el procesamiento
	channels := channels{
		done:  make(chan bool),
		start: make(chan bool),
	}
	// Representa la cantidad de workers max permitida
	buffer := 1
	// Representa los Workers activos
	workers := 0
	// ==============================================
	// Gestion de usuarios procesados
	// ==============================================
	usuariosProcesadosLock := sync.Mutex{}
	usuariosProcesados := 0
	// ==============================================

	usuarios := []Usuario{
		{
			document: "1018408050",
			nombres:  "usuario A |50",
			celular:  "3188725398",
			sleep:    50,
			estado:   "pendiente",
		},
		{
			document: "01231560164",
			nombres:  "usuario B |1",
			celular:  "3188725397",
			sleep:    1,
			estado:   "pendiente",
		},
		{
			document: "516515615564",
			nombres:  "usuario C |30",
			celular:  "3188725396",
			sleep:    30,
			estado:   "pendiente",
		},
		{
			document: "1561594564647",
			nombres:  "usuario D |50",
			celular:  "3188725396",
			sleep:    50,
			estado:   "pendiente",
		},
		{
			document: "15645646543",
			nombres:  "usuario 3 |20",
			celular:  "3188725394",
			sleep:    20,
			estado:   "pendiente",
		},
		{
			document: "25465465433",
			nombres:  "usuario 4 |30",
			celular:  "3188725393",
			sleep:    30,
			estado:   "pendiente",
		},
		{
			document: "78945612313",
			nombres:  "usuario 5 |4",
			celular:  "3188725392",
			sleep:    4,
			estado:   "pendiente",
		},
		{
			document: "4546456150",
			nombres:  "usuario 6 |1",
			celular:  "3188725391",
			sleep:    1,
			estado:   "pendiente",
		},
		{
			document: "16516551453",
			nombres:  "usuario 7 |2",
			celular:  "3188725390",
			sleep:    2,
			estado:   "pendiente",
		},
	}
	wg.Add(1)

	go func() {
		doneWorkerListener(&channels.done, &workers)
	}()
	go func() {
		startWorkerListener(&channels.start, &workers)
	}()

	go procesarUsuarios(&workers, buffer, &usuariosProcesadosLock, &usuariosProcesados, &usuarios, wg, channels)

	go mostrarProceso(&usuarios)
	wg.Wait()
}

func doneWorkerListener(done *chan bool, workers *int) {
	for termino := range *done {
		if termino {
			//color.Green("-Workers")
			*workers--
		}
	}

}

func startWorkerListener(start *chan bool, workers *int) {
	for inicio := range *start {
		if inicio {
			//color.Green("+Workers")
			*workers++
		}
	}
}

func procesarUsuarios(workers *int, buffer int, usuariosProcesadosLock *sync.Mutex, usuariosProcesados *int, usuarios *[]Usuario, wg *sync.WaitGroup, channels channels) {

	for {

		if *workers < buffer {
			usuariosProcesadosLock.Lock()
			if *usuariosProcesados >= len(*usuarios) {
				break
			}
			usuario := &(*usuarios)[*usuariosProcesados]
			*usuariosProcesados++
			wg.Add(1)
			go procesarRegistroUsuario(usuario, wg, channels.done)
			channels.start <- true
			usuariosProcesadosLock.Unlock()
		} else if *workers >= buffer {
			time.Sleep(2 * time.Second)
		}

	}
	time.Sleep(10 * time.Second)
	defer wg.Done()

}

func procesarRegistroUsuario(usuario *Usuario, wg *sync.WaitGroup, done chan bool) {
	//fmt.Println("Start\t [" + usuario.nombres + "\t| sleep:" + string(rune(usuario.sleep)) + "\t]")
	(*usuario).estado = "procesando"
	time.Sleep(time.Duration(usuario.sleep) * time.Second)
	(*usuario).estado = "procesado"
	//color.Red("End\t [" + usuario.nombres + "\t| sleep:" + string(rune(usuario.sleep)) + "\t]")
	done <- true
	defer wg.Done()

}

func mostrarProceso(usuarios *[]Usuario) {
	for {

		fmt.Print("\033[H\033[2J")
		for _, usuario := range *usuarios {
			//s, _ := json.MarshalIndent(usuario, "", "\t")
			if usuario.estado == "pendiente" {
				fmt.Println("[" + usuario.nombres + " " + usuario.estado + "] ")
			} else if usuario.estado == "procesando" {
				color.Blue("[*" + usuario.nombres + " " + usuario.estado + "] ")

			} else if usuario.estado == "procesado" {
				color.Green("[" + usuario.nombres + " " + usuario.estado + "] ")

			}
		}

		time.Sleep(50 * time.Millisecond)

	}
}
