import React from "react";
import {Outlet} from "react-router-dom"
import Footer from "./footer/footer";
import Header from "./header/header";



function Layout() {


    return (
        <>
            <Header/>
            <Outlet/>
            <Footer/>
        </>
    )
}

export default Layout;