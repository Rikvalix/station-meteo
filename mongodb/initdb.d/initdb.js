db = db.getSiblingDB("station_meteo")

db.createCollection("mesures")
db.createCollection("stations")
db.createCollection("utilisateurs")

// Utilisateur par défaut
// mp = beautifulPassword
db.utilisateurs.insertOne({
    username:"root",
    password:"$2a$12$JAxmd2Ha1c7iMAt2PDdtlORf/J9/kLvGgI/AN/53chAPlKmUaASja"
})

// Station par défaut
db.stations.insertOne({
    name: "Rasberry pi",
    auth_key: "ER&3423SD",
    location: "Chambre",
    address: "Quelque part dans le monde",
    components: [
    "DHT11"
],
    "enabled": true
})