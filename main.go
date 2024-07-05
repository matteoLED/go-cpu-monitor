package main

import (
	"go_sys_monitor/internal/monitor"
	"go_sys_monitor/internal/ui"
	"time"
)

func main() {
	// Créer une chaîne de communication pour synchroniser les mises à jour des FPS
	done := make(chan struct{})

	// Ticker pour mettre à jour les FPS
	frameTicker := time.NewTicker(time.Millisecond * 16) // Environ 60 FPS

	// Lancer les mises à jour des FPS dans une goroutine séparée
	go func() {
		for {
			select {
			case <-frameTicker.C:
				monitor.Monitor.UpdateFPS()
			case <-done:
				frameTicker.Stop()
				return
			}
		}
	}()

	// Lancer l'interface utilisateur dans la goroutine principale
	ui.RunUI()

	// Fermer la chaîne de communication pour arrêter la goroutine des mises à jour des FPS
	close(done)
}
