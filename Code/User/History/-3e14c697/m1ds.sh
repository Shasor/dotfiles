#!/bin/bash

# Exécuter une commande vide pour simuler une activité
xdotool key Return

# Attendre un certain temps (ajuster la valeur de SLEEP_TIME en secondes)
sleep 7200

# Facultatif : Afficher un message indiquant la fin de la période d'activité forcée
notify-send "Période d'activité forcée terminée"
