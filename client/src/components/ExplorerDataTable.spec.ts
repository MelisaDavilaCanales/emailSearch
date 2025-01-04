import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'
import ExplorerDataTable from './ExplorerDataTable.vue'

describe('ExplorerDataTable', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(ExplorerDataTable)
  })

  it('should mount the ExplorerDataTable component', () => {
    expect(ExplorerDataTable).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should render a <EmailTable> element whit initial state', () => {
    const emailTable = wrapper.findComponent({ name: 'EmailTable' })
    expect(emailTable.exists()).toBe(true)
  })

  it('should not render a <PersonTable> element whit initial state', () => {
    const personTable = wrapper.findComponent({ name: 'PersonTable' })
    expect(personTable.exists()).toBe(false)
  })
})
