import React from 'react';
import ReactDOM from 'react-dom/client';
import {createBrowserRouter, RouterProvider} from "react-router-dom";

import App from "./app";
import Calculator from "./calculator";
import Plot from "./plot";
import Reference from "./reference";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
    },
    {
        path: "/calculator",
        element: <Calculator/>
    },
    {
        path: "/plot",
        element: <Plot/>
    },
    {
        path: "/reference",
        element: <Reference/>
    }])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <RouterProvider router={router}/>
);