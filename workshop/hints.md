# Hints :)

## Go Unittests draaien
In principe is het draaien van unittests voor go zo eenvoudig als het uitvoeren van `go test`, maar omdat we in dit
project de unittests in een submodule hebben zitten moeten de tests uitgevoerd worden met `go test ./...` zodat er ook 
recursief in directories gezocht wordt.

## Go bouwen voor andere Operating Systems en architecturen.
`go build` zal standaard een applicatie bouwen die geschikt is voor jouw machine, maar soms wil je ook een applicatie 
bouwen voor een andere soort machine. Bijvoorbeeld een 32 bit Linux machine, of een RaspberryPi of een 64 bit Windows
machine. Dit kan je eenvoudig doen met behulp van omgevingsvariabelen. 

Een applicatie bouwen voor 32 bit windows doe je bijvoorbeeld zo: `GOOS=windows GOARCH=386 go build -o bin/HelloGo-386.exe`
en voor 64 bit windows `GOOS=windows GOARCH=AMD64 go build -o bin/HelloGo-386.exe`

## De applicatie draaien op een VM
Je hoeft niet uit te vogelen hoe je een applicatie installeert op een VM. Dat gaat te ver voor deze workshop, dus ik 
vertel het je hier. 
1. Kopieer de binary naar de VM
2. Kopieer de service file naar de VM
3. Maak op de vm de binary uitvoerbaar: `sudo chmod +x HelloGo-amd64`
4. Installeer de service file op de VM: `sudo cp hello-go.service /etc/systemd/system/hello-go.service && sudo systemctl daemon-reload`
5. (Her)start de service: `sudo service hello-go restart`
6. Verwijs poort 80 door naar poort 8080 `sudo iptables --flush -t nat && sudo iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 80 -j REDIRECT --to-port 8080 && sudo iptables-save`
7. Controleer of de applicatie draait en toegankelijk is: `curl --no-progress-meter $(curl --no-progress-meter https://api.ipify.org)/health-check`

Disclaimer: Dit is niet de netste manier om een applicatie te installeren, maar het werkt en dat is goed genoeg voor
deze workshop.

## SSH verbinding opzetten
Voor een deploy moet je een SSH verbinding maken met een server. In Voorbereiding.md lees je hoe je een Virtuele Machine
kan aanmaken in Azure. Je zal daarna een SSH verbinding moeten instellen in Azure Devops. Dat doe je onder 
*Project Settings -> Service Connections*. Deze SSH verbinding kan je daarna gebruiken in je pipeline om uit te rollen.
