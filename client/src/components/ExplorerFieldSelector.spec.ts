import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'

import ExplorerFieldSelector from './ExplorerFieldSelector.vue'

describe('ExplorerFieldSelector', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(ExplorerFieldSelector)
  })

  it('should mount the ExplorerFieldSelector component', () => {
    expect(ExplorerFieldSelector).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should render a <select> element whit default opcion Select field', () => {
    const selectElement = wrapper.find('select')
    expect(selectElement.exists()).toBe(true)
    expect(selectElement.text()).toContain('Select field')
  })
})
