import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'

import PersonDetail from './PersonDetail.vue'

describe('PersonDetail', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(PersonDetail)
  })

  it('should mount the PersonDetail component', () => {
    expect(PersonDetail).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should render at less 3 <buttons> elements', () => {
    const buttonElements = wrapper.findAll('button')
    expect(buttonElements.length).toBeGreaterThanOrEqual(3)
  })
})
