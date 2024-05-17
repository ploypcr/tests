export const Taxi = (kilo :number, minutes :number) => {
    return 4*kilo + minutes
}

export const roundMinimumFare = (fare : number) => {
    if (fare < 35){
        return 35
    }
    return fare
}

export const roundDistance = (distance : number) => {
    return Math.ceil(distance * 2)/2
}

export const roundWaitingTime = (minutes : number) => {
    return Math.ceil(minutes)
}

export const fare = (distance: number, minutes : number) => {
    return Taxi(roundDistance(distance), roundWaitingTime(minutes))
}