/*import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <div>
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
*/
import React, { useEffect, useState } from "react";
import { Layout, message } from "antd";
import axios from "axios";
import AppointmentForm from "./components/AppointmentForm";
import AppointmentTable from "./components/AppointmentTable";

const { Header, Content, Footer } = Layout;

const App: React.FC = () => {


  return (
    <Layout>
      <Header style={{ color: "white", textAlign: "center", fontSize: "24px" }}>
        Patient Appointment System
      </Header>
      <Content style={{ padding: "20px" }}>
        <AppointmentForm  />
        <AppointmentTable />
      </Content>
      <Footer style={{ textAlign: "center" }}>©2024 Patient Appointment Blockchain</Footer>
    </Layout>
  );
};

export default App;
