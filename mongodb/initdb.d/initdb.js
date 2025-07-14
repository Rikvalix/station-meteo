db = db.getSiblingDB("station_meteo")

db.createCollection("mesures")
db.createCollection("stations")
