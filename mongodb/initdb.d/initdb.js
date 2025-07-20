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