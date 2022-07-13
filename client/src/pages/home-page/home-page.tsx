import React from 'react';
import {Link} from "react-router-dom";
import Button, {ButtonVariant} from "../../component/ui/button/button";


const HomePage = () => {
    return (
        <div>
       <Link to={"/login"}><Button variant={ButtonVariant.primary}>Login</Button></Link>
        </div>
    );
};

export default HomePage;
