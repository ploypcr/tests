export const Taxi = (kilo :number, minutes :number) => {
    var sec = (minutes * 60) % 60
    var min = Math.floor(minutes)
    if(sec > 0){
        min += 1
    }
    var k = kilo % 1
    kilo = Math.floor(kilo)
    if(k > 0.5){
        kilo += 1
    }
    return 4 * kilo + 1*min
}
