import os
from pathlib import Path
from dotenv import load_dotenv

from rasberry.data.api import ApiClient


class Config:

    def __init__(self):
        self.apiToken = None
        self.apiUrl = None
        self.config = None

    def config_loader(self):
        """
        Charge le fichier .env pour la configuration
        :return:
        """
        env_path = Path(__file__).resolve().parent.parent / ".env"
        load_dotenv(dotenv_path=env_path)
        self.apiToken = os.getenv("API_TOKEN")
        if not self.apiToken:
            raise Exception("API_TOKEN is not set")
        self.apiUrl = os.getenv("API_URL")
        if not self.apiUrl:
            raise Exception("API_URL is not set")

        self.set_config()
        return True

    def set_config(self):
        """
        Lors du premier appel à l'API set le nom de la station, adresse, etc...
        :return:
        """
        apiClient = ApiClient(self.apiToken, self.apiUrl)
        apiClient.who_am_i()
