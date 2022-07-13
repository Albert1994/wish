import './app.scss';
import Layout from "../component/layout/layout";
import {Routes, Route} from 'react-router-dom'
import React, {useState} from "react"
import HomePage from '../pages/home-page/home-page';
import RegisterPage from "../pages/register-page/register-page";
import LoginPage from "../pages/login-page/login";


function App() {

    const [user, setUser] = useState(null)


        if(user) {
           return   <Routes>
                <Route path='/' element={<Layout/>}>
                    <Route path="/user" element={<div>Вы зашли</div>}/>
                </Route>
            </Routes>
        } else {
           return (
               <Routes>
                   <Route path='/' element={<Layout/>}>
                       <Route path="/" element={<HomePage/>}/>
                       <Route path="/login" element={<LoginPage/>}/>
                       <Route path="/register" element={<RegisterPage/>}/>
                   </Route>
               </Routes>
           )
        }
}


export default App;



