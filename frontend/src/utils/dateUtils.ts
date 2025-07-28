import type MeasureModel from "../model/MeasureModel.ts";
import {format} from "date-fns";

export const dateUtils = {
    groupDate : (measures : MeasureModel[]) => {
        const dates : string[] = [];
        measures.forEach(measure => {
            const tempDate = format(new Date(measure.date), 'dd-MM-yyyy')
            if (!dates.includes(tempDate)) {
                dates.push(tempDate);
            }
        })
        return dates;
    }
}