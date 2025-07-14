from .sensors.sensor_manager import SensorManager
from ..utils.logger import logger
from time import sleep
class WeatherStationApp:
    def __init__(self):
        self.sensors = SensorManager() # Gestion des capteurs
        self.api = None # Client de l'api
        self.config = None # Paramétre de la carte

    def run(self):
        """
        Runner de l'application
        :return:
        """

        # Ajouter le check de la config
        logger.info("Démarrage de la station météo")
        while True:
            try:
                self.sensors.read_all()

            except Exception as e:
                logger.error(f"[SensorManager] {e}")

            sleep(30)