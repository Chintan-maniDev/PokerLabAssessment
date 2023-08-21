import "./Dashboard.css";
import axios from "axios";
import React, { useState } from "react";
import Pokerlogo from "../asset/poker-logo.png";

const Dashboard = () => {
  const [data, SetData] = useState("");
  const [result, setResult] = useState([]);

  const submitHandler = (e) => {
    e.preventDefault();
    axios.post('http://localhost:8080/evaluate', {
        text: data
        }, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
      .then(response => {
        console.log("rest ", response.data.Responses);
        setResult(response.data.Responses)
        SetData("")
      })
      .catch(error => {
        console.error('Error:', error);
      });
  };


  const output = result.map((d, index) => {
    return (
      <tr key={index}>
        <th scope="row">{index+1}</th>
        <td>{d.Hand}</td>
        <td>{d.HandType}</td>
        <td>{d.Rank}</td>
        <td>{d.UniqueRank}</td>
      </tr>
    );
  });

  
  
  return (
    <div className="container">
      <img className="container-img" src={Pokerlogo} alt="poker-logo" />
      <h1>Poker Evaluator</h1>
      <input
        type="text"
        value={data}
        onChange={(e) => SetData(e.target.value)}
        placeholder="Enter comma separated cards number"
      />
      <button type="submit" onClick={submitHandler}>
        Evaluate
      </button>
        <table className="table table-striped">
          <thead>
            <tr>
              <th scope="col">Sno.</th>
              <th scope="col">Input Hand</th>
              <th scope="col">Hand Type</th>
              <th scope="col">Rank</th>
              <th scope="col">Unique Rank</th>
            </tr>
          </thead>
          <tbody>{output}</tbody>
        </table>
    </div>
  );
};

export default Dashboard;
