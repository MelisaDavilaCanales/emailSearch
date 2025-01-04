import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'

import BannerDataNotFound from './BannerDataNotFound.vue'

describe('BannerDataNotFound', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(BannerDataNotFound)
  })

  it('should mount the BannerDataNotFound component', () => {
    expect(BannerDataNotFound).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <div> element', () => {
    const divElement = wrapper.find('div')
    expect(divElement.exists()).toBe(true)
  })

  it('should render a <h4> element', () => {
    const h2Element = wrapper.find('h4')
    expect(h2Element.exists()).toBe(true)
  })

  it('should render a <p> element', () => {
    const pElement = wrapper.find('p')
    expect(pElement.exists()).toBe(true)
  })
})
