import { describe, it, expect, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'

import EmailDetailHeaders from './EmailDetailHeaders.vue'

describe('EmailDetailHeaders', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(EmailDetailHeaders, {
      props: {
        message_id: '1234',
        date: '2004-02-03 18:14',
        from: 'martin@example.com',
        subject: 'Meet about the project',
      },
    })
  })

  it('should mount the EmailDetailHeaders component', () => {
    expect(EmailDetailHeaders).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should render the message_id prop in the second <span> element with v-html', () => {
    const spanElements = wrapper.findAll('span')
    expect(spanElements.length).toBeGreaterThan(1)

    const secondSpan = spanElements.at(1)
    expect(secondSpan).toBeDefined()

    if (secondSpan) {
      expect(secondSpan.element.innerHTML).toBe('1234')
    }
  })
})
