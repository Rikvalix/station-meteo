import board

from .sensor_factory import SensorFactory
from ...utils.logger import logger


class SensorManager:
    """
    Manager de tous les capteurs connectés
    """

    def __init__(self):
        self.sensor_factory = SensorFactory()
        self.sensors = [
            self.sensor_factory.get_sensor("DHT11",board.D4)
        ]
        logger.info(f"[SensorManager] {len(self.sensors)} capteur(s) initialisé(s).")

    def read_all(self):
        """
        Lire toutes les données et renvoyer un dictionnaire
        :return:
        """
        data = {}
        for sensor in self.sensors:
            try :
                sensor_data = sensor.read()
                if sensor_data:
                    data[sensor.name] = sensor_data
                    logger.debug(f"[SensorManager] {sensor.name} : {sensor_data}")
                else:
                    logger.warning(f"[SensorManager] {sensor.name} n'a rien renvoyé.")
            except Exception as e:
                logger.error(f"[SensorManager] {sensor.name} : {e}")

        return data if data else None

