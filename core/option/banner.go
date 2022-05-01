package options

import "Zscan/core/logger"

const banner = `

 ████████  ████████   ██████      ██     ████     ██
░░░░░░██  ██░░░░░░   ██░░░░██    ████   ░██░██   ░██
     ██  ░██        ██    ░░    ██░░██  ░██░░██  ░██
    ██   ░█████████░██         ██  ░░██ ░██ ░░██ ░██
   ██    ░░░░░░░░██░██        ██████████░██  ░░██░██
  ██            ░██░░██    ██░██░░░░░░██░██   ░░████
 ████████ ████████  ░░██████ ░██     ░██░██    ░░███
░░░░░░░░ ░░░░░░░░    ░░░░░░  ░░      ░░ ░░      ░░░
`

const Version = `0.0.1`

func ShowBanner() {
	logger.Printf("%s\n", banner)
	logger.Printf("                                                        %s\n\n", Version)
}