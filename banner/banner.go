package banner

import (
	"github.com/projectdiscovery/gologger"
)

func Banner() {
	gologger.Print().Msgf(`
███████╗██╗   ██╗██████╗            ██╗                                                      
██╔════╝██║   ██║██╔══██╗        ██╗╚██╗                                                     
███████╗██║   ██║██████╔╝        ╚═╝ ██║                                                     
╚════██║██║   ██║██╔══██╗        ██╗ ██║                                                     
███████║╚██████╔╝██████╔╝        ╚═╝██╔╝                                                     
╚══════╝ ╚═════╝ ╚═════╝            ╚═╝                                                      
                                                                                             
███████╗███╗   ██╗██╗   ██╗███╗   ███╗███████╗██████╗  █████╗ ████████╗██╗ ██████╗ ███╗   ██╗
██╔════╝████╗  ██║██║   ██║████╗ ████║██╔════╝██╔══██╗██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║
█████╗  ██╔██╗ ██║██║   ██║██╔████╔██║█████╗  ██████╔╝███████║   ██║   ██║██║   ██║██╔██╗ ██║
██╔══╝  ██║╚██╗██║██║   ██║██║╚██╔╝██║██╔══╝  ██╔══██╗██╔══██║   ██║   ██║██║   ██║██║╚██╗██║
███████╗██║ ╚████║╚██████╔╝██║ ╚═╝ ██║███████╗██║  ██║██║  ██║   ██║   ██║╚██████╔╝██║ ╚████║
╚══════╝╚═╝  ╚═══╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝

`)
	gologger.Print().Msgf("   Created by Sharo_k_h :)\n\n")
}
