
import socket
import subprocess

# Adresse et port de Stratux
stratux_host = "192.168.10.1"  # Exemple, à confirmer
stratux_port = 2000  # Exemple, port NMEA

# Chemin vers l'exécutable Go
go_executable = "./parseAISmessage"

# Se connecter à Stratux
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((stratux_host, stratux_port))
    print("Connexion au flux Stratux...")
    while True:
        data = s.recv(1024)
        if data:
            try:
                # Décoder et afficher les données NMEA reçues
#                nmea_data = data.decode('ascii').strip()
#                print(f"Reçu de {stratux_host}: {nmea_data}")
#
                # Appeler l'exécutable Go et passer les données NMEA comme argument
                result = subprocess.run([go_executable, data], capture_output=True, text=True)
                
                # Afficher le résultat du programme Go
                print(f"Résultat du traitement Go : {result.stdout}")

            except Exception as e:
                print(f"Erreur de décodage : {e}")
# Se connecter à Stratux
with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((stratux_host, stratux_port))
    print("Connexion au flux Stratux...")
    while True:
        data = s.recv(1024)
        if data:
                try:
                        # Décoder et afficher les données NMEA reçues
                        parseAisMessage(data)
                        nmea_data = data.decode('ascii').strip()
                        print(f"Reçu de {addr}: {nmea_data}")

                        # Vous pouvez ici ajouter un traitement spécifique aux données NMEA
                        # Par exemple, traiter les messages NMEA, les stocker dans un fichier, etc.

                except Exception as e:
                        print(f"Erreur de décodage : {e}")data:
