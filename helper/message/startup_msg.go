package message

import (
	"fmt"

	"github.com/labstack/gommon/color"
	"github.com/ramabmtr/heimdall/config"
)

func ServiceNameMessage() {
	banner := `
###  ###        #               ##        ##  ##  
 #    #                          #         #   #  
 #    #                          #         #   #  
 ######   ###  ##  ########   ####  ####   #   #  
 #    #  #   #  #   #  #  #  #   #     #   #   #  
 #    #  #####  #   #  #  #  #   #  ####   #   #  
 #    #  #      #   #  #  #  #   #  #  #   #   #  
###  ###  #### ### #########  ##### ##### ### ###  
..................................................
	`

	fmt.Println(banner)
}

func EnvInfoMessage() {
	fmt.Print("» Environment: ")
	if config.AppMode == "production" {
		color.Println(color.Blue(config.AppMode))
	} else {
		color.Println(color.Red(config.AppMode))
	}

	fmt.Print("» Debug: ")
	if config.AppDebug {
		color.Println(color.Red("ON"))
	} else {
		color.Println(color.Blue("OFF"))
	}
}
