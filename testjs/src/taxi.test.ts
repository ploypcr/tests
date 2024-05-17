import {fare, roundDistance, roundMinimumFare, roundWaitingTime} from './taxi'

describe('Taxi', () => {
  it('should round distance less than or equal .5 to .5', () => {
    const result = roundDistance(1.4)
    expect(result).toEqual(1.5)
  })
  it('should round distance more than .5 to .0', () => {
    const result = roundDistance(1.6)
    expect(result).toEqual(2)
  })

  it('should not round times', () => {
    const result = roundWaitingTime(0)
    expect(result).toEqual(0)
  })
  it('should round times', () => {
    const result = roundWaitingTime(0.1)
    expect(result).toEqual(1)
  })

  it('test calculate fare', () => {
    const result = fare(1.5, 0)
    expect(result).toEqual(6)
  })

  it('should round minimum fare', () => {
    const result = roundMinimumFare(10)
    expect(result).toEqual(35)
  })

  it('should not round minimum fare', () => {
    const result = roundMinimumFare(36)
    expect(result).toEqual(36)
  })

  it('test all', () => {
    const result = roundMinimumFare(fare(8.2,2))
    expect(result).toEqual(36)
  })
})
