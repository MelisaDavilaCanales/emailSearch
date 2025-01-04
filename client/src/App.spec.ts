import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'
import App from './App.vue'

let wrapper: VueWrapper<any>

beforeEach(() => {
  wrapper = shallowMount(App)
})

describe('App.vue', () => {
  it('should mount the App component', () => {
    expect(App).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a header', () => {
    const header = wrapper.find('header')
    expect(header.exists()).toBe(true)
  })

  it('should render a nav inside the header', () => {
    const header = wrapper.find('header')
    const nav = header.find('nav')
    expect(nav.exists()).toBe(true)
  })

  it('should contain a RouterView placeholder', () => {
    const routerViewPlaceholder = wrapper.findComponent({ name: 'RouterView' })
    expect(routerViewPlaceholder.exists()).toBe(true)
  })
})
