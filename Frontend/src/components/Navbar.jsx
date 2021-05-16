import React from 'react'
import { Nav, Navbar as BootstrapNavbar } from 'react-bootstrap'
import NavbarLink from './NavbarLink'

const Navbar = ({ authenticated }) => (
    <BootstrapNavbar bg="dark" variant="dark">
        <Nav className="mr-auto">
            <NavbarLink text='Products' url='/dashboard/products'/>
            {
                authenticated &&
                <NavbarLink text='Reports' url='/dashboard/reports'/>
            }
        </Nav>
        <Nav>
            <NavbarLink text={authenticated ? 'Logout' : 'Login'} url={authenticated ? '/logout' : '/login'}/>
        </Nav>
    </BootstrapNavbar>
)

export default Navbar