import React from 'react'
import { Nav, Navbar as BootstrapNavbar } from 'react-bootstrap'
import NavbarLink from './NavbarLink'
import LoginButton from './LoginButton'
import LogoutButton from './LogoutButton'
import { useAuth0 } from '@auth0/auth0-react'

const Navbar = () => {
    const { isAuthenticated } = useAuth0()

    return (
        <BootstrapNavbar bg="dark" variant="dark">
            <Nav className="mr-auto">
                <NavbarLink text='Products' url='/dashboard/products'/>
                {
                    isAuthenticated &&
                    <NavbarLink text='Reports' url='/dashboard/reports'/>
                }
            </Nav>
            <Nav>
                {
                    isAuthenticated ? <LogoutButton/> : <LoginButton/>
                }
            </Nav>
        </BootstrapNavbar>
    )
}

export default Navbar