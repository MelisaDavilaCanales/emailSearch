import { describe, it, expect, beforeEach } from 'vitest'
import { VueWrapper, shallowMount } from '@vue/test-utils'

import EmailDetailMailList from './EmailDetailMailList.vue'

describe('EmailDetailMailList', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(EmailDetailMailList)
  })

  it('should mount EmailDetailMailList component', () => {
    expect(EmailDetailMailList).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })
})
