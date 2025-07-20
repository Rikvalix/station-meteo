import requests


class ApiClient:
    def __init__(self, apiToken, apiUrl):
        self._apiToken = apiToken
        self._apiUrl = apiUrl
        self.session = requests.Session()
        self.session.headers.update({
            "Authorization": self._apiToken
        })
    def who_am_i(self):
        try:
            response = self.session.get(f"{self._apiUrl}")
            if response.status_code != 200:
                raise Exception(response)
            return response.json()
        except Exception as e:
            raise Exception("Connection à l'API refusé")

    def send_data(self, data):
        try:
            # Construction de la data
            payload = {
                "temperature": data["DHT-11"]["temperature"],
                "humidity": data["DHT-11"]["humidity"],
            }
            response = self.session.post(f"{self._apiUrl}/api/v1/station/measurements", json=payload)
            if response.status_code != 200:
                raise Exception(response)
            return response.json()
        except Exception as e:
            raise Exception("Erreur de l'envoi des données")
