# kost84.github.io

# Portfolio Web Dynamique en Go

## Librairies utilisées
* **HTML5 / CSS3** : Structure sémantique et mise en page responsive.
* **Go (Golang)** : Création du serveur HTTP (utilisation exclusive de la librairie standard).
  * `net/http` : Gestion du serveur, des routes et des requêtes.
  * `html/template` : Moteur de rendu pour injecter les données dans le HTML.
  * `encoding/json` : Lecture et décodage de la "base de données" locale.

## Structure du projet
* `main.go` : Le cœur de l'application (serveur HTTP, routes et logique de décodage).
* `project.json` : Fichier de données contenant la liste et les détails de mes projets.
* `page/` : Dossier contenant les templates HTML (Accueil, Projets, Contact).
* `css/` : Feuilles de style pour le design du site.
* `imageetdoc/` : Fichiers statiques (images des projets, CV, etc.).

## Installation et Lancement

Pour tester ce projet en local sur votre machine, suivez ces étapes :

### 1. Prérequis
Assurez-vous d'avoir **Go** installé sur votre machine. 
Vous pouvez le vérifier en tapant `go version` dans votre terminal.

### 2. Démarrer le serveur
1. Clonez ce dépôt ou téléchargez les fichiers.
2. Ouvrez un terminal à la racine du projet (là où se trouve le fichier `main.go`).
3. Lancez la commande suivante pour démarrer le serveur :
   ```bash
   go run main.go