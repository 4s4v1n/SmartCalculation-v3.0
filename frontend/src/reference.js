import React from "react";

import "./refernce.css"

class Reference extends React.Component {

    render() {
        return(
            <div className="reference-container">
                <label className="reference-label">Reference</label>
                <br/>
                <label className="reference-text">
                    The 4th version of calculator<br/>
                    Core wrote on C++<br/>
                    Backend wrote on Go<br/>
                    Frontend wrote on React JS<br/>
                    <br/>
                    Using is native. Pages location:<br/>
                    "/" - menu<br/>
                    "/calculator" - calculator<br/>
                    "/plot"       - plot<br/>
                    "/reference"  - documentation<br/>
                    <br/>
                    Made by Anton Savin
                 </label>
            </div>
        )
    }
}

export default Reference;