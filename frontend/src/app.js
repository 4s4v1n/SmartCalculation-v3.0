import React from "react";
import {Link} from "react-router-dom";

import './app.css'

class App extends React.Component {

    render() {
        return (
            <div className="app-container">
                <Link to="/calculator">
                    <button className="button">calculator</button>
                </Link>
                <Link to="/plot">
                    <button className="button">plot</button>
                </Link>
                <Link to="/reference">
                    <button className="button">reference</button>
                </Link>
            </div>
        );
    }

}

export default App;