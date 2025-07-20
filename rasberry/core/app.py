from .sensors.sensor_manager import SensorManager
from rasberry.utils.logger import logger
from rasberry.config.config import Config
from time import sleep

from rasberry.data.api import ApiClient


class WeatherStationApp:
    def __init__(self):
        self.sensors = SensorManager() # Gestion des capteurs
        self.api = None # Client de l'api
        self.config = Config() # Paramétre de la carte

    def run(self):
        """
        Runner de l'application
        :return:
        """
        # Ajouter le check de la config
        logger.info("Démarrage de la station météo")
        # chargement config
        statusConfig =  self.config.config_loader()

        if not statusConfig:
            raise Exception("Erreur de chargement de la configuration")
        logger.info("Configuration chargé")

        # Instanciation du client api
        self.api = ApiClient(self.config.apiToken, self.config.apiUrl)
        logger.info("API connecté")

        while True:
            try:
                data = self.sensors.read_all()
                self.api.send_data(data)
            except Exception as e:
                logger.error(f"[SensorManager] {e}")

            sleep(1800)