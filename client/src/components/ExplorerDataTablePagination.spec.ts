import { it, describe, expect, beforeEach } from 'vitest'
import { shallowMount, VueWrapper } from '@vue/test-utils'
import ExplorerDataTablePagination from './ExplorerDataTablePagination.vue'

describe('ExplorerDataTablePagination', () => {
  let wrapper: VueWrapper<any>

  beforeEach(() => {
    wrapper = shallowMount(ExplorerDataTablePagination, {
      props: {
        currentPage: 1,
        totalPages: 10,
        pageSize: 10,
      },
    })
  })

  it('should mount the ExplorerDataTablePagination component', () => {
    expect(ExplorerDataTablePagination).toBeTruthy()
    expect(wrapper.exists()).toBe(true)
  })

  it('should render correctly with props', () => {
    expect(wrapper.props().currentPage).toBe(1)
    expect(wrapper.props().totalPages).toBe(10)
    expect(wrapper.props().pageSize).toBe(10)
  })
})
