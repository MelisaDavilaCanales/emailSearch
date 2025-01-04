import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'
import EmailTable from './EmailTable.vue'

describe('EmailTable', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(EmailTable)
  })

  it('should mount the EmailTable component', () => {
    expect(EmailTable).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })
})
