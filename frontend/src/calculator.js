import React from "react";

import "./calculator.css"

class Calculator extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            input: "0",
            x_input: "0"
        }
        this.ref_input = React.createRef()
        this.ref_x_input = React.createRef()
    }

    commonButtonClicked(value) {
        let ref = this.ref_input.current

        this.setState({
            input: this.state.input.value + value
        })

        if (ref.value === "0") {
            ref.value = ""
        }

        ref.value += value
    }

    async equalButtonClicked() {
        const response = await fetch("http://localhost:8080/api/calculator?" + new URLSearchParams({
            expression: this.ref_input.current.value,
            x: this.ref_x_input.current.value
        }))
        const json = await response.json()
        let ref = this.ref_input.current

        if (response.status !== 200) {
            this.setState({
                input: "invalid input",
                x_input: this.state.x_input
            })
            ref.value = "invalid input"

            return
        }

        this.setState({
            input: json["value"]
        })
        ref.value = json["value"]
    }

    clearButtonClicked(erase_all) {
        let ref = this.ref_input.current

        if (erase_all === true) {
            this.setState({
                input: "0",
                x_input: this.state.x_input
            })
            ref.value = "0"
            return
        }

        this.setState({
            input: ref.value.slice(0, -1),
            x_input: this.state.x_input
        })
        ref.value = ref.value.slice(0, -1)

        if (ref.value === "") {
            ref.value = "0"
        }
    }

    async previousExpressionClicked() {
        const response = await fetch("http://localhost:8080/api/previous_expression")
        const json = await response.json()

        if (response.status !== 200) {
            return
        }

        this.ref_input.current.value = json["expression"]
    }

    async clearHistoryClicked() {
         await fetch("http://localhost:8080/api/clear_history", {
             method: "PUT"
         })
    }

    render() {
        return (
            <div className="calculator-container">
                <input className="expression-input" ref={this.ref_input} type="text" defaultValue={this.state.input}/>
                <div>
                    <label className="x-label">x value:</label>
                    <input className="x-input" ref={this.ref_x_input} type="number"
                           defaultValue={this.state.x_input}></input>
                </div>
                <div>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(7)}>7</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(8)}>8</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(9)}>9</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked('+')}>+</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked('-')}>-</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked('*')}>*</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked('/')}>/</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(4)}>4</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(5)}>5</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(6)}>6</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked('^')}>^</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked('%')}>%</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked('(')}>(</button>
                    <button className="light-gray-button"
                            onClick={() => this.commonButtonClicked(')')}>)</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(1)}>1</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(2)}>2</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(3)}>3</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('sin')}>sin</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('cos')}>cos</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('tan')}>tan</button>
                    <button className="dark-gray-button"
                            onClick={() => this.clearButtonClicked(false)}>←</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked(0)}>0</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked('.')}>.</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked('x')}>x</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('asin')}>asin</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('acos')}>acos</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('atan')}>atan</button>
                    <button className="dark-gray-button"
                            onClick={() => this.clearButtonClicked(true)}>AC</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked('pi')}>π</button>
                    <button className="dark-gray-button"
                            onClick={() => this.commonButtonClicked('e')}>e</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('sqrt')}>√</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('ln')}>ln</button>
                    <button className="colored-button"
                            onClick={() => this.commonButtonClicked('lg')}>lg</button>
                    <button className="equal-button"
                            onClick={() => this.equalButtonClicked()}>=</button>
                    <button className="history-button"
                            onClick={() => this.previousExpressionClicked()}>previous expression</button>
                    <button className="history-button"
                            onClick={() => this.clearHistoryClicked()}>clear history</button>
                </div>
            </div>
        )
    }
}

export default Calculator;
