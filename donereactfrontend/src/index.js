import React, {Fragment} from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import Layout from "./Pages/home_page/layout/layout";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <Layout>
        <React.StrictMode>
            <App />
        </React.StrictMode>
    </Layout>

);
