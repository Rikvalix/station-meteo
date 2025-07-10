class BaseSensor:
    def __init__(self, name, pin):
        self.name = name
        self.pin = pin

    def read(self):
        """
        Lire une mesure
        :return:
        """
        raise NotImplementedError("Impémentation dans une sous classe")