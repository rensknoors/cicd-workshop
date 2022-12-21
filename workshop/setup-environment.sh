#!/usr/bin/env bash
set -e

echo "Wat is je naam? (bijvoorbeeld chris-jansen)"
read NAME
echo "Welk admin password wil je voor je VM?"
read ADMIN_PASS

if [ -z $NAME ]
then
  echo "Oops, je moet wel even je naam opgeven. Bijvoorbeeld 'chris-jansen'"
  exit 1
fi

if [ -z $ADMIN_PASS ]
then
  echo "Hey! geef eens een wachtwoord op."
  exit 1
fi

RG="ods-cicd-$NAME"

az group create --name "$RG" --location westeurope --output none
az deployment group create \
    --resource-group "$RG" \
    --template-file workshop/template/template.json \
    --parameters @workshop/template/parameters.json \
        virtualNetworkName="$RG-vnet" \
        virtualMachineRG="$RG" \
        adminPassword="$ADMIN_PASS" \
    --output none

echo "-----------------------"
echo "🎉 Ik heb zojuist een resource group '$RG' aangemaakt in azure met daar in een virtual machine waarop je je applicatie straks kan uitrollen."
echo "-----------------------"
echo "🚨🚨🚨🚨 LET OP! Deze vm is NIET beveiligd en hoewel een brute force aanval een tijdje zal duren omdat we geen standaard gebruikersnaam/password combinatie gebruiken is het prima mogelijk dat deze na verloop van tijd een bitcoin miner wordt. Met andere woorden:"
echo "Gooi je resource group even weg als je klaar bent met spelen"
echo "-----------------------"
echo "Ik zal even proberen het ip adres van je VM te achterhalen. lukt dat nou niet, kijk dan even in de web interface."
IP=$(az vm list-ip-addresses --resource-group "$RG" --name ods-cicd-vm --query "[].virtualMachine.network.publicIpAddresses[].ipAddress" --output table | sed '3q;d')
echo "Je IP is $IP. Je kan verbinding maken met je VM via het commando 'ssh ods-cicd-user@$IP' en dan wachtwoord $ADMIN_PASS opgeven"
