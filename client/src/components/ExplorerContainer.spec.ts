import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'
import ExplorerContainer from './ExplorerContainer.vue'

describe('ExplorerContainer', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(ExplorerContainer)
  })

  it('should mount the ExplorerContainer component', () => {
    expect(ExplorerContainer).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should contain a ExplorerSearcher component', () => {
    const explorerSearcher = wrapper.findComponent({ name: 'ExplorerSearcher' })
    expect(explorerSearcher.exists()).toBe(true)
  })

  it('should contain a ExplorerDataTable component', () => {
    const ExplorerDataTable = wrapper.findComponent({ name: 'ExplorerDataTable' })
    expect(ExplorerDataTable.exists()).toBe(true)
  })
})
