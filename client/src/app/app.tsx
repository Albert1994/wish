import './app.scss';
import Layout from "../component/layout/layout";
import {Routes, Route} from 'react-router-dom'
import React from "react"


function App() {
    return (
        <Routes>
            <Route path='/' element={<Layout/>}>
                <Route path="/" element={<div>HomePage</div>}/>
            </Route>
        </Routes>
    );
}

export default App;



