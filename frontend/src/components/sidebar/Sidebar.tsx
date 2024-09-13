import React from 'react'
import DesktopSidebar from './DesktopSidebar'
import MobileSidebar from './MobileSidebar'

type Props = {}

const Sidebar = (props: Props) => {
  return (
    <header className=''>
      <div className='hidden md:block'>
        <DesktopSidebar />
      </div>
      <div className='md:hidden'>
        <MobileSidebar />
      </div>
    </header>
  )
}

export default Sidebar