from .base_sensor import BaseSensor
import adafruit_dht
from ...utils.logger import logger

class DHT11(BaseSensor):
    def __init__(self, pin):
        super().__init__("DHT-11", pin)
        self.sensor = adafruit_dht.DHT11

    def read(self):
        try:
            temp = self.sensor.temperature or 0.0
            humidity = self.sensor.humidity or 0.0
            return {
                "temperature": temp,
                "humidity": humidity,
            }
        except Exception as e:
            logger.error(f"Sensor read failed: {e}")
            return {
                "temperature": None,
                "humidity": None
            }
