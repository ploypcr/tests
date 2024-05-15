import {Taxi} from './taxi'

describe('Taxi', () => {
  it('test kilo not rounded up and minutes rounded up', () => {
    const result = Taxi(1.5, 1.1)
    expect(result).toEqual(6)
  })
  it('test kilo and minutes rounded up', () => {
    const result = Taxi(1.6, 2.9)
    expect(result).toEqual(11)
  })
  it('test kilo and minutes not rounded up', () => {
    const result = Taxi(1.49, 1)
    expect(result).toEqual(5)
  })
  it('test kilo rounded up and minutes not rounded up', () => {
    const result = Taxi(1.51, 1)
    expect(result).toEqual(9)
  })
})
