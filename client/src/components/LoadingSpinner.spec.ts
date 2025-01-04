import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'

import LoadingSpinner from './LoadingSpinner.vue'

describe('LoadingSpinner', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(LoadingSpinner)
  })

  it('should mount the LoadingSpinner component', () => {
    expect(LoadingSpinner).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should render a <div> element with the role="status"', () => {
    const divElement = wrapper.find('div[role="status"]')
    expect(divElement.exists()).toBe(true)
    expect(divElement.attributes('role')).toBe('status')
  })

  it('should render a <svg> element', () => {
    const svgElement = wrapper.find('svg')
    expect(svgElement.exists()).toBe(true)
  })

  it('should render a <span> element with the text Loading...', () => {
    const spanElement = wrapper.find('span')
    expect(spanElement.exists()).toBe(true)
    expect(spanElement.text()).toContain('Loading...')
  })
})
