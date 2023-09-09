import {CartesianGrid, Legend, Line, LineChart, Tooltip, XAxis, YAxis} from "recharts";
import React from "react";

import './plot.css'

class Plot extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            expression: "0",
            begin: "0",
            end: "0",
            data: []
        }

        this.ref_expression = React.createRef()
        this.ref_begin = React.createRef()
        this.ref_end = React.createRef()
    }

    async plotButtonClicked() {
        const response = await fetch("http://localhost:8080/api/plot?" + new URLSearchParams({
            begin: this.ref_begin.current.value,
            end: this.ref_end.current.value,
            expression: this.ref_expression.current.value
        }))
        const json = await response.json()
        let ref = this.ref_expression.current

        if (response.status !== 200) {
            this.setState({
                expression: "invalid input"
            })
            ref.value = "invalid input"

            return
        }

        let abscissa = json["abscissa"]
        let ordinate = json["ordinate"]

        let data = []
        for (let i = 0; i < abscissa.length; i++) {
            data.push({
                name: abscissa[i],
                f: ordinate[i],
                amt: abscissa[i]
            })
        }

        this.setState({
            expression: this.state.expression,
            data: data
        })
    }


    render() {
        return (
            <div className="plot-container">
                <input className="plot-expression-input" ref={this.ref_expression} type="text"
                       defaultValue={this.state.expression}/>
                <br/>
                <label className="range-label">x begin:</label>
                <input className="range-input" ref={this.ref_begin} type="number"
                       defaultValue={this.state.begin}></input>
                <label className="range-label">x end:</label>
                <input className="range-input" ref={this.ref_end} type="number"
                       defaultValue={this.state.end}></input>
                <LineChart
                    width={600}
                    height={600}
                    margin={{
                        top: 20,
                    }}
                    data={this.state.data}
                >
                    <CartesianGrid strokeDasharray="3 3"/>
                    <XAxis dataKey="name"/>
                    <YAxis/>
                    <Tooltip/>
                    <Legend/>
                    <Line type="monotone" dataKey="f" stroke="#FF7F50" activeDot={{r: 8}}/>
                </LineChart>
                <br/>
                <button className="plot-button" onClick={() => this.plotButtonClicked()}>plot</button>
            </div>
        );
    }

}

export default Plot;