from .sensors.sensor_manager import SensorManager
from rasberry.utils.logger import logger
from rasberry.config.config import Config
from time import sleep
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

        while True:
            try:
                self.sensors.read_all()
            except Exception as e:
                logger.error(f"[SensorManager] {e}")

            sleep(30)