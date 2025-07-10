from .base_sensor import BaseSensor

class DHT11(BaseSensor):
    def __init__(self, pin):
        super().__init__("DHT-11", pin)
