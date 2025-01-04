import { it, expect, describe, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'
import MainView from './MainView.vue'

let wrapper: VueWrapper<any>

beforeEach(() => {
  wrapper = shallowMount(MainView)
})

describe('MainView', () => {
  it('should mount the MainView component', () => {
    expect(MainView).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render a <main> element', () => {
    const mainElement = wrapper.find('main')
    expect(mainElement.exists()).toBe(true)
  })
})
