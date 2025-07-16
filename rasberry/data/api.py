import requests
from urllib3 import HTTPConnectionPool

from rasberry.utils.logger import logger


class ApiClient:
    def __init__(self, apiToken, apiUrl):
        self._apiToken = apiToken
        self._apiUrl = apiUrl


    def who_am_i(self):
        try:
            response = requests.get(f"{self._apiUrl}/api/v1/system/me")
            if response.status_code != 200:
                raise Exception(response)
            return response.json()
        except Exception as e:
            raise Exception("Connection à l'API refusé")

