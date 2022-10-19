package main

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"io/ioutil"
	"log"
	"github.com/SharokhAtaie/subdomain-enumeration/banner"
	"github.com/SharokhAtaie/subdomain-enumeration/commands"
	"github.com/SharokhAtaie/subdomain-enumeration/functions"
)

type options struct {
	domain string
	output string
	silent bool
}

func main() {
	opt := &options{}

	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription("A tool for Subdomain Enumeration")
	flagSet.StringVarP(&opt.domain, "domain", "d", "", "domain for enumeration")
	flagSet.StringVarP(&opt.output, "output", "o", "", "filename for save output")
	flagSet.BoolVarP(&opt.silent, "silent", "s", false, "show silent output")

	if err := flagSet.Parse(); err != nil {
		log.Fatalf("Could not parse flags: %s\n", err)
	}

	if opt.domain == "" {
		banner.Banner()
		gologger.Print().Msgf("Flags:\n")
		gologger.Print().Msgf("\t-domain, -d 		Domain for enumeration")
		gologger.Print().Msgf("\t-output, -o 		Filename for save output")
		gologger.Print().Msgf("\t-silent, -s 		Show silent output")
		return
	}

	if opt.silent != true {
		banner.Banner()
	}

	data := FinalSubs(opt.domain)

	// save file
	if opt.output != "" {
		b := []byte(data)
		err1 := ioutil.WriteFile(opt.output, b, 0644)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	fmt.Println(data)
}

func FinalSubs(domain string) string {
	// crt.sh
	crt := functions.CRtsh("https://crt.sh/?q=%." + domain + "&output=json")

	// omni
	omn := functions.Omnisint("https://sonar.omnisint.io/subdomains/" + domain)

	// jldc.me
	jldc := functions.Jldcme("https://jldc.me/anubis/subdomains/" + domain)

	// abusedb.com
	abuse := functions.AbuseDb("https://www.abuseipdb.com/whois/" + domain)
	beautyabuse := commands.Sed2(abuse, domain)
	data := crt + beautyabuse + jldc + omn
	final := commands.Sort(data)
	return final
}
