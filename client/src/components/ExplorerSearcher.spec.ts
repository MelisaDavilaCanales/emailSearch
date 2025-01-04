import { it, describe, expect, beforeEach } from 'vitest'
import { VueWrapper, shallowMount } from '@vue/test-utils'

import ExplorerSearcher from './ExplorerSearcher.vue'

describe('ExplorerSearcher', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(ExplorerSearcher)
  })

  it('should mount the ExplorerSearcher component', () => {
    expect(ExplorerSearcher).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should render an <input> element of type="text"', () => {
    const inputElement = wrapper.find('input')
    expect(inputElement.exists()).toBe(true)
    expect(inputElement.attributes().type).toBe('text')
  })

  it('should render a <button> element with the text Search', () => {
    const buttonElement = wrapper.find('button')
    expect(buttonElement.exists()).toBe(true)
    expect(buttonElement.text()).toContain('Search')
  })

  it('should render at least 2 <input> elements of type="radio"', () => {
    const radioInputs = wrapper.findAll('input[type="radio"]')
    expect(radioInputs.length).toBeGreaterThanOrEqual(2)
  })
})
