# Go Phone Directory CLI

Ce programme est une petite application en ligne de commande en Go permettant de gérer un répertoire téléphonique simple (recherche, ajout, suppression, mise à jour, etc.).

## 📦 Fonctionnalités

Le programme prend en charge plusieurs actions via des **paramètres passés en ligne de commande** :

### 🔧 Actions disponibles

- `--action list`  
  Liste toutes les personnes présentes dans le répertoire.

- `--action search --name <nom>`  
  Recherche une personne par son nom.

- `--action add --name <nom> --phone <téléphone>`  
  Ajoute une nouvelle personne au répertoire.

- `--action update --name <nom> --phone <téléphone>`  
  Met à jour le numéro de téléphone d’une personne.

- `--action delete --name <nom>`  
  Supprime une personne du répertoire.

## 📥 Exemple d’utilisation

```bash
go run main.go --action add --name "Charlie" --phone "0811223344"
go run main.go --action list
go run main.go --action search --name "ziad"
go run main.go --action update --name "ziad" --phone "07
```
## 📌 Remarque
Nous avons tenté d'ajouter un paramètre --age pour stocker l'âge des personnes dans le répertoire, mais cela n’a pas fonctionné comme prévu.

De même, une fonction de vérification des paramètres (paramVerif) avait été implémentée pour valider les entrées (nom, numéro, âge), mais elle n’a pas été correctement intégrée dans le flux d’exécution. 

Ces fonctionnalités ont été mises de côté pour le moment.
## 👥 Auteurs
- Ziad Lahrouni

- Sabrina Tamda