import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'
import EmailDetail from './EmailDetail.vue'

describe('EmailDetail', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(EmailDetail)
  })

  it('should mount the EmailDetail component', () => {
    expect(EmailDetail).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  // it('should render a <span> element with the text Message ID', () => {
  //   const spanElement = wrapper.find('span')
  //   expect(spanElement.exists()).toBe(true)
  //   expect(spanElement.text()).toContain('Message ID')
  // })
})
