from .dht11 import DHT11


class SensorFactory:

    @staticmethod
    def get_sensor(sensor_name: str, pin: int):
        """
        Factory des capteurs
        :param sensor_name:
        :param pin:
        :return:
        """
        match sensor_name:
            case "DHT11":
                return DHT11(pin)
        return None