import {TemperatureEnum} from "../model/enum/TemperatureEnum.ts";
import {HumidityEnum} from "../model/enum/HumidityEnum.ts";

export const weatherUtils = {
    getTemperatureCategory : (temp: number) : TemperatureEnum => {
        switch (true) {
            case temp < 16:
                return TemperatureEnum.Cold;
            case temp < 19:
                return TemperatureEnum.Cool;
            case temp <= 22:
                return TemperatureEnum.Comfortable;
            default:
                return TemperatureEnum.TooWarm;
        }
    },
    getHumidityCategory(humidity: number): HumidityEnum {
        switch (true) {
            case humidity < 30:
                return HumidityEnum.TooDry;
            case humidity <= 60:
                return HumidityEnum.Comfortable;
            case humidity > 60:
                return HumidityEnum.Humid;
            default:
                return HumidityEnum.TooHumid;
        }
    }
}